package tmdbsearch_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/tmdbsearch"
	"github.com/krelinga/go-tmdb"
)

func TestFindSeries(t *testing.T) {
	ctx := context.Background()

	
	// Set up context with TMDB configuration
	tmdbCtx := tmdb.Context{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
		Client:          http.DefaultClient,
	}
	ctx = tmdb.SetContext(ctx, tmdbCtx)
options := tmdbsearch.FindSeriesOptions{

	}
	httpReply, err := tmdbsearch.FindSeries(ctx, "Breaking Bad", options)
	if err != nil {
		t.Fatalf("Series search failed: %v", err)
	}
	reply, err := tmdbsearch.ParseFindSeriesReply(httpReply)
	if err != nil {
		t.Fatalf("ParseSeriesReply failed: %v", err)
	}
	if reply == nil {
		t.Fatal("Series search reply is nil")
	}
	t.Logf("Series search reply: %+v", reply)
}