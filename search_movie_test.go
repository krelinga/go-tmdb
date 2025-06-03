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
	for movie, err := range tmdb.SearchMovie(ctx, client, "Inception", nil) {
		if err != nil {
			t.Errorf("SearchMovie failed: %v", err)
			return
		}
		if movie == nil {
			break
		}
		t.Logf("Found movie: %d", movie.Id())
	}
}