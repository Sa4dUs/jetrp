package metrics

import (
	"log"
	"time"
)

func LogRequest(method, path string, status int, duration time.Duration) {
	log.Printf("[%s] %s %d %s", method, path, status, duration)
}
