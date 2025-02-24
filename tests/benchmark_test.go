package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func BenchmarkReverseProxy(b *testing.B) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer backend.Close()

	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(backend.URL)
		if err != nil {
			b.Fatalf("Proxy request failed: %v", err)
		}
		defer resp.Body.Close()
		w.WriteHeader(resp.StatusCode)
	}))
	defer proxy.Close()

	client := &http.Client{Timeout: 2 * time.Second}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Get(proxy.URL)
		if err != nil {
			b.Fatalf("Request failed: %v", err)
		}
		resp.Body.Close()
	}
}
