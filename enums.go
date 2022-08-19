package main

import "github.com/quickfixgo/enum"

type OrderSide string

const (
	OrderSide_BUY  OrderSide = "1"
	OrderSide_SELL OrderSide = "2"
)

func (e OrderSide) getQFenum() (side enum.Side) {
	switch e {
	case OrderSide_BUY:
		return enum.Side_BUY
	case OrderSide_SELL:
		return enum.Side_SELL
	}

	panic("OrderSide enum did not translate to quickfix enum")
}

type OrderType string

const (
	OrderType_MARKET OrderType = "1"
	OrderType_LIMIT  OrderType = "2"
	OrderType_STOP   OrderType = "4"
)

func (e OrderType) getQFenum() (side enum.OrdType) {
	switch e {
	case OrderType_MARKET:
		return enum.OrdType_MARKET
	case OrderType_LIMIT:
		return enum.OrdType_LIMIT
	case OrderType_STOP:
		return enum.OrdType_STOP_LIMIT
	}

	panic("OrderType enum did not translate to quickfix enum")
}

type OrderTimeInForce string

const (
	OrderTimeInForce_GOOD_TILL_CANCEL    OrderTimeInForce = "1"
	OrderTimeInForce_IMMEDIATE_OR_CANCEL OrderTimeInForce = "3"
	OrderTimeInForce_FILL_OR_KILL        OrderTimeInForce = "4"
	OrderTimeInForce_POST_ONLY           OrderTimeInForce = "P"
)

type OrderSelfTradePrevention string

const (
	OrderTimeInForce_DECREMENT_AND_CANCEL OrderSelfTradePrevention = "D"
	OrderTimeInForce_CANCEL_RESTING       OrderSelfTradePrevention = "O"
	OrderTimeInForce_CANCEL_INCOMING      OrderSelfTradePrevention = "N"
	OrderTimeInForce_CANCEL_BOTH          OrderSelfTradePrevention = "B"
)
