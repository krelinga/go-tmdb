package tmdb_test

import (
	"context"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestSearchTv(t *testing.T) {
	client := getClient(t)
	ctx := context.Background()

	// Test with a simple query
	for result, err := range tmdb.SearchTv(ctx, client, "Trek") {
		if err != nil {
			t.Errorf("SearchTv failed: %v", err)
			return
		}
		if result == nil {
			t.Error("SearchTv returned nil result")
			return
		}
		if result.Tv == nil {
			t.Error("SearchTv returned result with nil Movie")
			return
		}
		t.Logf("Found Tv: %d", result.Tv.Key)
	}
}
