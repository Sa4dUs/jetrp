package proxy

import (
	"net/url"
	"sync/atomic"
)

type LoadBalancer struct {
	backends []*url.URL
	counter  uint32
}

func NewLoadBalancer(backendUrls []string) *LoadBalancer {
	var urls []*url.URL
	for _, backend := range backendUrls {
		parsedUrl, _ := url.Parse(backend)
		urls = append(urls, parsedUrl)
	}
	return &LoadBalancer{backends: urls}
}

func (lb *LoadBalancer) GetNextBackend() *url.URL {
	index := atomic.AddUint32(&lb.counter, 1)
	return lb.backends[index%uint32(len(lb.backends))]
}
