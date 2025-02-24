package main

import (
	"net/http"
	"testing"
)

func BenchmarkProxy(b *testing.B) {
	client := &http.Client{}
	for i := 0; i < b.N; i++ {
		_, err := client.Get("http://localhost:8080")
		if err != nil {
			b.Fatalf("Error en la solicitud: %v", err)
		}
	}
}
