package tmdbsearch_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	tmdbsearch "github.com/krelinga/go-tmdb/search"
)

func TestMovie(t *testing.T) {
	ctx := context.Background()

	options := tmdbsearch.FindMoviesOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	}
	httpReply, err := tmdbsearch.FindMovies(ctx, http.DefaultClient, "Star Wars", options)
	if err != nil {
		t.Fatalf("Movie search failed: %v", err)
	}
	reply, err := tmdbsearch.ParseFindMoviesReply(httpReply)
	if err != nil {
		t.Fatalf("ParseMovieReply failed: %v", err)
	}
	if reply == nil {
		t.Fatal("Movie search reply is nil")
	}
	t.Logf("Movie search reply: %+v", reply)
}
