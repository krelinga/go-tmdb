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
		g := &tmdb.Graph{}
		movie := result.Upsert(g)
		if movie == nil {
			t.Error("SearchMovie returned result that did not upsert a Movie")
			return
		}
		t.Logf("Found movie: %d", movie.Key)
	}
}