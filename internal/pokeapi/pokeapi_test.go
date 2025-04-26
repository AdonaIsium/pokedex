package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AdonaIsium/pokedex/internal/pokecache"
)

func TestClient_ListLocations_Cache(t *testing.T) {
	body := []byte(`{}`)

	callCount := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}))
	defer ts.Close()

	client := &Client{
		httpClient: *ts.Client(),
	}

	cache := pokecache.NewCache(1 * time.Minute)

	pageURL := ts.URL + "/location-area"

	if _, err := client.ListLocations(&pageURL, cache); err != nil {
		t.Fatalf("first ListLocations error: %v", err)
	}
	if callCount != 1 {
		t.Fatalf("after 1st call, expected 1 HTTP request, got %d", callCount)
	}

	if _, err := client.ListLocations(&pageURL, cache); err != nil {
		t.Fatalf("second ListLocations error: %v", err)
	}
	if callCount != 1 {
		t.Errorf("after 2nd call, expected still 1 HTTP request, got %d", callCount)
	}
}
