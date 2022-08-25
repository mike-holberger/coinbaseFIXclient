package main

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"coinbaseFIXclient/internal/enum"
	"coinbaseFIXclient/internal/field"
	fix42neworderbatch "coinbaseFIXclient/internal/fix42/neworderbatch"
	fix42neworderlist "coinbaseFIXclient/internal/fix42/neworderlist"
	fix42nos "coinbaseFIXclient/internal/fix42/newordersingle"
	fix42ordercancel "coinbaseFIXclient/internal/fix42/ordercancelrequest"

	"github.com/google/uuid"
	"github.com/quickfixgo/quickfix"
	"github.com/rs/zerolog/log"
)

const (
	cbHostName = "fix.exchange.coinbase.com"
	cbPort     = "4198"
	cbtarget   = "Coinbase"

	cfgFileName = "quickfix.cfg"
)

// CoinbaseFIXclient implements the quickfix.Application interface
type CoinbaseFIXclient struct {
	initiator         *quickfix.Initiator
	key, secret, pass string
	execReports       *execReportCallbacks
}

type execReportCallbacks struct {
	mu          sync.Mutex
	reportChans []execReportChan
}

type execReportChan struct {
	clientID   string
	callbackCh chan ExecutionReport
}

func (e *CoinbaseFIXclient) Logon(key, secret, passphrase string) (err error) {
	e.key = key
	e.secret = secret
	e.pass = passphrase
	e.execReports = &execReportCallbacks{}

	var appSettings *quickfix.Settings

	cfg, err := os.Open(cfgFileName)
	if err != nil {
		// Apply default settings
		log.Debug().Msgf("%s file not found. Using default settings.", cfgFileName)
		appSettings = e.getDefaultSettings()
	} else {
		// Open config file and apply settings
		log.Debug().Msgf("Applying settings from config file: %s.", cfgFileName)
		stringData, readErr := ioutil.ReadAll(cfg)
		if readErr != nil {
			return fmt.Errorf("Error reading cfg: %s,", readErr)
		}

		appSettings, err = quickfix.ParseSettings(bytes.NewReader(stringData))
		if err != nil {
			return fmt.Errorf("Error ParseSettings from cfg: %s,", err)
		}

		cfg.Close()
	}

	fileLogFactory, err := quickfix.NewFileLogFactory(appSettings)
	if err != nil {
		return fmt.Errorf("Error creating file log factory: %s,", err)
	}

	e.initiator, err = quickfix.NewInitiator(*e, quickfix.NewMemoryStoreFactory(), appSettings, fileLogFactory)
	if err != nil {
		return fmt.Errorf("Unable to create Initiator: %s\n", err)
	}

	err = e.initiator.Start()
	if err != nil {
		return fmt.Errorf("Unable to start Initiator: %s\n", err)
	}

	// TODO: Logon callback

	return
}

func (e CoinbaseFIXclient) Logout() {
	e.initiator.Stop()

	// TODO: Logout callback
}

func (e CoinbaseFIXclient) NewOrderSingle(order CoinbaseOrderFIX, waitForExecReport bool, ctx context.Context) (execReport ExecutionReport, err error) {
	if order.ClientID == "" {
		err = fmt.Errorf("Order must contain a ClientID")
		return
	}

	nos := fix42nos.New(
		field.NewClOrdID(order.ClientID),
		field.NewSymbol(strings.ToUpper(order.Symbol)),
		field.NewSide(enum.Side(order.Side)),
		field.NewTransactTime(time.Now().UTC()),
		field.NewOrdType(enum.OrdType(order.OrderType)),
	)

	// Set price and Qty
	switch order.OrderType {
	case OrdType_MARKET:
		if order.CashOrderQty != "" {
			nos.SetString(152, order.CashOrderQty)
		} else {
			nos.SetString(38, order.Qty)
		}
	case OrdType_LIMIT:
		nos.SetString(44, order.Price)
		nos.SetString(38, order.Qty)
	case OrdType_STOP_LIMIT:
		nos.SetString(99, order.Price)
		nos.SetString(38, order.Qty)
	}

	// Set options
	if order.TimeInForce != "" {
		nos.SetString(59, string(order.TimeInForce))
	}

	if order.SelfTradePrevention != "" {
		nos.SetString(7928, string(order.SelfTradePrevention))
	}

	// Finalize and send
	msg := nos.ToMessage()
	msg.Header.Set(field.NewSenderCompID(e.key))
	msg.Header.Set(field.NewTargetCompID(cbtarget))

	if !waitForExecReport {
		err = quickfix.Send(msg)
		return
	}

	// Create callback channel and Wait for ExecReport
	callbackChan := make(chan ExecutionReport)

	e.execReports.mu.Lock()
	e.execReports.reportChans = append(e.execReports.reportChans, execReportChan{
		clientID:   order.ClientID,
		callbackCh: callbackChan,
	})
	e.execReports.mu.Unlock()

	err = quickfix.Send(msg)
	if err != nil {
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	select {
	case <-ctx.Done():
		e.execReports.mu.Lock()
		for i, reportChans := range e.execReports.reportChans {
			// Remove ExecReport chan
			if reportChans.clientID == order.ClientID {
				close(reportChans.callbackCh)

				// Remove from callback chans
				e.execReports.reportChans[i] = e.execReports.reportChans[len(e.execReports.reportChans)-1]
				e.execReports.reportChans = e.execReports.reportChans[:len(e.execReports.reportChans)-1]
				break
			}
		}
		e.execReports.mu.Unlock()

		err = fmt.Errorf("Single Order Context Timout")
		return
	case execReport = <-callbackChan:
		if execReport.OrdStatus == "8" {
			err = fmt.Errorf("Single Order Rejected: %s", execReport.Text)
			return
		}

		return
	}
}

// Fastest way to cancel an order. Must include ClientID, Symbol, and Side
func (e CoinbaseFIXclient) OrderCancel(order CoinbaseOrderFIX) (err error) {
	cancelID := uuid.New().String()

	println("??????: " + cancelID)

	oc := fix42ordercancel.New(
		field.NewOrigClOrdID(order.ClientID),
		field.NewClOrdID(cancelID),
		field.NewSymbol(order.Symbol),
		field.NewSide(enum.Side(order.Side)),
		field.NewTransactTime(time.Now().UTC()),
	)

	// Finalize and send
	msg := oc.ToMessage()
	msg.Header.Set(field.NewSenderCompID(e.key))
	msg.Header.Set(field.NewTargetCompID(cbtarget))

	//if !waitForExecReport {
	err = quickfix.Send(msg)
	return
	//}

	// // Create callback channel and Wait for ExecReport
	// callbackChan := make(chan ExecutionReport)

	// e.execReports.mu.Lock()
	// e.execReports.reportChans = append(e.execReports.reportChans, execReportChan{
	// 	clientID:   cancelID,
	// 	callbackCh: callbackChan,
	// })
	// e.execReports.mu.Unlock()

	// err = quickfix.Send(msg)
	// if err != nil {
	// 	return
	// }

	// if ctx == nil {
	// 	ctx = context.Background()
	// }

	// select {
	// case <-ctx.Done():
	// 	e.execReports.mu.Lock()
	// 	for i, reportChans := range e.execReports.reportChans {
	// 		// Remove ExecReport chan
	// 		if reportChans.clientID == cancelID {
	// 			close(reportChans.callbackCh)

	// 			// Remove from callback chans
	// 			e.execReports.reportChans[i] = e.execReports.reportChans[len(e.execReports.reportChans)-1]
	// 			e.execReports.reportChans = e.execReports.reportChans[:len(e.execReports.reportChans)-1]
	// 			break
	// 		}
	// 	}
	// 	e.execReports.mu.Unlock()

	// 	err = fmt.Errorf("Cancel Order Context Timout")
	// 	return
	// case execReport = <-callbackChan:
	// 	if execReport.OrdStatus == "8" {
	// 		err = fmt.Errorf("Cancel Order Rejected: %s", execReport.Text)
	// 		return
	// 	}

	// 	return
	// }
}

func (e CoinbaseFIXclient) OrderStatus() (err error) {

	return
}

// Beta only?
func (e CoinbaseFIXclient) ModifyOrder() (err error) {

	return
}

func (e CoinbaseFIXclient) NewOrdersBatch(batchID string, orders []CoinbaseOrderFIX, waitForExecReports bool, ctx context.Context) (execReports []ExecutionReport, err error) {
	nob := fix42neworderbatch.New(field.NewBatchID(batchID))
	nob.SetString(73, fmt.Sprintf("%d", len(orders)))

	group := fix42neworderlist.NewNoOrdersRepeatingGroup()

	for _, ord := range orders {
		g := group.Add()
		g.Set(field.NewClOrdID(ord.ClientID))
		g.SetString(55, strings.ToUpper(ord.Symbol))
		g.SetString(54, string(ord.Side))
		g.Set(field.NewOrdType(enum.OrdType_LIMIT))
		g.SetString(44, ord.Price)
		g.SetString(38, ord.Qty)
	}

	nob.SetGroup(group)

	msg := nob.ToMessage()
	msg.Header.Set(field.NewSenderCompID(e.key))
	msg.Header.Set(field.NewTargetCompID(cbtarget))

	if !waitForExecReports {
		err = quickfix.Send(msg)
		return
	}

	batchRejectChan := make(chan ExecutionReport)
	chans := []chan ExecutionReport{}

	// Add batch reject channel
	e.execReports.mu.Lock()
	e.execReports.reportChans = append(e.execReports.reportChans, execReportChan{
		clientID:   batchID,
		callbackCh: batchRejectChan,
	})

	// Add order execReport callback channels
	for _, order := range orders {
		callbackChan := make(chan ExecutionReport)
		chans = append(chans, callbackChan)

		e.execReports.reportChans = append(e.execReports.reportChans, execReportChan{
			clientID:   order.ClientID,
			callbackCh: callbackChan,
		})
	}
	e.execReports.mu.Unlock()

	// Send Msg
	err = quickfix.Send(msg)
	if err != nil {
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	// Wait for ExecReports
	for _, callbackChan := range chans {
		select {
		// Batch successful - collect reports and remove reject channel
		case report := <-callbackChan:
			// Remove reject chan
			e.execReports.mu.Lock()
			for i, reportChans := range e.execReports.reportChans {
				if reportChans.clientID == batchID {
					// Remove from callback chans
					e.execReports.reportChans[i] = e.execReports.reportChans[len(e.execReports.reportChans)-1]
					e.execReports.reportChans = e.execReports.reportChans[:len(e.execReports.reportChans)-1]
					continue
				}
			}
			e.execReports.mu.Unlock()

			execReports = append(execReports, report)

		// Batch rejected - collect reject report and remove execReport channels
		case rejct := <-batchRejectChan:
			e.execReports.mu.Lock()
		REJECT_ORDERS:
			for _, ord := range orders {
				for i, reportChans := range e.execReports.reportChans {
					if reportChans.clientID == ord.ClientID {
						// Remove from callback chans
						e.execReports.reportChans[i] = e.execReports.reportChans[len(e.execReports.reportChans)-1]
						e.execReports.reportChans = e.execReports.reportChans[:len(e.execReports.reportChans)-1]
						continue REJECT_ORDERS
					}
				}
			}
			e.execReports.mu.Unlock()

			execReports = append(execReports, rejct)
			err = fmt.Errorf(rejct.Text)
			return

		// Batch Context Timeout - close and remove all callback channels
		case <-ctx.Done():
			e.execReports.mu.Lock()
		CTX_ORDERS:
			for _, ord := range orders {
				for i, reportChans := range e.execReports.reportChans {
					// Remove reject chan
					if reportChans.clientID == batchID {
						close(reportChans.callbackCh)

						// Remove from callback chans
						e.execReports.reportChans[i] = e.execReports.reportChans[len(e.execReports.reportChans)-1]
						e.execReports.reportChans = e.execReports.reportChans[:len(e.execReports.reportChans)-1]
						continue
					}

					// Remove ExecReport chans
					if reportChans.clientID == ord.ClientID {
						close(reportChans.callbackCh)

						// Remove from callback chans
						e.execReports.reportChans[i] = e.execReports.reportChans[len(e.execReports.reportChans)-1]
						e.execReports.reportChans = e.execReports.reportChans[:len(e.execReports.reportChans)-1]
						continue CTX_ORDERS
					}
				}
			}
			e.execReports.mu.Unlock()

			err = fmt.Errorf("Batch Order Context Timout")
			return
		}
	}

	return
}

func (e CoinbaseFIXclient) OrderCancelBatch() (err error) {

	return
}

func (e CoinbaseFIXclient) getDefaultSettings() (appSettings *quickfix.Settings) {
	appSettings = &quickfix.Settings{}

	globalSettings := appSettings.GlobalSettings()
	globalSettings.Set("FileLogPath", "logs")
	globalSettings.Set("HeartBtInt", "30")
	globalSettings.Set("ResetOnLogon", "Y")
	globalSettings.Set("SenderCompID", e.key)
	globalSettings.Set("SocketConnectHost", cbHostName)
	globalSettings.Set("SocketConnectPort", cbPort)
	globalSettings.Set("SocketUseSSL", "Y")
	globalSettings.Set("TargetCompID", cbtarget)

	sessionSettings := &quickfix.SessionSettings{}
	sessionSettings.Set("BeginString", "FIX.4.2")

	appSettings.AddSession(sessionSettings)

	return
}

func (e CoinbaseFIXclient) sign(presign string, secret string) (rawData string, err error) {
	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return
	}

	mac := hmac.New(sha256.New, key)
	_, err = mac.Write([]byte(presign))
	if err != nil {
		return
	}

	rawData = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return
}

type CoinbaseOrderFIX struct {
	// User defined order ID
	ClientID string
	// Example: "ETH-USD"
	Symbol string
	// Enum: OrderSide_BUY or OrderSide_SELL
	Side Side
	// Enum: OrderType_MARKET, OrderType_LIMIT, or OrderType_STOP
	OrderType OrdType
	Price     string
	Qty       string
	// Order size in quote units (e.g., USD) (Market order only)
	CashOrderQty string
	// OPTIONAL Enum: OrderTimeInForce_GOOD_TILL_CANCEL (default), OrderTimeInForce_IMMEDIATE_OR_CANCEL, OrderTimeInForce_FILL_OR_KILL, or OrderTimeInForce_POST_ONLY
	TimeInForce enum.TimeInForce
	// OPTIONAL Enum: OrderTimeInForce_DECREMENT_AND_CANCEL (default), OrderTimeInForce_CANCEL_RESTING, OrderTimeInForce_CANCEL_INCOMING, or OrderTimeInForce_CANCEL_BOTH
	SelfTradePrevention enum.SelfTradePrevention
}

func getMsgType(msgType string) string {
	switch msgType {
	case "A":
		return "Logon"
	case "D":
		return "Order-Single"
	case "8":
		return "ExecReport"
	case "9":
		return "CancelReject-Single"
	case "5":
		return "Logout"
	case "U6":
		return "Orders-Batch"
	case "U7":
		return "BatchRejectReport"
	default:
		return msgType
	}
}
