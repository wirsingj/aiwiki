package article

import "time"

type InfoboxRow struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Infobox struct {
	Heading string       `json:"heading"`
	Rows    []InfoboxRow `json:"rows"`
}

type Section struct {
	Heading    string   `json:"heading"`
	Paragraphs []string `json:"paragraphs"`
	Links      []string `json:"links"`
}

type Reference struct {
	Label string `json:"label"`
	Note  string `json:"note"`
}

type Article struct {
	Slug        string      `json:"slug"`
	Title       string      `json:"title"`
	Summary     string      `json:"summary"`
	Infobox     *Infobox    `json:"infobox,omitempty"`
	Sections    []Section   `json:"sections"`
	References  []Reference `json:"references"`
	SeeAlso     []string    `json:"seeAlso"`
	GeneratedAt time.Time   `json:"generatedAt"`
	Cached      bool        `json:"cached"`
}
