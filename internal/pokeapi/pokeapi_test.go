package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AdonaIsium/pokedex/internal/pokecache"
)

func TestClient_ListLocations_Cache(t *testing.T) {
	// Fake JSON body (empty object is valid for RespShallowLocations)
	body := []byte(`{}`)

	// Count how many times our test server is hit
	callCount := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}))
	defer ts.Close()

	// Build a Client that uses our test server and a fresh cache
	client := &Client{
		httpClient: *ts.Client(), // deref to match http.Client field
		cache:      *pokecache.NewCache(time.Minute),
	}

	// Use pageURL so ListLocations skips baseURL entirely
	pageURL := ts.URL + "/location-area"

	// First call: should hit the server exactly once
	if _, err := client.ListLocations(&pageURL); err != nil {
		t.Fatalf("first ListLocations error: %v", err)
	}
	if callCount != 1 {
		t.Fatalf("after 1st call, expected 1 HTTP request, got %d", callCount)
	}

	// Second call: should be served from cache (no new HTTP calls)
	if _, err := client.ListLocations(&pageURL); err != nil {
		t.Fatalf("second ListLocations error: %v", err)
	}
	if callCount != 1 {
		t.Errorf("after 2nd call, expected still 1 HTTP request, got %d", callCount)
	}
}
