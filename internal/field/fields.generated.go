package field

import (
	"coinbaseFIXclient/internal/enum"
	"coinbaseFIXclient/internal/tag"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"
)

// AccountField is a STRING field.
type AccountField struct{ quickfix.FIXString }

// Tag returns tag.Account (1).
func (f AccountField) Tag() quickfix.Tag { return tag.Account }

// NewAccount returns a new AccountField initialized with val.
func NewAccount(val string) AccountField {
	return AccountField{quickfix.FIXString(val)}
}

func (f AccountField) Value() string { return f.String() }

// AccruedInterestAmtField is a AMT field.
type AccruedInterestAmtField struct{ quickfix.FIXDecimal }

// Tag returns tag.AccruedInterestAmt (159).
func (f AccruedInterestAmtField) Tag() quickfix.Tag { return tag.AccruedInterestAmt }

// NewAccruedInterestAmt returns a new AccruedInterestAmtField initialized with val and scale.
func NewAccruedInterestAmt(val decimal.Decimal, scale int32) AccruedInterestAmtField {
	return AccruedInterestAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AccruedInterestAmtField) Value() (val decimal.Decimal) { return f.Decimal }

// AccruedInterestRateField is a FLOAT field.
type AccruedInterestRateField struct{ quickfix.FIXDecimal }

// Tag returns tag.AccruedInterestRate (158).
func (f AccruedInterestRateField) Tag() quickfix.Tag { return tag.AccruedInterestRate }

// NewAccruedInterestRate returns a new AccruedInterestRateField initialized with val and scale.
func NewAccruedInterestRate(val decimal.Decimal, scale int32) AccruedInterestRateField {
	return AccruedInterestRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AccruedInterestRateField) Value() (val decimal.Decimal) { return f.Decimal }

// AdjustmentField is a enum.Adjustment field.
type AdjustmentField struct{ quickfix.FIXString }

// Tag returns tag.Adjustment (334).
func (f AdjustmentField) Tag() quickfix.Tag { return tag.Adjustment }

func NewAdjustment(val enum.Adjustment) AdjustmentField {
	return AdjustmentField{quickfix.FIXString(val)}
}

func (f AdjustmentField) Value() enum.Adjustment { return enum.Adjustment(f.String()) }

// AdvIdField is a STRING field.
type AdvIdField struct{ quickfix.FIXString }

// Tag returns tag.AdvId (2).
func (f AdvIdField) Tag() quickfix.Tag { return tag.AdvId }

// NewAdvId returns a new AdvIdField initialized with val.
func NewAdvId(val string) AdvIdField {
	return AdvIdField{quickfix.FIXString(val)}
}

func (f AdvIdField) Value() string { return f.String() }

// AdvRefIDField is a STRING field.
type AdvRefIDField struct{ quickfix.FIXString }

// Tag returns tag.AdvRefID (3).
func (f AdvRefIDField) Tag() quickfix.Tag { return tag.AdvRefID }

// NewAdvRefID returns a new AdvRefIDField initialized with val.
func NewAdvRefID(val string) AdvRefIDField {
	return AdvRefIDField{quickfix.FIXString(val)}
}

func (f AdvRefIDField) Value() string { return f.String() }

// AdvSideField is a enum.AdvSide field.
type AdvSideField struct{ quickfix.FIXString }

// Tag returns tag.AdvSide (4).
func (f AdvSideField) Tag() quickfix.Tag { return tag.AdvSide }

func NewAdvSide(val enum.AdvSide) AdvSideField {
	return AdvSideField{quickfix.FIXString(val)}
}

func (f AdvSideField) Value() enum.AdvSide { return enum.AdvSide(f.String()) }

// AdvTransTypeField is a enum.AdvTransType field.
type AdvTransTypeField struct{ quickfix.FIXString }

// Tag returns tag.AdvTransType (5).
func (f AdvTransTypeField) Tag() quickfix.Tag { return tag.AdvTransType }

func NewAdvTransType(val enum.AdvTransType) AdvTransTypeField {
	return AdvTransTypeField{quickfix.FIXString(val)}
}

func (f AdvTransTypeField) Value() enum.AdvTransType { return enum.AdvTransType(f.String()) }

// AggregatedBookField is a BOOLEAN field.
type AggregatedBookField struct{ quickfix.FIXBoolean }

// Tag returns tag.AggregatedBook (266).
func (f AggregatedBookField) Tag() quickfix.Tag { return tag.AggregatedBook }

// NewAggregatedBook returns a new AggregatedBookField initialized with val.
func NewAggregatedBook(val bool) AggregatedBookField {
	return AggregatedBookField{quickfix.FIXBoolean(val)}
}

func (f AggregatedBookField) Value() bool { return f.Bool() }

// AllocAccountField is a STRING field.
type AllocAccountField struct{ quickfix.FIXString }

// Tag returns tag.AllocAccount (79).
func (f AllocAccountField) Tag() quickfix.Tag { return tag.AllocAccount }

// NewAllocAccount returns a new AllocAccountField initialized with val.
func NewAllocAccount(val string) AllocAccountField {
	return AllocAccountField{quickfix.FIXString(val)}
}

func (f AllocAccountField) Value() string { return f.String() }

// AllocAvgPxField is a PRICE field.
type AllocAvgPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.AllocAvgPx (153).
func (f AllocAvgPxField) Tag() quickfix.Tag { return tag.AllocAvgPx }

// NewAllocAvgPx returns a new AllocAvgPxField initialized with val and scale.
func NewAllocAvgPx(val decimal.Decimal, scale int32) AllocAvgPxField {
	return AllocAvgPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocAvgPxField) Value() (val decimal.Decimal) { return f.Decimal }

// AllocHandlInstField is a enum.AllocHandlInst field.
type AllocHandlInstField struct{ quickfix.FIXString }

// Tag returns tag.AllocHandlInst (209).
func (f AllocHandlInstField) Tag() quickfix.Tag { return tag.AllocHandlInst }

func NewAllocHandlInst(val enum.AllocHandlInst) AllocHandlInstField {
	return AllocHandlInstField{quickfix.FIXString(val)}
}

func (f AllocHandlInstField) Value() enum.AllocHandlInst { return enum.AllocHandlInst(f.String()) }

// AllocIDField is a STRING field.
type AllocIDField struct{ quickfix.FIXString }

// Tag returns tag.AllocID (70).
func (f AllocIDField) Tag() quickfix.Tag { return tag.AllocID }

// NewAllocID returns a new AllocIDField initialized with val.
func NewAllocID(val string) AllocIDField {
	return AllocIDField{quickfix.FIXString(val)}
}

func (f AllocIDField) Value() string { return f.String() }

// AllocLinkIDField is a STRING field.
type AllocLinkIDField struct{ quickfix.FIXString }

// Tag returns tag.AllocLinkID (196).
func (f AllocLinkIDField) Tag() quickfix.Tag { return tag.AllocLinkID }

// NewAllocLinkID returns a new AllocLinkIDField initialized with val.
func NewAllocLinkID(val string) AllocLinkIDField {
	return AllocLinkIDField{quickfix.FIXString(val)}
}

func (f AllocLinkIDField) Value() string { return f.String() }

// AllocLinkTypeField is a enum.AllocLinkType field.
type AllocLinkTypeField struct{ quickfix.FIXString }

// Tag returns tag.AllocLinkType (197).
func (f AllocLinkTypeField) Tag() quickfix.Tag { return tag.AllocLinkType }

func NewAllocLinkType(val enum.AllocLinkType) AllocLinkTypeField {
	return AllocLinkTypeField{quickfix.FIXString(val)}
}

func (f AllocLinkTypeField) Value() enum.AllocLinkType { return enum.AllocLinkType(f.String()) }

// AllocNetMoneyField is a AMT field.
type AllocNetMoneyField struct{ quickfix.FIXDecimal }

// Tag returns tag.AllocNetMoney (154).
func (f AllocNetMoneyField) Tag() quickfix.Tag { return tag.AllocNetMoney }

// NewAllocNetMoney returns a new AllocNetMoneyField initialized with val and scale.
func NewAllocNetMoney(val decimal.Decimal, scale int32) AllocNetMoneyField {
	return AllocNetMoneyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocNetMoneyField) Value() (val decimal.Decimal) { return f.Decimal }

// AllocPriceField is a PRICE field.
type AllocPriceField struct{ quickfix.FIXDecimal }

// Tag returns tag.AllocPrice (366).
func (f AllocPriceField) Tag() quickfix.Tag { return tag.AllocPrice }

// NewAllocPrice returns a new AllocPriceField initialized with val and scale.
func NewAllocPrice(val decimal.Decimal, scale int32) AllocPriceField {
	return AllocPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocPriceField) Value() (val decimal.Decimal) { return f.Decimal }

// AllocRejCodeField is a enum.AllocRejCode field.
type AllocRejCodeField struct{ quickfix.FIXString }

// Tag returns tag.AllocRejCode (88).
func (f AllocRejCodeField) Tag() quickfix.Tag { return tag.AllocRejCode }

func NewAllocRejCode(val enum.AllocRejCode) AllocRejCodeField {
	return AllocRejCodeField{quickfix.FIXString(val)}
}

func (f AllocRejCodeField) Value() enum.AllocRejCode { return enum.AllocRejCode(f.String()) }

// AllocSharesField is a QTY field.
type AllocSharesField struct{ quickfix.FIXDecimal }

// Tag returns tag.AllocShares (80).
func (f AllocSharesField) Tag() quickfix.Tag { return tag.AllocShares }

// NewAllocShares returns a new AllocSharesField initialized with val and scale.
func NewAllocShares(val decimal.Decimal, scale int32) AllocSharesField {
	return AllocSharesField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocSharesField) Value() (val decimal.Decimal) { return f.Decimal }

// AllocStatusField is a enum.AllocStatus field.
type AllocStatusField struct{ quickfix.FIXString }

// Tag returns tag.AllocStatus (87).
func (f AllocStatusField) Tag() quickfix.Tag { return tag.AllocStatus }

func NewAllocStatus(val enum.AllocStatus) AllocStatusField {
	return AllocStatusField{quickfix.FIXString(val)}
}

func (f AllocStatusField) Value() enum.AllocStatus { return enum.AllocStatus(f.String()) }

// AllocTextField is a STRING field.
type AllocTextField struct{ quickfix.FIXString }

// Tag returns tag.AllocText (161).
func (f AllocTextField) Tag() quickfix.Tag { return tag.AllocText }

// NewAllocText returns a new AllocTextField initialized with val.
func NewAllocText(val string) AllocTextField {
	return AllocTextField{quickfix.FIXString(val)}
}

func (f AllocTextField) Value() string { return f.String() }

// AllocTransTypeField is a enum.AllocTransType field.
type AllocTransTypeField struct{ quickfix.FIXString }

// Tag returns tag.AllocTransType (71).
func (f AllocTransTypeField) Tag() quickfix.Tag { return tag.AllocTransType }

func NewAllocTransType(val enum.AllocTransType) AllocTransTypeField {
	return AllocTransTypeField{quickfix.FIXString(val)}
}

func (f AllocTransTypeField) Value() enum.AllocTransType { return enum.AllocTransType(f.String()) }

// AvgPrxPrecisionField is a INT field.
type AvgPrxPrecisionField struct{ quickfix.FIXInt }

// Tag returns tag.AvgPrxPrecision (74).
func (f AvgPrxPrecisionField) Tag() quickfix.Tag { return tag.AvgPrxPrecision }

// NewAvgPrxPrecision returns a new AvgPrxPrecisionField initialized with val.
func NewAvgPrxPrecision(val int) AvgPrxPrecisionField {
	return AvgPrxPrecisionField{quickfix.FIXInt(val)}
}

func (f AvgPrxPrecisionField) Value() int { return f.Int() }

// AvgPxField is a PRICE field.
type AvgPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.AvgPx (6).
func (f AvgPxField) Tag() quickfix.Tag { return tag.AvgPx }

// NewAvgPx returns a new AvgPxField initialized with val and scale.
func NewAvgPx(val decimal.Decimal, scale int32) AvgPxField {
	return AvgPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AvgPxField) Value() (val decimal.Decimal) { return f.Decimal }

// BasisPxTypeField is a enum.BasisPxType field.
type BasisPxTypeField struct{ quickfix.FIXString }

// Tag returns tag.BasisPxType (419).
func (f BasisPxTypeField) Tag() quickfix.Tag { return tag.BasisPxType }

func NewBasisPxType(val enum.BasisPxType) BasisPxTypeField {
	return BasisPxTypeField{quickfix.FIXString(val)}
}

func (f BasisPxTypeField) Value() enum.BasisPxType { return enum.BasisPxType(f.String()) }

// BatchIDField is a STRING field.
type BatchIDField struct{ quickfix.FIXString }

// Tag returns tag.BatchID (8014).
func (f BatchIDField) Tag() quickfix.Tag { return tag.BatchID }

// NewBatchID returns a new BatchIDField initialized with val.
func NewBatchID(val string) BatchIDField {
	return BatchIDField{quickfix.FIXString(val)}
}

func (f BatchIDField) Value() string { return f.String() }

// BeginSeqNoField is a INT field.
type BeginSeqNoField struct{ quickfix.FIXInt }

// Tag returns tag.BeginSeqNo (7).
func (f BeginSeqNoField) Tag() quickfix.Tag { return tag.BeginSeqNo }

// NewBeginSeqNo returns a new BeginSeqNoField initialized with val.
func NewBeginSeqNo(val int) BeginSeqNoField {
	return BeginSeqNoField{quickfix.FIXInt(val)}
}

func (f BeginSeqNoField) Value() int { return f.Int() }

// BeginStringField is a STRING field.
type BeginStringField struct{ quickfix.FIXString }

// Tag returns tag.BeginString (8).
func (f BeginStringField) Tag() quickfix.Tag { return tag.BeginString }

// NewBeginString returns a new BeginStringField initialized with val.
func NewBeginString(val string) BeginStringField {
	return BeginStringField{quickfix.FIXString(val)}
}

func (f BeginStringField) Value() string { return f.String() }

// BenchmarkField is a enum.Benchmark field.
type BenchmarkField struct{ quickfix.FIXString }

// Tag returns tag.Benchmark (219).
func (f BenchmarkField) Tag() quickfix.Tag { return tag.Benchmark }

func NewBenchmark(val enum.Benchmark) BenchmarkField {
	return BenchmarkField{quickfix.FIXString(val)}
}

func (f BenchmarkField) Value() enum.Benchmark { return enum.Benchmark(f.String()) }

// BidDescriptorField is a STRING field.
type BidDescriptorField struct{ quickfix.FIXString }

// Tag returns tag.BidDescriptor (400).
func (f BidDescriptorField) Tag() quickfix.Tag { return tag.BidDescriptor }

// NewBidDescriptor returns a new BidDescriptorField initialized with val.
func NewBidDescriptor(val string) BidDescriptorField {
	return BidDescriptorField{quickfix.FIXString(val)}
}

func (f BidDescriptorField) Value() string { return f.String() }

// BidDescriptorTypeField is a INT field.
type BidDescriptorTypeField struct{ quickfix.FIXInt }

// Tag returns tag.BidDescriptorType (399).
func (f BidDescriptorTypeField) Tag() quickfix.Tag { return tag.BidDescriptorType }

// NewBidDescriptorType returns a new BidDescriptorTypeField initialized with val.
func NewBidDescriptorType(val int) BidDescriptorTypeField {
	return BidDescriptorTypeField{quickfix.FIXInt(val)}
}

func (f BidDescriptorTypeField) Value() int { return f.Int() }

// BidForwardPointsField is a PRICEOFFSET field.
type BidForwardPointsField struct{ quickfix.FIXDecimal }

// Tag returns tag.BidForwardPoints (189).
func (f BidForwardPointsField) Tag() quickfix.Tag { return tag.BidForwardPoints }

// NewBidForwardPoints returns a new BidForwardPointsField initialized with val and scale.
func NewBidForwardPoints(val decimal.Decimal, scale int32) BidForwardPointsField {
	return BidForwardPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidForwardPointsField) Value() (val decimal.Decimal) { return f.Decimal }

// BidIDField is a STRING field.
type BidIDField struct{ quickfix.FIXString }

// Tag returns tag.BidID (390).
func (f BidIDField) Tag() quickfix.Tag { return tag.BidID }

// NewBidID returns a new BidIDField initialized with val.
func NewBidID(val string) BidIDField {
	return BidIDField{quickfix.FIXString(val)}
}

func (f BidIDField) Value() string { return f.String() }

// BidPxField is a PRICE field.
type BidPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.BidPx (132).
func (f BidPxField) Tag() quickfix.Tag { return tag.BidPx }

// NewBidPx returns a new BidPxField initialized with val and scale.
func NewBidPx(val decimal.Decimal, scale int32) BidPxField {
	return BidPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidPxField) Value() (val decimal.Decimal) { return f.Decimal }

// BidRequestTransTypeField is a enum.BidRequestTransType field.
type BidRequestTransTypeField struct{ quickfix.FIXString }

// Tag returns tag.BidRequestTransType (374).
func (f BidRequestTransTypeField) Tag() quickfix.Tag { return tag.BidRequestTransType }

func NewBidRequestTransType(val enum.BidRequestTransType) BidRequestTransTypeField {
	return BidRequestTransTypeField{quickfix.FIXString(val)}
}

func (f BidRequestTransTypeField) Value() enum.BidRequestTransType {
	return enum.BidRequestTransType(f.String())
}

// BidSizeField is a QTY field.
type BidSizeField struct{ quickfix.FIXDecimal }

// Tag returns tag.BidSize (134).
func (f BidSizeField) Tag() quickfix.Tag { return tag.BidSize }

// NewBidSize returns a new BidSizeField initialized with val and scale.
func NewBidSize(val decimal.Decimal, scale int32) BidSizeField {
	return BidSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidSizeField) Value() (val decimal.Decimal) { return f.Decimal }

// BidSpotRateField is a PRICE field.
type BidSpotRateField struct{ quickfix.FIXDecimal }

// Tag returns tag.BidSpotRate (188).
func (f BidSpotRateField) Tag() quickfix.Tag { return tag.BidSpotRate }

// NewBidSpotRate returns a new BidSpotRateField initialized with val and scale.
func NewBidSpotRate(val decimal.Decimal, scale int32) BidSpotRateField {
	return BidSpotRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidSpotRateField) Value() (val decimal.Decimal) { return f.Decimal }

// BidTypeField is a INT field.
type BidTypeField struct{ quickfix.FIXInt }

// Tag returns tag.BidType (394).
func (f BidTypeField) Tag() quickfix.Tag { return tag.BidType }

// NewBidType returns a new BidTypeField initialized with val.
func NewBidType(val int) BidTypeField {
	return BidTypeField{quickfix.FIXInt(val)}
}

func (f BidTypeField) Value() int { return f.Int() }

// BodyLengthField is a INT field.
type BodyLengthField struct{ quickfix.FIXInt }

// Tag returns tag.BodyLength (9).
func (f BodyLengthField) Tag() quickfix.Tag { return tag.BodyLength }

// NewBodyLength returns a new BodyLengthField initialized with val.
func NewBodyLength(val int) BodyLengthField {
	return BodyLengthField{quickfix.FIXInt(val)}
}

func (f BodyLengthField) Value() int { return f.Int() }

// BrokerOfCreditField is a STRING field.
type BrokerOfCreditField struct{ quickfix.FIXString }

// Tag returns tag.BrokerOfCredit (92).
func (f BrokerOfCreditField) Tag() quickfix.Tag { return tag.BrokerOfCredit }

// NewBrokerOfCredit returns a new BrokerOfCreditField initialized with val.
func NewBrokerOfCredit(val string) BrokerOfCreditField {
	return BrokerOfCreditField{quickfix.FIXString(val)}
}

func (f BrokerOfCreditField) Value() string { return f.String() }

// BusinessRejectReasonField is a enum.BusinessRejectReason field.
type BusinessRejectReasonField struct{ quickfix.FIXString }

// Tag returns tag.BusinessRejectReason (380).
func (f BusinessRejectReasonField) Tag() quickfix.Tag { return tag.BusinessRejectReason }

func NewBusinessRejectReason(val enum.BusinessRejectReason) BusinessRejectReasonField {
	return BusinessRejectReasonField{quickfix.FIXString(val)}
}

func (f BusinessRejectReasonField) Value() enum.BusinessRejectReason {
	return enum.BusinessRejectReason(f.String())
}

// BusinessRejectRefIDField is a STRING field.
type BusinessRejectRefIDField struct{ quickfix.FIXString }

// Tag returns tag.BusinessRejectRefID (379).
func (f BusinessRejectRefIDField) Tag() quickfix.Tag { return tag.BusinessRejectRefID }

// NewBusinessRejectRefID returns a new BusinessRejectRefIDField initialized with val.
func NewBusinessRejectRefID(val string) BusinessRejectRefIDField {
	return BusinessRejectRefIDField{quickfix.FIXString(val)}
}

func (f BusinessRejectRefIDField) Value() string { return f.String() }

// BuyVolumeField is a QTY field.
type BuyVolumeField struct{ quickfix.FIXDecimal }

// Tag returns tag.BuyVolume (330).
func (f BuyVolumeField) Tag() quickfix.Tag { return tag.BuyVolume }

// NewBuyVolume returns a new BuyVolumeField initialized with val and scale.
func NewBuyVolume(val decimal.Decimal, scale int32) BuyVolumeField {
	return BuyVolumeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BuyVolumeField) Value() (val decimal.Decimal) { return f.Decimal }

// CancelOrdersOnDisconnectField is a enum.CancelOrdersOnDisconnect field.
type CancelOrdersOnDisconnectField struct{ quickfix.FIXString }

// Tag returns tag.CancelOrdersOnDisconnect (8013).
func (f CancelOrdersOnDisconnectField) Tag() quickfix.Tag { return tag.CancelOrdersOnDisconnect }

func NewCancelOrdersOnDisconnect(val enum.CancelOrdersOnDisconnect) CancelOrdersOnDisconnectField {
	return CancelOrdersOnDisconnectField{quickfix.FIXString(val)}
}

func (f CancelOrdersOnDisconnectField) Value() enum.CancelOrdersOnDisconnect {
	return enum.CancelOrdersOnDisconnect(f.String())
}

// CashOrderQtyField is a QTY field.
type CashOrderQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.CashOrderQty (152).
func (f CashOrderQtyField) Tag() quickfix.Tag { return tag.CashOrderQty }

// NewCashOrderQty returns a new CashOrderQtyField initialized with val and scale.
func NewCashOrderQty(val decimal.Decimal, scale int32) CashOrderQtyField {
	return CashOrderQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CashOrderQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// CashSettlAgentAcctNameField is a STRING field.
type CashSettlAgentAcctNameField struct{ quickfix.FIXString }

// Tag returns tag.CashSettlAgentAcctName (185).
func (f CashSettlAgentAcctNameField) Tag() quickfix.Tag { return tag.CashSettlAgentAcctName }

// NewCashSettlAgentAcctName returns a new CashSettlAgentAcctNameField initialized with val.
func NewCashSettlAgentAcctName(val string) CashSettlAgentAcctNameField {
	return CashSettlAgentAcctNameField{quickfix.FIXString(val)}
}

func (f CashSettlAgentAcctNameField) Value() string { return f.String() }

// CashSettlAgentAcctNumField is a STRING field.
type CashSettlAgentAcctNumField struct{ quickfix.FIXString }

// Tag returns tag.CashSettlAgentAcctNum (184).
func (f CashSettlAgentAcctNumField) Tag() quickfix.Tag { return tag.CashSettlAgentAcctNum }

// NewCashSettlAgentAcctNum returns a new CashSettlAgentAcctNumField initialized with val.
func NewCashSettlAgentAcctNum(val string) CashSettlAgentAcctNumField {
	return CashSettlAgentAcctNumField{quickfix.FIXString(val)}
}

func (f CashSettlAgentAcctNumField) Value() string { return f.String() }

// CashSettlAgentCodeField is a STRING field.
type CashSettlAgentCodeField struct{ quickfix.FIXString }

// Tag returns tag.CashSettlAgentCode (183).
func (f CashSettlAgentCodeField) Tag() quickfix.Tag { return tag.CashSettlAgentCode }

// NewCashSettlAgentCode returns a new CashSettlAgentCodeField initialized with val.
func NewCashSettlAgentCode(val string) CashSettlAgentCodeField {
	return CashSettlAgentCodeField{quickfix.FIXString(val)}
}

func (f CashSettlAgentCodeField) Value() string { return f.String() }

// CashSettlAgentContactNameField is a STRING field.
type CashSettlAgentContactNameField struct{ quickfix.FIXString }

// Tag returns tag.CashSettlAgentContactName (186).
func (f CashSettlAgentContactNameField) Tag() quickfix.Tag { return tag.CashSettlAgentContactName }

// NewCashSettlAgentContactName returns a new CashSettlAgentContactNameField initialized with val.
func NewCashSettlAgentContactName(val string) CashSettlAgentContactNameField {
	return CashSettlAgentContactNameField{quickfix.FIXString(val)}
}

func (f CashSettlAgentContactNameField) Value() string { return f.String() }

// CashSettlAgentContactPhoneField is a STRING field.
type CashSettlAgentContactPhoneField struct{ quickfix.FIXString }

// Tag returns tag.CashSettlAgentContactPhone (187).
func (f CashSettlAgentContactPhoneField) Tag() quickfix.Tag { return tag.CashSettlAgentContactPhone }

// NewCashSettlAgentContactPhone returns a new CashSettlAgentContactPhoneField initialized with val.
func NewCashSettlAgentContactPhone(val string) CashSettlAgentContactPhoneField {
	return CashSettlAgentContactPhoneField{quickfix.FIXString(val)}
}

func (f CashSettlAgentContactPhoneField) Value() string { return f.String() }

// CashSettlAgentNameField is a STRING field.
type CashSettlAgentNameField struct{ quickfix.FIXString }

// Tag returns tag.CashSettlAgentName (182).
func (f CashSettlAgentNameField) Tag() quickfix.Tag { return tag.CashSettlAgentName }

// NewCashSettlAgentName returns a new CashSettlAgentNameField initialized with val.
func NewCashSettlAgentName(val string) CashSettlAgentNameField {
	return CashSettlAgentNameField{quickfix.FIXString(val)}
}

func (f CashSettlAgentNameField) Value() string { return f.String() }

// CheckSumField is a STRING field.
type CheckSumField struct{ quickfix.FIXString }

// Tag returns tag.CheckSum (10).
func (f CheckSumField) Tag() quickfix.Tag { return tag.CheckSum }

// NewCheckSum returns a new CheckSumField initialized with val.
func NewCheckSum(val string) CheckSumField {
	return CheckSumField{quickfix.FIXString(val)}
}

func (f CheckSumField) Value() string { return f.String() }

// ClOrdIDField is a STRING field.
type ClOrdIDField struct{ quickfix.FIXString }

// Tag returns tag.ClOrdID (11).
func (f ClOrdIDField) Tag() quickfix.Tag { return tag.ClOrdID }

// NewClOrdID returns a new ClOrdIDField initialized with val.
func NewClOrdID(val string) ClOrdIDField {
	return ClOrdIDField{quickfix.FIXString(val)}
}

func (f ClOrdIDField) Value() string { return f.String() }

// ClearingAccountField is a STRING field.
type ClearingAccountField struct{ quickfix.FIXString }

// Tag returns tag.ClearingAccount (440).
func (f ClearingAccountField) Tag() quickfix.Tag { return tag.ClearingAccount }

// NewClearingAccount returns a new ClearingAccountField initialized with val.
func NewClearingAccount(val string) ClearingAccountField {
	return ClearingAccountField{quickfix.FIXString(val)}
}

func (f ClearingAccountField) Value() string { return f.String() }

// ClearingFirmField is a STRING field.
type ClearingFirmField struct{ quickfix.FIXString }

// Tag returns tag.ClearingFirm (439).
func (f ClearingFirmField) Tag() quickfix.Tag { return tag.ClearingFirm }

// NewClearingFirm returns a new ClearingFirmField initialized with val.
func NewClearingFirm(val string) ClearingFirmField {
	return ClearingFirmField{quickfix.FIXString(val)}
}

func (f ClearingFirmField) Value() string { return f.String() }

// ClientBidIDField is a STRING field.
type ClientBidIDField struct{ quickfix.FIXString }

// Tag returns tag.ClientBidID (391).
func (f ClientBidIDField) Tag() quickfix.Tag { return tag.ClientBidID }

// NewClientBidID returns a new ClientBidIDField initialized with val.
func NewClientBidID(val string) ClientBidIDField {
	return ClientBidIDField{quickfix.FIXString(val)}
}

func (f ClientBidIDField) Value() string { return f.String() }

// ClientIDField is a STRING field.
type ClientIDField struct{ quickfix.FIXString }

// Tag returns tag.ClientID (109).
func (f ClientIDField) Tag() quickfix.Tag { return tag.ClientID }

// NewClientID returns a new ClientIDField initialized with val.
func NewClientID(val string) ClientIDField {
	return ClientIDField{quickfix.FIXString(val)}
}

func (f ClientIDField) Value() string { return f.String() }

// CommTypeField is a enum.CommType field.
type CommTypeField struct{ quickfix.FIXString }

// Tag returns tag.CommType (13).
func (f CommTypeField) Tag() quickfix.Tag { return tag.CommType }

func NewCommType(val enum.CommType) CommTypeField {
	return CommTypeField{quickfix.FIXString(val)}
}

func (f CommTypeField) Value() enum.CommType { return enum.CommType(f.String()) }

// CommissionField is a AMT field.
type CommissionField struct{ quickfix.FIXDecimal }

// Tag returns tag.Commission (12).
func (f CommissionField) Tag() quickfix.Tag { return tag.Commission }

// NewCommission returns a new CommissionField initialized with val and scale.
func NewCommission(val decimal.Decimal, scale int32) CommissionField {
	return CommissionField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CommissionField) Value() (val decimal.Decimal) { return f.Decimal }

// ComplianceIDField is a STRING field.
type ComplianceIDField struct{ quickfix.FIXString }

// Tag returns tag.ComplianceID (376).
func (f ComplianceIDField) Tag() quickfix.Tag { return tag.ComplianceID }

// NewComplianceID returns a new ComplianceIDField initialized with val.
func NewComplianceID(val string) ComplianceIDField {
	return ComplianceIDField{quickfix.FIXString(val)}
}

func (f ComplianceIDField) Value() string { return f.String() }

// ContraBrokerField is a STRING field.
type ContraBrokerField struct{ quickfix.FIXString }

// Tag returns tag.ContraBroker (375).
func (f ContraBrokerField) Tag() quickfix.Tag { return tag.ContraBroker }

// NewContraBroker returns a new ContraBrokerField initialized with val.
func NewContraBroker(val string) ContraBrokerField {
	return ContraBrokerField{quickfix.FIXString(val)}
}

func (f ContraBrokerField) Value() string { return f.String() }

// ContraTradeQtyField is a QTY field.
type ContraTradeQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.ContraTradeQty (437).
func (f ContraTradeQtyField) Tag() quickfix.Tag { return tag.ContraTradeQty }

// NewContraTradeQty returns a new ContraTradeQtyField initialized with val and scale.
func NewContraTradeQty(val decimal.Decimal, scale int32) ContraTradeQtyField {
	return ContraTradeQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ContraTradeQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// ContraTradeTimeField is a UTCTIMESTAMP field.
type ContraTradeTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.ContraTradeTime (438).
func (f ContraTradeTimeField) Tag() quickfix.Tag { return tag.ContraTradeTime }

// NewContraTradeTime returns a new ContraTradeTimeField initialized with val.
func NewContraTradeTime(val time.Time) ContraTradeTimeField {
	return NewContraTradeTimeWithPrecision(val, quickfix.Millis)
}

// NewContraTradeTimeNoMillis returns a new ContraTradeTimeField initialized with val without millisecs.
func NewContraTradeTimeNoMillis(val time.Time) ContraTradeTimeField {
	return NewContraTradeTimeWithPrecision(val, quickfix.Seconds)
}

// NewContraTradeTimeWithPrecision returns a new ContraTradeTimeField initialized with val of specified precision.
func NewContraTradeTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) ContraTradeTimeField {
	return ContraTradeTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f ContraTradeTimeField) Value() time.Time { return f.Time }

// ContraTraderField is a STRING field.
type ContraTraderField struct{ quickfix.FIXString }

// Tag returns tag.ContraTrader (337).
func (f ContraTraderField) Tag() quickfix.Tag { return tag.ContraTrader }

// NewContraTrader returns a new ContraTraderField initialized with val.
func NewContraTrader(val string) ContraTraderField {
	return ContraTraderField{quickfix.FIXString(val)}
}

func (f ContraTraderField) Value() string { return f.String() }

// ContractMultiplierField is a FLOAT field.
type ContractMultiplierField struct{ quickfix.FIXDecimal }

// Tag returns tag.ContractMultiplier (231).
func (f ContractMultiplierField) Tag() quickfix.Tag { return tag.ContractMultiplier }

// NewContractMultiplier returns a new ContractMultiplierField initialized with val and scale.
func NewContractMultiplier(val decimal.Decimal, scale int32) ContractMultiplierField {
	return ContractMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ContractMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

// CorporateActionField is a enum.CorporateAction field.
type CorporateActionField struct{ quickfix.FIXString }

// Tag returns tag.CorporateAction (292).
func (f CorporateActionField) Tag() quickfix.Tag { return tag.CorporateAction }

func NewCorporateAction(val enum.CorporateAction) CorporateActionField {
	return CorporateActionField{quickfix.FIXString(val)}
}

func (f CorporateActionField) Value() enum.CorporateAction { return enum.CorporateAction(f.String()) }

// CountryField is a STRING field.
type CountryField struct{ quickfix.FIXString }

// Tag returns tag.Country (421).
func (f CountryField) Tag() quickfix.Tag { return tag.Country }

// NewCountry returns a new CountryField initialized with val.
func NewCountry(val string) CountryField {
	return CountryField{quickfix.FIXString(val)}
}

func (f CountryField) Value() string { return f.String() }

// CouponRateField is a FLOAT field.
type CouponRateField struct{ quickfix.FIXDecimal }

// Tag returns tag.CouponRate (223).
func (f CouponRateField) Tag() quickfix.Tag { return tag.CouponRate }

// NewCouponRate returns a new CouponRateField initialized with val and scale.
func NewCouponRate(val decimal.Decimal, scale int32) CouponRateField {
	return CouponRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CouponRateField) Value() (val decimal.Decimal) { return f.Decimal }

// CoveredOrUncoveredField is a enum.CoveredOrUncovered field.
type CoveredOrUncoveredField struct{ quickfix.FIXString }

// Tag returns tag.CoveredOrUncovered (203).
func (f CoveredOrUncoveredField) Tag() quickfix.Tag { return tag.CoveredOrUncovered }

func NewCoveredOrUncovered(val enum.CoveredOrUncovered) CoveredOrUncoveredField {
	return CoveredOrUncoveredField{quickfix.FIXString(val)}
}

func (f CoveredOrUncoveredField) Value() enum.CoveredOrUncovered {
	return enum.CoveredOrUncovered(f.String())
}

// CrossPercentField is a FLOAT field.
type CrossPercentField struct{ quickfix.FIXDecimal }

// Tag returns tag.CrossPercent (413).
func (f CrossPercentField) Tag() quickfix.Tag { return tag.CrossPercent }

// NewCrossPercent returns a new CrossPercentField initialized with val and scale.
func NewCrossPercent(val decimal.Decimal, scale int32) CrossPercentField {
	return CrossPercentField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CrossPercentField) Value() (val decimal.Decimal) { return f.Decimal }

// CumQtyField is a QTY field.
type CumQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.CumQty (14).
func (f CumQtyField) Tag() quickfix.Tag { return tag.CumQty }

// NewCumQty returns a new CumQtyField initialized with val and scale.
func NewCumQty(val decimal.Decimal, scale int32) CumQtyField {
	return CumQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CumQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// CurrencyField is a CURRENCY field.
type CurrencyField struct{ quickfix.FIXString }

// Tag returns tag.Currency (15).
func (f CurrencyField) Tag() quickfix.Tag { return tag.Currency }

// NewCurrency returns a new CurrencyField initialized with val.
func NewCurrency(val string) CurrencyField {
	return CurrencyField{quickfix.FIXString(val)}
}

func (f CurrencyField) Value() string { return f.String() }

// CustomerOrFirmField is a enum.CustomerOrFirm field.
type CustomerOrFirmField struct{ quickfix.FIXString }

// Tag returns tag.CustomerOrFirm (204).
func (f CustomerOrFirmField) Tag() quickfix.Tag { return tag.CustomerOrFirm }

func NewCustomerOrFirm(val enum.CustomerOrFirm) CustomerOrFirmField {
	return CustomerOrFirmField{quickfix.FIXString(val)}
}

func (f CustomerOrFirmField) Value() enum.CustomerOrFirm { return enum.CustomerOrFirm(f.String()) }

// CxlQtyField is a QTY field.
type CxlQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.CxlQty (84).
func (f CxlQtyField) Tag() quickfix.Tag { return tag.CxlQty }

// NewCxlQty returns a new CxlQtyField initialized with val and scale.
func NewCxlQty(val decimal.Decimal, scale int32) CxlQtyField {
	return CxlQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CxlQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// CxlRejReasonField is a enum.CxlRejReason field.
type CxlRejReasonField struct{ quickfix.FIXString }

// Tag returns tag.CxlRejReason (102).
func (f CxlRejReasonField) Tag() quickfix.Tag { return tag.CxlRejReason }

func NewCxlRejReason(val enum.CxlRejReason) CxlRejReasonField {
	return CxlRejReasonField{quickfix.FIXString(val)}
}

func (f CxlRejReasonField) Value() enum.CxlRejReason { return enum.CxlRejReason(f.String()) }

// CxlRejResponseToField is a enum.CxlRejResponseTo field.
type CxlRejResponseToField struct{ quickfix.FIXString }

// Tag returns tag.CxlRejResponseTo (434).
func (f CxlRejResponseToField) Tag() quickfix.Tag { return tag.CxlRejResponseTo }

func NewCxlRejResponseTo(val enum.CxlRejResponseTo) CxlRejResponseToField {
	return CxlRejResponseToField{quickfix.FIXString(val)}
}

func (f CxlRejResponseToField) Value() enum.CxlRejResponseTo {
	return enum.CxlRejResponseTo(f.String())
}

// CxlTypeField is a CHAR field.
type CxlTypeField struct{ quickfix.FIXString }

// Tag returns tag.CxlType (125).
func (f CxlTypeField) Tag() quickfix.Tag { return tag.CxlType }

// NewCxlType returns a new CxlTypeField initialized with val.
func NewCxlType(val string) CxlTypeField {
	return CxlTypeField{quickfix.FIXString(val)}
}

func (f CxlTypeField) Value() string { return f.String() }

// DKReasonField is a enum.DKReason field.
type DKReasonField struct{ quickfix.FIXString }

// Tag returns tag.DKReason (127).
func (f DKReasonField) Tag() quickfix.Tag { return tag.DKReason }

func NewDKReason(val enum.DKReason) DKReasonField {
	return DKReasonField{quickfix.FIXString(val)}
}

func (f DKReasonField) Value() enum.DKReason { return enum.DKReason(f.String()) }

// DayAvgPxField is a PRICE field.
type DayAvgPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.DayAvgPx (426).
func (f DayAvgPxField) Tag() quickfix.Tag { return tag.DayAvgPx }

// NewDayAvgPx returns a new DayAvgPxField initialized with val and scale.
func NewDayAvgPx(val decimal.Decimal, scale int32) DayAvgPxField {
	return DayAvgPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DayAvgPxField) Value() (val decimal.Decimal) { return f.Decimal }

// DayCumQtyField is a QTY field.
type DayCumQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.DayCumQty (425).
func (f DayCumQtyField) Tag() quickfix.Tag { return tag.DayCumQty }

// NewDayCumQty returns a new DayCumQtyField initialized with val and scale.
func NewDayCumQty(val decimal.Decimal, scale int32) DayCumQtyField {
	return DayCumQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DayCumQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// DayOrderQtyField is a QTY field.
type DayOrderQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.DayOrderQty (424).
func (f DayOrderQtyField) Tag() quickfix.Tag { return tag.DayOrderQty }

// NewDayOrderQty returns a new DayOrderQtyField initialized with val and scale.
func NewDayOrderQty(val decimal.Decimal, scale int32) DayOrderQtyField {
	return DayOrderQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DayOrderQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// DefBidSizeField is a QTY field.
type DefBidSizeField struct{ quickfix.FIXDecimal }

// Tag returns tag.DefBidSize (293).
func (f DefBidSizeField) Tag() quickfix.Tag { return tag.DefBidSize }

// NewDefBidSize returns a new DefBidSizeField initialized with val and scale.
func NewDefBidSize(val decimal.Decimal, scale int32) DefBidSizeField {
	return DefBidSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DefBidSizeField) Value() (val decimal.Decimal) { return f.Decimal }

// DefOfferSizeField is a QTY field.
type DefOfferSizeField struct{ quickfix.FIXDecimal }

// Tag returns tag.DefOfferSize (294).
func (f DefOfferSizeField) Tag() quickfix.Tag { return tag.DefOfferSize }

// NewDefOfferSize returns a new DefOfferSizeField initialized with val and scale.
func NewDefOfferSize(val decimal.Decimal, scale int32) DefOfferSizeField {
	return DefOfferSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DefOfferSizeField) Value() (val decimal.Decimal) { return f.Decimal }

// DeleteReasonField is a enum.DeleteReason field.
type DeleteReasonField struct{ quickfix.FIXString }

// Tag returns tag.DeleteReason (285).
func (f DeleteReasonField) Tag() quickfix.Tag { return tag.DeleteReason }

func NewDeleteReason(val enum.DeleteReason) DeleteReasonField {
	return DeleteReasonField{quickfix.FIXString(val)}
}

func (f DeleteReasonField) Value() enum.DeleteReason { return enum.DeleteReason(f.String()) }

// DeliverToCompIDField is a STRING field.
type DeliverToCompIDField struct{ quickfix.FIXString }

// Tag returns tag.DeliverToCompID (128).
func (f DeliverToCompIDField) Tag() quickfix.Tag { return tag.DeliverToCompID }

// NewDeliverToCompID returns a new DeliverToCompIDField initialized with val.
func NewDeliverToCompID(val string) DeliverToCompIDField {
	return DeliverToCompIDField{quickfix.FIXString(val)}
}

func (f DeliverToCompIDField) Value() string { return f.String() }

// DeliverToLocationIDField is a STRING field.
type DeliverToLocationIDField struct{ quickfix.FIXString }

// Tag returns tag.DeliverToLocationID (145).
func (f DeliverToLocationIDField) Tag() quickfix.Tag { return tag.DeliverToLocationID }

// NewDeliverToLocationID returns a new DeliverToLocationIDField initialized with val.
func NewDeliverToLocationID(val string) DeliverToLocationIDField {
	return DeliverToLocationIDField{quickfix.FIXString(val)}
}

func (f DeliverToLocationIDField) Value() string { return f.String() }

// DeliverToSubIDField is a STRING field.
type DeliverToSubIDField struct{ quickfix.FIXString }

// Tag returns tag.DeliverToSubID (129).
func (f DeliverToSubIDField) Tag() quickfix.Tag { return tag.DeliverToSubID }

// NewDeliverToSubID returns a new DeliverToSubIDField initialized with val.
func NewDeliverToSubID(val string) DeliverToSubIDField {
	return DeliverToSubIDField{quickfix.FIXString(val)}
}

func (f DeliverToSubIDField) Value() string { return f.String() }

// DeskIDField is a STRING field.
type DeskIDField struct{ quickfix.FIXString }

// Tag returns tag.DeskID (284).
func (f DeskIDField) Tag() quickfix.Tag { return tag.DeskID }

// NewDeskID returns a new DeskIDField initialized with val.
func NewDeskID(val string) DeskIDField {
	return DeskIDField{quickfix.FIXString(val)}
}

func (f DeskIDField) Value() string { return f.String() }

// DiscretionInstField is a enum.DiscretionInst field.
type DiscretionInstField struct{ quickfix.FIXString }

// Tag returns tag.DiscretionInst (388).
func (f DiscretionInstField) Tag() quickfix.Tag { return tag.DiscretionInst }

func NewDiscretionInst(val enum.DiscretionInst) DiscretionInstField {
	return DiscretionInstField{quickfix.FIXString(val)}
}

func (f DiscretionInstField) Value() enum.DiscretionInst { return enum.DiscretionInst(f.String()) }

// DiscretionOffsetField is a PRICEOFFSET field.
type DiscretionOffsetField struct{ quickfix.FIXDecimal }

// Tag returns tag.DiscretionOffset (389).
func (f DiscretionOffsetField) Tag() quickfix.Tag { return tag.DiscretionOffset }

// NewDiscretionOffset returns a new DiscretionOffsetField initialized with val and scale.
func NewDiscretionOffset(val decimal.Decimal, scale int32) DiscretionOffsetField {
	return DiscretionOffsetField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DiscretionOffsetField) Value() (val decimal.Decimal) { return f.Decimal }

// DlvyInstField is a STRING field.
type DlvyInstField struct{ quickfix.FIXString }

// Tag returns tag.DlvyInst (86).
func (f DlvyInstField) Tag() quickfix.Tag { return tag.DlvyInst }

// NewDlvyInst returns a new DlvyInstField initialized with val.
func NewDlvyInst(val string) DlvyInstField {
	return DlvyInstField{quickfix.FIXString(val)}
}

func (f DlvyInstField) Value() string { return f.String() }

// DropCopyFlagField is a enum.DropCopyFlag field.
type DropCopyFlagField struct{ quickfix.FIXString }

// Tag returns tag.DropCopyFlag (9406).
func (f DropCopyFlagField) Tag() quickfix.Tag { return tag.DropCopyFlag }

func NewDropCopyFlag(val enum.DropCopyFlag) DropCopyFlagField {
	return DropCopyFlagField{quickfix.FIXString(val)}
}

func (f DropCopyFlagField) Value() enum.DropCopyFlag { return enum.DropCopyFlag(f.String()) }

// DueToRelatedField is a BOOLEAN field.
type DueToRelatedField struct{ quickfix.FIXBoolean }

// Tag returns tag.DueToRelated (329).
func (f DueToRelatedField) Tag() quickfix.Tag { return tag.DueToRelated }

// NewDueToRelated returns a new DueToRelatedField initialized with val.
func NewDueToRelated(val bool) DueToRelatedField {
	return DueToRelatedField{quickfix.FIXBoolean(val)}
}

func (f DueToRelatedField) Value() bool { return f.Bool() }

// EFPTrackingErrorField is a FLOAT field.
type EFPTrackingErrorField struct{ quickfix.FIXDecimal }

// Tag returns tag.EFPTrackingError (405).
func (f EFPTrackingErrorField) Tag() quickfix.Tag { return tag.EFPTrackingError }

// NewEFPTrackingError returns a new EFPTrackingErrorField initialized with val and scale.
func NewEFPTrackingError(val decimal.Decimal, scale int32) EFPTrackingErrorField {
	return EFPTrackingErrorField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f EFPTrackingErrorField) Value() (val decimal.Decimal) { return f.Decimal }

// EffectiveTimeField is a UTCTIMESTAMP field.
type EffectiveTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.EffectiveTime (168).
func (f EffectiveTimeField) Tag() quickfix.Tag { return tag.EffectiveTime }

// NewEffectiveTime returns a new EffectiveTimeField initialized with val.
func NewEffectiveTime(val time.Time) EffectiveTimeField {
	return NewEffectiveTimeWithPrecision(val, quickfix.Millis)
}

// NewEffectiveTimeNoMillis returns a new EffectiveTimeField initialized with val without millisecs.
func NewEffectiveTimeNoMillis(val time.Time) EffectiveTimeField {
	return NewEffectiveTimeWithPrecision(val, quickfix.Seconds)
}

// NewEffectiveTimeWithPrecision returns a new EffectiveTimeField initialized with val of specified precision.
func NewEffectiveTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) EffectiveTimeField {
	return EffectiveTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f EffectiveTimeField) Value() time.Time { return f.Time }

// EmailThreadIDField is a STRING field.
type EmailThreadIDField struct{ quickfix.FIXString }

// Tag returns tag.EmailThreadID (164).
func (f EmailThreadIDField) Tag() quickfix.Tag { return tag.EmailThreadID }

// NewEmailThreadID returns a new EmailThreadIDField initialized with val.
func NewEmailThreadID(val string) EmailThreadIDField {
	return EmailThreadIDField{quickfix.FIXString(val)}
}

func (f EmailThreadIDField) Value() string { return f.String() }

// EmailTypeField is a enum.EmailType field.
type EmailTypeField struct{ quickfix.FIXString }

// Tag returns tag.EmailType (94).
func (f EmailTypeField) Tag() quickfix.Tag { return tag.EmailType }

func NewEmailType(val enum.EmailType) EmailTypeField {
	return EmailTypeField{quickfix.FIXString(val)}
}

func (f EmailTypeField) Value() enum.EmailType { return enum.EmailType(f.String()) }

// EncodedAllocTextField is a DATA field.
type EncodedAllocTextField struct{ quickfix.FIXString }

// Tag returns tag.EncodedAllocText (361).
func (f EncodedAllocTextField) Tag() quickfix.Tag { return tag.EncodedAllocText }

// NewEncodedAllocText returns a new EncodedAllocTextField initialized with val.
func NewEncodedAllocText(val string) EncodedAllocTextField {
	return EncodedAllocTextField{quickfix.FIXString(val)}
}

func (f EncodedAllocTextField) Value() string { return f.String() }

// EncodedAllocTextLenField is a LENGTH field.
type EncodedAllocTextLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedAllocTextLen (360).
func (f EncodedAllocTextLenField) Tag() quickfix.Tag { return tag.EncodedAllocTextLen }

// NewEncodedAllocTextLen returns a new EncodedAllocTextLenField initialized with val.
func NewEncodedAllocTextLen(val int) EncodedAllocTextLenField {
	return EncodedAllocTextLenField{quickfix.FIXInt(val)}
}

func (f EncodedAllocTextLenField) Value() int { return f.Int() }

// EncodedHeadlineField is a DATA field.
type EncodedHeadlineField struct{ quickfix.FIXString }

// Tag returns tag.EncodedHeadline (359).
func (f EncodedHeadlineField) Tag() quickfix.Tag { return tag.EncodedHeadline }

// NewEncodedHeadline returns a new EncodedHeadlineField initialized with val.
func NewEncodedHeadline(val string) EncodedHeadlineField {
	return EncodedHeadlineField{quickfix.FIXString(val)}
}

func (f EncodedHeadlineField) Value() string { return f.String() }

// EncodedHeadlineLenField is a LENGTH field.
type EncodedHeadlineLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedHeadlineLen (358).
func (f EncodedHeadlineLenField) Tag() quickfix.Tag { return tag.EncodedHeadlineLen }

// NewEncodedHeadlineLen returns a new EncodedHeadlineLenField initialized with val.
func NewEncodedHeadlineLen(val int) EncodedHeadlineLenField {
	return EncodedHeadlineLenField{quickfix.FIXInt(val)}
}

func (f EncodedHeadlineLenField) Value() int { return f.Int() }

// EncodedIssuerField is a DATA field.
type EncodedIssuerField struct{ quickfix.FIXString }

// Tag returns tag.EncodedIssuer (349).
func (f EncodedIssuerField) Tag() quickfix.Tag { return tag.EncodedIssuer }

// NewEncodedIssuer returns a new EncodedIssuerField initialized with val.
func NewEncodedIssuer(val string) EncodedIssuerField {
	return EncodedIssuerField{quickfix.FIXString(val)}
}

func (f EncodedIssuerField) Value() string { return f.String() }

// EncodedIssuerLenField is a LENGTH field.
type EncodedIssuerLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedIssuerLen (348).
func (f EncodedIssuerLenField) Tag() quickfix.Tag { return tag.EncodedIssuerLen }

// NewEncodedIssuerLen returns a new EncodedIssuerLenField initialized with val.
func NewEncodedIssuerLen(val int) EncodedIssuerLenField {
	return EncodedIssuerLenField{quickfix.FIXInt(val)}
}

func (f EncodedIssuerLenField) Value() int { return f.Int() }

// EncodedListExecInstField is a DATA field.
type EncodedListExecInstField struct{ quickfix.FIXString }

// Tag returns tag.EncodedListExecInst (353).
func (f EncodedListExecInstField) Tag() quickfix.Tag { return tag.EncodedListExecInst }

// NewEncodedListExecInst returns a new EncodedListExecInstField initialized with val.
func NewEncodedListExecInst(val string) EncodedListExecInstField {
	return EncodedListExecInstField{quickfix.FIXString(val)}
}

func (f EncodedListExecInstField) Value() string { return f.String() }

// EncodedListExecInstLenField is a LENGTH field.
type EncodedListExecInstLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedListExecInstLen (352).
func (f EncodedListExecInstLenField) Tag() quickfix.Tag { return tag.EncodedListExecInstLen }

// NewEncodedListExecInstLen returns a new EncodedListExecInstLenField initialized with val.
func NewEncodedListExecInstLen(val int) EncodedListExecInstLenField {
	return EncodedListExecInstLenField{quickfix.FIXInt(val)}
}

func (f EncodedListExecInstLenField) Value() int { return f.Int() }

// EncodedListStatusTextField is a DATA field.
type EncodedListStatusTextField struct{ quickfix.FIXString }

// Tag returns tag.EncodedListStatusText (446).
func (f EncodedListStatusTextField) Tag() quickfix.Tag { return tag.EncodedListStatusText }

// NewEncodedListStatusText returns a new EncodedListStatusTextField initialized with val.
func NewEncodedListStatusText(val string) EncodedListStatusTextField {
	return EncodedListStatusTextField{quickfix.FIXString(val)}
}

func (f EncodedListStatusTextField) Value() string { return f.String() }

// EncodedListStatusTextLenField is a LENGTH field.
type EncodedListStatusTextLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedListStatusTextLen (445).
func (f EncodedListStatusTextLenField) Tag() quickfix.Tag { return tag.EncodedListStatusTextLen }

// NewEncodedListStatusTextLen returns a new EncodedListStatusTextLenField initialized with val.
func NewEncodedListStatusTextLen(val int) EncodedListStatusTextLenField {
	return EncodedListStatusTextLenField{quickfix.FIXInt(val)}
}

func (f EncodedListStatusTextLenField) Value() int { return f.Int() }

// EncodedSecurityDescField is a DATA field.
type EncodedSecurityDescField struct{ quickfix.FIXString }

// Tag returns tag.EncodedSecurityDesc (351).
func (f EncodedSecurityDescField) Tag() quickfix.Tag { return tag.EncodedSecurityDesc }

// NewEncodedSecurityDesc returns a new EncodedSecurityDescField initialized with val.
func NewEncodedSecurityDesc(val string) EncodedSecurityDescField {
	return EncodedSecurityDescField{quickfix.FIXString(val)}
}

func (f EncodedSecurityDescField) Value() string { return f.String() }

// EncodedSecurityDescLenField is a LENGTH field.
type EncodedSecurityDescLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedSecurityDescLen (350).
func (f EncodedSecurityDescLenField) Tag() quickfix.Tag { return tag.EncodedSecurityDescLen }

// NewEncodedSecurityDescLen returns a new EncodedSecurityDescLenField initialized with val.
func NewEncodedSecurityDescLen(val int) EncodedSecurityDescLenField {
	return EncodedSecurityDescLenField{quickfix.FIXInt(val)}
}

func (f EncodedSecurityDescLenField) Value() int { return f.Int() }

// EncodedSubjectField is a DATA field.
type EncodedSubjectField struct{ quickfix.FIXString }

// Tag returns tag.EncodedSubject (357).
func (f EncodedSubjectField) Tag() quickfix.Tag { return tag.EncodedSubject }

// NewEncodedSubject returns a new EncodedSubjectField initialized with val.
func NewEncodedSubject(val string) EncodedSubjectField {
	return EncodedSubjectField{quickfix.FIXString(val)}
}

func (f EncodedSubjectField) Value() string { return f.String() }

// EncodedSubjectLenField is a LENGTH field.
type EncodedSubjectLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedSubjectLen (356).
func (f EncodedSubjectLenField) Tag() quickfix.Tag { return tag.EncodedSubjectLen }

// NewEncodedSubjectLen returns a new EncodedSubjectLenField initialized with val.
func NewEncodedSubjectLen(val int) EncodedSubjectLenField {
	return EncodedSubjectLenField{quickfix.FIXInt(val)}
}

func (f EncodedSubjectLenField) Value() int { return f.Int() }

// EncodedTextField is a DATA field.
type EncodedTextField struct{ quickfix.FIXString }

// Tag returns tag.EncodedText (355).
func (f EncodedTextField) Tag() quickfix.Tag { return tag.EncodedText }

// NewEncodedText returns a new EncodedTextField initialized with val.
func NewEncodedText(val string) EncodedTextField {
	return EncodedTextField{quickfix.FIXString(val)}
}

func (f EncodedTextField) Value() string { return f.String() }

// EncodedTextLenField is a LENGTH field.
type EncodedTextLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedTextLen (354).
func (f EncodedTextLenField) Tag() quickfix.Tag { return tag.EncodedTextLen }

// NewEncodedTextLen returns a new EncodedTextLenField initialized with val.
func NewEncodedTextLen(val int) EncodedTextLenField {
	return EncodedTextLenField{quickfix.FIXInt(val)}
}

func (f EncodedTextLenField) Value() int { return f.Int() }

// EncodedUnderlyingIssuerField is a DATA field.
type EncodedUnderlyingIssuerField struct{ quickfix.FIXString }

// Tag returns tag.EncodedUnderlyingIssuer (363).
func (f EncodedUnderlyingIssuerField) Tag() quickfix.Tag { return tag.EncodedUnderlyingIssuer }

// NewEncodedUnderlyingIssuer returns a new EncodedUnderlyingIssuerField initialized with val.
func NewEncodedUnderlyingIssuer(val string) EncodedUnderlyingIssuerField {
	return EncodedUnderlyingIssuerField{quickfix.FIXString(val)}
}

func (f EncodedUnderlyingIssuerField) Value() string { return f.String() }

// EncodedUnderlyingIssuerLenField is a LENGTH field.
type EncodedUnderlyingIssuerLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedUnderlyingIssuerLen (362).
func (f EncodedUnderlyingIssuerLenField) Tag() quickfix.Tag { return tag.EncodedUnderlyingIssuerLen }

// NewEncodedUnderlyingIssuerLen returns a new EncodedUnderlyingIssuerLenField initialized with val.
func NewEncodedUnderlyingIssuerLen(val int) EncodedUnderlyingIssuerLenField {
	return EncodedUnderlyingIssuerLenField{quickfix.FIXInt(val)}
}

func (f EncodedUnderlyingIssuerLenField) Value() int { return f.Int() }

// EncodedUnderlyingSecurityDescField is a DATA field.
type EncodedUnderlyingSecurityDescField struct{ quickfix.FIXString }

// Tag returns tag.EncodedUnderlyingSecurityDesc (365).
func (f EncodedUnderlyingSecurityDescField) Tag() quickfix.Tag {
	return tag.EncodedUnderlyingSecurityDesc
}

// NewEncodedUnderlyingSecurityDesc returns a new EncodedUnderlyingSecurityDescField initialized with val.
func NewEncodedUnderlyingSecurityDesc(val string) EncodedUnderlyingSecurityDescField {
	return EncodedUnderlyingSecurityDescField{quickfix.FIXString(val)}
}

func (f EncodedUnderlyingSecurityDescField) Value() string { return f.String() }

// EncodedUnderlyingSecurityDescLenField is a LENGTH field.
type EncodedUnderlyingSecurityDescLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedUnderlyingSecurityDescLen (364).
func (f EncodedUnderlyingSecurityDescLenField) Tag() quickfix.Tag {
	return tag.EncodedUnderlyingSecurityDescLen
}

// NewEncodedUnderlyingSecurityDescLen returns a new EncodedUnderlyingSecurityDescLenField initialized with val.
func NewEncodedUnderlyingSecurityDescLen(val int) EncodedUnderlyingSecurityDescLenField {
	return EncodedUnderlyingSecurityDescLenField{quickfix.FIXInt(val)}
}

func (f EncodedUnderlyingSecurityDescLenField) Value() int { return f.Int() }

// EncryptMethodField is a enum.EncryptMethod field.
type EncryptMethodField struct{ quickfix.FIXString }

// Tag returns tag.EncryptMethod (98).
func (f EncryptMethodField) Tag() quickfix.Tag { return tag.EncryptMethod }

func NewEncryptMethod(val enum.EncryptMethod) EncryptMethodField {
	return EncryptMethodField{quickfix.FIXString(val)}
}

func (f EncryptMethodField) Value() enum.EncryptMethod { return enum.EncryptMethod(f.String()) }

// EndSeqNoField is a INT field.
type EndSeqNoField struct{ quickfix.FIXInt }

// Tag returns tag.EndSeqNo (16).
func (f EndSeqNoField) Tag() quickfix.Tag { return tag.EndSeqNo }

// NewEndSeqNo returns a new EndSeqNoField initialized with val.
func NewEndSeqNo(val int) EndSeqNoField {
	return EndSeqNoField{quickfix.FIXInt(val)}
}

func (f EndSeqNoField) Value() int { return f.Int() }

// ExDestinationField is a EXCHANGE field.
type ExDestinationField struct{ quickfix.FIXString }

// Tag returns tag.ExDestination (100).
func (f ExDestinationField) Tag() quickfix.Tag { return tag.ExDestination }

// NewExDestination returns a new ExDestinationField initialized with val.
func NewExDestination(val string) ExDestinationField {
	return ExDestinationField{quickfix.FIXString(val)}
}

func (f ExDestinationField) Value() string { return f.String() }

// ExchangeForPhysicalField is a BOOLEAN field.
type ExchangeForPhysicalField struct{ quickfix.FIXBoolean }

// Tag returns tag.ExchangeForPhysical (411).
func (f ExchangeForPhysicalField) Tag() quickfix.Tag { return tag.ExchangeForPhysical }

// NewExchangeForPhysical returns a new ExchangeForPhysicalField initialized with val.
func NewExchangeForPhysical(val bool) ExchangeForPhysicalField {
	return ExchangeForPhysicalField{quickfix.FIXBoolean(val)}
}

func (f ExchangeForPhysicalField) Value() bool { return f.Bool() }

// ExecBrokerField is a STRING field.
type ExecBrokerField struct{ quickfix.FIXString }

// Tag returns tag.ExecBroker (76).
func (f ExecBrokerField) Tag() quickfix.Tag { return tag.ExecBroker }

// NewExecBroker returns a new ExecBrokerField initialized with val.
func NewExecBroker(val string) ExecBrokerField {
	return ExecBrokerField{quickfix.FIXString(val)}
}

func (f ExecBrokerField) Value() string { return f.String() }

// ExecIDField is a STRING field.
type ExecIDField struct{ quickfix.FIXString }

// Tag returns tag.ExecID (17).
func (f ExecIDField) Tag() quickfix.Tag { return tag.ExecID }

// NewExecID returns a new ExecIDField initialized with val.
func NewExecID(val string) ExecIDField {
	return ExecIDField{quickfix.FIXString(val)}
}

func (f ExecIDField) Value() string { return f.String() }

// ExecInstField is a enum.ExecInst field.
type ExecInstField struct{ quickfix.FIXString }

// Tag returns tag.ExecInst (18).
func (f ExecInstField) Tag() quickfix.Tag { return tag.ExecInst }

func NewExecInst(val enum.ExecInst) ExecInstField {
	return ExecInstField{quickfix.FIXString(val)}
}

func (f ExecInstField) Value() enum.ExecInst { return enum.ExecInst(f.String()) }

// ExecRefIDField is a STRING field.
type ExecRefIDField struct{ quickfix.FIXString }

// Tag returns tag.ExecRefID (19).
func (f ExecRefIDField) Tag() quickfix.Tag { return tag.ExecRefID }

// NewExecRefID returns a new ExecRefIDField initialized with val.
func NewExecRefID(val string) ExecRefIDField {
	return ExecRefIDField{quickfix.FIXString(val)}
}

func (f ExecRefIDField) Value() string { return f.String() }

// ExecRestatementReasonField is a enum.ExecRestatementReason field.
type ExecRestatementReasonField struct{ quickfix.FIXString }

// Tag returns tag.ExecRestatementReason (378).
func (f ExecRestatementReasonField) Tag() quickfix.Tag { return tag.ExecRestatementReason }

func NewExecRestatementReason(val enum.ExecRestatementReason) ExecRestatementReasonField {
	return ExecRestatementReasonField{quickfix.FIXString(val)}
}

func (f ExecRestatementReasonField) Value() enum.ExecRestatementReason {
	return enum.ExecRestatementReason(f.String())
}

// ExecTransTypeField is a enum.ExecTransType field.
type ExecTransTypeField struct{ quickfix.FIXString }

// Tag returns tag.ExecTransType (20).
func (f ExecTransTypeField) Tag() quickfix.Tag { return tag.ExecTransType }

func NewExecTransType(val enum.ExecTransType) ExecTransTypeField {
	return ExecTransTypeField{quickfix.FIXString(val)}
}

func (f ExecTransTypeField) Value() enum.ExecTransType { return enum.ExecTransType(f.String()) }

// ExecTypeField is a enum.ExecType field.
type ExecTypeField struct{ quickfix.FIXString }

// Tag returns tag.ExecType (150).
func (f ExecTypeField) Tag() quickfix.Tag { return tag.ExecType }

func NewExecType(val enum.ExecType) ExecTypeField {
	return ExecTypeField{quickfix.FIXString(val)}
}

func (f ExecTypeField) Value() enum.ExecType { return enum.ExecType(f.String()) }

// ExpireDateField is a LOCALMKTDATE field.
type ExpireDateField struct{ quickfix.FIXString }

// Tag returns tag.ExpireDate (432).
func (f ExpireDateField) Tag() quickfix.Tag { return tag.ExpireDate }

// NewExpireDate returns a new ExpireDateField initialized with val.
func NewExpireDate(val string) ExpireDateField {
	return ExpireDateField{quickfix.FIXString(val)}
}

func (f ExpireDateField) Value() string { return f.String() }

// ExpireTimeField is a UTCTIMESTAMP field.
type ExpireTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.ExpireTime (126).
func (f ExpireTimeField) Tag() quickfix.Tag { return tag.ExpireTime }

// NewExpireTime returns a new ExpireTimeField initialized with val.
func NewExpireTime(val time.Time) ExpireTimeField {
	return NewExpireTimeWithPrecision(val, quickfix.Millis)
}

// NewExpireTimeNoMillis returns a new ExpireTimeField initialized with val without millisecs.
func NewExpireTimeNoMillis(val time.Time) ExpireTimeField {
	return NewExpireTimeWithPrecision(val, quickfix.Seconds)
}

// NewExpireTimeWithPrecision returns a new ExpireTimeField initialized with val of specified precision.
func NewExpireTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) ExpireTimeField {
	return ExpireTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f ExpireTimeField) Value() time.Time { return f.Time }

// FairValueField is a AMT field.
type FairValueField struct{ quickfix.FIXDecimal }

// Tag returns tag.FairValue (406).
func (f FairValueField) Tag() quickfix.Tag { return tag.FairValue }

// NewFairValue returns a new FairValueField initialized with val and scale.
func NewFairValue(val decimal.Decimal, scale int32) FairValueField {
	return FairValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f FairValueField) Value() (val decimal.Decimal) { return f.Decimal }

// FinancialStatusField is a enum.FinancialStatus field.
type FinancialStatusField struct{ quickfix.FIXString }

// Tag returns tag.FinancialStatus (291).
func (f FinancialStatusField) Tag() quickfix.Tag { return tag.FinancialStatus }

func NewFinancialStatus(val enum.FinancialStatus) FinancialStatusField {
	return FinancialStatusField{quickfix.FIXString(val)}
}

func (f FinancialStatusField) Value() enum.FinancialStatus { return enum.FinancialStatus(f.String()) }

// ForexReqField is a BOOLEAN field.
type ForexReqField struct{ quickfix.FIXBoolean }

// Tag returns tag.ForexReq (121).
func (f ForexReqField) Tag() quickfix.Tag { return tag.ForexReq }

// NewForexReq returns a new ForexReqField initialized with val.
func NewForexReq(val bool) ForexReqField {
	return ForexReqField{quickfix.FIXBoolean(val)}
}

func (f ForexReqField) Value() bool { return f.Bool() }

// FutSettDateField is a LOCALMKTDATE field.
type FutSettDateField struct{ quickfix.FIXString }

// Tag returns tag.FutSettDate (64).
func (f FutSettDateField) Tag() quickfix.Tag { return tag.FutSettDate }

// NewFutSettDate returns a new FutSettDateField initialized with val.
func NewFutSettDate(val string) FutSettDateField {
	return FutSettDateField{quickfix.FIXString(val)}
}

func (f FutSettDateField) Value() string { return f.String() }

// FutSettDate2Field is a LOCALMKTDATE field.
type FutSettDate2Field struct{ quickfix.FIXString }

// Tag returns tag.FutSettDate2 (193).
func (f FutSettDate2Field) Tag() quickfix.Tag { return tag.FutSettDate2 }

// NewFutSettDate2 returns a new FutSettDate2Field initialized with val.
func NewFutSettDate2(val string) FutSettDate2Field {
	return FutSettDate2Field{quickfix.FIXString(val)}
}

func (f FutSettDate2Field) Value() string { return f.String() }

// GTBookingInstField is a enum.GTBookingInst field.
type GTBookingInstField struct{ quickfix.FIXString }

// Tag returns tag.GTBookingInst (427).
func (f GTBookingInstField) Tag() quickfix.Tag { return tag.GTBookingInst }

func NewGTBookingInst(val enum.GTBookingInst) GTBookingInstField {
	return GTBookingInstField{quickfix.FIXString(val)}
}

func (f GTBookingInstField) Value() enum.GTBookingInst { return enum.GTBookingInst(f.String()) }

// GapFillFlagField is a BOOLEAN field.
type GapFillFlagField struct{ quickfix.FIXBoolean }

// Tag returns tag.GapFillFlag (123).
func (f GapFillFlagField) Tag() quickfix.Tag { return tag.GapFillFlag }

// NewGapFillFlag returns a new GapFillFlagField initialized with val.
func NewGapFillFlag(val bool) GapFillFlagField {
	return GapFillFlagField{quickfix.FIXBoolean(val)}
}

func (f GapFillFlagField) Value() bool { return f.Bool() }

// GrossTradeAmtField is a AMT field.
type GrossTradeAmtField struct{ quickfix.FIXDecimal }

// Tag returns tag.GrossTradeAmt (381).
func (f GrossTradeAmtField) Tag() quickfix.Tag { return tag.GrossTradeAmt }

// NewGrossTradeAmt returns a new GrossTradeAmtField initialized with val and scale.
func NewGrossTradeAmt(val decimal.Decimal, scale int32) GrossTradeAmtField {
	return GrossTradeAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f GrossTradeAmtField) Value() (val decimal.Decimal) { return f.Decimal }

// HaltReasonCharField is a enum.HaltReasonChar field.
type HaltReasonCharField struct{ quickfix.FIXString }

// Tag returns tag.HaltReasonChar (327).
func (f HaltReasonCharField) Tag() quickfix.Tag { return tag.HaltReasonChar }

func NewHaltReasonChar(val enum.HaltReasonChar) HaltReasonCharField {
	return HaltReasonCharField{quickfix.FIXString(val)}
}

func (f HaltReasonCharField) Value() enum.HaltReasonChar { return enum.HaltReasonChar(f.String()) }

// HandlInstField is a enum.HandlInst field.
type HandlInstField struct{ quickfix.FIXString }

// Tag returns tag.HandlInst (21).
func (f HandlInstField) Tag() quickfix.Tag { return tag.HandlInst }

func NewHandlInst(val enum.HandlInst) HandlInstField {
	return HandlInstField{quickfix.FIXString(val)}
}

func (f HandlInstField) Value() enum.HandlInst { return enum.HandlInst(f.String()) }

// HeadlineField is a STRING field.
type HeadlineField struct{ quickfix.FIXString }

// Tag returns tag.Headline (148).
func (f HeadlineField) Tag() quickfix.Tag { return tag.Headline }

// NewHeadline returns a new HeadlineField initialized with val.
func NewHeadline(val string) HeadlineField {
	return HeadlineField{quickfix.FIXString(val)}
}

func (f HeadlineField) Value() string { return f.String() }

// HeartBtIntField is a INT field.
type HeartBtIntField struct{ quickfix.FIXInt }

// Tag returns tag.HeartBtInt (108).
func (f HeartBtIntField) Tag() quickfix.Tag { return tag.HeartBtInt }

// NewHeartBtInt returns a new HeartBtIntField initialized with val.
func NewHeartBtInt(val int) HeartBtIntField {
	return HeartBtIntField{quickfix.FIXInt(val)}
}

func (f HeartBtIntField) Value() int { return f.Int() }

// HighPxField is a PRICE field.
type HighPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.HighPx (332).
func (f HighPxField) Tag() quickfix.Tag { return tag.HighPx }

// NewHighPx returns a new HighPxField initialized with val and scale.
func NewHighPx(val decimal.Decimal, scale int32) HighPxField {
	return HighPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f HighPxField) Value() (val decimal.Decimal) { return f.Decimal }

// IDSourceField is a enum.IDSource field.
type IDSourceField struct{ quickfix.FIXString }

// Tag returns tag.IDSource (22).
func (f IDSourceField) Tag() quickfix.Tag { return tag.IDSource }

func NewIDSource(val enum.IDSource) IDSourceField {
	return IDSourceField{quickfix.FIXString(val)}
}

func (f IDSourceField) Value() enum.IDSource { return enum.IDSource(f.String()) }

// IOINaturalFlagField is a BOOLEAN field.
type IOINaturalFlagField struct{ quickfix.FIXBoolean }

// Tag returns tag.IOINaturalFlag (130).
func (f IOINaturalFlagField) Tag() quickfix.Tag { return tag.IOINaturalFlag }

// NewIOINaturalFlag returns a new IOINaturalFlagField initialized with val.
func NewIOINaturalFlag(val bool) IOINaturalFlagField {
	return IOINaturalFlagField{quickfix.FIXBoolean(val)}
}

func (f IOINaturalFlagField) Value() bool { return f.Bool() }

// IOIOthSvcField is a CHAR field.
type IOIOthSvcField struct{ quickfix.FIXString }

// Tag returns tag.IOIOthSvc (24).
func (f IOIOthSvcField) Tag() quickfix.Tag { return tag.IOIOthSvc }

// NewIOIOthSvc returns a new IOIOthSvcField initialized with val.
func NewIOIOthSvc(val string) IOIOthSvcField {
	return IOIOthSvcField{quickfix.FIXString(val)}
}

func (f IOIOthSvcField) Value() string { return f.String() }

// IOIQltyIndField is a enum.IOIQltyInd field.
type IOIQltyIndField struct{ quickfix.FIXString }

// Tag returns tag.IOIQltyInd (25).
func (f IOIQltyIndField) Tag() quickfix.Tag { return tag.IOIQltyInd }

func NewIOIQltyInd(val enum.IOIQltyInd) IOIQltyIndField {
	return IOIQltyIndField{quickfix.FIXString(val)}
}

func (f IOIQltyIndField) Value() enum.IOIQltyInd { return enum.IOIQltyInd(f.String()) }

// IOIQualifierField is a enum.IOIQualifier field.
type IOIQualifierField struct{ quickfix.FIXString }

// Tag returns tag.IOIQualifier (104).
func (f IOIQualifierField) Tag() quickfix.Tag { return tag.IOIQualifier }

func NewIOIQualifier(val enum.IOIQualifier) IOIQualifierField {
	return IOIQualifierField{quickfix.FIXString(val)}
}

func (f IOIQualifierField) Value() enum.IOIQualifier { return enum.IOIQualifier(f.String()) }

// IOIRefIDField is a STRING field.
type IOIRefIDField struct{ quickfix.FIXString }

// Tag returns tag.IOIRefID (26).
func (f IOIRefIDField) Tag() quickfix.Tag { return tag.IOIRefID }

// NewIOIRefID returns a new IOIRefIDField initialized with val.
func NewIOIRefID(val string) IOIRefIDField {
	return IOIRefIDField{quickfix.FIXString(val)}
}

func (f IOIRefIDField) Value() string { return f.String() }

// IOISharesField is a enum.IOIShares field.
type IOISharesField struct{ quickfix.FIXString }

// Tag returns tag.IOIShares (27).
func (f IOISharesField) Tag() quickfix.Tag { return tag.IOIShares }

func NewIOIShares(val enum.IOIShares) IOISharesField {
	return IOISharesField{quickfix.FIXString(val)}
}

func (f IOISharesField) Value() enum.IOIShares { return enum.IOIShares(f.String()) }

// IOITransTypeField is a enum.IOITransType field.
type IOITransTypeField struct{ quickfix.FIXString }

// Tag returns tag.IOITransType (28).
func (f IOITransTypeField) Tag() quickfix.Tag { return tag.IOITransType }

func NewIOITransType(val enum.IOITransType) IOITransTypeField {
	return IOITransTypeField{quickfix.FIXString(val)}
}

func (f IOITransTypeField) Value() enum.IOITransType { return enum.IOITransType(f.String()) }

// IOIidField is a STRING field.
type IOIidField struct{ quickfix.FIXString }

// Tag returns tag.IOIid (23).
func (f IOIidField) Tag() quickfix.Tag { return tag.IOIid }

// NewIOIid returns a new IOIidField initialized with val.
func NewIOIid(val string) IOIidField {
	return IOIidField{quickfix.FIXString(val)}
}

func (f IOIidField) Value() string { return f.String() }

// InViewOfCommonField is a BOOLEAN field.
type InViewOfCommonField struct{ quickfix.FIXBoolean }

// Tag returns tag.InViewOfCommon (328).
func (f InViewOfCommonField) Tag() quickfix.Tag { return tag.InViewOfCommon }

// NewInViewOfCommon returns a new InViewOfCommonField initialized with val.
func NewInViewOfCommon(val bool) InViewOfCommonField {
	return InViewOfCommonField{quickfix.FIXBoolean(val)}
}

func (f InViewOfCommonField) Value() bool { return f.Bool() }

// IncTaxIndField is a enum.IncTaxInd field.
type IncTaxIndField struct{ quickfix.FIXString }

// Tag returns tag.IncTaxInd (416).
func (f IncTaxIndField) Tag() quickfix.Tag { return tag.IncTaxInd }

func NewIncTaxInd(val enum.IncTaxInd) IncTaxIndField {
	return IncTaxIndField{quickfix.FIXString(val)}
}

func (f IncTaxIndField) Value() enum.IncTaxInd { return enum.IncTaxInd(f.String()) }

// IssuerField is a STRING field.
type IssuerField struct{ quickfix.FIXString }

// Tag returns tag.Issuer (106).
func (f IssuerField) Tag() quickfix.Tag { return tag.Issuer }

// NewIssuer returns a new IssuerField initialized with val.
func NewIssuer(val string) IssuerField {
	return IssuerField{quickfix.FIXString(val)}
}

func (f IssuerField) Value() string { return f.String() }

// LastCapacityField is a enum.LastCapacity field.
type LastCapacityField struct{ quickfix.FIXString }

// Tag returns tag.LastCapacity (29).
func (f LastCapacityField) Tag() quickfix.Tag { return tag.LastCapacity }

func NewLastCapacity(val enum.LastCapacity) LastCapacityField {
	return LastCapacityField{quickfix.FIXString(val)}
}

func (f LastCapacityField) Value() enum.LastCapacity { return enum.LastCapacity(f.String()) }

// LastForwardPointsField is a PRICEOFFSET field.
type LastForwardPointsField struct{ quickfix.FIXDecimal }

// Tag returns tag.LastForwardPoints (195).
func (f LastForwardPointsField) Tag() quickfix.Tag { return tag.LastForwardPoints }

// NewLastForwardPoints returns a new LastForwardPointsField initialized with val and scale.
func NewLastForwardPoints(val decimal.Decimal, scale int32) LastForwardPointsField {
	return LastForwardPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastForwardPointsField) Value() (val decimal.Decimal) { return f.Decimal }

// LastMktField is a EXCHANGE field.
type LastMktField struct{ quickfix.FIXString }

// Tag returns tag.LastMkt (30).
func (f LastMktField) Tag() quickfix.Tag { return tag.LastMkt }

// NewLastMkt returns a new LastMktField initialized with val.
func NewLastMkt(val string) LastMktField {
	return LastMktField{quickfix.FIXString(val)}
}

func (f LastMktField) Value() string { return f.String() }

// LastMsgSeqNumProcessedField is a INT field.
type LastMsgSeqNumProcessedField struct{ quickfix.FIXInt }

// Tag returns tag.LastMsgSeqNumProcessed (369).
func (f LastMsgSeqNumProcessedField) Tag() quickfix.Tag { return tag.LastMsgSeqNumProcessed }

// NewLastMsgSeqNumProcessed returns a new LastMsgSeqNumProcessedField initialized with val.
func NewLastMsgSeqNumProcessed(val int) LastMsgSeqNumProcessedField {
	return LastMsgSeqNumProcessedField{quickfix.FIXInt(val)}
}

func (f LastMsgSeqNumProcessedField) Value() int { return f.Int() }

// LastPxField is a PRICE field.
type LastPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.LastPx (31).
func (f LastPxField) Tag() quickfix.Tag { return tag.LastPx }

// NewLastPx returns a new LastPxField initialized with val and scale.
func NewLastPx(val decimal.Decimal, scale int32) LastPxField {
	return LastPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastPxField) Value() (val decimal.Decimal) { return f.Decimal }

// LastSharesField is a QTY field.
type LastSharesField struct{ quickfix.FIXDecimal }

// Tag returns tag.LastShares (32).
func (f LastSharesField) Tag() quickfix.Tag { return tag.LastShares }

// NewLastShares returns a new LastSharesField initialized with val and scale.
func NewLastShares(val decimal.Decimal, scale int32) LastSharesField {
	return LastSharesField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastSharesField) Value() (val decimal.Decimal) { return f.Decimal }

// LastSpotRateField is a PRICE field.
type LastSpotRateField struct{ quickfix.FIXDecimal }

// Tag returns tag.LastSpotRate (194).
func (f LastSpotRateField) Tag() quickfix.Tag { return tag.LastSpotRate }

// NewLastSpotRate returns a new LastSpotRateField initialized with val and scale.
func NewLastSpotRate(val decimal.Decimal, scale int32) LastSpotRateField {
	return LastSpotRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastSpotRateField) Value() (val decimal.Decimal) { return f.Decimal }

// LeavesQtyField is a QTY field.
type LeavesQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.LeavesQty (151).
func (f LeavesQtyField) Tag() quickfix.Tag { return tag.LeavesQty }

// NewLeavesQty returns a new LeavesQtyField initialized with val and scale.
func NewLeavesQty(val decimal.Decimal, scale int32) LeavesQtyField {
	return LeavesQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LeavesQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// LinesOfTextField is a INT field.
type LinesOfTextField struct{ quickfix.FIXInt }

// Tag returns tag.LinesOfText (33).
func (f LinesOfTextField) Tag() quickfix.Tag { return tag.LinesOfText }

// NewLinesOfText returns a new LinesOfTextField initialized with val.
func NewLinesOfText(val int) LinesOfTextField {
	return LinesOfTextField{quickfix.FIXInt(val)}
}

func (f LinesOfTextField) Value() int { return f.Int() }

// LiquidityIndTypeField is a enum.LiquidityIndType field.
type LiquidityIndTypeField struct{ quickfix.FIXString }

// Tag returns tag.LiquidityIndType (409).
func (f LiquidityIndTypeField) Tag() quickfix.Tag { return tag.LiquidityIndType }

func NewLiquidityIndType(val enum.LiquidityIndType) LiquidityIndTypeField {
	return LiquidityIndTypeField{quickfix.FIXString(val)}
}

func (f LiquidityIndTypeField) Value() enum.LiquidityIndType {
	return enum.LiquidityIndType(f.String())
}

// LiquidityNumSecuritiesField is a INT field.
type LiquidityNumSecuritiesField struct{ quickfix.FIXInt }

// Tag returns tag.LiquidityNumSecurities (441).
func (f LiquidityNumSecuritiesField) Tag() quickfix.Tag { return tag.LiquidityNumSecurities }

// NewLiquidityNumSecurities returns a new LiquidityNumSecuritiesField initialized with val.
func NewLiquidityNumSecurities(val int) LiquidityNumSecuritiesField {
	return LiquidityNumSecuritiesField{quickfix.FIXInt(val)}
}

func (f LiquidityNumSecuritiesField) Value() int { return f.Int() }

// LiquidityPctHighField is a FLOAT field.
type LiquidityPctHighField struct{ quickfix.FIXDecimal }

// Tag returns tag.LiquidityPctHigh (403).
func (f LiquidityPctHighField) Tag() quickfix.Tag { return tag.LiquidityPctHigh }

// NewLiquidityPctHigh returns a new LiquidityPctHighField initialized with val and scale.
func NewLiquidityPctHigh(val decimal.Decimal, scale int32) LiquidityPctHighField {
	return LiquidityPctHighField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LiquidityPctHighField) Value() (val decimal.Decimal) { return f.Decimal }

// LiquidityPctLowField is a FLOAT field.
type LiquidityPctLowField struct{ quickfix.FIXDecimal }

// Tag returns tag.LiquidityPctLow (402).
func (f LiquidityPctLowField) Tag() quickfix.Tag { return tag.LiquidityPctLow }

// NewLiquidityPctLow returns a new LiquidityPctLowField initialized with val and scale.
func NewLiquidityPctLow(val decimal.Decimal, scale int32) LiquidityPctLowField {
	return LiquidityPctLowField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LiquidityPctLowField) Value() (val decimal.Decimal) { return f.Decimal }

// LiquidityValueField is a AMT field.
type LiquidityValueField struct{ quickfix.FIXDecimal }

// Tag returns tag.LiquidityValue (404).
func (f LiquidityValueField) Tag() quickfix.Tag { return tag.LiquidityValue }

// NewLiquidityValue returns a new LiquidityValueField initialized with val and scale.
func NewLiquidityValue(val decimal.Decimal, scale int32) LiquidityValueField {
	return LiquidityValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LiquidityValueField) Value() (val decimal.Decimal) { return f.Decimal }

// ListExecInstField is a STRING field.
type ListExecInstField struct{ quickfix.FIXString }

// Tag returns tag.ListExecInst (69).
func (f ListExecInstField) Tag() quickfix.Tag { return tag.ListExecInst }

// NewListExecInst returns a new ListExecInstField initialized with val.
func NewListExecInst(val string) ListExecInstField {
	return ListExecInstField{quickfix.FIXString(val)}
}

func (f ListExecInstField) Value() string { return f.String() }

// ListExecInstTypeField is a enum.ListExecInstType field.
type ListExecInstTypeField struct{ quickfix.FIXString }

// Tag returns tag.ListExecInstType (433).
func (f ListExecInstTypeField) Tag() quickfix.Tag { return tag.ListExecInstType }

func NewListExecInstType(val enum.ListExecInstType) ListExecInstTypeField {
	return ListExecInstTypeField{quickfix.FIXString(val)}
}

func (f ListExecInstTypeField) Value() enum.ListExecInstType {
	return enum.ListExecInstType(f.String())
}

// ListIDField is a STRING field.
type ListIDField struct{ quickfix.FIXString }

// Tag returns tag.ListID (66).
func (f ListIDField) Tag() quickfix.Tag { return tag.ListID }

// NewListID returns a new ListIDField initialized with val.
func NewListID(val string) ListIDField {
	return ListIDField{quickfix.FIXString(val)}
}

func (f ListIDField) Value() string { return f.String() }

// ListNameField is a STRING field.
type ListNameField struct{ quickfix.FIXString }

// Tag returns tag.ListName (392).
func (f ListNameField) Tag() quickfix.Tag { return tag.ListName }

// NewListName returns a new ListNameField initialized with val.
func NewListName(val string) ListNameField {
	return ListNameField{quickfix.FIXString(val)}
}

func (f ListNameField) Value() string { return f.String() }

// ListOrderStatusField is a INT field.
type ListOrderStatusField struct{ quickfix.FIXInt }

// Tag returns tag.ListOrderStatus (431).
func (f ListOrderStatusField) Tag() quickfix.Tag { return tag.ListOrderStatus }

// NewListOrderStatus returns a new ListOrderStatusField initialized with val.
func NewListOrderStatus(val int) ListOrderStatusField {
	return ListOrderStatusField{quickfix.FIXInt(val)}
}

func (f ListOrderStatusField) Value() int { return f.Int() }

// ListSeqNoField is a INT field.
type ListSeqNoField struct{ quickfix.FIXInt }

// Tag returns tag.ListSeqNo (67).
func (f ListSeqNoField) Tag() quickfix.Tag { return tag.ListSeqNo }

// NewListSeqNo returns a new ListSeqNoField initialized with val.
func NewListSeqNo(val int) ListSeqNoField {
	return ListSeqNoField{quickfix.FIXInt(val)}
}

func (f ListSeqNoField) Value() int { return f.Int() }

// ListStatusTextField is a STRING field.
type ListStatusTextField struct{ quickfix.FIXString }

// Tag returns tag.ListStatusText (444).
func (f ListStatusTextField) Tag() quickfix.Tag { return tag.ListStatusText }

// NewListStatusText returns a new ListStatusTextField initialized with val.
func NewListStatusText(val string) ListStatusTextField {
	return ListStatusTextField{quickfix.FIXString(val)}
}

func (f ListStatusTextField) Value() string { return f.String() }

// ListStatusTypeField is a INT field.
type ListStatusTypeField struct{ quickfix.FIXInt }

// Tag returns tag.ListStatusType (429).
func (f ListStatusTypeField) Tag() quickfix.Tag { return tag.ListStatusType }

// NewListStatusType returns a new ListStatusTypeField initialized with val.
func NewListStatusType(val int) ListStatusTypeField {
	return ListStatusTypeField{quickfix.FIXInt(val)}
}

func (f ListStatusTypeField) Value() int { return f.Int() }

// LocateReqdField is a BOOLEAN field.
type LocateReqdField struct{ quickfix.FIXBoolean }

// Tag returns tag.LocateReqd (114).
func (f LocateReqdField) Tag() quickfix.Tag { return tag.LocateReqd }

// NewLocateReqd returns a new LocateReqdField initialized with val.
func NewLocateReqd(val bool) LocateReqdField {
	return LocateReqdField{quickfix.FIXBoolean(val)}
}

func (f LocateReqdField) Value() bool { return f.Bool() }

// LocationIDField is a STRING field.
type LocationIDField struct{ quickfix.FIXString }

// Tag returns tag.LocationID (283).
func (f LocationIDField) Tag() quickfix.Tag { return tag.LocationID }

// NewLocationID returns a new LocationIDField initialized with val.
func NewLocationID(val string) LocationIDField {
	return LocationIDField{quickfix.FIXString(val)}
}

func (f LocationIDField) Value() string { return f.String() }

// LowPxField is a PRICE field.
type LowPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.LowPx (333).
func (f LowPxField) Tag() quickfix.Tag { return tag.LowPx }

// NewLowPx returns a new LowPxField initialized with val and scale.
func NewLowPx(val decimal.Decimal, scale int32) LowPxField {
	return LowPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LowPxField) Value() (val decimal.Decimal) { return f.Decimal }

// MDEntryBuyerField is a STRING field.
type MDEntryBuyerField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryBuyer (288).
func (f MDEntryBuyerField) Tag() quickfix.Tag { return tag.MDEntryBuyer }

// NewMDEntryBuyer returns a new MDEntryBuyerField initialized with val.
func NewMDEntryBuyer(val string) MDEntryBuyerField {
	return MDEntryBuyerField{quickfix.FIXString(val)}
}

func (f MDEntryBuyerField) Value() string { return f.String() }

// MDEntryDateField is a UTCDATE field.
type MDEntryDateField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryDate (272).
func (f MDEntryDateField) Tag() quickfix.Tag { return tag.MDEntryDate }

// NewMDEntryDate returns a new MDEntryDateField initialized with val.
func NewMDEntryDate(val string) MDEntryDateField {
	return MDEntryDateField{quickfix.FIXString(val)}
}

func (f MDEntryDateField) Value() string { return f.String() }

// MDEntryIDField is a STRING field.
type MDEntryIDField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryID (278).
func (f MDEntryIDField) Tag() quickfix.Tag { return tag.MDEntryID }

// NewMDEntryID returns a new MDEntryIDField initialized with val.
func NewMDEntryID(val string) MDEntryIDField {
	return MDEntryIDField{quickfix.FIXString(val)}
}

func (f MDEntryIDField) Value() string { return f.String() }

// MDEntryOriginatorField is a STRING field.
type MDEntryOriginatorField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryOriginator (282).
func (f MDEntryOriginatorField) Tag() quickfix.Tag { return tag.MDEntryOriginator }

// NewMDEntryOriginator returns a new MDEntryOriginatorField initialized with val.
func NewMDEntryOriginator(val string) MDEntryOriginatorField {
	return MDEntryOriginatorField{quickfix.FIXString(val)}
}

func (f MDEntryOriginatorField) Value() string { return f.String() }

// MDEntryPositionNoField is a INT field.
type MDEntryPositionNoField struct{ quickfix.FIXInt }

// Tag returns tag.MDEntryPositionNo (290).
func (f MDEntryPositionNoField) Tag() quickfix.Tag { return tag.MDEntryPositionNo }

// NewMDEntryPositionNo returns a new MDEntryPositionNoField initialized with val.
func NewMDEntryPositionNo(val int) MDEntryPositionNoField {
	return MDEntryPositionNoField{quickfix.FIXInt(val)}
}

func (f MDEntryPositionNoField) Value() int { return f.Int() }

// MDEntryPxField is a PRICE field.
type MDEntryPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.MDEntryPx (270).
func (f MDEntryPxField) Tag() quickfix.Tag { return tag.MDEntryPx }

// NewMDEntryPx returns a new MDEntryPxField initialized with val and scale.
func NewMDEntryPx(val decimal.Decimal, scale int32) MDEntryPxField {
	return MDEntryPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MDEntryPxField) Value() (val decimal.Decimal) { return f.Decimal }

// MDEntryRefIDField is a STRING field.
type MDEntryRefIDField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryRefID (280).
func (f MDEntryRefIDField) Tag() quickfix.Tag { return tag.MDEntryRefID }

// NewMDEntryRefID returns a new MDEntryRefIDField initialized with val.
func NewMDEntryRefID(val string) MDEntryRefIDField {
	return MDEntryRefIDField{quickfix.FIXString(val)}
}

func (f MDEntryRefIDField) Value() string { return f.String() }

// MDEntrySellerField is a STRING field.
type MDEntrySellerField struct{ quickfix.FIXString }

// Tag returns tag.MDEntrySeller (289).
func (f MDEntrySellerField) Tag() quickfix.Tag { return tag.MDEntrySeller }

// NewMDEntrySeller returns a new MDEntrySellerField initialized with val.
func NewMDEntrySeller(val string) MDEntrySellerField {
	return MDEntrySellerField{quickfix.FIXString(val)}
}

func (f MDEntrySellerField) Value() string { return f.String() }

// MDEntrySizeField is a QTY field.
type MDEntrySizeField struct{ quickfix.FIXDecimal }

// Tag returns tag.MDEntrySize (271).
func (f MDEntrySizeField) Tag() quickfix.Tag { return tag.MDEntrySize }

// NewMDEntrySize returns a new MDEntrySizeField initialized with val and scale.
func NewMDEntrySize(val decimal.Decimal, scale int32) MDEntrySizeField {
	return MDEntrySizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MDEntrySizeField) Value() (val decimal.Decimal) { return f.Decimal }

// MDEntryTimeField is a UTCTIMEONLY field.
type MDEntryTimeField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryTime (273).
func (f MDEntryTimeField) Tag() quickfix.Tag { return tag.MDEntryTime }

// NewMDEntryTime returns a new MDEntryTimeField initialized with val.
func NewMDEntryTime(val string) MDEntryTimeField {
	return MDEntryTimeField{quickfix.FIXString(val)}
}

func (f MDEntryTimeField) Value() string { return f.String() }

// MDEntryTypeField is a enum.MDEntryType field.
type MDEntryTypeField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryType (269).
func (f MDEntryTypeField) Tag() quickfix.Tag { return tag.MDEntryType }

func NewMDEntryType(val enum.MDEntryType) MDEntryTypeField {
	return MDEntryTypeField{quickfix.FIXString(val)}
}

func (f MDEntryTypeField) Value() enum.MDEntryType { return enum.MDEntryType(f.String()) }

// MDMktField is a EXCHANGE field.
type MDMktField struct{ quickfix.FIXString }

// Tag returns tag.MDMkt (275).
func (f MDMktField) Tag() quickfix.Tag { return tag.MDMkt }

// NewMDMkt returns a new MDMktField initialized with val.
func NewMDMkt(val string) MDMktField {
	return MDMktField{quickfix.FIXString(val)}
}

func (f MDMktField) Value() string { return f.String() }

// MDReqIDField is a STRING field.
type MDReqIDField struct{ quickfix.FIXString }

// Tag returns tag.MDReqID (262).
func (f MDReqIDField) Tag() quickfix.Tag { return tag.MDReqID }

// NewMDReqID returns a new MDReqIDField initialized with val.
func NewMDReqID(val string) MDReqIDField {
	return MDReqIDField{quickfix.FIXString(val)}
}

func (f MDReqIDField) Value() string { return f.String() }

// MDReqRejReasonField is a enum.MDReqRejReason field.
type MDReqRejReasonField struct{ quickfix.FIXString }

// Tag returns tag.MDReqRejReason (281).
func (f MDReqRejReasonField) Tag() quickfix.Tag { return tag.MDReqRejReason }

func NewMDReqRejReason(val enum.MDReqRejReason) MDReqRejReasonField {
	return MDReqRejReasonField{quickfix.FIXString(val)}
}

func (f MDReqRejReasonField) Value() enum.MDReqRejReason { return enum.MDReqRejReason(f.String()) }

// MDUpdateActionField is a enum.MDUpdateAction field.
type MDUpdateActionField struct{ quickfix.FIXString }

// Tag returns tag.MDUpdateAction (279).
func (f MDUpdateActionField) Tag() quickfix.Tag { return tag.MDUpdateAction }

func NewMDUpdateAction(val enum.MDUpdateAction) MDUpdateActionField {
	return MDUpdateActionField{quickfix.FIXString(val)}
}

func (f MDUpdateActionField) Value() enum.MDUpdateAction { return enum.MDUpdateAction(f.String()) }

// MDUpdateTypeField is a enum.MDUpdateType field.
type MDUpdateTypeField struct{ quickfix.FIXString }

// Tag returns tag.MDUpdateType (265).
func (f MDUpdateTypeField) Tag() quickfix.Tag { return tag.MDUpdateType }

func NewMDUpdateType(val enum.MDUpdateType) MDUpdateTypeField {
	return MDUpdateTypeField{quickfix.FIXString(val)}
}

func (f MDUpdateTypeField) Value() enum.MDUpdateType { return enum.MDUpdateType(f.String()) }

// MarketDepthField is a INT field.
type MarketDepthField struct{ quickfix.FIXInt }

// Tag returns tag.MarketDepth (264).
func (f MarketDepthField) Tag() quickfix.Tag { return tag.MarketDepth }

// NewMarketDepth returns a new MarketDepthField initialized with val.
func NewMarketDepth(val int) MarketDepthField {
	return MarketDepthField{quickfix.FIXInt(val)}
}

func (f MarketDepthField) Value() int { return f.Int() }

// MaturityDayField is a DAYOFMONTH field.
type MaturityDayField struct{ quickfix.FIXInt }

// Tag returns tag.MaturityDay (205).
func (f MaturityDayField) Tag() quickfix.Tag { return tag.MaturityDay }

// NewMaturityDay returns a new MaturityDayField initialized with val.
func NewMaturityDay(val int) MaturityDayField {
	return MaturityDayField{quickfix.FIXInt(val)}
}

func (f MaturityDayField) Value() int { return f.Int() }

// MaturityMonthYearField is a MONTHYEAR field.
type MaturityMonthYearField struct{ quickfix.FIXString }

// Tag returns tag.MaturityMonthYear (200).
func (f MaturityMonthYearField) Tag() quickfix.Tag { return tag.MaturityMonthYear }

// NewMaturityMonthYear returns a new MaturityMonthYearField initialized with val.
func NewMaturityMonthYear(val string) MaturityMonthYearField {
	return MaturityMonthYearField{quickfix.FIXString(val)}
}

func (f MaturityMonthYearField) Value() string { return f.String() }

// MaxFloorField is a QTY field.
type MaxFloorField struct{ quickfix.FIXDecimal }

// Tag returns tag.MaxFloor (111).
func (f MaxFloorField) Tag() quickfix.Tag { return tag.MaxFloor }

// NewMaxFloor returns a new MaxFloorField initialized with val and scale.
func NewMaxFloor(val decimal.Decimal, scale int32) MaxFloorField {
	return MaxFloorField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MaxFloorField) Value() (val decimal.Decimal) { return f.Decimal }

// MaxMessageSizeField is a INT field.
type MaxMessageSizeField struct{ quickfix.FIXInt }

// Tag returns tag.MaxMessageSize (383).
func (f MaxMessageSizeField) Tag() quickfix.Tag { return tag.MaxMessageSize }

// NewMaxMessageSize returns a new MaxMessageSizeField initialized with val.
func NewMaxMessageSize(val int) MaxMessageSizeField {
	return MaxMessageSizeField{quickfix.FIXInt(val)}
}

func (f MaxMessageSizeField) Value() int { return f.Int() }

// MaxShowField is a QTY field.
type MaxShowField struct{ quickfix.FIXDecimal }

// Tag returns tag.MaxShow (210).
func (f MaxShowField) Tag() quickfix.Tag { return tag.MaxShow }

// NewMaxShow returns a new MaxShowField initialized with val and scale.
func NewMaxShow(val decimal.Decimal, scale int32) MaxShowField {
	return MaxShowField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MaxShowField) Value() (val decimal.Decimal) { return f.Decimal }

// MessageEncodingField is a enum.MessageEncoding field.
type MessageEncodingField struct{ quickfix.FIXString }

// Tag returns tag.MessageEncoding (347).
func (f MessageEncodingField) Tag() quickfix.Tag { return tag.MessageEncoding }

func NewMessageEncoding(val enum.MessageEncoding) MessageEncodingField {
	return MessageEncodingField{quickfix.FIXString(val)}
}

func (f MessageEncodingField) Value() enum.MessageEncoding { return enum.MessageEncoding(f.String()) }

// MinQtyField is a QTY field.
type MinQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.MinQty (110).
func (f MinQtyField) Tag() quickfix.Tag { return tag.MinQty }

// NewMinQty returns a new MinQtyField initialized with val and scale.
func NewMinQty(val decimal.Decimal, scale int32) MinQtyField {
	return MinQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MinQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// MiscFeeAmtField is a AMT field.
type MiscFeeAmtField struct{ quickfix.FIXDecimal }

// Tag returns tag.MiscFeeAmt (137).
func (f MiscFeeAmtField) Tag() quickfix.Tag { return tag.MiscFeeAmt }

// NewMiscFeeAmt returns a new MiscFeeAmtField initialized with val and scale.
func NewMiscFeeAmt(val decimal.Decimal, scale int32) MiscFeeAmtField {
	return MiscFeeAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MiscFeeAmtField) Value() (val decimal.Decimal) { return f.Decimal }

// MiscFeeCurrField is a CURRENCY field.
type MiscFeeCurrField struct{ quickfix.FIXString }

// Tag returns tag.MiscFeeCurr (138).
func (f MiscFeeCurrField) Tag() quickfix.Tag { return tag.MiscFeeCurr }

// NewMiscFeeCurr returns a new MiscFeeCurrField initialized with val.
func NewMiscFeeCurr(val string) MiscFeeCurrField {
	return MiscFeeCurrField{quickfix.FIXString(val)}
}

func (f MiscFeeCurrField) Value() string { return f.String() }

// MiscFeeTypeField is a enum.MiscFeeType field.
type MiscFeeTypeField struct{ quickfix.FIXString }

// Tag returns tag.MiscFeeType (139).
func (f MiscFeeTypeField) Tag() quickfix.Tag { return tag.MiscFeeType }

func NewMiscFeeType(val enum.MiscFeeType) MiscFeeTypeField {
	return MiscFeeTypeField{quickfix.FIXString(val)}
}

func (f MiscFeeTypeField) Value() enum.MiscFeeType { return enum.MiscFeeType(f.String()) }

// MsgDirectionField is a enum.MsgDirection field.
type MsgDirectionField struct{ quickfix.FIXString }

// Tag returns tag.MsgDirection (385).
func (f MsgDirectionField) Tag() quickfix.Tag { return tag.MsgDirection }

func NewMsgDirection(val enum.MsgDirection) MsgDirectionField {
	return MsgDirectionField{quickfix.FIXString(val)}
}

func (f MsgDirectionField) Value() enum.MsgDirection { return enum.MsgDirection(f.String()) }

// MsgSeqNumField is a INT field.
type MsgSeqNumField struct{ quickfix.FIXInt }

// Tag returns tag.MsgSeqNum (34).
func (f MsgSeqNumField) Tag() quickfix.Tag { return tag.MsgSeqNum }

// NewMsgSeqNum returns a new MsgSeqNumField initialized with val.
func NewMsgSeqNum(val int) MsgSeqNumField {
	return MsgSeqNumField{quickfix.FIXInt(val)}
}

func (f MsgSeqNumField) Value() int { return f.Int() }

// MsgTypeField is a enum.MsgType field.
type MsgTypeField struct{ quickfix.FIXString }

// Tag returns tag.MsgType (35).
func (f MsgTypeField) Tag() quickfix.Tag { return tag.MsgType }

func NewMsgType(val enum.MsgType) MsgTypeField {
	return MsgTypeField{quickfix.FIXString(val)}
}

func (f MsgTypeField) Value() enum.MsgType { return enum.MsgType(f.String()) }

// MultiLegReportingTypeField is a enum.MultiLegReportingType field.
type MultiLegReportingTypeField struct{ quickfix.FIXString }

// Tag returns tag.MultiLegReportingType (442).
func (f MultiLegReportingTypeField) Tag() quickfix.Tag { return tag.MultiLegReportingType }

func NewMultiLegReportingType(val enum.MultiLegReportingType) MultiLegReportingTypeField {
	return MultiLegReportingTypeField{quickfix.FIXString(val)}
}

func (f MultiLegReportingTypeField) Value() enum.MultiLegReportingType {
	return enum.MultiLegReportingType(f.String())
}

// NetGrossIndField is a enum.NetGrossInd field.
type NetGrossIndField struct{ quickfix.FIXString }

// Tag returns tag.NetGrossInd (430).
func (f NetGrossIndField) Tag() quickfix.Tag { return tag.NetGrossInd }

func NewNetGrossInd(val enum.NetGrossInd) NetGrossIndField {
	return NetGrossIndField{quickfix.FIXString(val)}
}

func (f NetGrossIndField) Value() enum.NetGrossInd { return enum.NetGrossInd(f.String()) }

// NetMoneyField is a AMT field.
type NetMoneyField struct{ quickfix.FIXDecimal }

// Tag returns tag.NetMoney (118).
func (f NetMoneyField) Tag() quickfix.Tag { return tag.NetMoney }

// NewNetMoney returns a new NetMoneyField initialized with val and scale.
func NewNetMoney(val decimal.Decimal, scale int32) NetMoneyField {
	return NetMoneyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f NetMoneyField) Value() (val decimal.Decimal) { return f.Decimal }

// NewSeqNoField is a INT field.
type NewSeqNoField struct{ quickfix.FIXInt }

// Tag returns tag.NewSeqNo (36).
func (f NewSeqNoField) Tag() quickfix.Tag { return tag.NewSeqNo }

// NewNewSeqNo returns a new NewSeqNoField initialized with val.
func NewNewSeqNo(val int) NewSeqNoField {
	return NewSeqNoField{quickfix.FIXInt(val)}
}

func (f NewSeqNoField) Value() int { return f.Int() }

// NoAllocsField is a INT field.
type NoAllocsField struct{ quickfix.FIXInt }

// Tag returns tag.NoAllocs (78).
func (f NoAllocsField) Tag() quickfix.Tag { return tag.NoAllocs }

// NewNoAllocs returns a new NoAllocsField initialized with val.
func NewNoAllocs(val int) NoAllocsField {
	return NoAllocsField{quickfix.FIXInt(val)}
}

func (f NoAllocsField) Value() int { return f.Int() }

// NoBidComponentsField is a INT field.
type NoBidComponentsField struct{ quickfix.FIXInt }

// Tag returns tag.NoBidComponents (420).
func (f NoBidComponentsField) Tag() quickfix.Tag { return tag.NoBidComponents }

// NewNoBidComponents returns a new NoBidComponentsField initialized with val.
func NewNoBidComponents(val int) NoBidComponentsField {
	return NoBidComponentsField{quickfix.FIXInt(val)}
}

func (f NoBidComponentsField) Value() int { return f.Int() }

// NoBidDescriptorsField is a INT field.
type NoBidDescriptorsField struct{ quickfix.FIXInt }

// Tag returns tag.NoBidDescriptors (398).
func (f NoBidDescriptorsField) Tag() quickfix.Tag { return tag.NoBidDescriptors }

// NewNoBidDescriptors returns a new NoBidDescriptorsField initialized with val.
func NewNoBidDescriptors(val int) NoBidDescriptorsField {
	return NoBidDescriptorsField{quickfix.FIXInt(val)}
}

func (f NoBidDescriptorsField) Value() int { return f.Int() }

// NoContraBrokersField is a INT field.
type NoContraBrokersField struct{ quickfix.FIXInt }

// Tag returns tag.NoContraBrokers (382).
func (f NoContraBrokersField) Tag() quickfix.Tag { return tag.NoContraBrokers }

// NewNoContraBrokers returns a new NoContraBrokersField initialized with val.
func NewNoContraBrokers(val int) NoContraBrokersField {
	return NoContraBrokersField{quickfix.FIXInt(val)}
}

func (f NoContraBrokersField) Value() int { return f.Int() }

// NoDlvyInstField is a INT field.
type NoDlvyInstField struct{ quickfix.FIXInt }

// Tag returns tag.NoDlvyInst (85).
func (f NoDlvyInstField) Tag() quickfix.Tag { return tag.NoDlvyInst }

// NewNoDlvyInst returns a new NoDlvyInstField initialized with val.
func NewNoDlvyInst(val int) NoDlvyInstField {
	return NoDlvyInstField{quickfix.FIXInt(val)}
}

func (f NoDlvyInstField) Value() int { return f.Int() }

// NoExecsField is a INT field.
type NoExecsField struct{ quickfix.FIXInt }

// Tag returns tag.NoExecs (124).
func (f NoExecsField) Tag() quickfix.Tag { return tag.NoExecs }

// NewNoExecs returns a new NoExecsField initialized with val.
func NewNoExecs(val int) NoExecsField {
	return NoExecsField{quickfix.FIXInt(val)}
}

func (f NoExecsField) Value() int { return f.Int() }

// NoIOIQualifiersField is a INT field.
type NoIOIQualifiersField struct{ quickfix.FIXInt }

// Tag returns tag.NoIOIQualifiers (199).
func (f NoIOIQualifiersField) Tag() quickfix.Tag { return tag.NoIOIQualifiers }

// NewNoIOIQualifiers returns a new NoIOIQualifiersField initialized with val.
func NewNoIOIQualifiers(val int) NoIOIQualifiersField {
	return NoIOIQualifiersField{quickfix.FIXInt(val)}
}

func (f NoIOIQualifiersField) Value() int { return f.Int() }

// NoMDEntriesField is a INT field.
type NoMDEntriesField struct{ quickfix.FIXInt }

// Tag returns tag.NoMDEntries (268).
func (f NoMDEntriesField) Tag() quickfix.Tag { return tag.NoMDEntries }

// NewNoMDEntries returns a new NoMDEntriesField initialized with val.
func NewNoMDEntries(val int) NoMDEntriesField {
	return NoMDEntriesField{quickfix.FIXInt(val)}
}

func (f NoMDEntriesField) Value() int { return f.Int() }

// NoMDEntryTypesField is a INT field.
type NoMDEntryTypesField struct{ quickfix.FIXInt }

// Tag returns tag.NoMDEntryTypes (267).
func (f NoMDEntryTypesField) Tag() quickfix.Tag { return tag.NoMDEntryTypes }

// NewNoMDEntryTypes returns a new NoMDEntryTypesField initialized with val.
func NewNoMDEntryTypes(val int) NoMDEntryTypesField {
	return NoMDEntryTypesField{quickfix.FIXInt(val)}
}

func (f NoMDEntryTypesField) Value() int { return f.Int() }

// NoMiscFeesField is a INT field.
type NoMiscFeesField struct{ quickfix.FIXInt }

// Tag returns tag.NoMiscFees (136).
func (f NoMiscFeesField) Tag() quickfix.Tag { return tag.NoMiscFees }

// NewNoMiscFees returns a new NoMiscFeesField initialized with val.
func NewNoMiscFees(val int) NoMiscFeesField {
	return NoMiscFeesField{quickfix.FIXInt(val)}
}

func (f NoMiscFeesField) Value() int { return f.Int() }

// NoMsgTypesField is a INT field.
type NoMsgTypesField struct{ quickfix.FIXInt }

// Tag returns tag.NoMsgTypes (384).
func (f NoMsgTypesField) Tag() quickfix.Tag { return tag.NoMsgTypes }

// NewNoMsgTypes returns a new NoMsgTypesField initialized with val.
func NewNoMsgTypes(val int) NoMsgTypesField {
	return NoMsgTypesField{quickfix.FIXInt(val)}
}

func (f NoMsgTypesField) Value() int { return f.Int() }

// NoOrdersField is a INT field.
type NoOrdersField struct{ quickfix.FIXInt }

// Tag returns tag.NoOrders (73).
func (f NoOrdersField) Tag() quickfix.Tag { return tag.NoOrders }

// NewNoOrders returns a new NoOrdersField initialized with val.
func NewNoOrders(val int) NoOrdersField {
	return NoOrdersField{quickfix.FIXInt(val)}
}

func (f NoOrdersField) Value() int { return f.Int() }

// NoQuoteEntriesField is a INT field.
type NoQuoteEntriesField struct{ quickfix.FIXInt }

// Tag returns tag.NoQuoteEntries (295).
func (f NoQuoteEntriesField) Tag() quickfix.Tag { return tag.NoQuoteEntries }

// NewNoQuoteEntries returns a new NoQuoteEntriesField initialized with val.
func NewNoQuoteEntries(val int) NoQuoteEntriesField {
	return NoQuoteEntriesField{quickfix.FIXInt(val)}
}

func (f NoQuoteEntriesField) Value() int { return f.Int() }

// NoQuoteSetsField is a INT field.
type NoQuoteSetsField struct{ quickfix.FIXInt }

// Tag returns tag.NoQuoteSets (296).
func (f NoQuoteSetsField) Tag() quickfix.Tag { return tag.NoQuoteSets }

// NewNoQuoteSets returns a new NoQuoteSetsField initialized with val.
func NewNoQuoteSets(val int) NoQuoteSetsField {
	return NoQuoteSetsField{quickfix.FIXInt(val)}
}

func (f NoQuoteSetsField) Value() int { return f.Int() }

// NoRelatedSymField is a INT field.
type NoRelatedSymField struct{ quickfix.FIXInt }

// Tag returns tag.NoRelatedSym (146).
func (f NoRelatedSymField) Tag() quickfix.Tag { return tag.NoRelatedSym }

// NewNoRelatedSym returns a new NoRelatedSymField initialized with val.
func NewNoRelatedSym(val int) NoRelatedSymField {
	return NoRelatedSymField{quickfix.FIXInt(val)}
}

func (f NoRelatedSymField) Value() int { return f.Int() }

// NoRoutingIDsField is a INT field.
type NoRoutingIDsField struct{ quickfix.FIXInt }

// Tag returns tag.NoRoutingIDs (215).
func (f NoRoutingIDsField) Tag() quickfix.Tag { return tag.NoRoutingIDs }

// NewNoRoutingIDs returns a new NoRoutingIDsField initialized with val.
func NewNoRoutingIDs(val int) NoRoutingIDsField {
	return NoRoutingIDsField{quickfix.FIXInt(val)}
}

func (f NoRoutingIDsField) Value() int { return f.Int() }

// NoRptsField is a INT field.
type NoRptsField struct{ quickfix.FIXInt }

// Tag returns tag.NoRpts (82).
func (f NoRptsField) Tag() quickfix.Tag { return tag.NoRpts }

// NewNoRpts returns a new NoRptsField initialized with val.
func NewNoRpts(val int) NoRptsField {
	return NoRptsField{quickfix.FIXInt(val)}
}

func (f NoRptsField) Value() int { return f.Int() }

// NoStrikesField is a INT field.
type NoStrikesField struct{ quickfix.FIXInt }

// Tag returns tag.NoStrikes (428).
func (f NoStrikesField) Tag() quickfix.Tag { return tag.NoStrikes }

// NewNoStrikes returns a new NoStrikesField initialized with val.
func NewNoStrikes(val int) NoStrikesField {
	return NoStrikesField{quickfix.FIXInt(val)}
}

func (f NoStrikesField) Value() int { return f.Int() }

// NoTradingSessionsField is a INT field.
type NoTradingSessionsField struct{ quickfix.FIXInt }

// Tag returns tag.NoTradingSessions (386).
func (f NoTradingSessionsField) Tag() quickfix.Tag { return tag.NoTradingSessions }

// NewNoTradingSessions returns a new NoTradingSessionsField initialized with val.
func NewNoTradingSessions(val int) NoTradingSessionsField {
	return NoTradingSessionsField{quickfix.FIXInt(val)}
}

func (f NoTradingSessionsField) Value() int { return f.Int() }

// NotifyBrokerOfCreditField is a BOOLEAN field.
type NotifyBrokerOfCreditField struct{ quickfix.FIXBoolean }

// Tag returns tag.NotifyBrokerOfCredit (208).
func (f NotifyBrokerOfCreditField) Tag() quickfix.Tag { return tag.NotifyBrokerOfCredit }

// NewNotifyBrokerOfCredit returns a new NotifyBrokerOfCreditField initialized with val.
func NewNotifyBrokerOfCredit(val bool) NotifyBrokerOfCreditField {
	return NotifyBrokerOfCreditField{quickfix.FIXBoolean(val)}
}

func (f NotifyBrokerOfCreditField) Value() bool { return f.Bool() }

// NumBiddersField is a INT field.
type NumBiddersField struct{ quickfix.FIXInt }

// Tag returns tag.NumBidders (417).
func (f NumBiddersField) Tag() quickfix.Tag { return tag.NumBidders }

// NewNumBidders returns a new NumBiddersField initialized with val.
func NewNumBidders(val int) NumBiddersField {
	return NumBiddersField{quickfix.FIXInt(val)}
}

func (f NumBiddersField) Value() int { return f.Int() }

// NumDaysInterestField is a INT field.
type NumDaysInterestField struct{ quickfix.FIXInt }

// Tag returns tag.NumDaysInterest (157).
func (f NumDaysInterestField) Tag() quickfix.Tag { return tag.NumDaysInterest }

// NewNumDaysInterest returns a new NumDaysInterestField initialized with val.
func NewNumDaysInterest(val int) NumDaysInterestField {
	return NumDaysInterestField{quickfix.FIXInt(val)}
}

func (f NumDaysInterestField) Value() int { return f.Int() }

// NumTicketsField is a INT field.
type NumTicketsField struct{ quickfix.FIXInt }

// Tag returns tag.NumTickets (395).
func (f NumTicketsField) Tag() quickfix.Tag { return tag.NumTickets }

// NewNumTickets returns a new NumTicketsField initialized with val.
func NewNumTickets(val int) NumTicketsField {
	return NumTicketsField{quickfix.FIXInt(val)}
}

func (f NumTicketsField) Value() int { return f.Int() }

// NumberOfOrdersField is a INT field.
type NumberOfOrdersField struct{ quickfix.FIXInt }

// Tag returns tag.NumberOfOrders (346).
func (f NumberOfOrdersField) Tag() quickfix.Tag { return tag.NumberOfOrders }

// NewNumberOfOrders returns a new NumberOfOrdersField initialized with val.
func NewNumberOfOrders(val int) NumberOfOrdersField {
	return NumberOfOrdersField{quickfix.FIXInt(val)}
}

func (f NumberOfOrdersField) Value() int { return f.Int() }

// OfferForwardPointsField is a PRICEOFFSET field.
type OfferForwardPointsField struct{ quickfix.FIXDecimal }

// Tag returns tag.OfferForwardPoints (191).
func (f OfferForwardPointsField) Tag() quickfix.Tag { return tag.OfferForwardPoints }

// NewOfferForwardPoints returns a new OfferForwardPointsField initialized with val and scale.
func NewOfferForwardPoints(val decimal.Decimal, scale int32) OfferForwardPointsField {
	return OfferForwardPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferForwardPointsField) Value() (val decimal.Decimal) { return f.Decimal }

// OfferPxField is a PRICE field.
type OfferPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.OfferPx (133).
func (f OfferPxField) Tag() quickfix.Tag { return tag.OfferPx }

// NewOfferPx returns a new OfferPxField initialized with val and scale.
func NewOfferPx(val decimal.Decimal, scale int32) OfferPxField {
	return OfferPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferPxField) Value() (val decimal.Decimal) { return f.Decimal }

// OfferSizeField is a QTY field.
type OfferSizeField struct{ quickfix.FIXDecimal }

// Tag returns tag.OfferSize (135).
func (f OfferSizeField) Tag() quickfix.Tag { return tag.OfferSize }

// NewOfferSize returns a new OfferSizeField initialized with val and scale.
func NewOfferSize(val decimal.Decimal, scale int32) OfferSizeField {
	return OfferSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferSizeField) Value() (val decimal.Decimal) { return f.Decimal }

// OfferSpotRateField is a PRICE field.
type OfferSpotRateField struct{ quickfix.FIXDecimal }

// Tag returns tag.OfferSpotRate (190).
func (f OfferSpotRateField) Tag() quickfix.Tag { return tag.OfferSpotRate }

// NewOfferSpotRate returns a new OfferSpotRateField initialized with val and scale.
func NewOfferSpotRate(val decimal.Decimal, scale int32) OfferSpotRateField {
	return OfferSpotRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferSpotRateField) Value() (val decimal.Decimal) { return f.Decimal }

// OnBehalfOfCompIDField is a STRING field.
type OnBehalfOfCompIDField struct{ quickfix.FIXString }

// Tag returns tag.OnBehalfOfCompID (115).
func (f OnBehalfOfCompIDField) Tag() quickfix.Tag { return tag.OnBehalfOfCompID }

// NewOnBehalfOfCompID returns a new OnBehalfOfCompIDField initialized with val.
func NewOnBehalfOfCompID(val string) OnBehalfOfCompIDField {
	return OnBehalfOfCompIDField{quickfix.FIXString(val)}
}

func (f OnBehalfOfCompIDField) Value() string { return f.String() }

// OnBehalfOfLocationIDField is a STRING field.
type OnBehalfOfLocationIDField struct{ quickfix.FIXString }

// Tag returns tag.OnBehalfOfLocationID (144).
func (f OnBehalfOfLocationIDField) Tag() quickfix.Tag { return tag.OnBehalfOfLocationID }

// NewOnBehalfOfLocationID returns a new OnBehalfOfLocationIDField initialized with val.
func NewOnBehalfOfLocationID(val string) OnBehalfOfLocationIDField {
	return OnBehalfOfLocationIDField{quickfix.FIXString(val)}
}

func (f OnBehalfOfLocationIDField) Value() string { return f.String() }

// OnBehalfOfSendingTimeField is a UTCTIMESTAMP field.
type OnBehalfOfSendingTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.OnBehalfOfSendingTime (370).
func (f OnBehalfOfSendingTimeField) Tag() quickfix.Tag { return tag.OnBehalfOfSendingTime }

// NewOnBehalfOfSendingTime returns a new OnBehalfOfSendingTimeField initialized with val.
func NewOnBehalfOfSendingTime(val time.Time) OnBehalfOfSendingTimeField {
	return NewOnBehalfOfSendingTimeWithPrecision(val, quickfix.Millis)
}

// NewOnBehalfOfSendingTimeNoMillis returns a new OnBehalfOfSendingTimeField initialized with val without millisecs.
func NewOnBehalfOfSendingTimeNoMillis(val time.Time) OnBehalfOfSendingTimeField {
	return NewOnBehalfOfSendingTimeWithPrecision(val, quickfix.Seconds)
}

// NewOnBehalfOfSendingTimeWithPrecision returns a new OnBehalfOfSendingTimeField initialized with val of specified precision.
func NewOnBehalfOfSendingTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) OnBehalfOfSendingTimeField {
	return OnBehalfOfSendingTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f OnBehalfOfSendingTimeField) Value() time.Time { return f.Time }

// OnBehalfOfSubIDField is a STRING field.
type OnBehalfOfSubIDField struct{ quickfix.FIXString }

// Tag returns tag.OnBehalfOfSubID (116).
func (f OnBehalfOfSubIDField) Tag() quickfix.Tag { return tag.OnBehalfOfSubID }

// NewOnBehalfOfSubID returns a new OnBehalfOfSubIDField initialized with val.
func NewOnBehalfOfSubID(val string) OnBehalfOfSubIDField {
	return OnBehalfOfSubIDField{quickfix.FIXString(val)}
}

func (f OnBehalfOfSubIDField) Value() string { return f.String() }

// OpenCloseField is a enum.OpenClose field.
type OpenCloseField struct{ quickfix.FIXString }

// Tag returns tag.OpenClose (77).
func (f OpenCloseField) Tag() quickfix.Tag { return tag.OpenClose }

func NewOpenClose(val enum.OpenClose) OpenCloseField {
	return OpenCloseField{quickfix.FIXString(val)}
}

func (f OpenCloseField) Value() enum.OpenClose { return enum.OpenClose(f.String()) }

// OpenCloseSettleFlagField is a enum.OpenCloseSettleFlag field.
type OpenCloseSettleFlagField struct{ quickfix.FIXString }

// Tag returns tag.OpenCloseSettleFlag (286).
func (f OpenCloseSettleFlagField) Tag() quickfix.Tag { return tag.OpenCloseSettleFlag }

func NewOpenCloseSettleFlag(val enum.OpenCloseSettleFlag) OpenCloseSettleFlagField {
	return OpenCloseSettleFlagField{quickfix.FIXString(val)}
}

func (f OpenCloseSettleFlagField) Value() enum.OpenCloseSettleFlag {
	return enum.OpenCloseSettleFlag(f.String())
}

// OptAttributeField is a CHAR field.
type OptAttributeField struct{ quickfix.FIXString }

// Tag returns tag.OptAttribute (206).
func (f OptAttributeField) Tag() quickfix.Tag { return tag.OptAttribute }

// NewOptAttribute returns a new OptAttributeField initialized with val.
func NewOptAttribute(val string) OptAttributeField {
	return OptAttributeField{quickfix.FIXString(val)}
}

func (f OptAttributeField) Value() string { return f.String() }

// OrdRejReasonField is a enum.OrdRejReason field.
type OrdRejReasonField struct{ quickfix.FIXString }

// Tag returns tag.OrdRejReason (103).
func (f OrdRejReasonField) Tag() quickfix.Tag { return tag.OrdRejReason }

func NewOrdRejReason(val enum.OrdRejReason) OrdRejReasonField {
	return OrdRejReasonField{quickfix.FIXString(val)}
}

func (f OrdRejReasonField) Value() enum.OrdRejReason { return enum.OrdRejReason(f.String()) }

// OrdStatusField is a enum.OrdStatus field.
type OrdStatusField struct{ quickfix.FIXString }

// Tag returns tag.OrdStatus (39).
func (f OrdStatusField) Tag() quickfix.Tag { return tag.OrdStatus }

func NewOrdStatus(val enum.OrdStatus) OrdStatusField {
	return OrdStatusField{quickfix.FIXString(val)}
}

func (f OrdStatusField) Value() enum.OrdStatus { return enum.OrdStatus(f.String()) }

// OrdTypeField is a enum.OrdType field.
type OrdTypeField struct{ quickfix.FIXString }

// Tag returns tag.OrdType (40).
func (f OrdTypeField) Tag() quickfix.Tag { return tag.OrdType }

func NewOrdType(val enum.OrdType) OrdTypeField {
	return OrdTypeField{quickfix.FIXString(val)}
}

func (f OrdTypeField) Value() enum.OrdType { return enum.OrdType(f.String()) }

// OrderIDField is a STRING field.
type OrderIDField struct{ quickfix.FIXString }

// Tag returns tag.OrderID (37).
func (f OrderIDField) Tag() quickfix.Tag { return tag.OrderID }

// NewOrderID returns a new OrderIDField initialized with val.
func NewOrderID(val string) OrderIDField {
	return OrderIDField{quickfix.FIXString(val)}
}

func (f OrderIDField) Value() string { return f.String() }

// OrderQtyField is a QTY field.
type OrderQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.OrderQty (38).
func (f OrderQtyField) Tag() quickfix.Tag { return tag.OrderQty }

// NewOrderQty returns a new OrderQtyField initialized with val and scale.
func NewOrderQty(val decimal.Decimal, scale int32) OrderQtyField {
	return OrderQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OrderQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// OrderQty2Field is a QTY field.
type OrderQty2Field struct{ quickfix.FIXDecimal }

// Tag returns tag.OrderQty2 (192).
func (f OrderQty2Field) Tag() quickfix.Tag { return tag.OrderQty2 }

// NewOrderQty2 returns a new OrderQty2Field initialized with val and scale.
func NewOrderQty2(val decimal.Decimal, scale int32) OrderQty2Field {
	return OrderQty2Field{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OrderQty2Field) Value() (val decimal.Decimal) { return f.Decimal }

// OrigClOrdIDField is a STRING field.
type OrigClOrdIDField struct{ quickfix.FIXString }

// Tag returns tag.OrigClOrdID (41).
func (f OrigClOrdIDField) Tag() quickfix.Tag { return tag.OrigClOrdID }

// NewOrigClOrdID returns a new OrigClOrdIDField initialized with val.
func NewOrigClOrdID(val string) OrigClOrdIDField {
	return OrigClOrdIDField{quickfix.FIXString(val)}
}

func (f OrigClOrdIDField) Value() string { return f.String() }

// OrigSendingTimeField is a UTCTIMESTAMP field.
type OrigSendingTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.OrigSendingTime (122).
func (f OrigSendingTimeField) Tag() quickfix.Tag { return tag.OrigSendingTime }

// NewOrigSendingTime returns a new OrigSendingTimeField initialized with val.
func NewOrigSendingTime(val time.Time) OrigSendingTimeField {
	return NewOrigSendingTimeWithPrecision(val, quickfix.Millis)
}

// NewOrigSendingTimeNoMillis returns a new OrigSendingTimeField initialized with val without millisecs.
func NewOrigSendingTimeNoMillis(val time.Time) OrigSendingTimeField {
	return NewOrigSendingTimeWithPrecision(val, quickfix.Seconds)
}

// NewOrigSendingTimeWithPrecision returns a new OrigSendingTimeField initialized with val of specified precision.
func NewOrigSendingTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) OrigSendingTimeField {
	return OrigSendingTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f OrigSendingTimeField) Value() time.Time { return f.Time }

// OrigTimeField is a UTCTIMESTAMP field.
type OrigTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.OrigTime (42).
func (f OrigTimeField) Tag() quickfix.Tag { return tag.OrigTime }

// NewOrigTime returns a new OrigTimeField initialized with val.
func NewOrigTime(val time.Time) OrigTimeField {
	return NewOrigTimeWithPrecision(val, quickfix.Millis)
}

// NewOrigTimeNoMillis returns a new OrigTimeField initialized with val without millisecs.
func NewOrigTimeNoMillis(val time.Time) OrigTimeField {
	return NewOrigTimeWithPrecision(val, quickfix.Seconds)
}

// NewOrigTimeWithPrecision returns a new OrigTimeField initialized with val of specified precision.
func NewOrigTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) OrigTimeField {
	return OrigTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f OrigTimeField) Value() time.Time { return f.Time }

// OutMainCntryUIndexField is a AMT field.
type OutMainCntryUIndexField struct{ quickfix.FIXDecimal }

// Tag returns tag.OutMainCntryUIndex (412).
func (f OutMainCntryUIndexField) Tag() quickfix.Tag { return tag.OutMainCntryUIndex }

// NewOutMainCntryUIndex returns a new OutMainCntryUIndexField initialized with val and scale.
func NewOutMainCntryUIndex(val decimal.Decimal, scale int32) OutMainCntryUIndexField {
	return OutMainCntryUIndexField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OutMainCntryUIndexField) Value() (val decimal.Decimal) { return f.Decimal }

// OutsideIndexPctField is a FLOAT field.
type OutsideIndexPctField struct{ quickfix.FIXDecimal }

// Tag returns tag.OutsideIndexPct (407).
func (f OutsideIndexPctField) Tag() quickfix.Tag { return tag.OutsideIndexPct }

// NewOutsideIndexPct returns a new OutsideIndexPctField initialized with val and scale.
func NewOutsideIndexPct(val decimal.Decimal, scale int32) OutsideIndexPctField {
	return OutsideIndexPctField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OutsideIndexPctField) Value() (val decimal.Decimal) { return f.Decimal }

// PasswordField is a STRING field.
type PasswordField struct{ quickfix.FIXString }

// Tag returns tag.Password (554).
func (f PasswordField) Tag() quickfix.Tag { return tag.Password }

// NewPassword returns a new PasswordField initialized with val.
func NewPassword(val string) PasswordField {
	return PasswordField{quickfix.FIXString(val)}
}

func (f PasswordField) Value() string { return f.String() }

// PegDifferenceField is a PRICEOFFSET field.
type PegDifferenceField struct{ quickfix.FIXDecimal }

// Tag returns tag.PegDifference (211).
func (f PegDifferenceField) Tag() quickfix.Tag { return tag.PegDifference }

// NewPegDifference returns a new PegDifferenceField initialized with val and scale.
func NewPegDifference(val decimal.Decimal, scale int32) PegDifferenceField {
	return PegDifferenceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PegDifferenceField) Value() (val decimal.Decimal) { return f.Decimal }

// PossDupFlagField is a BOOLEAN field.
type PossDupFlagField struct{ quickfix.FIXBoolean }

// Tag returns tag.PossDupFlag (43).
func (f PossDupFlagField) Tag() quickfix.Tag { return tag.PossDupFlag }

// NewPossDupFlag returns a new PossDupFlagField initialized with val.
func NewPossDupFlag(val bool) PossDupFlagField {
	return PossDupFlagField{quickfix.FIXBoolean(val)}
}

func (f PossDupFlagField) Value() bool { return f.Bool() }

// PossResendField is a BOOLEAN field.
type PossResendField struct{ quickfix.FIXBoolean }

// Tag returns tag.PossResend (97).
func (f PossResendField) Tag() quickfix.Tag { return tag.PossResend }

// NewPossResend returns a new PossResendField initialized with val.
func NewPossResend(val bool) PossResendField {
	return PossResendField{quickfix.FIXBoolean(val)}
}

func (f PossResendField) Value() bool { return f.Bool() }

// PrevClosePxField is a PRICE field.
type PrevClosePxField struct{ quickfix.FIXDecimal }

// Tag returns tag.PrevClosePx (140).
func (f PrevClosePxField) Tag() quickfix.Tag { return tag.PrevClosePx }

// NewPrevClosePx returns a new PrevClosePxField initialized with val and scale.
func NewPrevClosePx(val decimal.Decimal, scale int32) PrevClosePxField {
	return PrevClosePxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PrevClosePxField) Value() (val decimal.Decimal) { return f.Decimal }

// PriceField is a PRICE field.
type PriceField struct{ quickfix.FIXDecimal }

// Tag returns tag.Price (44).
func (f PriceField) Tag() quickfix.Tag { return tag.Price }

// NewPrice returns a new PriceField initialized with val and scale.
func NewPrice(val decimal.Decimal, scale int32) PriceField {
	return PriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PriceField) Value() (val decimal.Decimal) { return f.Decimal }

// PriceTypeField is a enum.PriceType field.
type PriceTypeField struct{ quickfix.FIXString }

// Tag returns tag.PriceType (423).
func (f PriceTypeField) Tag() quickfix.Tag { return tag.PriceType }

func NewPriceType(val enum.PriceType) PriceTypeField {
	return PriceTypeField{quickfix.FIXString(val)}
}

func (f PriceTypeField) Value() enum.PriceType { return enum.PriceType(f.String()) }

// ProcessCodeField is a enum.ProcessCode field.
type ProcessCodeField struct{ quickfix.FIXString }

// Tag returns tag.ProcessCode (81).
func (f ProcessCodeField) Tag() quickfix.Tag { return tag.ProcessCode }

func NewProcessCode(val enum.ProcessCode) ProcessCodeField {
	return ProcessCodeField{quickfix.FIXString(val)}
}

func (f ProcessCodeField) Value() enum.ProcessCode { return enum.ProcessCode(f.String()) }

// ProgPeriodIntervalField is a INT field.
type ProgPeriodIntervalField struct{ quickfix.FIXInt }

// Tag returns tag.ProgPeriodInterval (415).
func (f ProgPeriodIntervalField) Tag() quickfix.Tag { return tag.ProgPeriodInterval }

// NewProgPeriodInterval returns a new ProgPeriodIntervalField initialized with val.
func NewProgPeriodInterval(val int) ProgPeriodIntervalField {
	return ProgPeriodIntervalField{quickfix.FIXInt(val)}
}

func (f ProgPeriodIntervalField) Value() int { return f.Int() }

// ProgRptReqsField is a enum.ProgRptReqs field.
type ProgRptReqsField struct{ quickfix.FIXString }

// Tag returns tag.ProgRptReqs (414).
func (f ProgRptReqsField) Tag() quickfix.Tag { return tag.ProgRptReqs }

func NewProgRptReqs(val enum.ProgRptReqs) ProgRptReqsField {
	return ProgRptReqsField{quickfix.FIXString(val)}
}

func (f ProgRptReqsField) Value() enum.ProgRptReqs { return enum.ProgRptReqs(f.String()) }

// PutOrCallField is a enum.PutOrCall field.
type PutOrCallField struct{ quickfix.FIXString }

// Tag returns tag.PutOrCall (201).
func (f PutOrCallField) Tag() quickfix.Tag { return tag.PutOrCall }

func NewPutOrCall(val enum.PutOrCall) PutOrCallField {
	return PutOrCallField{quickfix.FIXString(val)}
}

func (f PutOrCallField) Value() enum.PutOrCall { return enum.PutOrCall(f.String()) }

// QuoteAckStatusField is a enum.QuoteAckStatus field.
type QuoteAckStatusField struct{ quickfix.FIXString }

// Tag returns tag.QuoteAckStatus (297).
func (f QuoteAckStatusField) Tag() quickfix.Tag { return tag.QuoteAckStatus }

func NewQuoteAckStatus(val enum.QuoteAckStatus) QuoteAckStatusField {
	return QuoteAckStatusField{quickfix.FIXString(val)}
}

func (f QuoteAckStatusField) Value() enum.QuoteAckStatus { return enum.QuoteAckStatus(f.String()) }

// QuoteCancelTypeField is a enum.QuoteCancelType field.
type QuoteCancelTypeField struct{ quickfix.FIXString }

// Tag returns tag.QuoteCancelType (298).
func (f QuoteCancelTypeField) Tag() quickfix.Tag { return tag.QuoteCancelType }

func NewQuoteCancelType(val enum.QuoteCancelType) QuoteCancelTypeField {
	return QuoteCancelTypeField{quickfix.FIXString(val)}
}

func (f QuoteCancelTypeField) Value() enum.QuoteCancelType { return enum.QuoteCancelType(f.String()) }

// QuoteConditionField is a enum.QuoteCondition field.
type QuoteConditionField struct{ quickfix.FIXString }

// Tag returns tag.QuoteCondition (276).
func (f QuoteConditionField) Tag() quickfix.Tag { return tag.QuoteCondition }

func NewQuoteCondition(val enum.QuoteCondition) QuoteConditionField {
	return QuoteConditionField{quickfix.FIXString(val)}
}

func (f QuoteConditionField) Value() enum.QuoteCondition { return enum.QuoteCondition(f.String()) }

// QuoteEntryIDField is a STRING field.
type QuoteEntryIDField struct{ quickfix.FIXString }

// Tag returns tag.QuoteEntryID (299).
func (f QuoteEntryIDField) Tag() quickfix.Tag { return tag.QuoteEntryID }

// NewQuoteEntryID returns a new QuoteEntryIDField initialized with val.
func NewQuoteEntryID(val string) QuoteEntryIDField {
	return QuoteEntryIDField{quickfix.FIXString(val)}
}

func (f QuoteEntryIDField) Value() string { return f.String() }

// QuoteEntryRejectReasonField is a enum.QuoteEntryRejectReason field.
type QuoteEntryRejectReasonField struct{ quickfix.FIXString }

// Tag returns tag.QuoteEntryRejectReason (368).
func (f QuoteEntryRejectReasonField) Tag() quickfix.Tag { return tag.QuoteEntryRejectReason }

func NewQuoteEntryRejectReason(val enum.QuoteEntryRejectReason) QuoteEntryRejectReasonField {
	return QuoteEntryRejectReasonField{quickfix.FIXString(val)}
}

func (f QuoteEntryRejectReasonField) Value() enum.QuoteEntryRejectReason {
	return enum.QuoteEntryRejectReason(f.String())
}

// QuoteIDField is a STRING field.
type QuoteIDField struct{ quickfix.FIXString }

// Tag returns tag.QuoteID (117).
func (f QuoteIDField) Tag() quickfix.Tag { return tag.QuoteID }

// NewQuoteID returns a new QuoteIDField initialized with val.
func NewQuoteID(val string) QuoteIDField {
	return QuoteIDField{quickfix.FIXString(val)}
}

func (f QuoteIDField) Value() string { return f.String() }

// QuoteRejectReasonField is a enum.QuoteRejectReason field.
type QuoteRejectReasonField struct{ quickfix.FIXString }

// Tag returns tag.QuoteRejectReason (300).
func (f QuoteRejectReasonField) Tag() quickfix.Tag { return tag.QuoteRejectReason }

func NewQuoteRejectReason(val enum.QuoteRejectReason) QuoteRejectReasonField {
	return QuoteRejectReasonField{quickfix.FIXString(val)}
}

func (f QuoteRejectReasonField) Value() enum.QuoteRejectReason {
	return enum.QuoteRejectReason(f.String())
}

// QuoteReqIDField is a STRING field.
type QuoteReqIDField struct{ quickfix.FIXString }

// Tag returns tag.QuoteReqID (131).
func (f QuoteReqIDField) Tag() quickfix.Tag { return tag.QuoteReqID }

// NewQuoteReqID returns a new QuoteReqIDField initialized with val.
func NewQuoteReqID(val string) QuoteReqIDField {
	return QuoteReqIDField{quickfix.FIXString(val)}
}

func (f QuoteReqIDField) Value() string { return f.String() }

// QuoteRequestTypeField is a enum.QuoteRequestType field.
type QuoteRequestTypeField struct{ quickfix.FIXString }

// Tag returns tag.QuoteRequestType (303).
func (f QuoteRequestTypeField) Tag() quickfix.Tag { return tag.QuoteRequestType }

func NewQuoteRequestType(val enum.QuoteRequestType) QuoteRequestTypeField {
	return QuoteRequestTypeField{quickfix.FIXString(val)}
}

func (f QuoteRequestTypeField) Value() enum.QuoteRequestType {
	return enum.QuoteRequestType(f.String())
}

// QuoteResponseLevelField is a enum.QuoteResponseLevel field.
type QuoteResponseLevelField struct{ quickfix.FIXString }

// Tag returns tag.QuoteResponseLevel (301).
func (f QuoteResponseLevelField) Tag() quickfix.Tag { return tag.QuoteResponseLevel }

func NewQuoteResponseLevel(val enum.QuoteResponseLevel) QuoteResponseLevelField {
	return QuoteResponseLevelField{quickfix.FIXString(val)}
}

func (f QuoteResponseLevelField) Value() enum.QuoteResponseLevel {
	return enum.QuoteResponseLevel(f.String())
}

// QuoteSetIDField is a STRING field.
type QuoteSetIDField struct{ quickfix.FIXString }

// Tag returns tag.QuoteSetID (302).
func (f QuoteSetIDField) Tag() quickfix.Tag { return tag.QuoteSetID }

// NewQuoteSetID returns a new QuoteSetIDField initialized with val.
func NewQuoteSetID(val string) QuoteSetIDField {
	return QuoteSetIDField{quickfix.FIXString(val)}
}

func (f QuoteSetIDField) Value() string { return f.String() }

// QuoteSetValidUntilTimeField is a UTCTIMESTAMP field.
type QuoteSetValidUntilTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.QuoteSetValidUntilTime (367).
func (f QuoteSetValidUntilTimeField) Tag() quickfix.Tag { return tag.QuoteSetValidUntilTime }

// NewQuoteSetValidUntilTime returns a new QuoteSetValidUntilTimeField initialized with val.
func NewQuoteSetValidUntilTime(val time.Time) QuoteSetValidUntilTimeField {
	return NewQuoteSetValidUntilTimeWithPrecision(val, quickfix.Millis)
}

// NewQuoteSetValidUntilTimeNoMillis returns a new QuoteSetValidUntilTimeField initialized with val without millisecs.
func NewQuoteSetValidUntilTimeNoMillis(val time.Time) QuoteSetValidUntilTimeField {
	return NewQuoteSetValidUntilTimeWithPrecision(val, quickfix.Seconds)
}

// NewQuoteSetValidUntilTimeWithPrecision returns a new QuoteSetValidUntilTimeField initialized with val of specified precision.
func NewQuoteSetValidUntilTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) QuoteSetValidUntilTimeField {
	return QuoteSetValidUntilTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f QuoteSetValidUntilTimeField) Value() time.Time { return f.Time }

// RatioQtyField is a QTY field.
type RatioQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.RatioQty (319).
func (f RatioQtyField) Tag() quickfix.Tag { return tag.RatioQty }

// NewRatioQty returns a new RatioQtyField initialized with val and scale.
func NewRatioQty(val decimal.Decimal, scale int32) RatioQtyField {
	return RatioQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RatioQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// RawDataField is a DATA field.
type RawDataField struct{ quickfix.FIXString }

// Tag returns tag.RawData (96).
func (f RawDataField) Tag() quickfix.Tag { return tag.RawData }

// NewRawData returns a new RawDataField initialized with val.
func NewRawData(val string) RawDataField {
	return RawDataField{quickfix.FIXString(val)}
}

func (f RawDataField) Value() string { return f.String() }

// RawDataLengthField is a LENGTH field.
type RawDataLengthField struct{ quickfix.FIXInt }

// Tag returns tag.RawDataLength (95).
func (f RawDataLengthField) Tag() quickfix.Tag { return tag.RawDataLength }

// NewRawDataLength returns a new RawDataLengthField initialized with val.
func NewRawDataLength(val int) RawDataLengthField {
	return RawDataLengthField{quickfix.FIXInt(val)}
}

func (f RawDataLengthField) Value() int { return f.Int() }

// RefAllocIDField is a STRING field.
type RefAllocIDField struct{ quickfix.FIXString }

// Tag returns tag.RefAllocID (72).
func (f RefAllocIDField) Tag() quickfix.Tag { return tag.RefAllocID }

// NewRefAllocID returns a new RefAllocIDField initialized with val.
func NewRefAllocID(val string) RefAllocIDField {
	return RefAllocIDField{quickfix.FIXString(val)}
}

func (f RefAllocIDField) Value() string { return f.String() }

// RefMsgTypeField is a STRING field.
type RefMsgTypeField struct{ quickfix.FIXString }

// Tag returns tag.RefMsgType (372).
func (f RefMsgTypeField) Tag() quickfix.Tag { return tag.RefMsgType }

// NewRefMsgType returns a new RefMsgTypeField initialized with val.
func NewRefMsgType(val string) RefMsgTypeField {
	return RefMsgTypeField{quickfix.FIXString(val)}
}

func (f RefMsgTypeField) Value() string { return f.String() }

// RefSeqNumField is a INT field.
type RefSeqNumField struct{ quickfix.FIXInt }

// Tag returns tag.RefSeqNum (45).
func (f RefSeqNumField) Tag() quickfix.Tag { return tag.RefSeqNum }

// NewRefSeqNum returns a new RefSeqNumField initialized with val.
func NewRefSeqNum(val int) RefSeqNumField {
	return RefSeqNumField{quickfix.FIXInt(val)}
}

func (f RefSeqNumField) Value() int { return f.Int() }

// RefTagIDField is a INT field.
type RefTagIDField struct{ quickfix.FIXInt }

// Tag returns tag.RefTagID (371).
func (f RefTagIDField) Tag() quickfix.Tag { return tag.RefTagID }

// NewRefTagID returns a new RefTagIDField initialized with val.
func NewRefTagID(val int) RefTagIDField {
	return RefTagIDField{quickfix.FIXInt(val)}
}

func (f RefTagIDField) Value() int { return f.Int() }

// RelatdSymField is a STRING field.
type RelatdSymField struct{ quickfix.FIXString }

// Tag returns tag.RelatdSym (46).
func (f RelatdSymField) Tag() quickfix.Tag { return tag.RelatdSym }

// NewRelatdSym returns a new RelatdSymField initialized with val.
func NewRelatdSym(val string) RelatdSymField {
	return RelatdSymField{quickfix.FIXString(val)}
}

func (f RelatdSymField) Value() string { return f.String() }

// ReportToExchField is a BOOLEAN field.
type ReportToExchField struct{ quickfix.FIXBoolean }

// Tag returns tag.ReportToExch (113).
func (f ReportToExchField) Tag() quickfix.Tag { return tag.ReportToExch }

// NewReportToExch returns a new ReportToExchField initialized with val.
func NewReportToExch(val bool) ReportToExchField {
	return ReportToExchField{quickfix.FIXBoolean(val)}
}

func (f ReportToExchField) Value() bool { return f.Bool() }

// ResetSeqNumFlagField is a BOOLEAN field.
type ResetSeqNumFlagField struct{ quickfix.FIXBoolean }

// Tag returns tag.ResetSeqNumFlag (141).
func (f ResetSeqNumFlagField) Tag() quickfix.Tag { return tag.ResetSeqNumFlag }

// NewResetSeqNumFlag returns a new ResetSeqNumFlagField initialized with val.
func NewResetSeqNumFlag(val bool) ResetSeqNumFlagField {
	return ResetSeqNumFlagField{quickfix.FIXBoolean(val)}
}

func (f ResetSeqNumFlagField) Value() bool { return f.Bool() }

// RoutingIDField is a STRING field.
type RoutingIDField struct{ quickfix.FIXString }

// Tag returns tag.RoutingID (217).
func (f RoutingIDField) Tag() quickfix.Tag { return tag.RoutingID }

// NewRoutingID returns a new RoutingIDField initialized with val.
func NewRoutingID(val string) RoutingIDField {
	return RoutingIDField{quickfix.FIXString(val)}
}

func (f RoutingIDField) Value() string { return f.String() }

// RoutingTypeField is a enum.RoutingType field.
type RoutingTypeField struct{ quickfix.FIXString }

// Tag returns tag.RoutingType (216).
func (f RoutingTypeField) Tag() quickfix.Tag { return tag.RoutingType }

func NewRoutingType(val enum.RoutingType) RoutingTypeField {
	return RoutingTypeField{quickfix.FIXString(val)}
}

func (f RoutingTypeField) Value() enum.RoutingType { return enum.RoutingType(f.String()) }

// RptSeqField is a INT field.
type RptSeqField struct{ quickfix.FIXInt }

// Tag returns tag.RptSeq (83).
func (f RptSeqField) Tag() quickfix.Tag { return tag.RptSeq }

// NewRptSeq returns a new RptSeqField initialized with val.
func NewRptSeq(val int) RptSeqField {
	return RptSeqField{quickfix.FIXInt(val)}
}

func (f RptSeqField) Value() int { return f.Int() }

// Rule80AField is a enum.Rule80A field.
type Rule80AField struct{ quickfix.FIXString }

// Tag returns tag.Rule80A (47).
func (f Rule80AField) Tag() quickfix.Tag { return tag.Rule80A }

func NewRule80A(val enum.Rule80A) Rule80AField {
	return Rule80AField{quickfix.FIXString(val)}
}

func (f Rule80AField) Value() enum.Rule80A { return enum.Rule80A(f.String()) }

// SecondaryOrderIDField is a STRING field.
type SecondaryOrderIDField struct{ quickfix.FIXString }

// Tag returns tag.SecondaryOrderID (198).
func (f SecondaryOrderIDField) Tag() quickfix.Tag { return tag.SecondaryOrderID }

// NewSecondaryOrderID returns a new SecondaryOrderIDField initialized with val.
func NewSecondaryOrderID(val string) SecondaryOrderIDField {
	return SecondaryOrderIDField{quickfix.FIXString(val)}
}

func (f SecondaryOrderIDField) Value() string { return f.String() }

// SecureDataField is a DATA field.
type SecureDataField struct{ quickfix.FIXString }

// Tag returns tag.SecureData (91).
func (f SecureDataField) Tag() quickfix.Tag { return tag.SecureData }

// NewSecureData returns a new SecureDataField initialized with val.
func NewSecureData(val string) SecureDataField {
	return SecureDataField{quickfix.FIXString(val)}
}

func (f SecureDataField) Value() string { return f.String() }

// SecureDataLenField is a LENGTH field.
type SecureDataLenField struct{ quickfix.FIXInt }

// Tag returns tag.SecureDataLen (90).
func (f SecureDataLenField) Tag() quickfix.Tag { return tag.SecureDataLen }

// NewSecureDataLen returns a new SecureDataLenField initialized with val.
func NewSecureDataLen(val int) SecureDataLenField {
	return SecureDataLenField{quickfix.FIXInt(val)}
}

func (f SecureDataLenField) Value() int { return f.Int() }

// SecurityDescField is a STRING field.
type SecurityDescField struct{ quickfix.FIXString }

// Tag returns tag.SecurityDesc (107).
func (f SecurityDescField) Tag() quickfix.Tag { return tag.SecurityDesc }

// NewSecurityDesc returns a new SecurityDescField initialized with val.
func NewSecurityDesc(val string) SecurityDescField {
	return SecurityDescField{quickfix.FIXString(val)}
}

func (f SecurityDescField) Value() string { return f.String() }

// SecurityExchangeField is a EXCHANGE field.
type SecurityExchangeField struct{ quickfix.FIXString }

// Tag returns tag.SecurityExchange (207).
func (f SecurityExchangeField) Tag() quickfix.Tag { return tag.SecurityExchange }

// NewSecurityExchange returns a new SecurityExchangeField initialized with val.
func NewSecurityExchange(val string) SecurityExchangeField {
	return SecurityExchangeField{quickfix.FIXString(val)}
}

func (f SecurityExchangeField) Value() string { return f.String() }

// SecurityIDField is a STRING field.
type SecurityIDField struct{ quickfix.FIXString }

// Tag returns tag.SecurityID (48).
func (f SecurityIDField) Tag() quickfix.Tag { return tag.SecurityID }

// NewSecurityID returns a new SecurityIDField initialized with val.
func NewSecurityID(val string) SecurityIDField {
	return SecurityIDField{quickfix.FIXString(val)}
}

func (f SecurityIDField) Value() string { return f.String() }

// SecurityReqIDField is a STRING field.
type SecurityReqIDField struct{ quickfix.FIXString }

// Tag returns tag.SecurityReqID (320).
func (f SecurityReqIDField) Tag() quickfix.Tag { return tag.SecurityReqID }

// NewSecurityReqID returns a new SecurityReqIDField initialized with val.
func NewSecurityReqID(val string) SecurityReqIDField {
	return SecurityReqIDField{quickfix.FIXString(val)}
}

func (f SecurityReqIDField) Value() string { return f.String() }

// SecurityRequestTypeField is a enum.SecurityRequestType field.
type SecurityRequestTypeField struct{ quickfix.FIXString }

// Tag returns tag.SecurityRequestType (321).
func (f SecurityRequestTypeField) Tag() quickfix.Tag { return tag.SecurityRequestType }

func NewSecurityRequestType(val enum.SecurityRequestType) SecurityRequestTypeField {
	return SecurityRequestTypeField{quickfix.FIXString(val)}
}

func (f SecurityRequestTypeField) Value() enum.SecurityRequestType {
	return enum.SecurityRequestType(f.String())
}

// SecurityResponseIDField is a STRING field.
type SecurityResponseIDField struct{ quickfix.FIXString }

// Tag returns tag.SecurityResponseID (322).
func (f SecurityResponseIDField) Tag() quickfix.Tag { return tag.SecurityResponseID }

// NewSecurityResponseID returns a new SecurityResponseIDField initialized with val.
func NewSecurityResponseID(val string) SecurityResponseIDField {
	return SecurityResponseIDField{quickfix.FIXString(val)}
}

func (f SecurityResponseIDField) Value() string { return f.String() }

// SecurityResponseTypeField is a enum.SecurityResponseType field.
type SecurityResponseTypeField struct{ quickfix.FIXString }

// Tag returns tag.SecurityResponseType (323).
func (f SecurityResponseTypeField) Tag() quickfix.Tag { return tag.SecurityResponseType }

func NewSecurityResponseType(val enum.SecurityResponseType) SecurityResponseTypeField {
	return SecurityResponseTypeField{quickfix.FIXString(val)}
}

func (f SecurityResponseTypeField) Value() enum.SecurityResponseType {
	return enum.SecurityResponseType(f.String())
}

// SecuritySettlAgentAcctNameField is a STRING field.
type SecuritySettlAgentAcctNameField struct{ quickfix.FIXString }

// Tag returns tag.SecuritySettlAgentAcctName (179).
func (f SecuritySettlAgentAcctNameField) Tag() quickfix.Tag { return tag.SecuritySettlAgentAcctName }

// NewSecuritySettlAgentAcctName returns a new SecuritySettlAgentAcctNameField initialized with val.
func NewSecuritySettlAgentAcctName(val string) SecuritySettlAgentAcctNameField {
	return SecuritySettlAgentAcctNameField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentAcctNameField) Value() string { return f.String() }

// SecuritySettlAgentAcctNumField is a STRING field.
type SecuritySettlAgentAcctNumField struct{ quickfix.FIXString }

// Tag returns tag.SecuritySettlAgentAcctNum (178).
func (f SecuritySettlAgentAcctNumField) Tag() quickfix.Tag { return tag.SecuritySettlAgentAcctNum }

// NewSecuritySettlAgentAcctNum returns a new SecuritySettlAgentAcctNumField initialized with val.
func NewSecuritySettlAgentAcctNum(val string) SecuritySettlAgentAcctNumField {
	return SecuritySettlAgentAcctNumField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentAcctNumField) Value() string { return f.String() }

// SecuritySettlAgentCodeField is a STRING field.
type SecuritySettlAgentCodeField struct{ quickfix.FIXString }

// Tag returns tag.SecuritySettlAgentCode (177).
func (f SecuritySettlAgentCodeField) Tag() quickfix.Tag { return tag.SecuritySettlAgentCode }

// NewSecuritySettlAgentCode returns a new SecuritySettlAgentCodeField initialized with val.
func NewSecuritySettlAgentCode(val string) SecuritySettlAgentCodeField {
	return SecuritySettlAgentCodeField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentCodeField) Value() string { return f.String() }

// SecuritySettlAgentContactNameField is a STRING field.
type SecuritySettlAgentContactNameField struct{ quickfix.FIXString }

// Tag returns tag.SecuritySettlAgentContactName (180).
func (f SecuritySettlAgentContactNameField) Tag() quickfix.Tag {
	return tag.SecuritySettlAgentContactName
}

// NewSecuritySettlAgentContactName returns a new SecuritySettlAgentContactNameField initialized with val.
func NewSecuritySettlAgentContactName(val string) SecuritySettlAgentContactNameField {
	return SecuritySettlAgentContactNameField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentContactNameField) Value() string { return f.String() }

// SecuritySettlAgentContactPhoneField is a STRING field.
type SecuritySettlAgentContactPhoneField struct{ quickfix.FIXString }

// Tag returns tag.SecuritySettlAgentContactPhone (181).
func (f SecuritySettlAgentContactPhoneField) Tag() quickfix.Tag {
	return tag.SecuritySettlAgentContactPhone
}

// NewSecuritySettlAgentContactPhone returns a new SecuritySettlAgentContactPhoneField initialized with val.
func NewSecuritySettlAgentContactPhone(val string) SecuritySettlAgentContactPhoneField {
	return SecuritySettlAgentContactPhoneField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentContactPhoneField) Value() string { return f.String() }

// SecuritySettlAgentNameField is a STRING field.
type SecuritySettlAgentNameField struct{ quickfix.FIXString }

// Tag returns tag.SecuritySettlAgentName (176).
func (f SecuritySettlAgentNameField) Tag() quickfix.Tag { return tag.SecuritySettlAgentName }

// NewSecuritySettlAgentName returns a new SecuritySettlAgentNameField initialized with val.
func NewSecuritySettlAgentName(val string) SecuritySettlAgentNameField {
	return SecuritySettlAgentNameField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentNameField) Value() string { return f.String() }

// SecurityStatusReqIDField is a STRING field.
type SecurityStatusReqIDField struct{ quickfix.FIXString }

// Tag returns tag.SecurityStatusReqID (324).
func (f SecurityStatusReqIDField) Tag() quickfix.Tag { return tag.SecurityStatusReqID }

// NewSecurityStatusReqID returns a new SecurityStatusReqIDField initialized with val.
func NewSecurityStatusReqID(val string) SecurityStatusReqIDField {
	return SecurityStatusReqIDField{quickfix.FIXString(val)}
}

func (f SecurityStatusReqIDField) Value() string { return f.String() }

// SecurityTradingStatusField is a enum.SecurityTradingStatus field.
type SecurityTradingStatusField struct{ quickfix.FIXString }

// Tag returns tag.SecurityTradingStatus (326).
func (f SecurityTradingStatusField) Tag() quickfix.Tag { return tag.SecurityTradingStatus }

func NewSecurityTradingStatus(val enum.SecurityTradingStatus) SecurityTradingStatusField {
	return SecurityTradingStatusField{quickfix.FIXString(val)}
}

func (f SecurityTradingStatusField) Value() enum.SecurityTradingStatus {
	return enum.SecurityTradingStatus(f.String())
}

// SecurityTypeField is a enum.SecurityType field.
type SecurityTypeField struct{ quickfix.FIXString }

// Tag returns tag.SecurityType (167).
func (f SecurityTypeField) Tag() quickfix.Tag { return tag.SecurityType }

func NewSecurityType(val enum.SecurityType) SecurityTypeField {
	return SecurityTypeField{quickfix.FIXString(val)}
}

func (f SecurityTypeField) Value() enum.SecurityType { return enum.SecurityType(f.String()) }

// SelfTradePreventionField is a enum.SelfTradePrevention field.
type SelfTradePreventionField struct{ quickfix.FIXString }

// Tag returns tag.SelfTradePrevention (7928).
func (f SelfTradePreventionField) Tag() quickfix.Tag { return tag.SelfTradePrevention }

func NewSelfTradePrevention(val enum.SelfTradePrevention) SelfTradePreventionField {
	return SelfTradePreventionField{quickfix.FIXString(val)}
}

func (f SelfTradePreventionField) Value() enum.SelfTradePrevention {
	return enum.SelfTradePrevention(f.String())
}

// SellVolumeField is a QTY field.
type SellVolumeField struct{ quickfix.FIXDecimal }

// Tag returns tag.SellVolume (331).
func (f SellVolumeField) Tag() quickfix.Tag { return tag.SellVolume }

// NewSellVolume returns a new SellVolumeField initialized with val and scale.
func NewSellVolume(val decimal.Decimal, scale int32) SellVolumeField {
	return SellVolumeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SellVolumeField) Value() (val decimal.Decimal) { return f.Decimal }

// SellerDaysField is a INT field.
type SellerDaysField struct{ quickfix.FIXInt }

// Tag returns tag.SellerDays (287).
func (f SellerDaysField) Tag() quickfix.Tag { return tag.SellerDays }

// NewSellerDays returns a new SellerDaysField initialized with val.
func NewSellerDays(val int) SellerDaysField {
	return SellerDaysField{quickfix.FIXInt(val)}
}

func (f SellerDaysField) Value() int { return f.Int() }

// SenderCompIDField is a STRING field.
type SenderCompIDField struct{ quickfix.FIXString }

// Tag returns tag.SenderCompID (49).
func (f SenderCompIDField) Tag() quickfix.Tag { return tag.SenderCompID }

// NewSenderCompID returns a new SenderCompIDField initialized with val.
func NewSenderCompID(val string) SenderCompIDField {
	return SenderCompIDField{quickfix.FIXString(val)}
}

func (f SenderCompIDField) Value() string { return f.String() }

// SenderLocationIDField is a STRING field.
type SenderLocationIDField struct{ quickfix.FIXString }

// Tag returns tag.SenderLocationID (142).
func (f SenderLocationIDField) Tag() quickfix.Tag { return tag.SenderLocationID }

// NewSenderLocationID returns a new SenderLocationIDField initialized with val.
func NewSenderLocationID(val string) SenderLocationIDField {
	return SenderLocationIDField{quickfix.FIXString(val)}
}

func (f SenderLocationIDField) Value() string { return f.String() }

// SenderSubIDField is a STRING field.
type SenderSubIDField struct{ quickfix.FIXString }

// Tag returns tag.SenderSubID (50).
func (f SenderSubIDField) Tag() quickfix.Tag { return tag.SenderSubID }

// NewSenderSubID returns a new SenderSubIDField initialized with val.
func NewSenderSubID(val string) SenderSubIDField {
	return SenderSubIDField{quickfix.FIXString(val)}
}

func (f SenderSubIDField) Value() string { return f.String() }

// SendingDateField is a LOCALMKTDATE field.
type SendingDateField struct{ quickfix.FIXString }

// Tag returns tag.SendingDate (51).
func (f SendingDateField) Tag() quickfix.Tag { return tag.SendingDate }

// NewSendingDate returns a new SendingDateField initialized with val.
func NewSendingDate(val string) SendingDateField {
	return SendingDateField{quickfix.FIXString(val)}
}

func (f SendingDateField) Value() string { return f.String() }

// SendingTimeField is a UTCTIMESTAMP field.
type SendingTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.SendingTime (52).
func (f SendingTimeField) Tag() quickfix.Tag { return tag.SendingTime }

// NewSendingTime returns a new SendingTimeField initialized with val.
func NewSendingTime(val time.Time) SendingTimeField {
	return NewSendingTimeWithPrecision(val, quickfix.Millis)
}

// NewSendingTimeNoMillis returns a new SendingTimeField initialized with val without millisecs.
func NewSendingTimeNoMillis(val time.Time) SendingTimeField {
	return NewSendingTimeWithPrecision(val, quickfix.Seconds)
}

// NewSendingTimeWithPrecision returns a new SendingTimeField initialized with val of specified precision.
func NewSendingTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) SendingTimeField {
	return SendingTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f SendingTimeField) Value() time.Time { return f.Time }

// SessionRejectReasonField is a enum.SessionRejectReason field.
type SessionRejectReasonField struct{ quickfix.FIXString }

// Tag returns tag.SessionRejectReason (373).
func (f SessionRejectReasonField) Tag() quickfix.Tag { return tag.SessionRejectReason }

func NewSessionRejectReason(val enum.SessionRejectReason) SessionRejectReasonField {
	return SessionRejectReasonField{quickfix.FIXString(val)}
}

func (f SessionRejectReasonField) Value() enum.SessionRejectReason {
	return enum.SessionRejectReason(f.String())
}

// SettlBrkrCodeField is a STRING field.
type SettlBrkrCodeField struct{ quickfix.FIXString }

// Tag returns tag.SettlBrkrCode (174).
func (f SettlBrkrCodeField) Tag() quickfix.Tag { return tag.SettlBrkrCode }

// NewSettlBrkrCode returns a new SettlBrkrCodeField initialized with val.
func NewSettlBrkrCode(val string) SettlBrkrCodeField {
	return SettlBrkrCodeField{quickfix.FIXString(val)}
}

func (f SettlBrkrCodeField) Value() string { return f.String() }

// SettlCurrAmtField is a AMT field.
type SettlCurrAmtField struct{ quickfix.FIXDecimal }

// Tag returns tag.SettlCurrAmt (119).
func (f SettlCurrAmtField) Tag() quickfix.Tag { return tag.SettlCurrAmt }

// NewSettlCurrAmt returns a new SettlCurrAmtField initialized with val and scale.
func NewSettlCurrAmt(val decimal.Decimal, scale int32) SettlCurrAmtField {
	return SettlCurrAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SettlCurrAmtField) Value() (val decimal.Decimal) { return f.Decimal }

// SettlCurrFxRateField is a FLOAT field.
type SettlCurrFxRateField struct{ quickfix.FIXDecimal }

// Tag returns tag.SettlCurrFxRate (155).
func (f SettlCurrFxRateField) Tag() quickfix.Tag { return tag.SettlCurrFxRate }

// NewSettlCurrFxRate returns a new SettlCurrFxRateField initialized with val and scale.
func NewSettlCurrFxRate(val decimal.Decimal, scale int32) SettlCurrFxRateField {
	return SettlCurrFxRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SettlCurrFxRateField) Value() (val decimal.Decimal) { return f.Decimal }

// SettlCurrFxRateCalcField is a enum.SettlCurrFxRateCalc field.
type SettlCurrFxRateCalcField struct{ quickfix.FIXString }

// Tag returns tag.SettlCurrFxRateCalc (156).
func (f SettlCurrFxRateCalcField) Tag() quickfix.Tag { return tag.SettlCurrFxRateCalc }

func NewSettlCurrFxRateCalc(val enum.SettlCurrFxRateCalc) SettlCurrFxRateCalcField {
	return SettlCurrFxRateCalcField{quickfix.FIXString(val)}
}

func (f SettlCurrFxRateCalcField) Value() enum.SettlCurrFxRateCalc {
	return enum.SettlCurrFxRateCalc(f.String())
}

// SettlCurrencyField is a CURRENCY field.
type SettlCurrencyField struct{ quickfix.FIXString }

// Tag returns tag.SettlCurrency (120).
func (f SettlCurrencyField) Tag() quickfix.Tag { return tag.SettlCurrency }

// NewSettlCurrency returns a new SettlCurrencyField initialized with val.
func NewSettlCurrency(val string) SettlCurrencyField {
	return SettlCurrencyField{quickfix.FIXString(val)}
}

func (f SettlCurrencyField) Value() string { return f.String() }

// SettlDeliveryTypeField is a INT field.
type SettlDeliveryTypeField struct{ quickfix.FIXInt }

// Tag returns tag.SettlDeliveryType (172).
func (f SettlDeliveryTypeField) Tag() quickfix.Tag { return tag.SettlDeliveryType }

// NewSettlDeliveryType returns a new SettlDeliveryTypeField initialized with val.
func NewSettlDeliveryType(val int) SettlDeliveryTypeField {
	return SettlDeliveryTypeField{quickfix.FIXInt(val)}
}

func (f SettlDeliveryTypeField) Value() int { return f.Int() }

// SettlDepositoryCodeField is a STRING field.
type SettlDepositoryCodeField struct{ quickfix.FIXString }

// Tag returns tag.SettlDepositoryCode (173).
func (f SettlDepositoryCodeField) Tag() quickfix.Tag { return tag.SettlDepositoryCode }

// NewSettlDepositoryCode returns a new SettlDepositoryCodeField initialized with val.
func NewSettlDepositoryCode(val string) SettlDepositoryCodeField {
	return SettlDepositoryCodeField{quickfix.FIXString(val)}
}

func (f SettlDepositoryCodeField) Value() string { return f.String() }

// SettlInstCodeField is a STRING field.
type SettlInstCodeField struct{ quickfix.FIXString }

// Tag returns tag.SettlInstCode (175).
func (f SettlInstCodeField) Tag() quickfix.Tag { return tag.SettlInstCode }

// NewSettlInstCode returns a new SettlInstCodeField initialized with val.
func NewSettlInstCode(val string) SettlInstCodeField {
	return SettlInstCodeField{quickfix.FIXString(val)}
}

func (f SettlInstCodeField) Value() string { return f.String() }

// SettlInstIDField is a STRING field.
type SettlInstIDField struct{ quickfix.FIXString }

// Tag returns tag.SettlInstID (162).
func (f SettlInstIDField) Tag() quickfix.Tag { return tag.SettlInstID }

// NewSettlInstID returns a new SettlInstIDField initialized with val.
func NewSettlInstID(val string) SettlInstIDField {
	return SettlInstIDField{quickfix.FIXString(val)}
}

func (f SettlInstIDField) Value() string { return f.String() }

// SettlInstModeField is a enum.SettlInstMode field.
type SettlInstModeField struct{ quickfix.FIXString }

// Tag returns tag.SettlInstMode (160).
func (f SettlInstModeField) Tag() quickfix.Tag { return tag.SettlInstMode }

func NewSettlInstMode(val enum.SettlInstMode) SettlInstModeField {
	return SettlInstModeField{quickfix.FIXString(val)}
}

func (f SettlInstModeField) Value() enum.SettlInstMode { return enum.SettlInstMode(f.String()) }

// SettlInstRefIDField is a STRING field.
type SettlInstRefIDField struct{ quickfix.FIXString }

// Tag returns tag.SettlInstRefID (214).
func (f SettlInstRefIDField) Tag() quickfix.Tag { return tag.SettlInstRefID }

// NewSettlInstRefID returns a new SettlInstRefIDField initialized with val.
func NewSettlInstRefID(val string) SettlInstRefIDField {
	return SettlInstRefIDField{quickfix.FIXString(val)}
}

func (f SettlInstRefIDField) Value() string { return f.String() }

// SettlInstSourceField is a enum.SettlInstSource field.
type SettlInstSourceField struct{ quickfix.FIXString }

// Tag returns tag.SettlInstSource (165).
func (f SettlInstSourceField) Tag() quickfix.Tag { return tag.SettlInstSource }

func NewSettlInstSource(val enum.SettlInstSource) SettlInstSourceField {
	return SettlInstSourceField{quickfix.FIXString(val)}
}

func (f SettlInstSourceField) Value() enum.SettlInstSource { return enum.SettlInstSource(f.String()) }

// SettlInstTransTypeField is a enum.SettlInstTransType field.
type SettlInstTransTypeField struct{ quickfix.FIXString }

// Tag returns tag.SettlInstTransType (163).
func (f SettlInstTransTypeField) Tag() quickfix.Tag { return tag.SettlInstTransType }

func NewSettlInstTransType(val enum.SettlInstTransType) SettlInstTransTypeField {
	return SettlInstTransTypeField{quickfix.FIXString(val)}
}

func (f SettlInstTransTypeField) Value() enum.SettlInstTransType {
	return enum.SettlInstTransType(f.String())
}

// SettlLocationField is a enum.SettlLocation field.
type SettlLocationField struct{ quickfix.FIXString }

// Tag returns tag.SettlLocation (166).
func (f SettlLocationField) Tag() quickfix.Tag { return tag.SettlLocation }

func NewSettlLocation(val enum.SettlLocation) SettlLocationField {
	return SettlLocationField{quickfix.FIXString(val)}
}

func (f SettlLocationField) Value() enum.SettlLocation { return enum.SettlLocation(f.String()) }

// SettlmntTypField is a enum.SettlmntTyp field.
type SettlmntTypField struct{ quickfix.FIXString }

// Tag returns tag.SettlmntTyp (63).
func (f SettlmntTypField) Tag() quickfix.Tag { return tag.SettlmntTyp }

func NewSettlmntTyp(val enum.SettlmntTyp) SettlmntTypField {
	return SettlmntTypField{quickfix.FIXString(val)}
}

func (f SettlmntTypField) Value() enum.SettlmntTyp { return enum.SettlmntTyp(f.String()) }

// SharesField is a QTY field.
type SharesField struct{ quickfix.FIXDecimal }

// Tag returns tag.Shares (53).
func (f SharesField) Tag() quickfix.Tag { return tag.Shares }

// NewShares returns a new SharesField initialized with val and scale.
func NewShares(val decimal.Decimal, scale int32) SharesField {
	return SharesField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SharesField) Value() (val decimal.Decimal) { return f.Decimal }

// SideField is a enum.Side field.
type SideField struct{ quickfix.FIXString }

// Tag returns tag.Side (54).
func (f SideField) Tag() quickfix.Tag { return tag.Side }

func NewSide(val enum.Side) SideField {
	return SideField{quickfix.FIXString(val)}
}

func (f SideField) Value() enum.Side { return enum.Side(f.String()) }

// SideValue1Field is a AMT field.
type SideValue1Field struct{ quickfix.FIXDecimal }

// Tag returns tag.SideValue1 (396).
func (f SideValue1Field) Tag() quickfix.Tag { return tag.SideValue1 }

// NewSideValue1 returns a new SideValue1Field initialized with val and scale.
func NewSideValue1(val decimal.Decimal, scale int32) SideValue1Field {
	return SideValue1Field{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SideValue1Field) Value() (val decimal.Decimal) { return f.Decimal }

// SideValue2Field is a AMT field.
type SideValue2Field struct{ quickfix.FIXDecimal }

// Tag returns tag.SideValue2 (397).
func (f SideValue2Field) Tag() quickfix.Tag { return tag.SideValue2 }

// NewSideValue2 returns a new SideValue2Field initialized with val and scale.
func NewSideValue2(val decimal.Decimal, scale int32) SideValue2Field {
	return SideValue2Field{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SideValue2Field) Value() (val decimal.Decimal) { return f.Decimal }

// SideValueIndField is a INT field.
type SideValueIndField struct{ quickfix.FIXInt }

// Tag returns tag.SideValueInd (401).
func (f SideValueIndField) Tag() quickfix.Tag { return tag.SideValueInd }

// NewSideValueInd returns a new SideValueIndField initialized with val.
func NewSideValueInd(val int) SideValueIndField {
	return SideValueIndField{quickfix.FIXInt(val)}
}

func (f SideValueIndField) Value() int { return f.Int() }

// SignatureField is a DATA field.
type SignatureField struct{ quickfix.FIXString }

// Tag returns tag.Signature (89).
func (f SignatureField) Tag() quickfix.Tag { return tag.Signature }

// NewSignature returns a new SignatureField initialized with val.
func NewSignature(val string) SignatureField {
	return SignatureField{quickfix.FIXString(val)}
}

func (f SignatureField) Value() string { return f.String() }

// SignatureLengthField is a LENGTH field.
type SignatureLengthField struct{ quickfix.FIXInt }

// Tag returns tag.SignatureLength (93).
func (f SignatureLengthField) Tag() quickfix.Tag { return tag.SignatureLength }

// NewSignatureLength returns a new SignatureLengthField initialized with val.
func NewSignatureLength(val int) SignatureLengthField {
	return SignatureLengthField{quickfix.FIXInt(val)}
}

func (f SignatureLengthField) Value() int { return f.Int() }

// SolicitedFlagField is a BOOLEAN field.
type SolicitedFlagField struct{ quickfix.FIXBoolean }

// Tag returns tag.SolicitedFlag (377).
func (f SolicitedFlagField) Tag() quickfix.Tag { return tag.SolicitedFlag }

// NewSolicitedFlag returns a new SolicitedFlagField initialized with val.
func NewSolicitedFlag(val bool) SolicitedFlagField {
	return SolicitedFlagField{quickfix.FIXBoolean(val)}
}

func (f SolicitedFlagField) Value() bool { return f.Bool() }

// SpreadToBenchmarkField is a PRICEOFFSET field.
type SpreadToBenchmarkField struct{ quickfix.FIXDecimal }

// Tag returns tag.SpreadToBenchmark (218).
func (f SpreadToBenchmarkField) Tag() quickfix.Tag { return tag.SpreadToBenchmark }

// NewSpreadToBenchmark returns a new SpreadToBenchmarkField initialized with val and scale.
func NewSpreadToBenchmark(val decimal.Decimal, scale int32) SpreadToBenchmarkField {
	return SpreadToBenchmarkField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SpreadToBenchmarkField) Value() (val decimal.Decimal) { return f.Decimal }

// StandInstDbIDField is a STRING field.
type StandInstDbIDField struct{ quickfix.FIXString }

// Tag returns tag.StandInstDbID (171).
func (f StandInstDbIDField) Tag() quickfix.Tag { return tag.StandInstDbID }

// NewStandInstDbID returns a new StandInstDbIDField initialized with val.
func NewStandInstDbID(val string) StandInstDbIDField {
	return StandInstDbIDField{quickfix.FIXString(val)}
}

func (f StandInstDbIDField) Value() string { return f.String() }

// StandInstDbNameField is a STRING field.
type StandInstDbNameField struct{ quickfix.FIXString }

// Tag returns tag.StandInstDbName (170).
func (f StandInstDbNameField) Tag() quickfix.Tag { return tag.StandInstDbName }

// NewStandInstDbName returns a new StandInstDbNameField initialized with val.
func NewStandInstDbName(val string) StandInstDbNameField {
	return StandInstDbNameField{quickfix.FIXString(val)}
}

func (f StandInstDbNameField) Value() string { return f.String() }

// StandInstDbTypeField is a enum.StandInstDbType field.
type StandInstDbTypeField struct{ quickfix.FIXString }

// Tag returns tag.StandInstDbType (169).
func (f StandInstDbTypeField) Tag() quickfix.Tag { return tag.StandInstDbType }

func NewStandInstDbType(val enum.StandInstDbType) StandInstDbTypeField {
	return StandInstDbTypeField{quickfix.FIXString(val)}
}

func (f StandInstDbTypeField) Value() enum.StandInstDbType { return enum.StandInstDbType(f.String()) }

// StopPxField is a PRICE field.
type StopPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.StopPx (99).
func (f StopPxField) Tag() quickfix.Tag { return tag.StopPx }

// NewStopPx returns a new StopPxField initialized with val and scale.
func NewStopPx(val decimal.Decimal, scale int32) StopPxField {
	return StopPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StopPxField) Value() (val decimal.Decimal) { return f.Decimal }

// StrikePriceField is a PRICE field.
type StrikePriceField struct{ quickfix.FIXDecimal }

// Tag returns tag.StrikePrice (202).
func (f StrikePriceField) Tag() quickfix.Tag { return tag.StrikePrice }

// NewStrikePrice returns a new StrikePriceField initialized with val and scale.
func NewStrikePrice(val decimal.Decimal, scale int32) StrikePriceField {
	return StrikePriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StrikePriceField) Value() (val decimal.Decimal) { return f.Decimal }

// StrikeTimeField is a UTCTIMESTAMP field.
type StrikeTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.StrikeTime (443).
func (f StrikeTimeField) Tag() quickfix.Tag { return tag.StrikeTime }

// NewStrikeTime returns a new StrikeTimeField initialized with val.
func NewStrikeTime(val time.Time) StrikeTimeField {
	return NewStrikeTimeWithPrecision(val, quickfix.Millis)
}

// NewStrikeTimeNoMillis returns a new StrikeTimeField initialized with val without millisecs.
func NewStrikeTimeNoMillis(val time.Time) StrikeTimeField {
	return NewStrikeTimeWithPrecision(val, quickfix.Seconds)
}

// NewStrikeTimeWithPrecision returns a new StrikeTimeField initialized with val of specified precision.
func NewStrikeTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) StrikeTimeField {
	return StrikeTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f StrikeTimeField) Value() time.Time { return f.Time }

// SubjectField is a STRING field.
type SubjectField struct{ quickfix.FIXString }

// Tag returns tag.Subject (147).
func (f SubjectField) Tag() quickfix.Tag { return tag.Subject }

// NewSubject returns a new SubjectField initialized with val.
func NewSubject(val string) SubjectField {
	return SubjectField{quickfix.FIXString(val)}
}

func (f SubjectField) Value() string { return f.String() }

// SubscriptionRequestTypeField is a enum.SubscriptionRequestType field.
type SubscriptionRequestTypeField struct{ quickfix.FIXString }

// Tag returns tag.SubscriptionRequestType (263).
func (f SubscriptionRequestTypeField) Tag() quickfix.Tag { return tag.SubscriptionRequestType }

func NewSubscriptionRequestType(val enum.SubscriptionRequestType) SubscriptionRequestTypeField {
	return SubscriptionRequestTypeField{quickfix.FIXString(val)}
}

func (f SubscriptionRequestTypeField) Value() enum.SubscriptionRequestType {
	return enum.SubscriptionRequestType(f.String())
}

// SymbolField is a STRING field.
type SymbolField struct{ quickfix.FIXString }

// Tag returns tag.Symbol (55).
func (f SymbolField) Tag() quickfix.Tag { return tag.Symbol }

// NewSymbol returns a new SymbolField initialized with val.
func NewSymbol(val string) SymbolField {
	return SymbolField{quickfix.FIXString(val)}
}

func (f SymbolField) Value() string { return f.String() }

// SymbolSfxField is a STRING field.
type SymbolSfxField struct{ quickfix.FIXString }

// Tag returns tag.SymbolSfx (65).
func (f SymbolSfxField) Tag() quickfix.Tag { return tag.SymbolSfx }

// NewSymbolSfx returns a new SymbolSfxField initialized with val.
func NewSymbolSfx(val string) SymbolSfxField {
	return SymbolSfxField{quickfix.FIXString(val)}
}

func (f SymbolSfxField) Value() string { return f.String() }

// TargetCompIDField is a STRING field.
type TargetCompIDField struct{ quickfix.FIXString }

// Tag returns tag.TargetCompID (56).
func (f TargetCompIDField) Tag() quickfix.Tag { return tag.TargetCompID }

// NewTargetCompID returns a new TargetCompIDField initialized with val.
func NewTargetCompID(val string) TargetCompIDField {
	return TargetCompIDField{quickfix.FIXString(val)}
}

func (f TargetCompIDField) Value() string { return f.String() }

// TargetLocationIDField is a STRING field.
type TargetLocationIDField struct{ quickfix.FIXString }

// Tag returns tag.TargetLocationID (143).
func (f TargetLocationIDField) Tag() quickfix.Tag { return tag.TargetLocationID }

// NewTargetLocationID returns a new TargetLocationIDField initialized with val.
func NewTargetLocationID(val string) TargetLocationIDField {
	return TargetLocationIDField{quickfix.FIXString(val)}
}

func (f TargetLocationIDField) Value() string { return f.String() }

// TargetSubIDField is a STRING field.
type TargetSubIDField struct{ quickfix.FIXString }

// Tag returns tag.TargetSubID (57).
func (f TargetSubIDField) Tag() quickfix.Tag { return tag.TargetSubID }

// NewTargetSubID returns a new TargetSubIDField initialized with val.
func NewTargetSubID(val string) TargetSubIDField {
	return TargetSubIDField{quickfix.FIXString(val)}
}

func (f TargetSubIDField) Value() string { return f.String() }

// TestReqIDField is a STRING field.
type TestReqIDField struct{ quickfix.FIXString }

// Tag returns tag.TestReqID (112).
func (f TestReqIDField) Tag() quickfix.Tag { return tag.TestReqID }

// NewTestReqID returns a new TestReqIDField initialized with val.
func NewTestReqID(val string) TestReqIDField {
	return TestReqIDField{quickfix.FIXString(val)}
}

func (f TestReqIDField) Value() string { return f.String() }

// TextField is a STRING field.
type TextField struct{ quickfix.FIXString }

// Tag returns tag.Text (58).
func (f TextField) Tag() quickfix.Tag { return tag.Text }

// NewText returns a new TextField initialized with val.
func NewText(val string) TextField {
	return TextField{quickfix.FIXString(val)}
}

func (f TextField) Value() string { return f.String() }

// TickDirectionField is a enum.TickDirection field.
type TickDirectionField struct{ quickfix.FIXString }

// Tag returns tag.TickDirection (274).
func (f TickDirectionField) Tag() quickfix.Tag { return tag.TickDirection }

func NewTickDirection(val enum.TickDirection) TickDirectionField {
	return TickDirectionField{quickfix.FIXString(val)}
}

func (f TickDirectionField) Value() enum.TickDirection { return enum.TickDirection(f.String()) }

// TimeInForceField is a enum.TimeInForce field.
type TimeInForceField struct{ quickfix.FIXString }

// Tag returns tag.TimeInForce (59).
func (f TimeInForceField) Tag() quickfix.Tag { return tag.TimeInForce }

func NewTimeInForce(val enum.TimeInForce) TimeInForceField {
	return TimeInForceField{quickfix.FIXString(val)}
}

func (f TimeInForceField) Value() enum.TimeInForce { return enum.TimeInForce(f.String()) }

// TotNoOrdersField is a INT field.
type TotNoOrdersField struct{ quickfix.FIXInt }

// Tag returns tag.TotNoOrders (68).
func (f TotNoOrdersField) Tag() quickfix.Tag { return tag.TotNoOrders }

// NewTotNoOrders returns a new TotNoOrdersField initialized with val.
func NewTotNoOrders(val int) TotNoOrdersField {
	return TotNoOrdersField{quickfix.FIXInt(val)}
}

func (f TotNoOrdersField) Value() int { return f.Int() }

// TotNoStrikesField is a INT field.
type TotNoStrikesField struct{ quickfix.FIXInt }

// Tag returns tag.TotNoStrikes (422).
func (f TotNoStrikesField) Tag() quickfix.Tag { return tag.TotNoStrikes }

// NewTotNoStrikes returns a new TotNoStrikesField initialized with val.
func NewTotNoStrikes(val int) TotNoStrikesField {
	return TotNoStrikesField{quickfix.FIXInt(val)}
}

func (f TotNoStrikesField) Value() int { return f.Int() }

// TotQuoteEntriesField is a INT field.
type TotQuoteEntriesField struct{ quickfix.FIXInt }

// Tag returns tag.TotQuoteEntries (304).
func (f TotQuoteEntriesField) Tag() quickfix.Tag { return tag.TotQuoteEntries }

// NewTotQuoteEntries returns a new TotQuoteEntriesField initialized with val.
func NewTotQuoteEntries(val int) TotQuoteEntriesField {
	return TotQuoteEntriesField{quickfix.FIXInt(val)}
}

func (f TotQuoteEntriesField) Value() int { return f.Int() }

// TotalNumSecuritiesField is a INT field.
type TotalNumSecuritiesField struct{ quickfix.FIXInt }

// Tag returns tag.TotalNumSecurities (393).
func (f TotalNumSecuritiesField) Tag() quickfix.Tag { return tag.TotalNumSecurities }

// NewTotalNumSecurities returns a new TotalNumSecuritiesField initialized with val.
func NewTotalNumSecurities(val int) TotalNumSecuritiesField {
	return TotalNumSecuritiesField{quickfix.FIXInt(val)}
}

func (f TotalNumSecuritiesField) Value() int { return f.Int() }

// TotalVolumeTradedField is a QTY field.
type TotalVolumeTradedField struct{ quickfix.FIXDecimal }

// Tag returns tag.TotalVolumeTraded (387).
func (f TotalVolumeTradedField) Tag() quickfix.Tag { return tag.TotalVolumeTraded }

// NewTotalVolumeTraded returns a new TotalVolumeTradedField initialized with val and scale.
func NewTotalVolumeTraded(val decimal.Decimal, scale int32) TotalVolumeTradedField {
	return TotalVolumeTradedField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TotalVolumeTradedField) Value() (val decimal.Decimal) { return f.Decimal }

// TradSesCloseTimeField is a UTCTIMESTAMP field.
type TradSesCloseTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.TradSesCloseTime (344).
func (f TradSesCloseTimeField) Tag() quickfix.Tag { return tag.TradSesCloseTime }

// NewTradSesCloseTime returns a new TradSesCloseTimeField initialized with val.
func NewTradSesCloseTime(val time.Time) TradSesCloseTimeField {
	return NewTradSesCloseTimeWithPrecision(val, quickfix.Millis)
}

// NewTradSesCloseTimeNoMillis returns a new TradSesCloseTimeField initialized with val without millisecs.
func NewTradSesCloseTimeNoMillis(val time.Time) TradSesCloseTimeField {
	return NewTradSesCloseTimeWithPrecision(val, quickfix.Seconds)
}

// NewTradSesCloseTimeWithPrecision returns a new TradSesCloseTimeField initialized with val of specified precision.
func NewTradSesCloseTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) TradSesCloseTimeField {
	return TradSesCloseTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f TradSesCloseTimeField) Value() time.Time { return f.Time }

// TradSesEndTimeField is a UTCTIMESTAMP field.
type TradSesEndTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.TradSesEndTime (345).
func (f TradSesEndTimeField) Tag() quickfix.Tag { return tag.TradSesEndTime }

// NewTradSesEndTime returns a new TradSesEndTimeField initialized with val.
func NewTradSesEndTime(val time.Time) TradSesEndTimeField {
	return NewTradSesEndTimeWithPrecision(val, quickfix.Millis)
}

// NewTradSesEndTimeNoMillis returns a new TradSesEndTimeField initialized with val without millisecs.
func NewTradSesEndTimeNoMillis(val time.Time) TradSesEndTimeField {
	return NewTradSesEndTimeWithPrecision(val, quickfix.Seconds)
}

// NewTradSesEndTimeWithPrecision returns a new TradSesEndTimeField initialized with val of specified precision.
func NewTradSesEndTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) TradSesEndTimeField {
	return TradSesEndTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f TradSesEndTimeField) Value() time.Time { return f.Time }

// TradSesMethodField is a enum.TradSesMethod field.
type TradSesMethodField struct{ quickfix.FIXString }

// Tag returns tag.TradSesMethod (338).
func (f TradSesMethodField) Tag() quickfix.Tag { return tag.TradSesMethod }

func NewTradSesMethod(val enum.TradSesMethod) TradSesMethodField {
	return TradSesMethodField{quickfix.FIXString(val)}
}

func (f TradSesMethodField) Value() enum.TradSesMethod { return enum.TradSesMethod(f.String()) }

// TradSesModeField is a enum.TradSesMode field.
type TradSesModeField struct{ quickfix.FIXString }

// Tag returns tag.TradSesMode (339).
func (f TradSesModeField) Tag() quickfix.Tag { return tag.TradSesMode }

func NewTradSesMode(val enum.TradSesMode) TradSesModeField {
	return TradSesModeField{quickfix.FIXString(val)}
}

func (f TradSesModeField) Value() enum.TradSesMode { return enum.TradSesMode(f.String()) }

// TradSesOpenTimeField is a UTCTIMESTAMP field.
type TradSesOpenTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.TradSesOpenTime (342).
func (f TradSesOpenTimeField) Tag() quickfix.Tag { return tag.TradSesOpenTime }

// NewTradSesOpenTime returns a new TradSesOpenTimeField initialized with val.
func NewTradSesOpenTime(val time.Time) TradSesOpenTimeField {
	return NewTradSesOpenTimeWithPrecision(val, quickfix.Millis)
}

// NewTradSesOpenTimeNoMillis returns a new TradSesOpenTimeField initialized with val without millisecs.
func NewTradSesOpenTimeNoMillis(val time.Time) TradSesOpenTimeField {
	return NewTradSesOpenTimeWithPrecision(val, quickfix.Seconds)
}

// NewTradSesOpenTimeWithPrecision returns a new TradSesOpenTimeField initialized with val of specified precision.
func NewTradSesOpenTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) TradSesOpenTimeField {
	return TradSesOpenTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f TradSesOpenTimeField) Value() time.Time { return f.Time }

// TradSesPreCloseTimeField is a UTCTIMESTAMP field.
type TradSesPreCloseTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.TradSesPreCloseTime (343).
func (f TradSesPreCloseTimeField) Tag() quickfix.Tag { return tag.TradSesPreCloseTime }

// NewTradSesPreCloseTime returns a new TradSesPreCloseTimeField initialized with val.
func NewTradSesPreCloseTime(val time.Time) TradSesPreCloseTimeField {
	return NewTradSesPreCloseTimeWithPrecision(val, quickfix.Millis)
}

// NewTradSesPreCloseTimeNoMillis returns a new TradSesPreCloseTimeField initialized with val without millisecs.
func NewTradSesPreCloseTimeNoMillis(val time.Time) TradSesPreCloseTimeField {
	return NewTradSesPreCloseTimeWithPrecision(val, quickfix.Seconds)
}

// NewTradSesPreCloseTimeWithPrecision returns a new TradSesPreCloseTimeField initialized with val of specified precision.
func NewTradSesPreCloseTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) TradSesPreCloseTimeField {
	return TradSesPreCloseTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f TradSesPreCloseTimeField) Value() time.Time { return f.Time }

// TradSesReqIDField is a STRING field.
type TradSesReqIDField struct{ quickfix.FIXString }

// Tag returns tag.TradSesReqID (335).
func (f TradSesReqIDField) Tag() quickfix.Tag { return tag.TradSesReqID }

// NewTradSesReqID returns a new TradSesReqIDField initialized with val.
func NewTradSesReqID(val string) TradSesReqIDField {
	return TradSesReqIDField{quickfix.FIXString(val)}
}

func (f TradSesReqIDField) Value() string { return f.String() }

// TradSesStartTimeField is a UTCTIMESTAMP field.
type TradSesStartTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.TradSesStartTime (341).
func (f TradSesStartTimeField) Tag() quickfix.Tag { return tag.TradSesStartTime }

// NewTradSesStartTime returns a new TradSesStartTimeField initialized with val.
func NewTradSesStartTime(val time.Time) TradSesStartTimeField {
	return NewTradSesStartTimeWithPrecision(val, quickfix.Millis)
}

// NewTradSesStartTimeNoMillis returns a new TradSesStartTimeField initialized with val without millisecs.
func NewTradSesStartTimeNoMillis(val time.Time) TradSesStartTimeField {
	return NewTradSesStartTimeWithPrecision(val, quickfix.Seconds)
}

// NewTradSesStartTimeWithPrecision returns a new TradSesStartTimeField initialized with val of specified precision.
func NewTradSesStartTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) TradSesStartTimeField {
	return TradSesStartTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f TradSesStartTimeField) Value() time.Time { return f.Time }

// TradSesStatusField is a enum.TradSesStatus field.
type TradSesStatusField struct{ quickfix.FIXString }

// Tag returns tag.TradSesStatus (340).
func (f TradSesStatusField) Tag() quickfix.Tag { return tag.TradSesStatus }

func NewTradSesStatus(val enum.TradSesStatus) TradSesStatusField {
	return TradSesStatusField{quickfix.FIXString(val)}
}

func (f TradSesStatusField) Value() enum.TradSesStatus { return enum.TradSesStatus(f.String()) }

// TradeConditionField is a enum.TradeCondition field.
type TradeConditionField struct{ quickfix.FIXString }

// Tag returns tag.TradeCondition (277).
func (f TradeConditionField) Tag() quickfix.Tag { return tag.TradeCondition }

func NewTradeCondition(val enum.TradeCondition) TradeConditionField {
	return TradeConditionField{quickfix.FIXString(val)}
}

func (f TradeConditionField) Value() enum.TradeCondition { return enum.TradeCondition(f.String()) }

// TradeDateField is a LOCALMKTDATE field.
type TradeDateField struct{ quickfix.FIXString }

// Tag returns tag.TradeDate (75).
func (f TradeDateField) Tag() quickfix.Tag { return tag.TradeDate }

// NewTradeDate returns a new TradeDateField initialized with val.
func NewTradeDate(val string) TradeDateField {
	return TradeDateField{quickfix.FIXString(val)}
}

func (f TradeDateField) Value() string { return f.String() }

// TradeTypeField is a enum.TradeType field.
type TradeTypeField struct{ quickfix.FIXString }

// Tag returns tag.TradeType (418).
func (f TradeTypeField) Tag() quickfix.Tag { return tag.TradeType }

func NewTradeType(val enum.TradeType) TradeTypeField {
	return TradeTypeField{quickfix.FIXString(val)}
}

func (f TradeTypeField) Value() enum.TradeType { return enum.TradeType(f.String()) }

// TradingSessionIDField is a STRING field.
type TradingSessionIDField struct{ quickfix.FIXString }

// Tag returns tag.TradingSessionID (336).
func (f TradingSessionIDField) Tag() quickfix.Tag { return tag.TradingSessionID }

// NewTradingSessionID returns a new TradingSessionIDField initialized with val.
func NewTradingSessionID(val string) TradingSessionIDField {
	return TradingSessionIDField{quickfix.FIXString(val)}
}

func (f TradingSessionIDField) Value() string { return f.String() }

// TransactTimeField is a UTCTIMESTAMP field.
type TransactTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.TransactTime (60).
func (f TransactTimeField) Tag() quickfix.Tag { return tag.TransactTime }

// NewTransactTime returns a new TransactTimeField initialized with val.
func NewTransactTime(val time.Time) TransactTimeField {
	return NewTransactTimeWithPrecision(val, quickfix.Millis)
}

// NewTransactTimeNoMillis returns a new TransactTimeField initialized with val without millisecs.
func NewTransactTimeNoMillis(val time.Time) TransactTimeField {
	return NewTransactTimeWithPrecision(val, quickfix.Seconds)
}

// NewTransactTimeWithPrecision returns a new TransactTimeField initialized with val of specified precision.
func NewTransactTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) TransactTimeField {
	return TransactTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f TransactTimeField) Value() time.Time { return f.Time }

// TriggerPriceDirectionField is a enum.TriggerPriceDirection field.
type TriggerPriceDirectionField struct{ quickfix.FIXString }

// Tag returns tag.TriggerPriceDirection (1109).
func (f TriggerPriceDirectionField) Tag() quickfix.Tag { return tag.TriggerPriceDirection }

func NewTriggerPriceDirection(val enum.TriggerPriceDirection) TriggerPriceDirectionField {
	return TriggerPriceDirectionField{quickfix.FIXString(val)}
}

func (f TriggerPriceDirectionField) Value() enum.TriggerPriceDirection {
	return enum.TriggerPriceDirection(f.String())
}

// URLLinkField is a STRING field.
type URLLinkField struct{ quickfix.FIXString }

// Tag returns tag.URLLink (149).
func (f URLLinkField) Tag() quickfix.Tag { return tag.URLLink }

// NewURLLink returns a new URLLinkField initialized with val.
func NewURLLink(val string) URLLinkField {
	return URLLinkField{quickfix.FIXString(val)}
}

func (f URLLinkField) Value() string { return f.String() }

// UnderlyingContractMultiplierField is a FLOAT field.
type UnderlyingContractMultiplierField struct{ quickfix.FIXDecimal }

// Tag returns tag.UnderlyingContractMultiplier (436).
func (f UnderlyingContractMultiplierField) Tag() quickfix.Tag {
	return tag.UnderlyingContractMultiplier
}

// NewUnderlyingContractMultiplier returns a new UnderlyingContractMultiplierField initialized with val and scale.
func NewUnderlyingContractMultiplier(val decimal.Decimal, scale int32) UnderlyingContractMultiplierField {
	return UnderlyingContractMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingContractMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

// UnderlyingCouponRateField is a FLOAT field.
type UnderlyingCouponRateField struct{ quickfix.FIXDecimal }

// Tag returns tag.UnderlyingCouponRate (435).
func (f UnderlyingCouponRateField) Tag() quickfix.Tag { return tag.UnderlyingCouponRate }

// NewUnderlyingCouponRate returns a new UnderlyingCouponRateField initialized with val and scale.
func NewUnderlyingCouponRate(val decimal.Decimal, scale int32) UnderlyingCouponRateField {
	return UnderlyingCouponRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingCouponRateField) Value() (val decimal.Decimal) { return f.Decimal }

// UnderlyingCurrencyField is a CURRENCY field.
type UnderlyingCurrencyField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingCurrency (318).
func (f UnderlyingCurrencyField) Tag() quickfix.Tag { return tag.UnderlyingCurrency }

// NewUnderlyingCurrency returns a new UnderlyingCurrencyField initialized with val.
func NewUnderlyingCurrency(val string) UnderlyingCurrencyField {
	return UnderlyingCurrencyField{quickfix.FIXString(val)}
}

func (f UnderlyingCurrencyField) Value() string { return f.String() }

// UnderlyingIDSourceField is a STRING field.
type UnderlyingIDSourceField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingIDSource (305).
func (f UnderlyingIDSourceField) Tag() quickfix.Tag { return tag.UnderlyingIDSource }

// NewUnderlyingIDSource returns a new UnderlyingIDSourceField initialized with val.
func NewUnderlyingIDSource(val string) UnderlyingIDSourceField {
	return UnderlyingIDSourceField{quickfix.FIXString(val)}
}

func (f UnderlyingIDSourceField) Value() string { return f.String() }

// UnderlyingIssuerField is a STRING field.
type UnderlyingIssuerField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingIssuer (306).
func (f UnderlyingIssuerField) Tag() quickfix.Tag { return tag.UnderlyingIssuer }

// NewUnderlyingIssuer returns a new UnderlyingIssuerField initialized with val.
func NewUnderlyingIssuer(val string) UnderlyingIssuerField {
	return UnderlyingIssuerField{quickfix.FIXString(val)}
}

func (f UnderlyingIssuerField) Value() string { return f.String() }

// UnderlyingMaturityDayField is a DAYOFMONTH field.
type UnderlyingMaturityDayField struct{ quickfix.FIXInt }

// Tag returns tag.UnderlyingMaturityDay (314).
func (f UnderlyingMaturityDayField) Tag() quickfix.Tag { return tag.UnderlyingMaturityDay }

// NewUnderlyingMaturityDay returns a new UnderlyingMaturityDayField initialized with val.
func NewUnderlyingMaturityDay(val int) UnderlyingMaturityDayField {
	return UnderlyingMaturityDayField{quickfix.FIXInt(val)}
}

func (f UnderlyingMaturityDayField) Value() int { return f.Int() }

// UnderlyingMaturityMonthYearField is a MONTHYEAR field.
type UnderlyingMaturityMonthYearField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingMaturityMonthYear (313).
func (f UnderlyingMaturityMonthYearField) Tag() quickfix.Tag { return tag.UnderlyingMaturityMonthYear }

// NewUnderlyingMaturityMonthYear returns a new UnderlyingMaturityMonthYearField initialized with val.
func NewUnderlyingMaturityMonthYear(val string) UnderlyingMaturityMonthYearField {
	return UnderlyingMaturityMonthYearField{quickfix.FIXString(val)}
}

func (f UnderlyingMaturityMonthYearField) Value() string { return f.String() }

// UnderlyingOptAttributeField is a CHAR field.
type UnderlyingOptAttributeField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingOptAttribute (317).
func (f UnderlyingOptAttributeField) Tag() quickfix.Tag { return tag.UnderlyingOptAttribute }

// NewUnderlyingOptAttribute returns a new UnderlyingOptAttributeField initialized with val.
func NewUnderlyingOptAttribute(val string) UnderlyingOptAttributeField {
	return UnderlyingOptAttributeField{quickfix.FIXString(val)}
}

func (f UnderlyingOptAttributeField) Value() string { return f.String() }

// UnderlyingPutOrCallField is a INT field.
type UnderlyingPutOrCallField struct{ quickfix.FIXInt }

// Tag returns tag.UnderlyingPutOrCall (315).
func (f UnderlyingPutOrCallField) Tag() quickfix.Tag { return tag.UnderlyingPutOrCall }

// NewUnderlyingPutOrCall returns a new UnderlyingPutOrCallField initialized with val.
func NewUnderlyingPutOrCall(val int) UnderlyingPutOrCallField {
	return UnderlyingPutOrCallField{quickfix.FIXInt(val)}
}

func (f UnderlyingPutOrCallField) Value() int { return f.Int() }

// UnderlyingSecurityDescField is a STRING field.
type UnderlyingSecurityDescField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingSecurityDesc (307).
func (f UnderlyingSecurityDescField) Tag() quickfix.Tag { return tag.UnderlyingSecurityDesc }

// NewUnderlyingSecurityDesc returns a new UnderlyingSecurityDescField initialized with val.
func NewUnderlyingSecurityDesc(val string) UnderlyingSecurityDescField {
	return UnderlyingSecurityDescField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityDescField) Value() string { return f.String() }

// UnderlyingSecurityExchangeField is a EXCHANGE field.
type UnderlyingSecurityExchangeField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingSecurityExchange (308).
func (f UnderlyingSecurityExchangeField) Tag() quickfix.Tag { return tag.UnderlyingSecurityExchange }

// NewUnderlyingSecurityExchange returns a new UnderlyingSecurityExchangeField initialized with val.
func NewUnderlyingSecurityExchange(val string) UnderlyingSecurityExchangeField {
	return UnderlyingSecurityExchangeField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityExchangeField) Value() string { return f.String() }

// UnderlyingSecurityIDField is a STRING field.
type UnderlyingSecurityIDField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingSecurityID (309).
func (f UnderlyingSecurityIDField) Tag() quickfix.Tag { return tag.UnderlyingSecurityID }

// NewUnderlyingSecurityID returns a new UnderlyingSecurityIDField initialized with val.
func NewUnderlyingSecurityID(val string) UnderlyingSecurityIDField {
	return UnderlyingSecurityIDField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityIDField) Value() string { return f.String() }

// UnderlyingSecurityTypeField is a STRING field.
type UnderlyingSecurityTypeField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingSecurityType (310).
func (f UnderlyingSecurityTypeField) Tag() quickfix.Tag { return tag.UnderlyingSecurityType }

// NewUnderlyingSecurityType returns a new UnderlyingSecurityTypeField initialized with val.
func NewUnderlyingSecurityType(val string) UnderlyingSecurityTypeField {
	return UnderlyingSecurityTypeField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityTypeField) Value() string { return f.String() }

// UnderlyingStrikePriceField is a PRICE field.
type UnderlyingStrikePriceField struct{ quickfix.FIXDecimal }

// Tag returns tag.UnderlyingStrikePrice (316).
func (f UnderlyingStrikePriceField) Tag() quickfix.Tag { return tag.UnderlyingStrikePrice }

// NewUnderlyingStrikePrice returns a new UnderlyingStrikePriceField initialized with val and scale.
func NewUnderlyingStrikePrice(val decimal.Decimal, scale int32) UnderlyingStrikePriceField {
	return UnderlyingStrikePriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingStrikePriceField) Value() (val decimal.Decimal) { return f.Decimal }

// UnderlyingSymbolField is a STRING field.
type UnderlyingSymbolField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingSymbol (311).
func (f UnderlyingSymbolField) Tag() quickfix.Tag { return tag.UnderlyingSymbol }

// NewUnderlyingSymbol returns a new UnderlyingSymbolField initialized with val.
func NewUnderlyingSymbol(val string) UnderlyingSymbolField {
	return UnderlyingSymbolField{quickfix.FIXString(val)}
}

func (f UnderlyingSymbolField) Value() string { return f.String() }

// UnderlyingSymbolSfxField is a STRING field.
type UnderlyingSymbolSfxField struct{ quickfix.FIXString }

// Tag returns tag.UnderlyingSymbolSfx (312).
func (f UnderlyingSymbolSfxField) Tag() quickfix.Tag { return tag.UnderlyingSymbolSfx }

// NewUnderlyingSymbolSfx returns a new UnderlyingSymbolSfxField initialized with val.
func NewUnderlyingSymbolSfx(val string) UnderlyingSymbolSfxField {
	return UnderlyingSymbolSfxField{quickfix.FIXString(val)}
}

func (f UnderlyingSymbolSfxField) Value() string { return f.String() }

// UnsolicitedIndicatorField is a BOOLEAN field.
type UnsolicitedIndicatorField struct{ quickfix.FIXBoolean }

// Tag returns tag.UnsolicitedIndicator (325).
func (f UnsolicitedIndicatorField) Tag() quickfix.Tag { return tag.UnsolicitedIndicator }

// NewUnsolicitedIndicator returns a new UnsolicitedIndicatorField initialized with val.
func NewUnsolicitedIndicator(val bool) UnsolicitedIndicatorField {
	return UnsolicitedIndicatorField{quickfix.FIXBoolean(val)}
}

func (f UnsolicitedIndicatorField) Value() bool { return f.Bool() }

// UrgencyField is a enum.Urgency field.
type UrgencyField struct{ quickfix.FIXString }

// Tag returns tag.Urgency (61).
func (f UrgencyField) Tag() quickfix.Tag { return tag.Urgency }

func NewUrgency(val enum.Urgency) UrgencyField {
	return UrgencyField{quickfix.FIXString(val)}
}

func (f UrgencyField) Value() enum.Urgency { return enum.Urgency(f.String()) }

// ValidUntilTimeField is a UTCTIMESTAMP field.
type ValidUntilTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.ValidUntilTime (62).
func (f ValidUntilTimeField) Tag() quickfix.Tag { return tag.ValidUntilTime }

// NewValidUntilTime returns a new ValidUntilTimeField initialized with val.
func NewValidUntilTime(val time.Time) ValidUntilTimeField {
	return NewValidUntilTimeWithPrecision(val, quickfix.Millis)
}

// NewValidUntilTimeNoMillis returns a new ValidUntilTimeField initialized with val without millisecs.
func NewValidUntilTimeNoMillis(val time.Time) ValidUntilTimeField {
	return NewValidUntilTimeWithPrecision(val, quickfix.Seconds)
}

// NewValidUntilTimeWithPrecision returns a new ValidUntilTimeField initialized with val of specified precision.
func NewValidUntilTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) ValidUntilTimeField {
	return ValidUntilTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f ValidUntilTimeField) Value() time.Time { return f.Time }

// ValueOfFuturesField is a AMT field.
type ValueOfFuturesField struct{ quickfix.FIXDecimal }

// Tag returns tag.ValueOfFutures (408).
func (f ValueOfFuturesField) Tag() quickfix.Tag { return tag.ValueOfFutures }

// NewValueOfFutures returns a new ValueOfFuturesField initialized with val and scale.
func NewValueOfFutures(val decimal.Decimal, scale int32) ValueOfFuturesField {
	return ValueOfFuturesField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ValueOfFuturesField) Value() (val decimal.Decimal) { return f.Decimal }

// WaveNoField is a STRING field.
type WaveNoField struct{ quickfix.FIXString }

// Tag returns tag.WaveNo (105).
func (f WaveNoField) Tag() quickfix.Tag { return tag.WaveNo }

// NewWaveNo returns a new WaveNoField initialized with val.
func NewWaveNo(val string) WaveNoField {
	return WaveNoField{quickfix.FIXString(val)}
}

func (f WaveNoField) Value() string { return f.String() }

// WtAverageLiquidityField is a FLOAT field.
type WtAverageLiquidityField struct{ quickfix.FIXDecimal }

// Tag returns tag.WtAverageLiquidity (410).
func (f WtAverageLiquidityField) Tag() quickfix.Tag { return tag.WtAverageLiquidity }

// NewWtAverageLiquidity returns a new WtAverageLiquidityField initialized with val and scale.
func NewWtAverageLiquidity(val decimal.Decimal, scale int32) WtAverageLiquidityField {
	return WtAverageLiquidityField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f WtAverageLiquidityField) Value() (val decimal.Decimal) { return f.Decimal }

// XmlDataField is a DATA field.
type XmlDataField struct{ quickfix.FIXString }

// Tag returns tag.XmlData (213).
func (f XmlDataField) Tag() quickfix.Tag { return tag.XmlData }

// NewXmlData returns a new XmlDataField initialized with val.
func NewXmlData(val string) XmlDataField {
	return XmlDataField{quickfix.FIXString(val)}
}

func (f XmlDataField) Value() string { return f.String() }

// XmlDataLenField is a LENGTH field.
type XmlDataLenField struct{ quickfix.FIXInt }

// Tag returns tag.XmlDataLen (212).
func (f XmlDataLenField) Tag() quickfix.Tag { return tag.XmlDataLen }

// NewXmlDataLen returns a new XmlDataLenField initialized with val.
func NewXmlDataLen(val int) XmlDataLenField {
	return XmlDataLenField{quickfix.FIXInt(val)}
}

func (f XmlDataLenField) Value() int { return f.Int() }
