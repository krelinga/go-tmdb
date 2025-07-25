package tmdbsearch_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/search"
)

func TestFindSeries(t *testing.T) {
	ctx := context.Background()

	options := tmdbsearch.FindSeriesOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	}
	httpReply, err := tmdbsearch.FindSeries(ctx, http.DefaultClient, "Breaking Bad", options)
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