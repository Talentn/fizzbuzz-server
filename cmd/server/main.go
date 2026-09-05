package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/talentn/fizzbuzz-server/internal/api"
	"github.com/talentn/fizzbuzz-server/internal/stats"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      api.NewHandler(stats.New()),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("server listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}