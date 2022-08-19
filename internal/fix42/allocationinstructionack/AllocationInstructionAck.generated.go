package allocationinstructionack

import (
	"time"

	"coinbaseFIXclient/internal/enum"
	"coinbaseFIXclient/internal/field"
	"coinbaseFIXclient/internal/fix42"
	"coinbaseFIXclient/internal/tag"
	"github.com/quickfixgo/quickfix"
)

// AllocationInstructionAck is the fix42 AllocationInstructionAck type, MsgType = P.
type AllocationInstructionAck struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

// FromMessage creates a AllocationInstructionAck from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) AllocationInstructionAck {
	return AllocationInstructionAck{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m AllocationInstructionAck) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a AllocationInstructionAck initialized with the required fields for AllocationInstructionAck.
func New(allocid field.AllocIDField, tradedate field.TradeDateField, allocstatus field.AllocStatusField) (m AllocationInstructionAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("P"))
	m.Set(allocid)
	m.Set(tradedate)
	m.Set(allocstatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg AllocationInstructionAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "P", r
}

// SetText sets Text, Tag 58.
func (m AllocationInstructionAck) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m AllocationInstructionAck) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetAllocID sets AllocID, Tag 70.
func (m AllocationInstructionAck) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

// SetTradeDate sets TradeDate, Tag 75.
func (m AllocationInstructionAck) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

// SetExecBroker sets ExecBroker, Tag 76.
func (m AllocationInstructionAck) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

// SetAllocStatus sets AllocStatus, Tag 87.
func (m AllocationInstructionAck) SetAllocStatus(v enum.AllocStatus) {
	m.Set(field.NewAllocStatus(v))
}

// SetAllocRejCode sets AllocRejCode, Tag 88.
func (m AllocationInstructionAck) SetAllocRejCode(v enum.AllocRejCode) {
	m.Set(field.NewAllocRejCode(v))
}

// SetClientID sets ClientID, Tag 109.
func (m AllocationInstructionAck) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m AllocationInstructionAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m AllocationInstructionAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// GetText gets Text, Tag 58.
func (m AllocationInstructionAck) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m AllocationInstructionAck) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocID gets AllocID, Tag 70.
func (m AllocationInstructionAck) GetAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeDate gets TradeDate, Tag 75.
func (m AllocationInstructionAck) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExecBroker gets ExecBroker, Tag 76.
func (m AllocationInstructionAck) GetExecBroker() (v string, err quickfix.MessageRejectError) {
	var f field.ExecBrokerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocStatus gets AllocStatus, Tag 87.
func (m AllocationInstructionAck) GetAllocStatus() (v enum.AllocStatus, err quickfix.MessageRejectError) {
	var f field.AllocStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocRejCode gets AllocRejCode, Tag 88.
func (m AllocationInstructionAck) GetAllocRejCode() (v enum.AllocRejCode, err quickfix.MessageRejectError) {
	var f field.AllocRejCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClientID gets ClientID, Tag 109.
func (m AllocationInstructionAck) GetClientID() (v string, err quickfix.MessageRejectError) {
	var f field.ClientIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m AllocationInstructionAck) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m AllocationInstructionAck) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m AllocationInstructionAck) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m AllocationInstructionAck) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasAllocID returns true if AllocID is present, Tag 70.
func (m AllocationInstructionAck) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

// HasTradeDate returns true if TradeDate is present, Tag 75.
func (m AllocationInstructionAck) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

// HasExecBroker returns true if ExecBroker is present, Tag 76.
func (m AllocationInstructionAck) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

// HasAllocStatus returns true if AllocStatus is present, Tag 87.
func (m AllocationInstructionAck) HasAllocStatus() bool {
	return m.Has(tag.AllocStatus)
}

// HasAllocRejCode returns true if AllocRejCode is present, Tag 88.
func (m AllocationInstructionAck) HasAllocRejCode() bool {
	return m.Has(tag.AllocRejCode)
}

// HasClientID returns true if ClientID is present, Tag 109.
func (m AllocationInstructionAck) HasClientID() bool {
	return m.Has(tag.ClientID)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m AllocationInstructionAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m AllocationInstructionAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}
