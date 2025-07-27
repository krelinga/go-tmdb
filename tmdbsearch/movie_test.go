package tmdbsearch_test

import (
	"context"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/internal/util"
	"github.com/krelinga/go-tmdb/tmdbsearch"
)

func TestMovie(t *testing.T) {
	ctx := util.ContextWithAPIReadAccessToken(context.Background(), os.Getenv("TMDB_READ_ACCESS_TOKEN"))

	options := tmdbsearch.FindMoviesOptions{}
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
