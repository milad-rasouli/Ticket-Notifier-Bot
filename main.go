package main

import (
	"fmt"

	"github.com/Milad75Rasouli/Ticket-Notifier-Bot/internal"
	"github.com/Milad75Rasouli/Ticket-Notifier-Bot/internal/fetcher"
	"go.uber.org/zap"
)

func main() {
	var (
		logger     *zap.Logger
		err        error
		busFetcher *fetcher.BusTicketFetcher
		cfg        internal.Config
	)

	{
		err = cfg.Read()
		if err != nil {
			panic(err)
		}
	}

	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infoln("server is running")
	sugar.Infof("config %+v", cfg)

	{
		busFetcher = fetcher.NewBusTicketFetcher(logger.Sugar().Named("bus fetcher"), cfg.URL, cfg.ContentType)
	}

	// var jsonStr = []byte(`{
	// 	"from": 31310000,
	// 	"to": 11320000,
	// 	"date": "2024-06-03",
	// 	"includeClosed": true,
	// 	"includePromotions": true,
	// 	"loadFromDbOnUnavailability": true,
	// 	"includeUnderDevelopment": true
	// }`)
	bus, err := busFetcher.FetchBusTicket(31310000, 11320000, "2024-06-03")
	if err != nil {
		sugar.Fatalln(err)
	}
	fmt.Printf("%+v", bus)
}
