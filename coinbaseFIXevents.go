package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/rs/zerolog/log"
)

// OnCreate implemented as part of Application interface
func (e CoinbaseFIXclient) OnCreate(sessionID quickfix.SessionID) {
	log.Debug().Interface("OnCreate", sessionID).Send()
}

// OnLogon implemented as part of Application interface
func (e CoinbaseFIXclient) OnLogon(sessionID quickfix.SessionID) {
	// Look for clientID in callback channels
	e.execReports.mu.Lock()
	for i, callback := range e.execReports.reportChans {
		if callback.clientID == "Logon" {
			callback.callbackCh <- ExecutionReport{}
			close(callback.callbackCh)

			// Remove from callback chans
			e.execReports.reportChans[i] = e.execReports.reportChans[len(e.execReports.reportChans)-1]
			e.execReports.reportChans = e.execReports.reportChans[:len(e.execReports.reportChans)-1]
			break
		}
	}
	e.execReports.mu.Unlock()

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

		rawData, err := sign(presign, e.secret)
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

	var clientID string

	switch msgType {
	case "8":
		// Order Execution Report
		clientID, _ = msg.Body.GetString(11)
	case "U7":
		// Batch Execution Report
		clientID, _ = msg.Body.GetString(8014)
	}

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

	return
}

func sign(presign string, secret string) (rawData string, err error) {
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
