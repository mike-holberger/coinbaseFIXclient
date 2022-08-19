package ordercancelbatchrequest

import (
	"time"

	"coinbaseFIXclient/internal/field"
	"coinbaseFIXclient/internal/fix42"
	"coinbaseFIXclient/internal/tag"
	"github.com/quickfixgo/quickfix"
)

// OrderCancelBatchRequest is the fix42 OrderCancelBatchRequest type, MsgType = U4.
type OrderCancelBatchRequest struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

// FromMessage creates a OrderCancelBatchRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) OrderCancelBatchRequest {
	return OrderCancelBatchRequest{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m OrderCancelBatchRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a OrderCancelBatchRequest initialized with the required fields for OrderCancelBatchRequest.
func New(batchid field.BatchIDField) (m OrderCancelBatchRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("U4"))
	m.Set(batchid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg OrderCancelBatchRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "U4", r
}

// SetNoOrders sets NoOrders, Tag 73.
func (m OrderCancelBatchRequest) SetNoOrders(f NoOrdersRepeatingGroup) {
	m.SetGroup(f)
}

// SetBatchID sets BatchID, Tag 8014.
func (m OrderCancelBatchRequest) SetBatchID(v string) {
	m.Set(field.NewBatchID(v))
}

// GetNoOrders gets NoOrders, Tag 73.
func (m OrderCancelBatchRequest) GetNoOrders() (NoOrdersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetBatchID gets BatchID, Tag 8014.
func (m OrderCancelBatchRequest) GetBatchID() (v string, err quickfix.MessageRejectError) {
	var f field.BatchIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasNoOrders returns true if NoOrders is present, Tag 73.
func (m OrderCancelBatchRequest) HasNoOrders() bool {
	return m.Has(tag.NoOrders)
}

// HasBatchID returns true if BatchID is present, Tag 8014.
func (m OrderCancelBatchRequest) HasBatchID() bool {
	return m.Has(tag.BatchID)
}

// NoOrders is a repeating group element, Tag 73.
type NoOrders struct {
	*quickfix.Group
}

// SetOrigClOrdID sets OrigClOrdID, Tag 41.
func (m NoOrders) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m NoOrders) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetOrderID sets OrderID, Tag 37.
func (m NoOrders) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

// SetClOrdID sets ClOrdID, Tag 11.
func (m NoOrders) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m NoOrders) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// GetOrigClOrdID gets OrigClOrdID, Tag 41.
func (m NoOrders) GetOrigClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.OrigClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m NoOrders) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderID gets OrderID, Tag 37.
func (m NoOrders) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClOrdID gets ClOrdID, Tag 11.
func (m NoOrders) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m NoOrders) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41.
func (m NoOrders) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m NoOrders) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasOrderID returns true if OrderID is present, Tag 37.
func (m NoOrders) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

// HasClOrdID returns true if ClOrdID is present, Tag 11.
func (m NoOrders) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m NoOrders) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// NoOrdersRepeatingGroup is a repeating group, Tag 73.
type NoOrdersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoOrdersRepeatingGroup returns an initialized, NoOrdersRepeatingGroup.
func NewNoOrdersRepeatingGroup() NoOrdersRepeatingGroup {
	return NoOrdersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOrders,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.OrigClOrdID), quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.TransactTime)})}
}

// Add create and append a new NoOrders to this group.
func (m NoOrdersRepeatingGroup) Add() NoOrders {
	g := m.RepeatingGroup.Add()
	return NoOrders{g}
}

// Get returns the ith NoOrders in the NoOrdersRepeatinGroup.
func (m NoOrdersRepeatingGroup) Get(i int) NoOrders {
	return NoOrders{m.RepeatingGroup.Get(i)}
}
