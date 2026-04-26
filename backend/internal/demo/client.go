package demo

import (
	"context"
	"encoding/json"
	"strings"

	"aiwiki/backend/internal/article"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Provider() string {
	return "demo"
}

func (c *Client) Model() string {
	return "aiwiki-template-demo"
}

func (c *Client) Health(context.Context) error {
	return nil
}

func (c *Client) RepairArticleJSON(_ context.Context, raw string) (string, error) {
	return raw, nil
}

func (c *Client) GenerateArticleJSON(_ context.Context, topic string) (string, error) {
	title := strings.TrimSpace(topic)
	if title == "" {
		title = "Untitled Topic"
	}

	a := article.Article{
		Title:   title,
		Summary: "This demo article is generated from a deterministic AIWIKI template rather than a live language model. It is intended for low-cost public deployments where the application flow, layout, routing, caching, and safety disclaimers need to be visible without running local inference. The subject is treated as untrusted data and is used only as the article topic.",
		Infobox: &article.Infobox{
			Heading: "AIWIKI demo profile",
			Rows: []article.InfoboxRow{
				{Label: "Topic", Value: title},
				{Label: "Generation mode", Value: "Template demo"},
				{Label: "Verification", Value: "Not independently verified"},
			},
		},
		Sections: []article.Section{
			{
				Heading: "Role in the AIWIKI demo",
				Paragraphs: []string{
					"AIWIKI presents the requested topic as a compact encyclopedia-style entry. In demo mode, the article is assembled instantly so a public site can show search, routing, article layout, table of contents, infobox data, generated notes, and internal links without paying for model inference.",
					"The text in this mode should not be treated as a factual account of the topic. It is controlled explanatory material designed to exercise the same rendering path used by live model output.",
				},
				Links: []string{"AI-generated content", "Encyclopedia style", "Public demo"},
			},
			{
				Heading: "Deployment behavior",
				Paragraphs: []string{
					"A public AWS free-tier deployment can use demo mode to stay responsive and avoid hosting a large model. The backend still validates topics, rate-limits generation requests, applies cache behavior, and returns the same JSON structure as the live generator.",
					"When a hosted or local model is configured later, this same page can be generated through that provider without changing the frontend contract.",
				},
				Links: []string{"AWS Free Tier", "Hosted model API", "Local Ollama"},
			},
			{
				Heading: "Verification limits",
				Paragraphs: []string{
					"Demo mode does not research, reason about, or summarize the requested topic. It intentionally avoids factual claims beyond describing the application behavior.",
					"For a real public AI demo, connect AIWIKI to an OpenAI-compatible hosted provider from the backend and keep provider keys out of the browser.",
				},
				Links: []string{"Model licensing", "Prompt injection", "Rate limiting"},
			},
		},
		References: []article.Reference{{
			Label: "Generated demo note",
			Note:  "This is an AIWIKI template-generated note and has not been independently verified.",
		}},
		SeeAlso: []string{"AI-generated content", "Local-first software", "Hosted model API"},
	}

	body, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
