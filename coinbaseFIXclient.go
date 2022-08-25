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
	"strconv"
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

// OnCreate implemented as part of Application interface
func (e CoinbaseFIXclient) OnCreate(sessionID quickfix.SessionID) {
	log.Debug().Interface("OnCreate", sessionID).Send()
}

// OnLogon implemented as part of Application interface
func (e CoinbaseFIXclient) OnLogon(sessionID quickfix.SessionID) {
	log.Info().Interface("OnLogon", sessionID).Send()
}

// OnLogout implemented as part of Application interface
func (e CoinbaseFIXclient) OnLogout(sessionID quickfix.SessionID) {
	log.Info().Interface("OnLogout", sessionID).Send()
}

// FromAdmin implemented as part of Application interface
func (e CoinbaseFIXclient) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	msgType, err := msg.Header.GetString(35)
	if err != nil {
		log.Error().Err(err).Msg("ToApp msg.MsgType()")
	}
	log.Debug().Str("MsgType", getMsgType(msgType)).Str("FromAdmin", strings.Replace(msg.String(), "\x01", " ", -1)).Send()

	return nil
}

// ToAdmin implemented as part of Application interface
func (e CoinbaseFIXclient) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	msgType, err := msg.Header.GetString(35)
	if err != nil {
		log.Error().Err(err).Msg("ToApp msg.MsgType()")
	}

	// Custom Fields for Logon Msg
	if msgType == "A" {
		// 98	EncryptMethod	Must be 0 (None)
		msg.Body.SetInt(98, 0)
		// 108	HeartBtInt	Must be 30 (seconds)
		msg.Body.SetInt(108, 30)
		// 554	Password	Client API passphrase
		msg.Body.SetString(554, e.pass)
		// 96	RawData	Client message signature (see below)
		// 8013	CancelOrdersOnDisconnect	Y: Cancel all open orders for the current profile; S: Cancel open orders placed during session
		msg.Body.SetString(8013, "S")
		// 9406	DropCopyFlag	If set to Y, execution reports will be generated for all user orders (defaults to Y)
		msg.Body.SetString(9406, "Y")

		// Override the time
		sendingTime := strconv.FormatInt(time.Now().Unix(), 10)
		msg.Header.SetString(52, sendingTime)

		presign := strings.Join([]string{
			sendingTime,
			"A",
			"1",
			e.key,
			cbtarget,
			e.pass,
		}, "\x01")

		rawData, err := e.sign(presign, e.secret)
		if err != nil {
			log.Error().Err(err).Msg("ToApp Logon Signature")
			return
		}
		msg.Body.SetString(96, rawData)
	}

	log.Debug().Str("MsgType", getMsgType(msgType)).Str("ToAdmin", strings.Replace(msg.String(), "\x01", " ", -1)).Send()
}

// ToApp implemented as part of Application interface
func (e CoinbaseFIXclient) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) (err error) {
	msgType, err := msg.Header.GetString(35)
	if err != nil {
		log.Error().Err(err).Msg("ToApp msg.MsgType()")
	}
	msgType = getMsgType(msgType)

	log.Debug().Str("MsgType", msgType).Str("ToApp", strings.Replace(msg.String(), "\x01", " ", -1)).Send()
	return
}

// FromApp implemented as part of Application interface. This is the callback for all Application level messages from the counter party.
func (e CoinbaseFIXclient) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	// Check for Execution Reports to send to response callback chans
	msgType, err := msg.Header.GetString(35)
	if err != nil {
		log.Error().Err(err).Msg("ToApp msg.MsgType()")
	}
	log.Debug().Str("MsgType", getMsgType(msgType)).Str("FromApp", strings.Replace(msg.String(), "\x01", " ", -1)).Send()

	switch msgType {
	case "8": // "9":
		// Single Order Execution Report
		clientID, _ := msg.Body.GetString(11)
		if err == nil {
			// Look for clientID in callback channels
			e.execReports.mu.Lock()
			for i, callback := range e.execReports.reportChans {
				if callback.clientID == clientID {
					// Unmarshal Execution report and send to chan for that clientID
					execReport := e.UnmarshalExecReport(msg.String())
					callback.callbackCh <- execReport
					close(callback.callbackCh)

					// Remove from callback chans
					e.execReports.reportChans[i] = e.execReports.reportChans[len(e.execReports.reportChans)-1]
					e.execReports.reportChans = e.execReports.reportChans[:len(e.execReports.reportChans)-1]
					break
				}
			}
			e.execReports.mu.Unlock()
		}
	case "U7":
		// Batch Execution Report
		batchID, _ := msg.Body.GetString(8014)
		// Look for batchID in reject channels
		e.execReports.mu.Lock()
		for i, callback := range e.execReports.reportChans {
			if callback.clientID == batchID {
				// Unmarshal Execution report and send to chan for that clientID
				rejectReport := e.UnmarshalExecReport(msg.String())
				callback.callbackCh <- rejectReport
				close(callback.callbackCh)

				// Remove from callback chans
				e.execReports.reportChans[i] = e.execReports.reportChans[len(e.execReports.reportChans)-1]
				e.execReports.reportChans = e.execReports.reportChans[:len(e.execReports.reportChans)-1]
				break
			}
		}
		e.execReports.mu.Unlock()
	}

	return
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
