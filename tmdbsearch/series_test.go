package tmdbsearch_test

import (
	"context"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/internal/util"
	"github.com/krelinga/go-tmdb/tmdbsearch"
)

func TestFindSeries(t *testing.T) {
	ctx := util.ContextWithAPIReadAccessToken(context.Background(), os.Getenv("TMDB_READ_ACCESS_TOKEN"))

	options := tmdbsearch.FindSeriesOptions{}
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
