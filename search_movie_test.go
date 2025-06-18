package tmdb_test

import (
	"context"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestSearchMovie(t *testing.T) {
	client := getClient(t)
	ctx := context.Background()

	// Test with a simple query
	for result, err := range tmdb.SearchMovie(ctx, client, "Inception") {
		if err != nil {
			t.Errorf("SearchMovie failed: %v", err)
			return
		}
		if result == nil {
			t.Error("SearchMovie returned nil result")
			return
		}
		if result.Movie == nil {
			t.Error("SearchMovie returned result with nil Movie")
			return
		}
		t.Logf("Found movie: %d", result.Movie.Key)
	}
}