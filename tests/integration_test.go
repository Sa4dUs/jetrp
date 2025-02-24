package tests

import (
	"net/http"
	"testing"
)

func TestProxy(t *testing.T) {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatalf("Request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
}
