package demo

import (
	"context"
	"strings"
	"testing"

	"aiwiki/backend/internal/article"
)

func TestDemoClientReturnsValidArticleJSON(t *testing.T) {
	client := NewClient()
	raw, err := client.GenerateArticleJSON(context.Background(), `<script>alert(1)</script>`)
	if err != nil {
		t.Fatalf("GenerateArticleJSON returned error: %v", err)
	}

	parsed, err := article.ParseArticleJSON(raw)
	if err != nil {
		t.Fatalf("demo JSON should parse: %v", err)
	}
	if parsed.Title != `<script>alert(1)</script>` {
		t.Fatalf("title = %q", parsed.Title)
	}
	if len(parsed.Sections) == 0 {
		t.Fatal("expected demo sections")
	}
	for _, section := range parsed.Sections {
		if section.Heading == "Overview" || section.Heading == "Limitations" {
			t.Fatalf("expected more specific demo heading, got %q", section.Heading)
		}
	}
	if strings.Contains(raw, "<script>alert(1)</script></h1>") {
		t.Fatal("demo output should remain structured JSON, not HTML")
	}
}
