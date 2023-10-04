package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"present/present/config"
	"present/present/pkg/postgres"
	"syscall"
)

func Run(cfg *config.Config) {
	// logger
	// repository
	pg, err := postgres.New(
		postgres.GetDSN(cfg),
		postgres.MaxPoolSize(cfg.PG.PoolMax),
	)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// use case
	// rabbitmq rpc server
	// http server
	// waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
		//case s := <-interrupt:
		//	l.Info("app - Run - signal: " + s.String())
		//case err = <-httpServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		//case err = <-rmqServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// shutdown
	//err = httpServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	//}
	//
	//err = rmqServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	//}
}
