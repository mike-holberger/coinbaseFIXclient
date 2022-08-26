package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	logtimeFormat := "2006-01-02 15:04:05.000"
	zerolog.TimeFieldFormat = logtimeFormat
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: logtimeFormat})
}

func main() {
	key := os.Getenv("COINBASE_PRO_KEY")
	if key == "" {
		log.Fatal().Msgf("No COINBASE_PRO_KEY in env vars")
	}
	secret := os.Getenv("COINBASE_PRO_SECRET")
	if secret == "" {
		log.Fatal().Msgf("No COINBASE_PRO_SECRET in env vars")
	}
	pass := os.Getenv("COINBASE_PRO_PASSPHRASE")
	if pass == "" {
		log.Fatal().Msgf("No COINBASE_PRO_PASSPHRASE in env vars")
	}

	var cbFIXclient CoinbaseFIXclient

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := cbFIXclient.Logon(key, secret, pass, ctx)
	if err != nil {
		log.Fatal().Err(err).Msgf("Logon Error")
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	execReport, err := cbFIXclient.NewOrderSingleLIMIT(CoinbaseFIXorderLIMIT{
		ClientID: "509e971a-1e70-11ed-861d-0242ac120001",
		Symbol:   "ETH-USD",
		Side:     Side_BUY,
		Price:    "1602",
		Qty:      "0.02",
	}, true, ctx)
	if err != nil {
		log.Error().Err(err).Msgf("NewOrderSingle Error")
	} else {
		log.Info().Interface("ExecReport", execReport).Send()
	}

	execReports, err := cbFIXclient.NewOrdersBatch("BATCH-TEST", []CoinbaseFIXorderLIMIT{
		{
			ClientID: "509e971a-1e70-11ed-861d-0242ac120002",
			Symbol:   "ETH-USD",
			Side:     Side_BUY,
			Price:    "1550",
			Qty:      "0.02",
		},
		{
			ClientID: "509e971a-1e70-11ed-861d-0242ac120055",
			Symbol:   "ETH-USD",
			Side:     Side_BUY,
			Price:    "1600",
			Qty:      "0.02",
		},
	}, true, ctx)
	if err != nil {
		log.Error().Err(err).Msgf("NewOrdersBatch Error")
	} else {
		log.Info().Interface("ExecReports", execReports).Send()
	}

	err = cbFIXclient.OrderCancelByClientID(ClientIDandSymbol{
		ClientID: "509e971a-1e70-11ed-861d-0242ac120055",
		Symbol:   "ETH-USD",
	})
	if err != nil {
		log.Error().Err(err).Msgf("Cancel Single Order Error")
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	execReport, err = cbFIXclient.NewOrderSingleLIMIT(CoinbaseFIXorderLIMIT{
		ClientID: "509e971a-1e70-11ed-861d-0242ac120003",
		Symbol:   "ETH-USD",
		Side:     Side_BUY,
		Price:    "1603",
		Qty:      "0.02",
	}, true, ctx)
	if err != nil {
		log.Error().Err(err).Msgf("NewOrderSingle Error")
	} else {
		log.Info().Interface("ExecReport", execReport).Send()
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	execReport, err = cbFIXclient.OrderStatusByClientID(ClientIDandSymbol{
		ClientID: "509e971a-1e70-11ed-861d-0242ac120003",
		Symbol:   "ETH-USD",
	}, ctx)
	if err != nil {
		log.Error().Err(err).Msgf("Order Status Error")
	} else {
		log.Info().Interface("StatusReport", execReport).Send()
	}

	time.Sleep(2 * time.Second)

	err = cbFIXclient.OrderCancelBatchByClientID([]ClientIDandSymbol{
		{
			ClientID: "509e971a-1e70-11ed-861d-0242ac120001",
			Symbol:   "ETH-USD",
		},
		{
			ClientID: "509e971a-1e70-11ed-861d-0242ac120002",
			Symbol:   "ETH-USD",
		},
		{
			ClientID: "509e971a-1e70-11ed-861d-0242ac120003",
			Symbol:   "ETH-USD",
		},
	})
	if err != nil {
		log.Error().Err(err).Msgf("Cancel Single Order Error")
	}

	time.Sleep(6 * time.Second)

	cbFIXclient.Logout()

	println(len(cbFIXclient.execReports.reportChans))
}
