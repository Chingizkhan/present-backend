package app

import (
	"log/slog"
	"os"
	"os/signal"
	"present/present/config"
	v1 "present/present/internal/controller/http/v1"
	"present/present/internal/usecase"
	"present/present/internal/usecase/repo"
	"present/present/pkg/httpserver"
	"present/present/pkg/logger"
	"present/present/pkg/postgres"
	"syscall"
)

func Run(cfg *config.Config) {
	// logger
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting url shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// repository
	pg, err := postgres.New(
		postgres.GetDSN(cfg),
		postgres.MaxPoolSize(cfg.PG.PoolMax),
	)
	if err != nil {
		log.Error("app - Run - postgres.New:", logger.Err(err))
		os.Exit(1)
	}
	defer pg.Close()

	// use case
	productUseCase := usecase.New(repo.New(pg))

	// rabbitmq rpc server

	// HTTP server
	router := v1.NewRouter(cfg, log, productUseCase)
	log.Info("starting server", slog.String("address", cfg.HTTP.Host+":"+cfg.HTTP.Port))

	httpServer := httpserver.New(router, httpserver.Addr(cfg.HTTP.Host, cfg.HTTP.Port))

	// waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: ", "signal", s.String())
	case err = <-httpServer.Notify():
		log.Error("app - Run - httpServer.Notify:", logger.Err(err))
		//case err = <-rmqServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error("app - Run - httpServer.Shutdown:", logger.Err(err))
	}

	//err = rmqServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	//}
}
