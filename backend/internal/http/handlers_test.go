package httpapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"aiwiki/backend/internal/article"
)

func TestArticlePostRejectsOversizedBody(t *testing.T) {
	t.Setenv("AIWIKI_RATE_LIMIT_PER_MINUTE", "100")
	service := article.NewService(&testGenerator{}, 10*time.Minute)
	handler := NewHandler(service).Routes()

	body := `{"topic":"` + strings.Repeat("a", maxRequestBodyBytes+1) + `"}`
	req := httptest.NewRequest(http.MethodPost, "/api/article", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("status = %d, want %d; body=%s", rec.Code, http.StatusRequestEntityTooLarge, rec.Body.String())
	}
}

func TestArticlePostRateLimitsPerIP(t *testing.T) {
	t.Setenv("AIWIKI_RATE_LIMIT_PER_MINUTE", "1")
	t.Setenv("AIWIKI_RATE_LIMIT_BURST", "1")
	service := article.NewService(&testGenerator{}, 10*time.Minute)
	handler := NewHandler(service).Routes()

	for i := 0; i < 2; i++ {
		req := httptest.NewRequest(http.MethodPost, "/api/article", strings.NewReader(`{"topic":"Rate Limit"}`))
		req.RemoteAddr = "203.0.113.10:1234"
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if i == 0 && rec.Code != http.StatusOK {
			t.Fatalf("first request status = %d, want 200; body=%s", rec.Code, rec.Body.String())
		}
		if i == 1 && rec.Code != http.StatusTooManyRequests {
			t.Fatalf("second request status = %d, want 429; body=%s", rec.Code, rec.Body.String())
		}
	}
}

func TestCORSAllowsOnlyConfiguredOrigins(t *testing.T) {
	t.Setenv("AIWIKI_ALLOWED_ORIGINS", "https://aiwiki.example.com")
	service := article.NewService(&testGenerator{}, 10*time.Minute)
	handler := NewHandler(service).Routes()

	allowedReq := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	allowedReq.Header.Set("Origin", "https://aiwiki.example.com")
	allowedRec := httptest.NewRecorder()
	handler.ServeHTTP(allowedRec, allowedReq)

	if got := allowedRec.Header().Get("Access-Control-Allow-Origin"); got != "https://aiwiki.example.com" {
		t.Fatalf("allowed origin header = %q", got)
	}

	blockedReq := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	blockedReq.Header.Set("Origin", "https://evil.example.com")
	blockedRec := httptest.NewRecorder()
	handler.ServeHTTP(blockedRec, blockedReq)

	if got := blockedRec.Header().Get("Access-Control-Allow-Origin"); got != "" {
		t.Fatalf("blocked origin should not get CORS header, got %q", got)
	}
}

type testGenerator struct{}

func (g *testGenerator) GenerateArticleJSON(context.Context, string) (string, error) {
	return `{"title":"Test","summary":"Summary","sections":[],"references":[],"seeAlso":[]}`, nil
}

func (g *testGenerator) RepairArticleJSON(_ context.Context, raw string) (string, error) {
	return raw, nil
}

func (g *testGenerator) Health(context.Context) error {
	return nil
}

func (g *testGenerator) Model() string {
	return "test"
}
