package enum

// Adjustment field enumeration values.
type Adjustment string

const (
	Adjustment_CANCEL     Adjustment = "1"
	Adjustment_ERROR      Adjustment = "2"
	Adjustment_CORRECTION Adjustment = "3"
)

// AdvSide field enumeration values.
type AdvSide string

const (
	AdvSide_BUY   AdvSide = "B"
	AdvSide_SELL  AdvSide = "S"
	AdvSide_TRADE AdvSide = "T"
	AdvSide_CROSS AdvSide = "X"
)

// AdvTransType field enumeration values.
type AdvTransType string

const (
	AdvTransType_CANCEL  AdvTransType = "C"
	AdvTransType_NEW     AdvTransType = "N"
	AdvTransType_REPLACE AdvTransType = "R"
)

// AggregatedBook field enumeration values.
type AggregatedBook string

const (
	AggregatedBook_NO  AggregatedBook = "N"
	AggregatedBook_YES AggregatedBook = "Y"
)

// AllocHandlInst field enumeration values.
type AllocHandlInst string

const (
	AllocHandlInst_MATCH             AllocHandlInst = "1"
	AllocHandlInst_FORWARD           AllocHandlInst = "2"
	AllocHandlInst_FORWARD_AND_MATCH AllocHandlInst = "3"
)

// AllocLinkType field enumeration values.
type AllocLinkType string

const (
	AllocLinkType_F_X_NETTING AllocLinkType = "0"
	AllocLinkType_F_X_SWAP    AllocLinkType = "1"
)

// AllocRejCode field enumeration values.
type AllocRejCode string

const (
	AllocRejCode_UNKNOWN_ACCOUNT                   AllocRejCode = "0"
	AllocRejCode_INCORRECT_QUANTITY                AllocRejCode = "1"
	AllocRejCode_INCORRECT_AVERAGE_PRICE           AllocRejCode = "2"
	AllocRejCode_UNKNOWN_EXECUTING_BROKER_MNEMONIC AllocRejCode = "3"
	AllocRejCode_COMMISSION_DIFFERENCE             AllocRejCode = "4"
	AllocRejCode_UNKNOWN_ORDERID                   AllocRejCode = "5"
	AllocRejCode_UNKNOWN_LISTID                    AllocRejCode = "6"
	AllocRejCode_OTHER                             AllocRejCode = "7"
)

// AllocStatus field enumeration values.
type AllocStatus string

const (
	AllocStatus_ACCEPTED       AllocStatus = "0"
	AllocStatus_REJECTED       AllocStatus = "1"
	AllocStatus_PARTIAL_ACCEPT AllocStatus = "2"
	AllocStatus_RECEIVED       AllocStatus = "3"
)

// AllocTransType field enumeration values.
type AllocTransType string

const (
	AllocTransType_NEW                            AllocTransType = "0"
	AllocTransType_REPLACE                        AllocTransType = "1"
	AllocTransType_CANCEL                         AllocTransType = "2"
	AllocTransType_PRELIMINARY                    AllocTransType = "3"
	AllocTransType_CALCULATED                     AllocTransType = "4"
	AllocTransType_CALCULATED_WITHOUT_PRELIMINARY AllocTransType = "5"
)

// BasisPxType field enumeration values.
type BasisPxType string

const (
	BasisPxType_CLOSING_PRICE_AT_MORNING_SESSION              BasisPxType = "2"
	BasisPxType_CLOSING_PRICE                                 BasisPxType = "3"
	BasisPxType_CURRENT_PRICE                                 BasisPxType = "4"
	BasisPxType_SQ                                            BasisPxType = "5"
	BasisPxType_VWAP_THROUGH_A_DAY                            BasisPxType = "6"
	BasisPxType_VWAP_THROUGH_A_MORNING_SESSION                BasisPxType = "7"
	BasisPxType_VWAP_THROUGH_AN_AFTERNOON_SESSION             BasisPxType = "8"
	BasisPxType_VWAP_THROUGH_A_DAY_EXCEPT_YORI                BasisPxType = "9"
	BasisPxType_VWAP_THROUGH_A_MORNING_SESSION_EXCEPT_YORI    BasisPxType = "A"
	BasisPxType_VWAP_THROUGH_AN_AFTERNOON_SESSION_EXCEPT_YORI BasisPxType = "B"
	BasisPxType_STRIKE                                        BasisPxType = "C"
	BasisPxType_OPEN                                          BasisPxType = "D"
	BasisPxType_OTHERS                                        BasisPxType = "Z"
)

// Benchmark field enumeration values.
type Benchmark string

const (
	Benchmark_CURVE      Benchmark = "1"
	Benchmark_5_YR       Benchmark = "2"
	Benchmark_OLD_5      Benchmark = "3"
	Benchmark_10_YR      Benchmark = "4"
	Benchmark_OLD_10     Benchmark = "5"
	Benchmark_30_YR      Benchmark = "6"
	Benchmark_OLD_30     Benchmark = "7"
	Benchmark_3_MO_LIBOR Benchmark = "8"
	Benchmark_6_MO_LIBOR Benchmark = "9"
)

// BidRequestTransType field enumeration values.
type BidRequestTransType string

const (
	BidRequestTransType_CANCEL BidRequestTransType = "C"
	BidRequestTransType_NO     BidRequestTransType = "N"
)

// BusinessRejectReason field enumeration values.
type BusinessRejectReason string

const (
	BusinessRejectReason_OTHER                                BusinessRejectReason = "0"
	BusinessRejectReason_UNKOWN_ID                            BusinessRejectReason = "1"
	BusinessRejectReason_UNKNOWN_SECURITY                     BusinessRejectReason = "2"
	BusinessRejectReason_UNSUPPORTED_MESSAGE_TYPE             BusinessRejectReason = "3"
	BusinessRejectReason_APPLICATION_NOT_AVAILABLE            BusinessRejectReason = "4"
	BusinessRejectReason_CONDITIONALLY_REQUIRED_FIELD_MISSING BusinessRejectReason = "5"
)

// CancelOrdersOnDisconnect field enumeration values.
type CancelOrdersOnDisconnect string

const (
	CancelOrdersOnDisconnect_SESSION CancelOrdersOnDisconnect = "S"
	CancelOrdersOnDisconnect_PROFILE CancelOrdersOnDisconnect = "Y"
)

// CommType field enumeration values.
type CommType string

const (
	CommType_PER_SHARE  CommType = "1"
	CommType_PERCENTAGE CommType = "2"
	CommType_ABSOLUTE   CommType = "3"
)

// CorporateAction field enumeration values.
type CorporateAction string

const (
	CorporateAction_EX_DIVIDEND     CorporateAction = "A"
	CorporateAction_EX_DISTRIBUTION CorporateAction = "B"
	CorporateAction_EX_RIGHTS       CorporateAction = "C"
	CorporateAction_NEW             CorporateAction = "D"
	CorporateAction_EX_INTEREST     CorporateAction = "E"
)

// CoveredOrUncovered field enumeration values.
type CoveredOrUncovered string

const (
	CoveredOrUncovered_COVERED   CoveredOrUncovered = "0"
	CoveredOrUncovered_UNCOVERED CoveredOrUncovered = "1"
)

// CustomerOrFirm field enumeration values.
type CustomerOrFirm string

const (
	CustomerOrFirm_CUSTOMER CustomerOrFirm = "0"
	CustomerOrFirm_FIRM     CustomerOrFirm = "1"
)

// CxlRejReason field enumeration values.
type CxlRejReason string

const (
	CxlRejReason_TOO_LATE_TO_CANCEL                                        CxlRejReason = "0"
	CxlRejReason_UNKNOWN_ORDER                                             CxlRejReason = "1"
	CxlRejReason_BROKER_OPTION                                             CxlRejReason = "2"
	CxlRejReason_ORDER_ALREADY_IN_PENDING_CANCEL_OR_PENDING_REPLACE_STATUS CxlRejReason = "3"
)

// CxlRejResponseTo field enumeration values.
type CxlRejResponseTo string

const (
	CxlRejResponseTo_ORDER_CANCEL_REQUEST         CxlRejResponseTo = "1"
	CxlRejResponseTo_ORDER_CANCEL_REPLACE_REQUEST CxlRejResponseTo = "2"
)

// DKReason field enumeration values.
type DKReason string

const (
	DKReason_UNKNOWN_SYMBOL         DKReason = "A"
	DKReason_WRONG_SIDE             DKReason = "B"
	DKReason_QUANTITY_EXCEEDS_ORDER DKReason = "C"
	DKReason_NO_MATCHING_ORDER      DKReason = "D"
	DKReason_PRICE_EXCEEDS_LIMIT    DKReason = "E"
	DKReason_OTHER                  DKReason = "Z"
)

// DeleteReason field enumeration values.
type DeleteReason string

const (
	DeleteReason_CANCELATION DeleteReason = "0"
	DeleteReason_ERROR       DeleteReason = "1"
)

// DiscretionInst field enumeration values.
type DiscretionInst string

const (
	DiscretionInst_RELATED_TO_DISPLAYED_PRICE     DiscretionInst = "0"
	DiscretionInst_RELATED_TO_MARKET_PRICE        DiscretionInst = "1"
	DiscretionInst_RELATED_TO_PRIMARY_PRICE       DiscretionInst = "2"
	DiscretionInst_RELATED_TO_LOCAL_PRIMARY_PRICE DiscretionInst = "3"
	DiscretionInst_RELATED_TO_MIDPOINT_PRICE      DiscretionInst = "4"
	DiscretionInst_RELATED_TO_LAST_TRADE_PRICE    DiscretionInst = "5"
)

// DropCopyFlag field enumeration values.
type DropCopyFlag string

const (
	DropCopyFlag_NO  DropCopyFlag = "N"
	DropCopyFlag_YES DropCopyFlag = "Y"
)

// DueToRelated field enumeration values.
type DueToRelated string

const (
	DueToRelated_NO  DueToRelated = "N"
	DueToRelated_YES DueToRelated = "Y"
)

// EmailType field enumeration values.
type EmailType string

const (
	EmailType_NEW         EmailType = "0"
	EmailType_REPLY       EmailType = "1"
	EmailType_ADMIN_REPLY EmailType = "2"
)

// EncryptMethod field enumeration values.
type EncryptMethod string

const (
	EncryptMethod_NONE        EncryptMethod = "0"
	EncryptMethod_PKCS        EncryptMethod = "1"
	EncryptMethod_DES         EncryptMethod = "2"
	EncryptMethod_PKCS_DES    EncryptMethod = "3"
	EncryptMethod_PGP_DES     EncryptMethod = "4"
	EncryptMethod_PGP_DES_MD5 EncryptMethod = "5"
	EncryptMethod_PEM_DES_MD5 EncryptMethod = "6"
)

// ExchangeForPhysical field enumeration values.
type ExchangeForPhysical string

const (
	ExchangeForPhysical_NO  ExchangeForPhysical = "N"
	ExchangeForPhysical_YES ExchangeForPhysical = "Y"
)

// ExecInst field enumeration values.
type ExecInst string

const (
	ExecInst_STAY_ON_OFFERSIDE                                     ExecInst = "0"
	ExecInst_NOT_HELD                                              ExecInst = "1"
	ExecInst_WORK                                                  ExecInst = "2"
	ExecInst_GO_ALONG                                              ExecInst = "3"
	ExecInst_OVER_THE_DAY                                          ExecInst = "4"
	ExecInst_HELD                                                  ExecInst = "5"
	ExecInst_PARTICIPATE_DONT_INITIATE                             ExecInst = "6"
	ExecInst_STRICT_SCALE                                          ExecInst = "7"
	ExecInst_TRY_TO_SCALE                                          ExecInst = "8"
	ExecInst_STAY_ON_BIDSIDE                                       ExecInst = "9"
	ExecInst_NO_CROSS                                              ExecInst = "A"
	ExecInst_OK_TO_CROSS                                           ExecInst = "B"
	ExecInst_CALL_FIRST                                            ExecInst = "C"
	ExecInst_PERCENT_OF_VOLUME                                     ExecInst = "D"
	ExecInst_DO_NOT_INCREASE                                       ExecInst = "E"
	ExecInst_DO_NOT_REDUCE                                         ExecInst = "F"
	ExecInst_ALL_OR_NONE                                           ExecInst = "G"
	ExecInst_INSTITUTIONS_ONLY                                     ExecInst = "I"
	ExecInst_LAST_PEG                                              ExecInst = "L"
	ExecInst_MID_PRICE_PEG                                         ExecInst = "M"
	ExecInst_NON_NEGOTIABLE                                        ExecInst = "N"
	ExecInst_OPENING_PEG                                           ExecInst = "O"
	ExecInst_MARKET_PEG                                            ExecInst = "P"
	ExecInst_PRIMARY_PEG                                           ExecInst = "R"
	ExecInst_SUSPEND                                               ExecInst = "S"
	ExecInst_FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER ExecInst = "T"
	ExecInst_CUSTOMER_DISPLAY_INSTRUCTION                          ExecInst = "U"
	ExecInst_NETTING                                               ExecInst = "V"
	ExecInst_PEG_TO_VWAP                                           ExecInst = "W"
)

// ExecRestatementReason field enumeration values.
type ExecRestatementReason string

const (
	ExecRestatementReason_GT_CORPORATE_ACTION         ExecRestatementReason = "0"
	ExecRestatementReason_GT_RENEWAL                  ExecRestatementReason = "1"
	ExecRestatementReason_VERBAL_CHANGE               ExecRestatementReason = "2"
	ExecRestatementReason_REPRICING_OF_ORDER          ExecRestatementReason = "3"
	ExecRestatementReason_BROKER_OPTION               ExecRestatementReason = "4"
	ExecRestatementReason_PARTIAL_DECLINE_OF_ORDERQTY ExecRestatementReason = "5"
)

// ExecTransType field enumeration values.
type ExecTransType string

const (
	ExecTransType_NEW     ExecTransType = "0"
	ExecTransType_CANCEL  ExecTransType = "1"
	ExecTransType_CORRECT ExecTransType = "2"
	ExecTransType_STATUS  ExecTransType = "3"
)

// ExecType field enumeration values.
type ExecType string

const (
	ExecType_NEW             ExecType = "0"
	ExecType_PARTIAL_FILL    ExecType = "1"
	ExecType_FILL            ExecType = "2"
	ExecType_DONE_FOR_DAY    ExecType = "3"
	ExecType_CANCELED        ExecType = "4"
	ExecType_REPLACE         ExecType = "5"
	ExecType_PENDING_CANCEL  ExecType = "6"
	ExecType_STOPPED         ExecType = "7"
	ExecType_REJECTED        ExecType = "8"
	ExecType_SUSPENDED       ExecType = "9"
	ExecType_PENDING_NEW     ExecType = "A"
	ExecType_CALCULATED      ExecType = "B"
	ExecType_EXPIRED         ExecType = "C"
	ExecType_RESTATED        ExecType = "D"
	ExecType_PENDING_REPLACE ExecType = "E"
)

// FinancialStatus field enumeration values.
type FinancialStatus string

const (
	FinancialStatus_BANKRUPT FinancialStatus = "1"
)

// ForexReq field enumeration values.
type ForexReq string

const (
	ForexReq_NO  ForexReq = "N"
	ForexReq_YES ForexReq = "Y"
)

// GTBookingInst field enumeration values.
type GTBookingInst string

const (
	GTBookingInst_BOOK_OUT_ALL_TRADES_ON_DAY_OF_EXECUTION                GTBookingInst = "0"
	GTBookingInst_ACCUMULATE_EXECUTIONS_UNTIL_ORDER_IS_FILLED_OR_EXPIRES GTBookingInst = "1"
	GTBookingInst_ACCUMULATE_UNTIL_VERBALLY_NOTIFIED_OTHERWISE           GTBookingInst = "2"
)

// GapFillFlag field enumeration values.
type GapFillFlag string

const (
	GapFillFlag_NO  GapFillFlag = "N"
	GapFillFlag_YES GapFillFlag = "Y"
)

// HaltReasonChar field enumeration values.
type HaltReasonChar string

const (
	HaltReasonChar_NEWS_DISSEMINATION     HaltReasonChar = "D"
	HaltReasonChar_ORDER_INFLUX           HaltReasonChar = "E"
	HaltReasonChar_ORDER_IMBALANCE        HaltReasonChar = "I"
	HaltReasonChar_ADDITIONAL_INFORMATION HaltReasonChar = "M"
	HaltReasonChar_NEWS_PENDING           HaltReasonChar = "P"
	HaltReasonChar_EQUIPMENT_CHANGEOVER   HaltReasonChar = "X"
)

// HandlInst field enumeration values.
type HandlInst string

const (
	HandlInst_AUTOMATED_EXECUTION_ORDER_PRIVATE_NO_BROKER_INTERVENTION HandlInst = "1"
	HandlInst_AUTOMATED_EXECUTION_ORDER_PUBLIC_BROKER_INTERVENTION_OK  HandlInst = "2"
	HandlInst_MANUAL_ORDER_BEST_EXECUTION                              HandlInst = "3"
)

// IDSource field enumeration values.
type IDSource string

const (
	IDSource_CUSIP                         IDSource = "1"
	IDSource_SEDOL                         IDSource = "2"
	IDSource_QUIK                          IDSource = "3"
	IDSource_ISIN_NUMBER                   IDSource = "4"
	IDSource_RIC_CODE                      IDSource = "5"
	IDSource_ISO_CURRENCY_CODE             IDSource = "6"
	IDSource_ISO_COUNTRY_CODE              IDSource = "7"
	IDSource_EXCHANGE_SYMBOL               IDSource = "8"
	IDSource_CONSOLIDATED_TAPE_ASSOCIATION IDSource = "9"
)

// IOINaturalFlag field enumeration values.
type IOINaturalFlag string

const (
	IOINaturalFlag_NO  IOINaturalFlag = "N"
	IOINaturalFlag_YES IOINaturalFlag = "Y"
)

// IOIQltyInd field enumeration values.
type IOIQltyInd string

const (
	IOIQltyInd_HIGH   IOIQltyInd = "H"
	IOIQltyInd_LOW    IOIQltyInd = "L"
	IOIQltyInd_MEDIUM IOIQltyInd = "M"
)

// IOIQualifier field enumeration values.
type IOIQualifier string

const (
	IOIQualifier_ALL_OR_NONE          IOIQualifier = "A"
	IOIQualifier_AT_THE_CLOSE         IOIQualifier = "C"
	IOIQualifier_IN_TOUCH_WITH        IOIQualifier = "I"
	IOIQualifier_LIMIT                IOIQualifier = "L"
	IOIQualifier_MORE_BEHIND          IOIQualifier = "M"
	IOIQualifier_AT_THE_OPEN          IOIQualifier = "O"
	IOIQualifier_TAKING_A_POSITION    IOIQualifier = "P"
	IOIQualifier_AT_THE_MARKET        IOIQualifier = "Q"
	IOIQualifier_READY_TO_TRADE       IOIQualifier = "R"
	IOIQualifier_PORTFOLIO_SHOW_N     IOIQualifier = "S"
	IOIQualifier_THROUGH_THE_DAY      IOIQualifier = "T"
	IOIQualifier_VERSUS               IOIQualifier = "V"
	IOIQualifier_INDICATION           IOIQualifier = "W"
	IOIQualifier_CROSSING_OPPORTUNITY IOIQualifier = "X"
	IOIQualifier_AT_THE_MIDPOINT      IOIQualifier = "Y"
	IOIQualifier_PRE_OPEN             IOIQualifier = "Z"
)

// IOIShares field enumeration values.
type IOIShares string

const (
	IOIShares_LARGE  IOIShares = "L"
	IOIShares_MEDIUM IOIShares = "M"
	IOIShares_SMALL  IOIShares = "S"
)

// IOITransType field enumeration values.
type IOITransType string

const (
	IOITransType_CANCEL  IOITransType = "C"
	IOITransType_NEW     IOITransType = "N"
	IOITransType_REPLACE IOITransType = "R"
)

// InViewOfCommon field enumeration values.
type InViewOfCommon string

const (
	InViewOfCommon_NO  InViewOfCommon = "N"
	InViewOfCommon_YES InViewOfCommon = "Y"
)

// IncTaxInd field enumeration values.
type IncTaxInd string

const (
	IncTaxInd_NET   IncTaxInd = "1"
	IncTaxInd_GROSS IncTaxInd = "2"
)

// LastCapacity field enumeration values.
type LastCapacity string

const (
	LastCapacity_AGENT              LastCapacity = "1"
	LastCapacity_CROSS_AS_AGENT     LastCapacity = "2"
	LastCapacity_CROSS_AS_PRINCIPAL LastCapacity = "3"
	LastCapacity_PRINCIPAL          LastCapacity = "4"
)

// LiquidityIndType field enumeration values.
type LiquidityIndType string

const (
	LiquidityIndType_5_DAY_MOVING_AVERAGE  LiquidityIndType = "1"
	LiquidityIndType_20_DAY_MOVING_AVERAGE LiquidityIndType = "2"
	LiquidityIndType_NORMAL_MARKET_SIZE    LiquidityIndType = "3"
	LiquidityIndType_OTHER                 LiquidityIndType = "4"
)

// ListExecInstType field enumeration values.
type ListExecInstType string

const (
	ListExecInstType_IMMEDIATE                    ListExecInstType = "1"
	ListExecInstType_WAIT_FOR_EXECUTE_INSTRUCTION ListExecInstType = "2"
)

// LocateReqd field enumeration values.
type LocateReqd string

const (
	LocateReqd_NO  LocateReqd = "N"
	LocateReqd_YES LocateReqd = "Y"
)

// MDEntryType field enumeration values.
type MDEntryType string

const (
	MDEntryType_BID                        MDEntryType = "0"
	MDEntryType_OFFER                      MDEntryType = "1"
	MDEntryType_TRADE                      MDEntryType = "2"
	MDEntryType_INDEX_VALUE                MDEntryType = "3"
	MDEntryType_OPENING_PRICE              MDEntryType = "4"
	MDEntryType_CLOSING_PRICE              MDEntryType = "5"
	MDEntryType_SETTLEMENT_PRICE           MDEntryType = "6"
	MDEntryType_TRADING_SESSION_HIGH_PRICE MDEntryType = "7"
	MDEntryType_TRADING_SESSION_LOW_PRICE  MDEntryType = "8"
	MDEntryType_TRADING_SESSION_VWAP_PRICE MDEntryType = "9"
)

// MDReqRejReason field enumeration values.
type MDReqRejReason string

const (
	MDReqRejReason_UNKNOWN_SYMBOL                      MDReqRejReason = "0"
	MDReqRejReason_DUPLICATE_MDREQID                   MDReqRejReason = "1"
	MDReqRejReason_INSUFFICIENT_BANDWIDTH              MDReqRejReason = "2"
	MDReqRejReason_INSUFFICIENT_PERMISSIONS            MDReqRejReason = "3"
	MDReqRejReason_UNSUPPORTED_SUBSCRIPTIONREQUESTTYPE MDReqRejReason = "4"
	MDReqRejReason_UNSUPPORTED_MARKETDEPTH             MDReqRejReason = "5"
	MDReqRejReason_UNSUPPORTED_MDUPDATETYPE            MDReqRejReason = "6"
	MDReqRejReason_UNSUPPORTED_AGGREGATEDBOOK          MDReqRejReason = "7"
	MDReqRejReason_UNSUPPORTED_MDENTRYTYPE             MDReqRejReason = "8"
)

// MDUpdateAction field enumeration values.
type MDUpdateAction string

const (
	MDUpdateAction_NEW    MDUpdateAction = "0"
	MDUpdateAction_CHANGE MDUpdateAction = "1"
	MDUpdateAction_DELETE MDUpdateAction = "2"
)

// MDUpdateType field enumeration values.
type MDUpdateType string

const (
	MDUpdateType_FULL_REFRESH        MDUpdateType = "0"
	MDUpdateType_INCREMENTAL_REFRESH MDUpdateType = "1"
)

// MessageEncoding field enumeration values.
type MessageEncoding string

const (
	MessageEncoding_EUC_JP      MessageEncoding = "EUC-JP"
	MessageEncoding_ISO_2022_JP MessageEncoding = "ISO-2022-JP"
	MessageEncoding_SHIFT_JIS   MessageEncoding = "SHIFT_JIS"
	MessageEncoding_UTF_8       MessageEncoding = "UTF-8"
)

// MiscFeeType field enumeration values.
type MiscFeeType string

const (
	MiscFeeType_REGULATORY       MiscFeeType = "1"
	MiscFeeType_TAX              MiscFeeType = "2"
	MiscFeeType_LOCAL_COMMISSION MiscFeeType = "3"
	MiscFeeType_EXCHANGE_FEES    MiscFeeType = "4"
	MiscFeeType_STAMP            MiscFeeType = "5"
	MiscFeeType_LEVY             MiscFeeType = "6"
	MiscFeeType_OTHER            MiscFeeType = "7"
	MiscFeeType_MARKUP           MiscFeeType = "8"
	MiscFeeType_CONSUMPTION_TAX  MiscFeeType = "9"
)

// MsgDirection field enumeration values.
type MsgDirection string

const (
	MsgDirection_RECEIVE MsgDirection = "R"
	MsgDirection_SEND    MsgDirection = "S"
)

// MsgType field enumeration values.
type MsgType string

const (
	MsgType_HEARTBEAT                         MsgType = "0"
	MsgType_TEST_REQUEST                      MsgType = "1"
	MsgType_RESEND_REQUEST                    MsgType = "2"
	MsgType_REJECT                            MsgType = "3"
	MsgType_SEQUENCE_RESET                    MsgType = "4"
	MsgType_LOGOUT                            MsgType = "5"
	MsgType_INDICATION_OF_INTEREST            MsgType = "6"
	MsgType_ADVERTISEMENT                     MsgType = "7"
	MsgType_EXECUTION_REPORT                  MsgType = "8"
	MsgType_ORDER_CANCEL_REJECT               MsgType = "9"
	MsgType_LOGON                             MsgType = "A"
	MsgType_NEWS                              MsgType = "B"
	MsgType_EMAIL                             MsgType = "C"
	MsgType_ORDER_SINGLE                      MsgType = "D"
	MsgType_ORDER_LIST                        MsgType = "E"
	MsgType_ORDER_CANCEL_REQUEST              MsgType = "F"
	MsgType_ORDER_CANCEL_REPLACE_REQUEST      MsgType = "G"
	MsgType_ORDER_STATUS_REQUEST              MsgType = "H"
	MsgType_ALLOCATION                        MsgType = "J"
	MsgType_LIST_CANCEL_REQUEST               MsgType = "K"
	MsgType_LIST_EXECUTE                      MsgType = "L"
	MsgType_LIST_STATUS_REQUEST               MsgType = "M"
	MsgType_LIST_STATUS                       MsgType = "N"
	MsgType_ALLOCATION_ACK                    MsgType = "P"
	MsgType_DONT_KNOW_TRADE                   MsgType = "Q"
	MsgType_QUOTE_REQUEST                     MsgType = "R"
	MsgType_QUOTE                             MsgType = "S"
	MsgType_SETTLEMENT_INSTRUCTIONS           MsgType = "T"
	MsgType_ORDER_CANCEL_BATCH_REQUEST        MsgType = "U4"
	MsgType_ORDER_CANCEL_BATCH_REJECT         MsgType = "U5"
	MsgType_ORDER_BATCH                       MsgType = "U6"
	MsgType_ORDER_BATCH_REJECT                MsgType = "U7"
	MsgType_MARKET_DATA_REQUEST               MsgType = "V"
	MsgType_MARKET_DATA_SNAPSHOT_FULL_REFRESH MsgType = "W"
	MsgType_MARKET_DATA_INCREMENTAL_REFRESH   MsgType = "X"
	MsgType_MARKET_DATA_REQUEST_REJECT        MsgType = "Y"
	MsgType_QUOTE_CANCEL                      MsgType = "Z"
	MsgType_QUOTE_STATUS_REQUEST              MsgType = "a"
	MsgType_QUOTE_ACKNOWLEDGEMENT             MsgType = "b"
	MsgType_SECURITY_DEFINITION_REQUEST       MsgType = "c"
	MsgType_SECURITY_DEFINITION               MsgType = "d"
	MsgType_SECURITY_STATUS_REQUEST           MsgType = "e"
	MsgType_SECURITY_STATUS                   MsgType = "f"
	MsgType_TRADING_SESSION_STATUS_REQUEST    MsgType = "g"
	MsgType_TRADING_SESSION_STATUS            MsgType = "h"
	MsgType_MASS_QUOTE                        MsgType = "i"
	MsgType_BUSINESS_MESSAGE_REJECT           MsgType = "j"
	MsgType_BID_REQUEST                       MsgType = "k"
	MsgType_BID_RESPONSE                      MsgType = "l"
	MsgType_LIST_STRIKE_PRICE                 MsgType = "m"
)

// MultiLegReportingType field enumeration values.
type MultiLegReportingType string

const (
	MultiLegReportingType_SINGLE_SECURITY                        MultiLegReportingType = "1"
	MultiLegReportingType_INDIVIDUAL_LEG_OF_A_MULTI_LEG_SECURITY MultiLegReportingType = "2"
	MultiLegReportingType_MULTI_LEG_SECURITY                     MultiLegReportingType = "3"
)

// NetGrossInd field enumeration values.
type NetGrossInd string

const (
	NetGrossInd_NET   NetGrossInd = "1"
	NetGrossInd_GROSS NetGrossInd = "2"
)

// NotifyBrokerOfCredit field enumeration values.
type NotifyBrokerOfCredit string

const (
	NotifyBrokerOfCredit_NO  NotifyBrokerOfCredit = "N"
	NotifyBrokerOfCredit_YES NotifyBrokerOfCredit = "Y"
)

// OpenClose field enumeration values.
type OpenClose string

const (
	OpenClose_CLOSE OpenClose = "C"
	OpenClose_OPEN  OpenClose = "O"
)

// OpenCloseSettleFlag field enumeration values.
type OpenCloseSettleFlag string

const (
	OpenCloseSettleFlag_DAILY_OPEN                OpenCloseSettleFlag = "0"
	OpenCloseSettleFlag_SESSION_OPEN              OpenCloseSettleFlag = "1"
	OpenCloseSettleFlag_DELIVERY_SETTLEMENT_PRICE OpenCloseSettleFlag = "2"
)

// OrdRejReason field enumeration values.
type OrdRejReason string

const (
	OrdRejReason_BROKER_OPTION                              OrdRejReason = "0"
	OrdRejReason_UNKNOWN_SYMBOL                             OrdRejReason = "1"
	OrdRejReason_EXCHANGE_CLOSED                            OrdRejReason = "2"
	OrdRejReason_ORDER_EXCEEDS_LIMIT                        OrdRejReason = "3"
	OrdRejReason_TOO_LATE_TO_ENTER                          OrdRejReason = "4"
	OrdRejReason_UNKNOWN_ORDER                              OrdRejReason = "5"
	OrdRejReason_DUPLICATE_ORDER                            OrdRejReason = "6"
	OrdRejReason_DUPLICATE_OF_A_VERBALLY_COMMUNICATED_ORDER OrdRejReason = "7"
	OrdRejReason_STALE_ORDER                                OrdRejReason = "8"
)

// OrdStatus field enumeration values.
type OrdStatus string

const (
	OrdStatus_NEW                  OrdStatus = "0"
	OrdStatus_PARTIALLY_FILLED     OrdStatus = "1"
	OrdStatus_FILLED               OrdStatus = "2"
	OrdStatus_DONE_FOR_DAY         OrdStatus = "3"
	OrdStatus_CANCELED             OrdStatus = "4"
	OrdStatus_REPLACED             OrdStatus = "5"
	OrdStatus_PENDING_CANCEL       OrdStatus = "6"
	OrdStatus_STOPPED              OrdStatus = "7"
	OrdStatus_REJECTED             OrdStatus = "8"
	OrdStatus_SUSPENDED            OrdStatus = "9"
	OrdStatus_PENDING_NEW          OrdStatus = "A"
	OrdStatus_CALCULATED           OrdStatus = "B"
	OrdStatus_EXPIRED              OrdStatus = "C"
	OrdStatus_ACCEPTED_FOR_BIDDING OrdStatus = "D"
	OrdStatus_PENDING_REPLACE      OrdStatus = "E"
)

// OrdType field enumeration values.
type OrdType string

const (
	OrdType_MARKET                OrdType = "1"
	OrdType_LIMIT                 OrdType = "2"
	OrdType_STOP                  OrdType = "3"
	OrdType_STOP_LIMIT            OrdType = "4"
	OrdType_MARKET_ON_CLOSE       OrdType = "5"
	OrdType_WITH_OR_WITHOUT       OrdType = "6"
	OrdType_LIMIT_OR_BETTER       OrdType = "7"
	OrdType_LIMIT_WITH_OR_WITHOUT OrdType = "8"
	OrdType_ON_BASIS              OrdType = "9"
	OrdType_ON_CLOSE              OrdType = "A"
	OrdType_LIMIT_ON_CLOSE        OrdType = "B"
	OrdType_FOREX_C               OrdType = "C"
	OrdType_PREVIOUSLY_QUOTED     OrdType = "D"
	OrdType_PREVIOUSLY_INDICATED  OrdType = "E"
	OrdType_FOREX_F               OrdType = "F"
	OrdType_FOREX_G               OrdType = "G"
	OrdType_FOREX_H               OrdType = "H"
	OrdType_FUNARI                OrdType = "I"
	OrdType_PEGGED                OrdType = "P"
)

// PossDupFlag field enumeration values.
type PossDupFlag string

const (
	PossDupFlag_NO  PossDupFlag = "N"
	PossDupFlag_YES PossDupFlag = "Y"
)

// PossResend field enumeration values.
type PossResend string

const (
	PossResend_NO  PossResend = "N"
	PossResend_YES PossResend = "Y"
)

// PriceType field enumeration values.
type PriceType string

const (
	PriceType_PERCENTAGE   PriceType = "1"
	PriceType_PER_SHARE    PriceType = "2"
	PriceType_FIXED_AMOUNT PriceType = "3"
)

// ProcessCode field enumeration values.
type ProcessCode string

const (
	ProcessCode_REGULAR              ProcessCode = "0"
	ProcessCode_SOFT_DOLLAR          ProcessCode = "1"
	ProcessCode_STEP_IN              ProcessCode = "2"
	ProcessCode_STEP_OUT             ProcessCode = "3"
	ProcessCode_SOFT_DOLLAR_STEP_IN  ProcessCode = "4"
	ProcessCode_SOFT_DOLLAR_STEP_OUT ProcessCode = "5"
	ProcessCode_PLAN_SPONSOR         ProcessCode = "6"
)

// ProgRptReqs field enumeration values.
type ProgRptReqs string

const (
	ProgRptReqs_BUYSIDE_EXPLICITLY_REQUESTS_STATUS_USING_STATUSREQUEST                                            ProgRptReqs = "1"
	ProgRptReqs_SELLSIDE_PERIODICALLY_SENDS_STATUS_USING_LISTSTATUS_PERIOD_OPTIONALLY_SPECIFIED_IN_PROGRESSPERIOD ProgRptReqs = "2"
	ProgRptReqs_REAL_TIME_EXECUTION_REPORTS                                                                       ProgRptReqs = "3"
)

// PutOrCall field enumeration values.
type PutOrCall string

const (
	PutOrCall_PUT  PutOrCall = "0"
	PutOrCall_CALL PutOrCall = "1"
)

// QuoteAckStatus field enumeration values.
type QuoteAckStatus string

const (
	QuoteAckStatus_ACCEPTED                   QuoteAckStatus = "0"
	QuoteAckStatus_CANCELED_FOR_SYMBOL        QuoteAckStatus = "1"
	QuoteAckStatus_CANCELED_FOR_SECURITY_TYPE QuoteAckStatus = "2"
	QuoteAckStatus_CANCELED_FOR_UNDERLYING    QuoteAckStatus = "3"
	QuoteAckStatus_CANCELED_ALL               QuoteAckStatus = "4"
	QuoteAckStatus_REJECTED                   QuoteAckStatus = "5"
)

// QuoteCancelType field enumeration values.
type QuoteCancelType string

const (
	QuoteCancelType_CANCEL_FOR_SYMBOL            QuoteCancelType = "1"
	QuoteCancelType_CANCEL_FOR_SECURITY_TYPE     QuoteCancelType = "2"
	QuoteCancelType_CANCEL_FOR_UNDERLYING_SYMBOL QuoteCancelType = "3"
	QuoteCancelType_CANCEL_FOR_ALL_QUOTES        QuoteCancelType = "4"
)

// QuoteCondition field enumeration values.
type QuoteCondition string

const (
	QuoteCondition_OPEN              QuoteCondition = "A"
	QuoteCondition_CLOSED            QuoteCondition = "B"
	QuoteCondition_EXCHANGE_BEST     QuoteCondition = "C"
	QuoteCondition_CONSOLIDATED_BEST QuoteCondition = "D"
	QuoteCondition_LOCKED            QuoteCondition = "E"
	QuoteCondition_CROSSED           QuoteCondition = "F"
	QuoteCondition_DEPTH             QuoteCondition = "G"
	QuoteCondition_FAST_TRADING      QuoteCondition = "H"
	QuoteCondition_NON_FIRM          QuoteCondition = "I"
)

// QuoteEntryRejectReason field enumeration values.
type QuoteEntryRejectReason string

const (
	QuoteEntryRejectReason_UNKNOWN_SYMBOL                   QuoteEntryRejectReason = "1"
	QuoteEntryRejectReason_EXCHANGE                         QuoteEntryRejectReason = "2"
	QuoteEntryRejectReason_QUOTE_EXCEEDS_LIMIT              QuoteEntryRejectReason = "3"
	QuoteEntryRejectReason_TOO_LATE_TO_ENTER                QuoteEntryRejectReason = "4"
	QuoteEntryRejectReason_UNKNOWN_QUOTE                    QuoteEntryRejectReason = "5"
	QuoteEntryRejectReason_DUPLICATE_QUOTE                  QuoteEntryRejectReason = "6"
	QuoteEntryRejectReason_INVALID_BID_ASK_SPREAD           QuoteEntryRejectReason = "7"
	QuoteEntryRejectReason_INVALID_PRICE                    QuoteEntryRejectReason = "8"
	QuoteEntryRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY QuoteEntryRejectReason = "9"
)

// QuoteRejectReason field enumeration values.
type QuoteRejectReason string

const (
	QuoteRejectReason_UNKNOWN_SYMBOL                   QuoteRejectReason = "1"
	QuoteRejectReason_EXCHANGE                         QuoteRejectReason = "2"
	QuoteRejectReason_QUOTE_REQUEST_EXCEEDS_LIMIT      QuoteRejectReason = "3"
	QuoteRejectReason_TOO_LATE_TO_ENTER                QuoteRejectReason = "4"
	QuoteRejectReason_UNKNOWN_QUOTE                    QuoteRejectReason = "5"
	QuoteRejectReason_DUPLICATE_QUOTE                  QuoteRejectReason = "6"
	QuoteRejectReason_INVALID_BID_ASK_SPREAD           QuoteRejectReason = "7"
	QuoteRejectReason_INVALID_PRICE                    QuoteRejectReason = "8"
	QuoteRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY QuoteRejectReason = "9"
)

// QuoteRequestType field enumeration values.
type QuoteRequestType string

const (
	QuoteRequestType_MANUAL    QuoteRequestType = "1"
	QuoteRequestType_AUTOMATIC QuoteRequestType = "2"
)

// QuoteResponseLevel field enumeration values.
type QuoteResponseLevel string

const (
	QuoteResponseLevel_NO_ACKNOWLEDGEMENT                            QuoteResponseLevel = "0"
	QuoteResponseLevel_ACKNOWLEDGE_ONLY_NEGATIVE_OR_ERRONEOUS_QUOTES QuoteResponseLevel = "1"
	QuoteResponseLevel_ACKNOWLEDGE_EACH_QUOTE_MESSAGES               QuoteResponseLevel = "2"
)

// ReportToExch field enumeration values.
type ReportToExch string

const (
	ReportToExch_NO  ReportToExch = "N"
	ReportToExch_YES ReportToExch = "Y"
)

// ResetSeqNumFlag field enumeration values.
type ResetSeqNumFlag string

const (
	ResetSeqNumFlag_NO  ResetSeqNumFlag = "N"
	ResetSeqNumFlag_YES ResetSeqNumFlag = "Y"
)

// RoutingType field enumeration values.
type RoutingType string

const (
	RoutingType_TARGET_FIRM RoutingType = "1"
	RoutingType_TARGET_LIST RoutingType = "2"
	RoutingType_BLOCK_FIRM  RoutingType = "3"
	RoutingType_BLOCK_LIST  RoutingType = "4"
)

// Rule80A field enumeration values.
type Rule80A string

const (
	Rule80A_AGENCY_SINGLE_ORDER                                                                                        Rule80A = "A"
	Rule80A_SHORT_EXEMPT_TRANSACTION_B                                                                                 Rule80A = "B"
	Rule80A_PROGRAM_ORDER_NON_INDEX_ARB_FOR_MEMBER_FIRM_ORG                                                            Rule80A = "C"
	Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_MEMBER_FIRM_ORG                                                                Rule80A = "D"
	Rule80A_REGISTERED_EQUITY_MARKET_MAKER_TRADES                                                                      Rule80A = "E"
	Rule80A_SHORT_EXEMPT_TRANSACTION_F                                                                                 Rule80A = "F"
	Rule80A_SHORT_EXEMPT_TRANSACTION_H                                                                                 Rule80A = "H"
	Rule80A_INDIVIDUAL_INVESTOR_SINGLE_ORDER                                                                           Rule80A = "I"
	Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_INDIVIDUAL_CUSTOMER                                                            Rule80A = "J"
	Rule80A_PROGRAM_ORDER_NON_INDEX_ARB_FOR_INDIVIDUAL_CUSTOMER                                                        Rule80A = "K"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_AFFILIATED_WITH_THE_FIRM_CLEARING_THE_TRADE     Rule80A = "L"
	Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_OTHER_MEMBER                                                                   Rule80A = "M"
	Rule80A_PROGRAM_ORDER_NON_INDEX_ARB_FOR_OTHER_MEMBER                                                               Rule80A = "N"
	Rule80A_COMPETING_DEALER_TRADES_O                                                                                  Rule80A = "O"
	Rule80A_PRINCIPAL                                                                                                  Rule80A = "P"
	Rule80A_COMPETING_DEALER_TRADES_R                                                                                  Rule80A = "R"
	Rule80A_SPECIALIST_TRADES                                                                                          Rule80A = "S"
	Rule80A_COMPETING_DEALER_TRADES_T                                                                                  Rule80A = "T"
	Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_OTHER_AGENCY                                                                   Rule80A = "U"
	Rule80A_ALL_OTHER_ORDERS_AS_AGENT_FOR_OTHER_MEMBER                                                                 Rule80A = "W"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_NOT_AFFILIATED_WITH_THE_FIRM_CLEARING_THE_TRADE Rule80A = "X"
	Rule80A_PROGRAM_ORDER_NON_INDEX_ARB_FOR_OTHER_AGENCY                                                               Rule80A = "Y"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_NON_MEMBER_COMPETING_MARKET_MAKER                                             Rule80A = "Z"
)

// SecurityRequestType field enumeration values.
type SecurityRequestType string

const (
	SecurityRequestType_REQUEST_SECURITY_IDENTITY_AND_SPECIFICATIONS              SecurityRequestType = "0"
	SecurityRequestType_REQUEST_SECURITY_IDENTITY_FOR_THE_SPECIFICATIONS_PROVIDED SecurityRequestType = "1"
	SecurityRequestType_REQUEST_LIST_SECURITY_TYPES                               SecurityRequestType = "2"
	SecurityRequestType_REQUEST_LIST_SECURITIES                                   SecurityRequestType = "3"
)

// SecurityResponseType field enumeration values.
type SecurityResponseType string

const (
	SecurityResponseType_ACCEPT_SECURITY_PROPOSAL_AS_IS                                      SecurityResponseType = "1"
	SecurityResponseType_ACCEPT_SECURITY_PROPOSAL_WITH_REVISIONS_AS_INDICATED_IN_THE_MESSAGE SecurityResponseType = "2"
	SecurityResponseType_LIST_OF_SECURITY_TYPES_RETURNED_PER_REQUEST                         SecurityResponseType = "3"
	SecurityResponseType_LIST_OF_SECURITIES_RETURNED_PER_REQUEST                             SecurityResponseType = "4"
	SecurityResponseType_REJECT_SECURITY_PROPOSAL                                            SecurityResponseType = "5"
	SecurityResponseType_CAN_NOT_MATCH_SELECTION_CRITERIA                                    SecurityResponseType = "6"
)

// SecurityTradingStatus field enumeration values.
type SecurityTradingStatus string

const (
	SecurityTradingStatus_OPENING_DELAY                  SecurityTradingStatus = "1"
	SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_SELL SecurityTradingStatus = "10"
	SecurityTradingStatus_NO_MARKET_IMBALANCE            SecurityTradingStatus = "12"
	SecurityTradingStatus_NO_MARKET_ON_CLOSE_IMBALANCE   SecurityTradingStatus = "13"
	SecurityTradingStatus_ITS_PRE_OPENING                SecurityTradingStatus = "14"
	SecurityTradingStatus_NEW_PRICE_INDICATION           SecurityTradingStatus = "15"
	SecurityTradingStatus_TRADE_DISSEMINATION_TIME       SecurityTradingStatus = "16"
	SecurityTradingStatus_READY_TO_TRADE                 SecurityTradingStatus = "17"
	SecurityTradingStatus_NOT_AVAILABLE_FOR_TRADING      SecurityTradingStatus = "18"
	SecurityTradingStatus_NOT_TRADED_ON_THIS_MARKET      SecurityTradingStatus = "19"
	SecurityTradingStatus_TRADING_HALT                   SecurityTradingStatus = "2"
	SecurityTradingStatus_UNKNOWN_OR_INVALID             SecurityTradingStatus = "20"
	SecurityTradingStatus_RESUME                         SecurityTradingStatus = "3"
	SecurityTradingStatus_NO_OPEN_NO_RESUME              SecurityTradingStatus = "4"
	SecurityTradingStatus_PRICE_INDICATION               SecurityTradingStatus = "5"
	SecurityTradingStatus_TRADING_RANGE_INDICATION       SecurityTradingStatus = "6"
	SecurityTradingStatus_MARKET_IMBALANCE_BUY           SecurityTradingStatus = "7"
	SecurityTradingStatus_MARKET_IMBALANCE_SELL          SecurityTradingStatus = "8"
	SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_BUY  SecurityTradingStatus = "9"
)

// SecurityType field enumeration values.
type SecurityType string

const (
	SecurityType_WILDCARD_ENTRY                           SecurityType = "?"
	SecurityType_BANKERS_ACCEPTANCE                       SecurityType = "BA"
	SecurityType_CONVERTIBLE_BOND                         SecurityType = "CB"
	SecurityType_CERTIFICATE_OF_DEPOSIT                   SecurityType = "CD"
	SecurityType_COLLATERALIZE_MORTGAGE_OBLIGATION        SecurityType = "CMO"
	SecurityType_CORPORATE_BOND                           SecurityType = "CORP"
	SecurityType_COMMERCIAL_PAPER                         SecurityType = "CP"
	SecurityType_CORPORATE_PRIVATE_PLACEMENT              SecurityType = "CPP"
	SecurityType_COMMON_STOCK                             SecurityType = "CS"
	SecurityType_FEDERAL_HOUSING_AUTHORITY                SecurityType = "FHA"
	SecurityType_FEDERAL_HOME_LOAN                        SecurityType = "FHL"
	SecurityType_FEDERAL_NATIONAL_MORTGAGE_ASSOCIATION    SecurityType = "FN"
	SecurityType_FOREIGN_EXCHANGE_CONTRACT                SecurityType = "FOR"
	SecurityType_FUTURE                                   SecurityType = "FUT"
	SecurityType_GOVERNMENT_NATIONAL_MORTGAGE_ASSOCIATION SecurityType = "GN"
	SecurityType_TREASURIES_PLUS_AGENCY_DEBENTURE         SecurityType = "GOVT"
	SecurityType_MORTGAGE_IOETTE                          SecurityType = "IET"
	SecurityType_MUTUAL_FUND                              SecurityType = "MF"
	SecurityType_MORTGAGE_INTEREST_ONLY                   SecurityType = "MIO"
	SecurityType_MORTGAGE_PRINCIPAL_ONLY                  SecurityType = "MPO"
	SecurityType_MORTGAGE_PRIVATE_PLACEMENT               SecurityType = "MPP"
	SecurityType_MISCELLANEOUS_PASS_THRU                  SecurityType = "MPT"
	SecurityType_MUNICIPAL_BOND                           SecurityType = "MUNI"
	SecurityType_NO_ISITC_SECURITY_TYPE                   SecurityType = "NONE"
	SecurityType_OPTION                                   SecurityType = "OPT"
	SecurityType_PREFERRED_STOCK                          SecurityType = "PS"
	SecurityType_REPURCHASE_AGREEMENT                     SecurityType = "RP"
	SecurityType_REVERSE_REPURCHASE_AGREEMENT             SecurityType = "RVRP"
	SecurityType_STUDENT_LOAN_MARKETING_ASSOCIATION       SecurityType = "SL"
	SecurityType_TIME_DEPOSIT                             SecurityType = "TD"
	SecurityType_US_TREASURY_BILL                         SecurityType = "USTB"
	SecurityType_WARRANT                                  SecurityType = "WAR"
	SecurityType_CATS_TIGERS_LIONS                        SecurityType = "ZOO"
)

// SelfTradePrevention field enumeration values.
type SelfTradePrevention string

const (
	SelfTradePrevention_CANCEL_BOTH          SelfTradePrevention = "B"
	SelfTradePrevention_DECREMENT_AND_CANCEL SelfTradePrevention = "D"
	SelfTradePrevention_CANCEL_NEWEST        SelfTradePrevention = "N"
	SelfTradePrevention_CANCEL_OLDEST        SelfTradePrevention = "O"
)

// SessionRejectReason field enumeration values.
type SessionRejectReason string

const (
	SessionRejectReason_INVALID_TAG_NUMBER                    SessionRejectReason = "0"
	SessionRejectReason_REQUIRED_TAG_MISSING                  SessionRejectReason = "1"
	SessionRejectReason_SENDINGTIME_ACCURACY_PROBLEM          SessionRejectReason = "10"
	SessionRejectReason_INVALID_MSGTYPE                       SessionRejectReason = "11"
	SessionRejectReason_TAG_NOT_DEFINED_FOR_THIS_MESSAGE_TYPE SessionRejectReason = "2"
	SessionRejectReason_UNDEFINED_TAG                         SessionRejectReason = "3"
	SessionRejectReason_TAG_SPECIFIED_WITHOUT_A_VALUE         SessionRejectReason = "4"
	SessionRejectReason_VALUE_IS_INCORRECT                    SessionRejectReason = "5"
	SessionRejectReason_INCORRECT_DATA_FORMAT_FOR_VALUE       SessionRejectReason = "6"
	SessionRejectReason_DECRYPTION_PROBLEM                    SessionRejectReason = "7"
	SessionRejectReason_SIGNATURE_PROBLEM                     SessionRejectReason = "8"
	SessionRejectReason_COMPID_PROBLEM                        SessionRejectReason = "9"
)

// SettlCurrFxRateCalc field enumeration values.
type SettlCurrFxRateCalc string

const (
	SettlCurrFxRateCalc_DIVIDE   SettlCurrFxRateCalc = "D"
	SettlCurrFxRateCalc_MULTIPLY SettlCurrFxRateCalc = "M"
)

// SettlInstMode field enumeration values.
type SettlInstMode string

const (
	SettlInstMode_DEFAULT                                SettlInstMode = "0"
	SettlInstMode_STANDING_INSTRUCTIONS_PROVIDED         SettlInstMode = "1"
	SettlInstMode_SPECIFIC_ALLOCATION_ACCOUNT_OVERRIDING SettlInstMode = "2"
	SettlInstMode_SPECIFIC_ALLOCATION_ACCOUNT_STANDING   SettlInstMode = "3"
)

// SettlInstSource field enumeration values.
type SettlInstSource string

const (
	SettlInstSource_BROKERS_INSTRUCTIONS      SettlInstSource = "1"
	SettlInstSource_INSTITUTIONS_INSTRUCTIONS SettlInstSource = "2"
)

// SettlInstTransType field enumeration values.
type SettlInstTransType string

const (
	SettlInstTransType_CANCEL  SettlInstTransType = "C"
	SettlInstTransType_NEW     SettlInstTransType = "N"
	SettlInstTransType_REPLACE SettlInstTransType = "R"
)

// SettlLocation field enumeration values.
type SettlLocation string

const (
	SettlLocation_CEDEL                        SettlLocation = "CED"
	SettlLocation_DEPOSITORY_TRUST_COMPANY     SettlLocation = "DTC"
	SettlLocation_EUROCLEAR                    SettlLocation = "EUR"
	SettlLocation_FEDERAL_BOOK_ENTRY           SettlLocation = "FED"
	SettlLocation_LOCAL_MARKET_SETTLE_LOCATION SettlLocation = "ISO Country Code"
	SettlLocation_PHYSICAL                     SettlLocation = "PNY"
	SettlLocation_PARTICIPANT_TRUST_COMPANY    SettlLocation = "PTC"
)

// SettlmntTyp field enumeration values.
type SettlmntTyp string

const (
	SettlmntTyp_REGULAR        SettlmntTyp = "0"
	SettlmntTyp_CASH           SettlmntTyp = "1"
	SettlmntTyp_NEXT_DAY       SettlmntTyp = "2"
	SettlmntTyp_T_PLUS_2       SettlmntTyp = "3"
	SettlmntTyp_T_PLUS_3       SettlmntTyp = "4"
	SettlmntTyp_T_PLUS_4       SettlmntTyp = "5"
	SettlmntTyp_FUTURE         SettlmntTyp = "6"
	SettlmntTyp_WHEN_ISSUED    SettlmntTyp = "7"
	SettlmntTyp_SELLERS_OPTION SettlmntTyp = "8"
	SettlmntTyp_T_PLUS_5       SettlmntTyp = "9"
)

// Side field enumeration values.
type Side string

const (
	Side_BUY               Side = "1"
	Side_SELL              Side = "2"
	Side_BUY_MINUS         Side = "3"
	Side_SELL_PLUS         Side = "4"
	Side_SELL_SHORT        Side = "5"
	Side_SELL_SHORT_EXEMPT Side = "6"
	Side_UNDISCLOSED       Side = "7"
	Side_CROSS             Side = "8"
	Side_CROSS_SHORT       Side = "9"
)

// SolicitedFlag field enumeration values.
type SolicitedFlag string

const (
	SolicitedFlag_NO  SolicitedFlag = "N"
	SolicitedFlag_YES SolicitedFlag = "Y"
)

// StandInstDbType field enumeration values.
type StandInstDbType string

const (
	StandInstDbType_OTHER              StandInstDbType = "0"
	StandInstDbType_DTC_SID            StandInstDbType = "1"
	StandInstDbType_THOMSON_ALERT      StandInstDbType = "2"
	StandInstDbType_A_GLOBAL_CUSTODIAN StandInstDbType = "3"
)

// SubscriptionRequestType field enumeration values.
type SubscriptionRequestType string

const (
	SubscriptionRequestType_SNAPSHOT                                      SubscriptionRequestType = "0"
	SubscriptionRequestType_SNAPSHOT_PLUS_UPDATES                         SubscriptionRequestType = "1"
	SubscriptionRequestType_DISABLE_PREVIOUS_SNAPSHOT_PLUS_UPDATE_REQUEST SubscriptionRequestType = "2"
)

// TickDirection field enumeration values.
type TickDirection string

const (
	TickDirection_PLUS_TICK       TickDirection = "0"
	TickDirection_ZERO_PLUS_TICK  TickDirection = "1"
	TickDirection_MINUS_TICK      TickDirection = "2"
	TickDirection_ZERO_MINUS_TICK TickDirection = "3"
)

// TimeInForce field enumeration values.
type TimeInForce string

const (
	TimeInForce_DAY                 TimeInForce = "0"
	TimeInForce_GOOD_TILL_CANCEL    TimeInForce = "1"
	TimeInForce_AT_THE_OPENING      TimeInForce = "2"
	TimeInForce_IMMEDIATE_OR_CANCEL TimeInForce = "3"
	TimeInForce_FILL_OR_KILL        TimeInForce = "4"
	TimeInForce_GOOD_TILL_CROSSING  TimeInForce = "5"
	TimeInForce_GOOD_TILL_DATE      TimeInForce = "6"
	TimeInForce_POST_ONLY           TimeInForce = "P"
)

// TradSesMethod field enumeration values.
type TradSesMethod string

const (
	TradSesMethod_ELECTRONIC  TradSesMethod = "1"
	TradSesMethod_OPEN_OUTCRY TradSesMethod = "2"
	TradSesMethod_TWO_PARTY   TradSesMethod = "3"
)

// TradSesMode field enumeration values.
type TradSesMode string

const (
	TradSesMode_TESTING    TradSesMode = "1"
	TradSesMode_SIMULATED  TradSesMode = "2"
	TradSesMode_PRODUCTION TradSesMode = "3"
)

// TradSesStatus field enumeration values.
type TradSesStatus string

const (
	TradSesStatus_HALTED    TradSesStatus = "1"
	TradSesStatus_OPEN      TradSesStatus = "2"
	TradSesStatus_CLOSED    TradSesStatus = "3"
	TradSesStatus_PRE_OPEN  TradSesStatus = "4"
	TradSesStatus_PRE_CLOSE TradSesStatus = "5"
)

// TradeCondition field enumeration values.
type TradeCondition string

const (
	TradeCondition_CASH                  TradeCondition = "A"
	TradeCondition_AVERAGE_PRICE_TRADE   TradeCondition = "B"
	TradeCondition_CASH_TRADE            TradeCondition = "C"
	TradeCondition_NEXT_DAY              TradeCondition = "D"
	TradeCondition_OPENING               TradeCondition = "E"
	TradeCondition_INTRADAY_TRADE_DETAIL TradeCondition = "F"
	TradeCondition_RULE_127_TRADE        TradeCondition = "G"
	TradeCondition_RULE_155_TRADE        TradeCondition = "H"
	TradeCondition_SOLD_LAST             TradeCondition = "I"
	TradeCondition_NEXT_DAY_TRADE        TradeCondition = "J"
	TradeCondition_OPENED                TradeCondition = "K"
	TradeCondition_SELLER                TradeCondition = "L"
	TradeCondition_SOLD                  TradeCondition = "M"
	TradeCondition_STOPPED_STOCK         TradeCondition = "N"
)

// TradeType field enumeration values.
type TradeType string

const (
	TradeType_AGENCY           TradeType = "A"
	TradeType_VWAP_GUARANTEE   TradeType = "G"
	TradeType_GUARANTEED_CLOSE TradeType = "J"
	TradeType_RISK_TRADE       TradeType = "R"
)

// TriggerPriceDirection field enumeration values.
type TriggerPriceDirection string

const (
	TriggerPriceDirection_DOWN TriggerPriceDirection = "D"
	TriggerPriceDirection_UP   TriggerPriceDirection = "U"
)

// UnsolicitedIndicator field enumeration values.
type UnsolicitedIndicator string

const (
	UnsolicitedIndicator_NO  UnsolicitedIndicator = "N"
	UnsolicitedIndicator_YES UnsolicitedIndicator = "Y"
)

// Urgency field enumeration values.
type Urgency string

const (
	Urgency_NORMAL     Urgency = "0"
	Urgency_FLASH      Urgency = "1"
	Urgency_BACKGROUND Urgency = "2"
)
