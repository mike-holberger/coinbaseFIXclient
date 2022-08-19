package ioi

import (
	"github.com/shopspring/decimal"
	"time"

	"coinbaseFIXclient/internal/enum"
	"coinbaseFIXclient/internal/field"
	"coinbaseFIXclient/internal/fix42"
	"coinbaseFIXclient/internal/tag"
	"github.com/quickfixgo/quickfix"
)

// IOI is the fix42 IOI type, MsgType = 6.
type IOI struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

// FromMessage creates a IOI from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) IOI {
	return IOI{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m IOI) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a IOI initialized with the required fields for IOI.
func New(ioiid field.IOIidField, ioitranstype field.IOITransTypeField, symbol field.SymbolField, side field.SideField, ioishares field.IOISharesField) (m IOI) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("6"))
	m.Set(ioiid)
	m.Set(ioitranstype)
	m.Set(symbol)
	m.Set(side)
	m.Set(ioishares)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg IOI, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "6", r
}

// SetCurrency sets Currency, Tag 15.
func (m IOI) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetIDSource sets IDSource, Tag 22.
func (m IOI) SetIDSource(v enum.IDSource) {
	m.Set(field.NewIDSource(v))
}

// SetIOIid sets IOIid, Tag 23.
func (m IOI) SetIOIid(v string) {
	m.Set(field.NewIOIid(v))
}

// SetIOIQltyInd sets IOIQltyInd, Tag 25.
func (m IOI) SetIOIQltyInd(v enum.IOIQltyInd) {
	m.Set(field.NewIOIQltyInd(v))
}

// SetIOIRefID sets IOIRefID, Tag 26.
func (m IOI) SetIOIRefID(v string) {
	m.Set(field.NewIOIRefID(v))
}

// SetIOIShares sets IOIShares, Tag 27.
func (m IOI) SetIOIShares(v enum.IOIShares) {
	m.Set(field.NewIOIShares(v))
}

// SetIOITransType sets IOITransType, Tag 28.
func (m IOI) SetIOITransType(v enum.IOITransType) {
	m.Set(field.NewIOITransType(v))
}

// SetPrice sets Price, Tag 44.
func (m IOI) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m IOI) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSide sets Side, Tag 54.
func (m IOI) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m IOI) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetText sets Text, Tag 58.
func (m IOI) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m IOI) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetValidUntilTime sets ValidUntilTime, Tag 62.
func (m IOI) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65.
func (m IOI) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

// SetIssuer sets Issuer, Tag 106.
func (m IOI) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107.
func (m IOI) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// SetIOINaturalFlag sets IOINaturalFlag, Tag 130.
func (m IOI) SetIOINaturalFlag(v bool) {
	m.Set(field.NewIOINaturalFlag(v))
}

// SetURLLink sets URLLink, Tag 149.
func (m IOI) SetURLLink(v string) {
	m.Set(field.NewURLLink(v))
}

// SetSecurityType sets SecurityType, Tag 167.
func (m IOI) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetNoIOIQualifiers sets NoIOIQualifiers, Tag 199.
func (m IOI) SetNoIOIQualifiers(f NoIOIQualifiersRepeatingGroup) {
	m.SetGroup(f)
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200.
func (m IOI) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetPutOrCall sets PutOrCall, Tag 201.
func (m IOI) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

// SetStrikePrice sets StrikePrice, Tag 202.
func (m IOI) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetMaturityDay sets MaturityDay, Tag 205.
func (m IOI) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

// SetOptAttribute sets OptAttribute, Tag 206.
func (m IOI) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m IOI) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetNoRoutingIDs sets NoRoutingIDs, Tag 215.
func (m IOI) SetNoRoutingIDs(f NoRoutingIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetSpreadToBenchmark sets SpreadToBenchmark, Tag 218.
func (m IOI) SetSpreadToBenchmark(value decimal.Decimal, scale int32) {
	m.Set(field.NewSpreadToBenchmark(value, scale))
}

// SetBenchmark sets Benchmark, Tag 219.
func (m IOI) SetBenchmark(v enum.Benchmark) {
	m.Set(field.NewBenchmark(v))
}

// SetCouponRate sets CouponRate, Tag 223.
func (m IOI) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

// SetContractMultiplier sets ContractMultiplier, Tag 231.
func (m IOI) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

// SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348.
func (m IOI) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

// SetEncodedIssuer sets EncodedIssuer, Tag 349.
func (m IOI) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

// SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350.
func (m IOI) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

// SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351.
func (m IOI) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m IOI) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m IOI) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// GetCurrency gets Currency, Tag 15.
func (m IOI) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIDSource gets IDSource, Tag 22.
func (m IOI) GetIDSource() (v enum.IDSource, err quickfix.MessageRejectError) {
	var f field.IDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIOIid gets IOIid, Tag 23.
func (m IOI) GetIOIid() (v string, err quickfix.MessageRejectError) {
	var f field.IOIidField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIOIQltyInd gets IOIQltyInd, Tag 25.
func (m IOI) GetIOIQltyInd() (v enum.IOIQltyInd, err quickfix.MessageRejectError) {
	var f field.IOIQltyIndField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIOIRefID gets IOIRefID, Tag 26.
func (m IOI) GetIOIRefID() (v string, err quickfix.MessageRejectError) {
	var f field.IOIRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIOIShares gets IOIShares, Tag 27.
func (m IOI) GetIOIShares() (v enum.IOIShares, err quickfix.MessageRejectError) {
	var f field.IOISharesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIOITransType gets IOITransType, Tag 28.
func (m IOI) GetIOITransType() (v enum.IOITransType, err quickfix.MessageRejectError) {
	var f field.IOITransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPrice gets Price, Tag 44.
func (m IOI) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m IOI) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSide gets Side, Tag 54.
func (m IOI) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m IOI) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m IOI) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m IOI) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetValidUntilTime gets ValidUntilTime, Tag 62.
func (m IOI) GetValidUntilTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ValidUntilTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65.
func (m IOI) GetSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106.
func (m IOI) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107.
func (m IOI) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIOINaturalFlag gets IOINaturalFlag, Tag 130.
func (m IOI) GetIOINaturalFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.IOINaturalFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetURLLink gets URLLink, Tag 149.
func (m IOI) GetURLLink() (v string, err quickfix.MessageRejectError) {
	var f field.URLLinkField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167.
func (m IOI) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoIOIQualifiers gets NoIOIQualifiers, Tag 199.
func (m IOI) GetNoIOIQualifiers() (NoIOIQualifiersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoIOIQualifiersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200.
func (m IOI) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPutOrCall gets PutOrCall, Tag 201.
func (m IOI) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202.
func (m IOI) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDay gets MaturityDay, Tag 205.
func (m IOI) GetMaturityDay() (v int, err quickfix.MessageRejectError) {
	var f field.MaturityDayField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206.
func (m IOI) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m IOI) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRoutingIDs gets NoRoutingIDs, Tag 215.
func (m IOI) GetNoRoutingIDs() (NoRoutingIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRoutingIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetSpreadToBenchmark gets SpreadToBenchmark, Tag 218.
func (m IOI) GetSpreadToBenchmark() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SpreadToBenchmarkField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmark gets Benchmark, Tag 219.
func (m IOI) GetBenchmark() (v enum.Benchmark, err quickfix.MessageRejectError) {
	var f field.BenchmarkField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponRate gets CouponRate, Tag 223.
func (m IOI) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplier gets ContractMultiplier, Tag 231.
func (m IOI) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348.
func (m IOI) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuer gets EncodedIssuer, Tag 349.
func (m IOI) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350.
func (m IOI) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351.
func (m IOI) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m IOI) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m IOI) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m IOI) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasIDSource returns true if IDSource is present, Tag 22.
func (m IOI) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

// HasIOIid returns true if IOIid is present, Tag 23.
func (m IOI) HasIOIid() bool {
	return m.Has(tag.IOIid)
}

// HasIOIQltyInd returns true if IOIQltyInd is present, Tag 25.
func (m IOI) HasIOIQltyInd() bool {
	return m.Has(tag.IOIQltyInd)
}

// HasIOIRefID returns true if IOIRefID is present, Tag 26.
func (m IOI) HasIOIRefID() bool {
	return m.Has(tag.IOIRefID)
}

// HasIOIShares returns true if IOIShares is present, Tag 27.
func (m IOI) HasIOIShares() bool {
	return m.Has(tag.IOIShares)
}

// HasIOITransType returns true if IOITransType is present, Tag 28.
func (m IOI) HasIOITransType() bool {
	return m.Has(tag.IOITransType)
}

// HasPrice returns true if Price is present, Tag 44.
func (m IOI) HasPrice() bool {
	return m.Has(tag.Price)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m IOI) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSide returns true if Side is present, Tag 54.
func (m IOI) HasSide() bool {
	return m.Has(tag.Side)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m IOI) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasText returns true if Text is present, Tag 58.
func (m IOI) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m IOI) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasValidUntilTime returns true if ValidUntilTime is present, Tag 62.
func (m IOI) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65.
func (m IOI) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasIssuer returns true if Issuer is present, Tag 106.
func (m IOI) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107.
func (m IOI) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasIOINaturalFlag returns true if IOINaturalFlag is present, Tag 130.
func (m IOI) HasIOINaturalFlag() bool {
	return m.Has(tag.IOINaturalFlag)
}

// HasURLLink returns true if URLLink is present, Tag 149.
func (m IOI) HasURLLink() bool {
	return m.Has(tag.URLLink)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m IOI) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasNoIOIQualifiers returns true if NoIOIQualifiers is present, Tag 199.
func (m IOI) HasNoIOIQualifiers() bool {
	return m.Has(tag.NoIOIQualifiers)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200.
func (m IOI) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasPutOrCall returns true if PutOrCall is present, Tag 201.
func (m IOI) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202.
func (m IOI) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasMaturityDay returns true if MaturityDay is present, Tag 205.
func (m IOI) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206.
func (m IOI) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m IOI) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasNoRoutingIDs returns true if NoRoutingIDs is present, Tag 215.
func (m IOI) HasNoRoutingIDs() bool {
	return m.Has(tag.NoRoutingIDs)
}

// HasSpreadToBenchmark returns true if SpreadToBenchmark is present, Tag 218.
func (m IOI) HasSpreadToBenchmark() bool {
	return m.Has(tag.SpreadToBenchmark)
}

// HasBenchmark returns true if Benchmark is present, Tag 219.
func (m IOI) HasBenchmark() bool {
	return m.Has(tag.Benchmark)
}

// HasCouponRate returns true if CouponRate is present, Tag 223.
func (m IOI) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

// HasContractMultiplier returns true if ContractMultiplier is present, Tag 231.
func (m IOI) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

// HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348.
func (m IOI) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

// HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349.
func (m IOI) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

// HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350.
func (m IOI) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

// HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351.
func (m IOI) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m IOI) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m IOI) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// NoIOIQualifiers is a repeating group element, Tag 199.
type NoIOIQualifiers struct {
	*quickfix.Group
}

// SetIOIQualifier sets IOIQualifier, Tag 104.
func (m NoIOIQualifiers) SetIOIQualifier(v enum.IOIQualifier) {
	m.Set(field.NewIOIQualifier(v))
}

// GetIOIQualifier gets IOIQualifier, Tag 104.
func (m NoIOIQualifiers) GetIOIQualifier() (v enum.IOIQualifier, err quickfix.MessageRejectError) {
	var f field.IOIQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasIOIQualifier returns true if IOIQualifier is present, Tag 104.
func (m NoIOIQualifiers) HasIOIQualifier() bool {
	return m.Has(tag.IOIQualifier)
}

// NoIOIQualifiersRepeatingGroup is a repeating group, Tag 199.
type NoIOIQualifiersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoIOIQualifiersRepeatingGroup returns an initialized, NoIOIQualifiersRepeatingGroup.
func NewNoIOIQualifiersRepeatingGroup() NoIOIQualifiersRepeatingGroup {
	return NoIOIQualifiersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoIOIQualifiers,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.IOIQualifier)})}
}

// Add create and append a new NoIOIQualifiers to this group.
func (m NoIOIQualifiersRepeatingGroup) Add() NoIOIQualifiers {
	g := m.RepeatingGroup.Add()
	return NoIOIQualifiers{g}
}

// Get returns the ith NoIOIQualifiers in the NoIOIQualifiersRepeatinGroup.
func (m NoIOIQualifiersRepeatingGroup) Get(i int) NoIOIQualifiers {
	return NoIOIQualifiers{m.RepeatingGroup.Get(i)}
}

// NoRoutingIDs is a repeating group element, Tag 215.
type NoRoutingIDs struct {
	*quickfix.Group
}

// SetRoutingType sets RoutingType, Tag 216.
func (m NoRoutingIDs) SetRoutingType(v enum.RoutingType) {
	m.Set(field.NewRoutingType(v))
}

// SetRoutingID sets RoutingID, Tag 217.
func (m NoRoutingIDs) SetRoutingID(v string) {
	m.Set(field.NewRoutingID(v))
}

// GetRoutingType gets RoutingType, Tag 216.
func (m NoRoutingIDs) GetRoutingType() (v enum.RoutingType, err quickfix.MessageRejectError) {
	var f field.RoutingTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRoutingID gets RoutingID, Tag 217.
func (m NoRoutingIDs) GetRoutingID() (v string, err quickfix.MessageRejectError) {
	var f field.RoutingIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRoutingType returns true if RoutingType is present, Tag 216.
func (m NoRoutingIDs) HasRoutingType() bool {
	return m.Has(tag.RoutingType)
}

// HasRoutingID returns true if RoutingID is present, Tag 217.
func (m NoRoutingIDs) HasRoutingID() bool {
	return m.Has(tag.RoutingID)
}

// NoRoutingIDsRepeatingGroup is a repeating group, Tag 215.
type NoRoutingIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRoutingIDsRepeatingGroup returns an initialized, NoRoutingIDsRepeatingGroup.
func NewNoRoutingIDsRepeatingGroup() NoRoutingIDsRepeatingGroup {
	return NoRoutingIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRoutingIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RoutingType), quickfix.GroupElement(tag.RoutingID)})}
}

// Add create and append a new NoRoutingIDs to this group.
func (m NoRoutingIDsRepeatingGroup) Add() NoRoutingIDs {
	g := m.RepeatingGroup.Add()
	return NoRoutingIDs{g}
}

// Get returns the ith NoRoutingIDs in the NoRoutingIDsRepeatinGroup.
func (m NoRoutingIDsRepeatingGroup) Get(i int) NoRoutingIDs {
	return NoRoutingIDs{m.RepeatingGroup.Get(i)}
}
