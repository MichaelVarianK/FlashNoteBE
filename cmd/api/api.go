package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
	// dbConfig
	// rateLimit
}

func (app *application) run() error {
	mux := http.NewServeMux()

	srv := &http.Server {
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Println("Started an application with port", app.config.addr)

	return srv.ListenAndServe()
}