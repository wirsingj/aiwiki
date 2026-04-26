package article

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	MinTopicLength = 1
	MaxTopicLength = 120
)

var (
	ErrTopicRequired          = errors.New("topic is required")
	ErrTopicTooLong           = errors.New("topic must be 120 characters or fewer")
	ErrTopicControlCharacters = errors.New("topic contains unsupported control characters")
)

func NormalizeTopic(input string) (string, error) {
	var b strings.Builder
	previousWhitespace := false

	for _, r := range input {
		if unicode.IsSpace(r) {
			if b.Len() > 0 && !previousWhitespace {
				b.WriteByte(' ')
				previousWhitespace = true
			}
			continue
		}
		if unicode.IsControl(r) || r == utf8.RuneError {
			return "", ErrTopicControlCharacters
		}
		b.WriteRune(r)
		previousWhitespace = false
	}

	topic := strings.TrimSpace(b.String())
	length := utf8.RuneCountInString(topic)
	if length < MinTopicLength {
		return "", ErrTopicRequired
	}
	if length > MaxTopicLength {
		return "", ErrTopicTooLong
	}
	return topic, nil
}

func TopicFromSlug(slug string) string {
	parts := strings.Fields(strings.ReplaceAll(slug, "-", " "))
	for i, part := range parts {
		runes := []rune(part)
		if len(runes) == 0 {
			continue
		}
		runes[0] = unicode.ToUpper(runes[0])
		parts[i] = string(runes)
	}
	return strings.Join(parts, " ")
}
