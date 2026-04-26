package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"aiwiki/backend/internal/article"
)

const maxRequestBodyBytes = 4096

type Handler struct {
	service     *article.Service
	rateLimiter *rateLimiter
	random      []string
}

func NewHandler(service *article.Service) *Handler {
	return &Handler{
		service: service,
		rateLimiter: newRateLimiter(
			intFromEnv(os.Getenv("AIWIKI_RATE_LIMIT_PER_MINUTE"), defaultRateLimitPerMinute),
			intFromEnv(os.Getenv("AIWIKI_RATE_LIMIT_BURST"), defaultRateLimitBurst),
		),
		random: []string{
			"Local artificial intelligence",
			"History of imaginary maps",
			"Coffeehouse computing",
			"Solar punk architecture",
			"Forgotten programming languages",
			"Mountain observatories",
			"Desk plants",
			"Public domain folklore",
		},
	}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", h.handleHealth)
	mux.HandleFunc("/api/random", h.handleRandom)
	mux.HandleFunc("/api/article/", h.handleArticleBySlug)
	mux.HandleFunc("/api/article", h.handleArticlePost)
	return cors(mux)
}

func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	writeJSON(w, http.StatusOK, h.service.Health(ctx))
}

func (h *Handler) handleRandom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}
	topic := h.random[rand.Intn(len(h.random))]
	slug := article.Slugify(topic)
	http.Redirect(w, r, "/api/article/"+slug, http.StatusTemporaryRedirect)
}

func (h *Handler) handleArticleBySlug(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}

	slug := strings.TrimPrefix(r.URL.Path, "/api/article/")
	if strings.TrimSpace(slug) == "" {
		writeError(w, http.StatusBadRequest, "article slug is required")
		return
	}
	if !h.rateLimiter.allow(clientIP(r)) {
		writeError(w, http.StatusTooManyRequests, "too many article generation requests; please wait and try again")
		return
	}

	bypassCache := r.URL.Query().Get("refresh") == "1"
	ctx, cancel := context.WithTimeout(r.Context(), 210*time.Second)
	defer cancel()

	result, err := h.service.Get(ctx, slug, bypassCache)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *Handler) handleArticlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}

	var body struct {
		Topic   string `json:"topic"`
		Refresh bool   `json:"refresh"`
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			writeError(w, http.StatusRequestEntityTooLarge, "request body is too large")
			return
		}
		writeError(w, http.StatusBadRequest, "request body must be JSON with a topic")
		return
	}
	if strings.TrimSpace(body.Topic) == "" {
		writeError(w, http.StatusBadRequest, "topic is required")
		return
	}
	if !h.rateLimiter.allow(clientIP(r)) {
		writeError(w, http.StatusTooManyRequests, "too many article generation requests; please wait and try again")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 210*time.Second)
	defer cancel()

	result, err := h.service.Create(ctx, body.Topic, body.Refresh)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func cors(next http.Handler) http.Handler {
	allowedOrigins := configuredAllowedOrigins()
	allowAnyOrigin := allowedOrigins["*"]

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Vary", "Origin")
		if !allowAnyOrigin && allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func configuredAllowedOrigins() map[string]bool {
	raw := strings.TrimSpace(os.Getenv("AIWIKI_ALLOWED_ORIGINS"))
	if raw == "" {
		raw = "http://localhost:5173,http://127.0.0.1:5173"
	}

	allowed := make(map[string]bool)
	for _, origin := range strings.Split(raw, ",") {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			allowed[origin] = true
		}
	}
	return allowed
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, article.ErrTopicRequired),
		errors.Is(err, article.ErrTopicTooLong),
		errors.Is(err, article.ErrTopicControlCharacters):
		writeError(w, http.StatusBadRequest, err.Error())
	default:
		writeError(w, http.StatusBadGateway, err.Error())
	}
}

func writeMethodNotAllowed(w http.ResponseWriter) {
	writeError(w, http.StatusMethodNotAllowed, "method not allowed")
}
