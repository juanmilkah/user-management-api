package middleware

import (
    "net/http"
    "sync"
    "golang.org/x/time/rate"
    "go-server/utils"
)

type RateLimiter struct {
    visitors map[string]*rate.Limiter
    mu       sync.RWMutex
    rate     rate.Limit
    burst    int
    errorHandler *utils.ErrorHandler
}

func NewRateLimiter(requestsPerMinute float64, errorHandler *utils.ErrorHandler) *RateLimiter {
    return &RateLimiter{
        visitors:     make(map[string]*rate.Limiter),
        rate:        rate.Limit(requestsPerMinute / 60), // convert to per-second
        burst:       int(requestsPerMinute / 30),        // allow bursts of up to 30 seconds worth
        errorHandler: errorHandler,
    }
}

func (rl *RateLimiter) getLimiter(ip string) *rate.Limiter {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    limiter, exists := rl.visitors[ip]
    if !exists {
        limiter = rate.NewLimiter(rl.rate, rl.burst)
        rl.visitors[ip] = limiter
    }

    return limiter
}

func (rl *RateLimiter) RateLimit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr // In production, you might want to use X-Forwarded-For or similar

        limiter := rl.getLimiter(ip)
        if !limiter.Allow() {
            rl.errorHandler.RespondWithError(w, http.StatusTooManyRequests, 
                "Rate limit exceeded. Please try again later.", nil)
            return
        }

        next.ServeHTTP(w, r)
    })
}
