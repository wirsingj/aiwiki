package httpapi

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultRateLimitPerMinute = 12
	defaultRateLimitBurst     = 4
)

type rateLimiter struct {
	mu        sync.Mutex
	limit     float64
	burst     float64
	visitors  map[string]*visitor
	lastSweep time.Time
}

type visitor struct {
	tokens float64
	last   time.Time
}

func newRateLimiter(limitPerMinute int, burst int) *rateLimiter {
	if limitPerMinute <= 0 {
		limitPerMinute = defaultRateLimitPerMinute
	}
	if burst <= 0 {
		burst = defaultRateLimitBurst
	}
	return &rateLimiter{
		limit:     float64(limitPerMinute),
		burst:     float64(burst),
		visitors:  make(map[string]*visitor),
		lastSweep: time.Now(),
	}
}

func (l *rateLimiter) allow(key string) bool {
	now := time.Now()

	l.mu.Lock()
	defer l.mu.Unlock()

	if now.Sub(l.lastSweep) > 5*time.Minute {
		for ip, v := range l.visitors {
			if now.Sub(v.last) > 10*time.Minute {
				delete(l.visitors, ip)
			}
		}
		l.lastSweep = now
	}

	v, ok := l.visitors[key]
	if !ok {
		l.visitors[key] = &visitor{tokens: l.burst - 1, last: now}
		return true
	}

	elapsed := now.Sub(v.last).Minutes()
	v.tokens += elapsed * l.limit
	if v.tokens > l.burst {
		v.tokens = l.burst
	}
	v.last = now

	if v.tokens < 1 {
		return false
	}
	v.tokens--
	return true
}

func clientIP(r *http.Request) string {
	if forwarded := strings.TrimSpace(r.Header.Get("X-Forwarded-For")); forwarded != "" {
		first := strings.TrimSpace(strings.Split(forwarded, ",")[0])
		if parsed := net.ParseIP(first); parsed != nil {
			return parsed.String()
		}
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		if parsed := net.ParseIP(host); parsed != nil {
			return parsed.String()
		}
		return host
	}
	return r.RemoteAddr
}

func intFromEnv(value string, fallback int) int {
	parsed, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}
