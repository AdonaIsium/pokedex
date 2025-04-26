package pokecache

import (
	"testing"
	"time"
)

func TestReapLoop(t *testing.T) {
	interval := 50 * time.Millisecond
	c := NewCache(interval)

	c.Add("foo", []byte("bar"))
	if v, ok := c.Get("foo"); !ok {
		t.Fatalf("expected key %q to exist immediately, got ok=false", "foo")
	} else if string(v) != "bar" {
		t.Fatalf("expected value %q, got %q", "bar", string(v))
	}

	time.Sleep(interval + 20*time.Millisecond)

	if _, ok := c.Get("foo"); ok {
		t.Errorf("expected key %q to be expired after %v, but it still exists", "foo", interval)
	}
}
