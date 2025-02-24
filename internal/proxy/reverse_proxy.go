package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync/atomic"
)

type ReverseProxy struct {
	backends []*url.URL
	counter  uint32
}

func NewReverseProxy(backendUrls []string) *ReverseProxy {
	var urls []*url.URL
	for _, backend := range backendUrls {
		backend = strings.TrimSpace(backend)
		parsedUrl, err := url.Parse(backend)
		if err != nil {
			log.Fatalf("Error parsing backend URL %s: %v", backend, err)
		}
		urls = append(urls, parsedUrl)
	}
	return &ReverseProxy{backends: urls}
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := rp.getNextBackend()
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

func (rp *ReverseProxy) getNextBackend() *url.URL {
	index := atomic.AddUint32(&rp.counter, 1)
	return rp.backends[index%uint32(len(rp.backends))]
}
