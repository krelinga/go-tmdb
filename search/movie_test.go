package tmdbsearch_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/search"
)

func TestMovie(t *testing.T) {
	ctx := context.Background()

	options := tmdbsearch.MovieOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	}
	reply, err := tmdbsearch.Movie(ctx, http.DefaultClient, "Star Wars", options)
	if err != nil {
		t.Fatalf("Movie search failed: %v", err)
	}
	if reply == nil {
		t.Fatal("Movie search reply is nil")
	}
	t.Logf("Movie search reply: %+v", reply)
}