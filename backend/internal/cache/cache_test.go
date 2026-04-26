package cache

import (
	"testing"
	"time"
)

func TestCacheExpires(t *testing.T) {
	c := New[string](20 * time.Millisecond)
	c.Set("topic", "article")

	if got, ok := c.Get("topic"); !ok || got != "article" {
		t.Fatalf("expected cached value, got %q ok=%v", got, ok)
	}

	time.Sleep(30 * time.Millisecond)
	if _, ok := c.Get("topic"); ok {
		t.Fatal("expected cached value to expire")
	}
}
