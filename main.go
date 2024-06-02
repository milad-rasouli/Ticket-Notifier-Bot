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
	return
	{
		busFetcher = fetcher.NewBusTicketFetcher(logger.Sugar().Named("bus fetcher"), "url", "contectType")
	}

	bus, err := busFetcher.FetchBusTicket(1, 1, "")
	if err != nil {
		sugar.Fatalln(err)
	}
	fmt.Printf("%+v", bus)
}
