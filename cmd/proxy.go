package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
)

var backends = []string {
	"http://localhost:8081",
	"http://localhost:8082",
}

var backendIndex = 0;
var mu sync.Mutex

func newReverseProxy(target string) (*httputil.ReverseProxy, error) {
	parsedURL, err := url.Parse(target)
	if err != nil {
		return nil, fmt.Errorf("failed to parse target URL: %v", err)
	}
	return httputil.NewSingleHostReverseProxy(parsedURL), nil
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	backendURL := backends[backendIndex]
	
	log.Printf("Proxying request to: %s %s -> %s", r.Method, r.URL.Path, backendURL)

	backendIndex = (backendIndex + 1) % len(backends)

	proxy, err := newReverseProxy(backendURL)
	if err != nil {
		log.Printf("Error creating reverse proxy: %v", err)
		http.Error(w, fmt.Sprintf("Error creating reverse proxy: %v", err), http.StatusInternalServerError)
		return
	}

	proxy.ServeHTTP(w, r)
}

func startProxyServer(port string) {
	http.HandleFunc("/", proxyHandler)

	log.Printf("Reverse proxy server running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	startProxyServer(port)
}
