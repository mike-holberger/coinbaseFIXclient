package main

// Side field enumeration values.
type Side string

const (
	Side_BUY  Side = "1"
	Side_SELL Side = "2"
)

// OrdType field enumeration values.
type OrdType string

const (
	OrdType_MARKET     OrdType = "1"
	OrdType_LIMIT      OrdType = "2"
	OrdType_STOP_LIMIT OrdType = "4"
)
