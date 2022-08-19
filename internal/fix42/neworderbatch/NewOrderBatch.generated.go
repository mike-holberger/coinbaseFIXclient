package neworderbatch

import (
	"github.com/shopspring/decimal"
	"time"

	"coinbaseFIXclient/internal/enum"
	"coinbaseFIXclient/internal/field"
	"coinbaseFIXclient/internal/fix42"
	"coinbaseFIXclient/internal/tag"
	"github.com/quickfixgo/quickfix"
)

// NewOrderBatch is the fix42 NewOrderBatch type, MsgType = U6.
type NewOrderBatch struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

// FromMessage creates a NewOrderBatch from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) NewOrderBatch {
	return NewOrderBatch{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m NewOrderBatch) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a NewOrderBatch initialized with the required fields for NewOrderBatch.
func New(batchid field.BatchIDField) (m NewOrderBatch) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("U6"))
	m.Set(batchid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg NewOrderBatch, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "U6", r
}

// SetNoOrders sets NoOrders, Tag 73.
func (m NewOrderBatch) SetNoOrders(f NoOrdersRepeatingGroup) {
	m.SetGroup(f)
}

// SetBatchID sets BatchID, Tag 8014.
func (m NewOrderBatch) SetBatchID(v string) {
	m.Set(field.NewBatchID(v))
}

// GetNoOrders gets NoOrders, Tag 73.
func (m NewOrderBatch) GetNoOrders() (NoOrdersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetBatchID gets BatchID, Tag 8014.
func (m NewOrderBatch) GetBatchID() (v string, err quickfix.MessageRejectError) {
	var f field.BatchIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasNoOrders returns true if NoOrders is present, Tag 73.
func (m NewOrderBatch) HasNoOrders() bool {
	return m.Has(tag.NoOrders)
}

// HasBatchID returns true if BatchID is present, Tag 8014.
func (m NewOrderBatch) HasBatchID() bool {
	return m.Has(tag.BatchID)
}

// NoOrders is a repeating group element, Tag 73.
type NoOrders struct {
	*quickfix.Group
}

// SetClOrdID sets ClOrdID, Tag 11.
func (m NoOrders) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

// SetHandlInst sets HandlInst, Tag 21.
func (m NoOrders) SetHandlInst(v enum.HandlInst) {
	m.Set(field.NewHandlInst(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m NoOrders) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetSide sets Side, Tag 54.
func (m NoOrders) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

// SetPrice sets Price, Tag 44.
func (m NoOrders) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

// SetOrderQty sets OrderQty, Tag 38.
func (m NoOrders) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

// SetOrdType sets OrdType, Tag 40.
func (m NoOrders) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

// SetTimeInForce sets TimeInForce, Tag 59.
func (m NoOrders) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

// SetSelfTradePrevention sets SelfTradePrevention, Tag 7928.
func (m NoOrders) SetSelfTradePrevention(v enum.SelfTradePrevention) {
	m.Set(field.NewSelfTradePrevention(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m NoOrders) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// GetClOrdID gets ClOrdID, Tag 11.
func (m NoOrders) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetHandlInst gets HandlInst, Tag 21.
func (m NoOrders) GetHandlInst() (v enum.HandlInst, err quickfix.MessageRejectError) {
	var f field.HandlInstField
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

// GetSide gets Side, Tag 54.
func (m NoOrders) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPrice gets Price, Tag 44.
func (m NoOrders) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderQty gets OrderQty, Tag 38.
func (m NoOrders) GetOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrdType gets OrdType, Tag 40.
func (m NoOrders) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTimeInForce gets TimeInForce, Tag 59.
func (m NoOrders) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSelfTradePrevention gets SelfTradePrevention, Tag 7928.
func (m NoOrders) GetSelfTradePrevention() (v enum.SelfTradePrevention, err quickfix.MessageRejectError) {
	var f field.SelfTradePreventionField
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

// HasClOrdID returns true if ClOrdID is present, Tag 11.
func (m NoOrders) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

// HasHandlInst returns true if HandlInst is present, Tag 21.
func (m NoOrders) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m NoOrders) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasSide returns true if Side is present, Tag 54.
func (m NoOrders) HasSide() bool {
	return m.Has(tag.Side)
}

// HasPrice returns true if Price is present, Tag 44.
func (m NoOrders) HasPrice() bool {
	return m.Has(tag.Price)
}

// HasOrderQty returns true if OrderQty is present, Tag 38.
func (m NoOrders) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

// HasOrdType returns true if OrdType is present, Tag 40.
func (m NoOrders) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

// HasTimeInForce returns true if TimeInForce is present, Tag 59.
func (m NoOrders) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

// HasSelfTradePrevention returns true if SelfTradePrevention is present, Tag 7928.
func (m NoOrders) HasSelfTradePrevention() bool {
	return m.Has(tag.SelfTradePrevention)
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.HandlInst), quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.Price), quickfix.GroupElement(tag.OrderQty), quickfix.GroupElement(tag.OrdType), quickfix.GroupElement(tag.TimeInForce), quickfix.GroupElement(tag.SelfTradePrevention), quickfix.GroupElement(tag.TransactTime)})}
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
