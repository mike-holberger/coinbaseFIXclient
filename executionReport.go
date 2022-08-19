package main

import "strings"

type ExecutionReport struct {
	// 6 - Calculated average price of all fills on this order
	AvgPx string
	// 11 - Only present on order acknowledgements, ExecType=New (150=0)
	ClOrdID string
	// 14 - Currently executed quantity for chain of orders
	CumQty string
	// 37 -	OrderID from the ExecutionReport with ExecType=New (150=0)
	OrderID string
	// 55 -	Symbol of the original order
	Symbol string
	// 54 -	Must be 1 to buy or 2 to sell
	Side string
	// 32 -	Amount filled (if ExecType=1). Also called LastQty as of FIX 4.3
	LastShares string
	// 44 -	Price of the fill if ExecType indicates a fill, otherwise the order price
	Price string
	// 38 -	OrderQty as accepted (may be less than requested upon self-trade prevention)
	OrderQty string
	// 151 - Quantity open for further execution
	LeavesQty string
	// 152 - Order size in quote units (e.g., USD) (Market order only)
	CashOrderQty string
	// 60 - Time the event occurred
	TransactTime string
	// 150 - May be "1" (Partial fill) for fills, D for self-trade prevention, etc.
	ExecType string
	// 39 - Order status as of the current message
	OrdStatus string
	// 103 - Insufficient funds=3, Post-only=8, Unknown error=0
	OrdRejReason string
	// 422 - "1" (Order Status Request responses and fill reports)
	NoMiscFees string
	// 137 - Fee amount (absolute value for Order Status Request responses, percentage value for fill reports)
	MiscFeeAmt string
	// 138 - Fee currency
	MiscFeeCurr string
	// 139 - "4" (Exchange fees) (Order Status Request responses and fill reports)
	MiscFeeType string
	// 891 - "2" (Percentage fee basis) (fill report only)
	MiscFeeBasis string
	// 1003 - Product unique trade id
	TradeID string
	// 1057 - Y for taker orders, N for maker orders
	AggressorIndicator string
	// 58 - Text
	Text string
}

func (e CoinbaseFIXclient) UnmarshalExecReport(message string) (execReport ExecutionReport) {
	fields := strings.Split(message, "\x01")
	for _, f := range fields {
		fieldVal := strings.Split(f, "=")
		switch fieldVal[0] {
		case "6":
			execReport.AvgPx = fieldVal[1]
			continue
		case "11":
			execReport.ClOrdID = fieldVal[1]
			continue
		case "14":
			execReport.CumQty = fieldVal[1]
			continue
		case "37":
			execReport.OrderID = fieldVal[1]
			continue
		case "55":
			execReport.Symbol = fieldVal[1]
			continue
		case "54":
			execReport.Side = fieldVal[1]
			continue
		case "32":
			execReport.LastShares = fieldVal[1]
			continue
		case "44":
			execReport.Price = fieldVal[1]
			continue
		case "38":
			execReport.OrderQty = fieldVal[1]
			continue
		case "151":
			execReport.LeavesQty = fieldVal[1]
			continue
		case "152":
			execReport.CashOrderQty = fieldVal[1]
			continue
		case "60":
			execReport.TransactTime = fieldVal[1]
			continue
		case "150":
			execReport.ExecType = fieldVal[1]
			continue
		case "39":
			execReport.OrdStatus = fieldVal[1]
			continue
		case "103":
			execReport.OrdRejReason = fieldVal[1]
			continue
		case "136":
			execReport.NoMiscFees = fieldVal[1]
			continue
		case "137":
			execReport.MiscFeeAmt = fieldVal[1]
			continue
		case "138":
			execReport.MiscFeeCurr = fieldVal[1]
			continue
		case "139":
			execReport.MiscFeeType = fieldVal[1]
			continue
		case "891":
			execReport.MiscFeeBasis = fieldVal[1]
			continue
		case "1003":
			execReport.TradeID = fieldVal[1]
			continue
		case "1057":
			execReport.AggressorIndicator = fieldVal[1]
			continue
		case "58":
			execReport.Text = fieldVal[1]
			continue
		}
	}

	return
}
