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

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	fix42neworderlist "github.com/quickfixgo/fix42/neworderlist"
	fix42nos "github.com/quickfixgo/fix42/newordersingle"
	"github.com/quickfixgo/quickfix"
	"github.com/rs/zerolog/log"
)

const (
	cbHostName = "fix.exchange.coinbase.com"
	cbPort     = "4198"
	cbtarget   = "Coinbase"

	cfgFileName = "quickfix.cfg"
)

var messageDict = map[string]string{
	"A": "Logon",
	"D": "Order-Single",
	"8": "ExecReport",
	"5": "Logout",
}

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
	log.Info().Interface("OnCreate", sessionID).Send()
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
	msgType, err := msg.MsgType()
	if err != nil {
		log.Error().Err(err).Msg("ToApp msg.MsgType()")
	} else if t, ok := messageDict[msgType]; ok {
		msgType = t
	}
	log.Info().Str("MsgType", msgType).Str("FromAdmin", strings.Replace(msg.String(), "\x01", " ", -1)).Send()

	return nil
}

// ToAdmin implemented as part of Application interface
func (e CoinbaseFIXclient) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	// Custom Fields for Logon Msg
	if msg.IsMsgTypeOf("A") {
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

	msgType, err := msg.MsgType()
	if err != nil {
		log.Error().Err(err).Msg("ToApp msg.MsgType()")
	} else if t, ok := messageDict[msgType]; ok {
		msgType = t
	}

	log.Info().Str("MsgType", msgType).Str("ToAdmin", strings.Replace(msg.String(), "\x01", " ", -1)).Send()
}

// ToApp implemented as part of Application interface
func (e CoinbaseFIXclient) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) (err error) {
	msgType, err := msg.MsgType()
	if err != nil {
		log.Error().Err(err).Msg("ToApp msg.MsgType()")
	} else if t, ok := messageDict[msgType]; ok {
		msgType = t
	}

	log.Info().Str("MsgType", msgType).Str("ToApp", strings.Replace(msg.String(), "\x01", " ", -1)).Send()
	return
}

// FromApp implemented as part of Application interface. This is the callback for all Application level messages from the counter party.
func (e CoinbaseFIXclient) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	// Check for Execution Reports to send to response callback chans
	clientID, err := msg.Body.GetString(11)
	if msg.IsMsgTypeOf("8") && err == nil {
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
			}
		}
		e.execReports.mu.Unlock()
	}

	msgType, err := msg.MsgType()
	if err != nil {
		log.Error().Err(err).Msg("FromApp msg.MsgType()")
	} else if t, ok := messageDict[msgType]; ok {
		msgType = t
	}

	log.Info().Str("MsgType", msgType).Str("FromApp", strings.Replace(msg.String(), "\x01", " ", -1)).Send()
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

	return
}

func (e CoinbaseFIXclient) Logout() {
	e.initiator.Stop()
}

func (e CoinbaseFIXclient) NewOrderSingle(order CoinbaseOrderFIX, waitForExecReport bool, ctx context.Context) (execReport ExecutionReport, err error) {
	if order.ClientID == "" {
		err = fmt.Errorf("Order must contain a ClientID")
		return
	}

	orderSide := order.Side.getQFenum()
	ordtype := order.OrderType.getQFenum()

	nos := fix42nos.New(
		field.NewClOrdID(order.ClientID),
		field.NewHandlInst(enum.HandlInst_MANUAL_ORDER_BEST_EXECUTION),
		field.NewSymbol(strings.ToUpper(order.Symbol)),
		field.NewSide(orderSide),
		field.NewTransactTime(time.Now().UTC()),
		field.NewOrdType(ordtype),
	)

	// Set price and Qty
	switch {
	case order.OrderType == OrderType_MARKET && order.CashOrderQty != "":
		nos.SetString(152, order.CashOrderQty)
	case order.OrderType == OrderType_MARKET && order.CashOrderQty == "":
		nos.SetString(38, order.Qty)
	case order.OrderType == OrderType_LIMIT:
		nos.SetString(44, order.Price)
		nos.SetString(38, order.Qty)
	case order.OrderType == OrderType_STOP:
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
		err = fmt.Errorf("ExecutionReport Callback Timout")
		return
	case report := <-callbackChan:
		return report, nil
	}
}

func (e CoinbaseFIXclient) OrderCancel() (err error) {

	return
}

func (e CoinbaseFIXclient) OrderStatus() (err error) {

	return
}

// Beta only?
func (e CoinbaseFIXclient) ModifyOrder() (err error) {

	return
}

func (e CoinbaseFIXclient) NewOrdersBatch(orders []CoinbaseOrderFIX) (err error) {
	listID := field.ListIDField{}
	bidType := field.BidTypeField{}
	totnoOrds := field.TotNoOrdersField{}

	nol := fix42neworderlist.New(listID, bidType, totnoOrds)

	var group fix42neworderlist.NoOrdersRepeatingGroup

	for _, ord := range orders {
		g := group.Add()
		g.Set(field.NewClOrdID(ord.ClientID))

		nol.SetGroup(group)
	}

	///

	msg := nol.ToMessage()
	msg.Header.Set(field.NewSenderCompID(e.key))
	msg.Header.Set(field.NewTargetCompID(cbtarget))

	return quickfix.Send(msg)
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
	Side OrderSide
	// Enum: OrderType_MARKET, OrderType_LIMIT, or OrderType_STOP
	OrderType OrderType
	Price     string
	Qty       string
	// Order size in quote units (e.g., USD) (Market order only)
	CashOrderQty string
	// OPTIONAL Enum: OrderTimeInForce_GOOD_TILL_CANCEL (default), OrderTimeInForce_IMMEDIATE_OR_CANCEL, OrderTimeInForce_FILL_OR_KILL, or OrderTimeInForce_POST_ONLY
	TimeInForce OrderTimeInForce
	// OPTIONAL Enum: OrderTimeInForce_DECREMENT_AND_CANCEL (default), OrderTimeInForce_CANCEL_RESTING, OrderTimeInForce_CANCEL_INCOMING, or OrderTimeInForce_CANCEL_BOTH
	SelfTradePrevention OrderSelfTradePrevention
}
