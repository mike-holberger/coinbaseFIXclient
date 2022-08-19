package neworderbatchreject

import (
	"coinbaseFIXclient/internal/field"
	"coinbaseFIXclient/internal/fix42"
	"coinbaseFIXclient/internal/tag"
	"github.com/quickfixgo/quickfix"
)

// NewOrderBatchReject is the fix42 NewOrderBatchReject type, MsgType = U7.
type NewOrderBatchReject struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

// FromMessage creates a NewOrderBatchReject from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) NewOrderBatchReject {
	return NewOrderBatchReject{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m NewOrderBatchReject) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a NewOrderBatchReject initialized with the required fields for NewOrderBatchReject.
func New(batchid field.BatchIDField, text field.TextField) (m NewOrderBatchReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("U7"))
	m.Set(batchid)
	m.Set(text)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg NewOrderBatchReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "U7", r
}

// SetText sets Text, Tag 58.
func (m NewOrderBatchReject) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetBatchID sets BatchID, Tag 8014.
func (m NewOrderBatchReject) SetBatchID(v string) {
	m.Set(field.NewBatchID(v))
}

// GetText gets Text, Tag 58.
func (m NewOrderBatchReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBatchID gets BatchID, Tag 8014.
func (m NewOrderBatchReject) GetBatchID() (v string, err quickfix.MessageRejectError) {
	var f field.BatchIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m NewOrderBatchReject) HasText() bool {
	return m.Has(tag.Text)
}

// HasBatchID returns true if BatchID is present, Tag 8014.
func (m NewOrderBatchReject) HasBatchID() bool {
	return m.Has(tag.BatchID)
}
