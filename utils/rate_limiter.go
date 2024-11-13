package utils

import (
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

var clients = make(map[string]*rate.Limiter)
var mu sync.Mutex

// Retrieve or create rate limiter for a client
func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	limiter, exists := clients[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 5) 
		clients[ip] = limiter
		go cleanupClients()
	}
	return limiter
}

// Middleware for rate limiting
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		limiter := getLimiter(ip)
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Clean up old entries in clients map
func cleanupClients() {
	for {
		time.Sleep(1 * time.Minute)
		mu.Lock()
		for ip, limiter := range clients {
			if limiter.Burst() == 0 {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}
