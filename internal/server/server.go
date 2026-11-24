package server

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

func SServer(logger *log.Logger, indexHandler, uploadHandler http.HandlerFunc) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)

	srvr := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPServer: srvr,
	}
}
