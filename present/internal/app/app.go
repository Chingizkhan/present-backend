package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"present/present/config"
	"present/present/pkg/httpserver"
	"present/present/pkg/postgres"
	"syscall"
)

func createRouter(cfg *config.Config) http.Handler {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/url", func(r chi.Router) {
		r.Use(middleware.BasicAuth("present-backend", map[string]string{
			cfg.HTTP.User: cfg.HTTP.Password,
		}))
		//r.Post("/", save.New(log, storage))
	})

	//router.Get("/{alias}", redirect.New(log, storage))

	return router
}

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

	// HTTP server
	router := createRouter(cfg)
	//log.Info("starting server", slog.String("address", cfg.Address))
	log.Println("starting server", slog.String("address", cfg.HTTP.Host+":"+cfg.HTTP.Port))

	httpServer := httpserver.New(router, httpserver.Addr(cfg.HTTP.Host, cfg.HTTP.Port))

	// waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	//case s := <-interrupt:
	//	l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		//l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		//case err = <-rmqServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
		//l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	//err = rmqServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	//}
}
