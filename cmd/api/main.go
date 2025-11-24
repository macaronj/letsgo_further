package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// INFO: Version number, will be generated at build time later on
const version = "1.0.0"

// INFO: Struct to hold configuration info
type config struct {
	port int
	env  string
}

// INFO: Struct to hold apllication dependencies
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	// Reading flags
	flag.IntVar(&cfg.port, "port", 4000, "API Server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (developemt|staging|production)")
	flag.Parse()

	// Initialite a new structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Declare instance of application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare a new servemux and a /v1/healthcheck route
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Declare http server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Start http server
	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
