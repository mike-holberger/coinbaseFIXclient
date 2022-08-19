package ordercancelbatchreject

import (
	"coinbaseFIXclient/internal/field"
	"coinbaseFIXclient/internal/fix42"
	"coinbaseFIXclient/internal/tag"
	"github.com/quickfixgo/quickfix"
)

// OrderCancelBatchReject is the fix42 OrderCancelBatchReject type, MsgType = U5.
type OrderCancelBatchReject struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

// FromMessage creates a OrderCancelBatchReject from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) OrderCancelBatchReject {
	return OrderCancelBatchReject{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m OrderCancelBatchReject) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a OrderCancelBatchReject initialized with the required fields for OrderCancelBatchReject.
func New(batchid field.BatchIDField) (m OrderCancelBatchReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("U5"))
	m.Set(batchid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg OrderCancelBatchReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "U5", r
}

// SetText sets Text, Tag 58.
func (m OrderCancelBatchReject) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetBatchID sets BatchID, Tag 8014.
func (m OrderCancelBatchReject) SetBatchID(v string) {
	m.Set(field.NewBatchID(v))
}

// GetText gets Text, Tag 58.
func (m OrderCancelBatchReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBatchID gets BatchID, Tag 8014.
func (m OrderCancelBatchReject) GetBatchID() (v string, err quickfix.MessageRejectError) {
	var f field.BatchIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m OrderCancelBatchReject) HasText() bool {
	return m.Has(tag.Text)
}

// HasBatchID returns true if BatchID is present, Tag 8014.
func (m OrderCancelBatchReject) HasBatchID() bool {
	return m.Has(tag.BatchID)
}
