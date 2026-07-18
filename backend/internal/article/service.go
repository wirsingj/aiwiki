package article

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"

	"aiwiki/backend/internal/cache"
)

type Generator interface {
	GenerateArticleJSON(ctx context.Context, topic string) (string, error)
	RepairArticleJSON(ctx context.Context, raw string) (string, error)
	Health(ctx context.Context) error
	Model() string
}

type ProviderNamer interface {
	Provider() string
}

type Service struct {
	generator Generator
	cache     *cache.Cache[Article]
}

func NewService(generator Generator, ttl time.Duration) *Service {
	return &Service{
		generator: generator,
		cache:     cache.New[Article](ttl),
	}
}

func (s *Service) Get(ctx context.Context, slug string, bypassCache bool) (Article, error) {
	normalizedSlug := Slugify(slug)
	if normalizedSlug == "" {
		return Article{}, ErrTopicRequired
	}

	topic, err := NormalizeTopic(TopicFromSlug(normalizedSlug))
	if err != nil {
		return Article{}, err
	}

	return s.generate(ctx, normalizedSlug, topic, bypassCache)
}

func (s *Service) Create(ctx context.Context, topic string, bypassCache bool) (Article, error) {
	normalizedTopic, err := NormalizeTopic(topic)
	if err != nil {
		return Article{}, err
	}
	return s.generate(ctx, Slugify(normalizedTopic), normalizedTopic, bypassCache)
}

func (s *Service) generate(ctx context.Context, normalizedSlug string, topic string, bypassCache bool) (Article, error) {
	if !bypassCache {
		if cached, ok := s.cache.Get(normalizedSlug); ok {
			cached.Cached = true
			return cached, nil
		}
	}

	raw, err := s.generator.GenerateArticleJSON(ctx, topic)
	if err != nil {
		return Article{}, err
	}

	parsed, parseErr := ParseArticleJSON(raw)
	if parseErr != nil {
		log.Printf("generator returned malformed JSON for slug %q: %v", normalizedSlug, parseErr)
		repaired, repairErr := s.generator.RepairArticleJSON(ctx, raw)
		if repairErr != nil {
			log.Printf("generator JSON repair failed for slug %q: %v", normalizedSlug, repairErr)
			return Article{}, errors.New("article generator returned malformed JSON and the repair attempt failed")
		}

		parsed, parseErr = ParseArticleJSON(repaired)
		if parseErr != nil {
			log.Printf("repaired generator JSON still invalid for slug %q: %v", normalizedSlug, parseErr)
			return Article{}, errors.New("article generator returned malformed JSON after repair")
		}
	}

	article := normalizeArticle(parsed, normalizedSlug)
	s.cache.Set(normalizedSlug, article)
	return article, nil
}

func (s *Service) Health(ctx context.Context) map[string]any {
	provider := "model"
	if named, ok := s.generator.(ProviderNamer); ok {
		provider = named.Provider()
	}

	status := map[string]any{
		"ok":       true,
		"service":  "aiwiki-backend",
		"provider": provider,
		"model":    s.generator.Model(),
	}
	if provider == "ollama" {
		status["ollamaModel"] = s.generator.Model()
	}
	if err := s.generator.Health(ctx); err != nil {
		status["ok"] = false
		status["modelStatus"] = "offline"
		if provider == "ollama" {
			status["ollama"] = "offline"
		}
		status["error"] = err.Error()
		return status
	}
	status["modelStatus"] = "online"
	if provider == "ollama" {
		status["ollama"] = "online"
	}
	return status
}

func ParseArticleJSON(raw string) (Article, error) {
	candidates := []string{strings.TrimSpace(raw)}
	if fenced := extractFencedJSON(raw); fenced != "" {
		candidates = append(candidates, fenced)
	}
	if object := extractJSONObject(raw); object != "" {
		candidates = append(candidates, object)
	}

	var lastErr error
	for _, candidate := range candidates {
		if candidate == "" {
			continue
		}
		var parsed Article
		if err := json.Unmarshal([]byte(candidate), &parsed); err != nil {
			lastErr = err
			continue
		}
		if strings.TrimSpace(parsed.Title) == "" || strings.TrimSpace(parsed.Summary) == "" {
			lastErr = errors.New("article JSON is missing title or summary")
			continue
		}
		return parsed, nil
	}

	if lastErr == nil {
		lastErr = errors.New("no JSON object found")
	}
	return Article{}, fmt.Errorf("invalid article JSON: %w", lastErr)
}

func Slugify(input string) string {
	input = strings.TrimSpace(strings.ToLower(input))
	var b strings.Builder
	lastDash := false

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			lastDash = false
			continue
		}
		if !lastDash && b.Len() > 0 {
			b.WriteRune('-')
			lastDash = true
		}
	}

	slug := strings.Trim(b.String(), "-")
	if slug == "" {
		return "untitled"
	}
	return slug
}

func normalizeArticle(a Article, slug string) Article {
	a.Slug = slug
	a.Cached = false
	a.GeneratedAt = time.Now().UTC()
	a.Title = fallback(a.Title, TopicFromSlug(slug))
	a.Summary = fallback(a.Summary, "AI-generated summary unavailable.")

	if a.Infobox != nil {
		a.Infobox.Heading = fallback(a.Infobox.Heading, a.Title)
		if a.Infobox.Rows == nil {
			a.Infobox.Rows = []InfoboxRow{}
		}
	}
	if a.Sections == nil {
		a.Sections = []Section{}
	}
	for i := range a.Sections {
		if a.Sections[i].Paragraphs == nil {
			a.Sections[i].Paragraphs = []string{}
		}
		if a.Sections[i].Links == nil {
			a.Sections[i].Links = []string{}
		}
	}
	if len(a.References) == 0 {
		a.References = []Reference{{
			Label: "AI-generated note",
			Note:  "This article was generated by a local AI model and has not been independently verified.",
		}}
	}
	for i := range a.References {
		a.References[i].Label = fallback(a.References[i].Label, "AI-generated note")
		a.References[i].Note = ensureGeneratedReferenceNote(a.References[i].Note)
	}
	if a.SeeAlso == nil {
		a.SeeAlso = []string{}
	}
	return a
}

func fallback(value, replacement string) string {
	if strings.TrimSpace(value) == "" {
		return replacement
	}
	return strings.TrimSpace(value)
}

func ensureGeneratedReferenceNote(note string) string {
	note = strings.TrimSpace(note)
	disclaimer := "This is an AI-generated reference note and has not been independently verified."
	if note == "" {
		return disclaimer
	}

	lower := strings.ToLower(note)
	if strings.Contains(lower, "ai-generated") && strings.Contains(lower, "independently verified") {
		return note
	}
	return note + " " + disclaimer
}

func extractFencedJSON(raw string) string {
	start := strings.Index(raw, "```")
	if start == -1 {
		return ""
	}
	rest := raw[start+3:]
	end := strings.Index(rest, "```")
	if end == -1 {
		return ""
	}
	block := strings.TrimSpace(rest[:end])
	if newline := strings.IndexByte(block, '\n'); newline != -1 {
		firstLine := strings.TrimSpace(strings.ToLower(block[:newline]))
		if firstLine == "json" || firstLine == "javascript" {
			block = block[newline+1:]
		}
	}
	return strings.TrimSpace(block)
}

func extractJSONObject(raw string) string {
	start := strings.Index(raw, "{")
	end := strings.LastIndex(raw, "}")
	if start == -1 || end == -1 || end <= start {
		return ""
	}
	return strings.TrimSpace(raw[start : end+1])
}
