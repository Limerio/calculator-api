package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Limerio/calculator-api/server/handlers"
	v1Handler "github.com/Limerio/calculator-api/server/handlers/v1"
	v2Handler "github.com/Limerio/calculator-api/server/handlers/v2"
	"github.com/Limerio/calculator-api/server/middlewares"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run(port string, wait time.Duration) {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/ping", handlers.Ping).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())

	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.Use(middlewares.Headers)
	v1Router.HandleFunc("/add", v1Handler.Add).Methods("POST")
	v1Router.HandleFunc("/divide", v1Handler.Divide).Methods("POST")
	v1Router.HandleFunc("/multiply", v1Handler.Multiply).Methods("POST")
	v1Router.HandleFunc("/subtract", v1Handler.Subtract).Methods("POST")

	v2Router := router.PathPrefix("/v2").Subrouter()
	v2Router.Use(middlewares.Headers, middlewares.FormatJson)
	v2Router.HandleFunc("/add", v2Handler.Add).Methods("POST")
	v2Router.HandleFunc("/divide", v2Handler.Divide).Methods("POST")
	v2Router.HandleFunc("/multiply", v2Handler.Multiply).Methods("POST")
	v2Router.HandleFunc("/subtract", v2Handler.Subtract).Methods("POST")

	stack := middlewares.Stacker(
		middlewares.Limit,
		middlewares.Logging,
		middlewares.Cors,
	)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      stack(router),
	}

	go func() {
		slog.Info("Server listening", "port", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	slog.Info("shutting down")
	os.Exit(0)
}
