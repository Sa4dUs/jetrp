package main

import (
	"jetrp/internal/config"
	"jetrp/internal/proxy"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	rp := proxy.NewReverseProxy(cfg.Backends)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: rp,
	}

	log.Printf("Reverse proxy running on port %s", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
