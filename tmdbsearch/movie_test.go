package tmdbsearch_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/tmdbsearch"
	"github.com/krelinga/go-tmdb"
)

func TestMovie(t *testing.T) {
	ctx := context.Background()

	
	// Set up context with TMDB configuration
	tmdbCtx := tmdb.Context{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
		Client:          http.DefaultClient,
	}
	ctx = tmdb.SetContext(ctx, tmdbCtx)
options := tmdbsearch.FindMoviesOptions{

	}
	httpReply, err := tmdbsearch.FindMovies(ctx, "Star Wars", options)
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
