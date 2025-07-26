package tmdbseries_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/tmdbseries"
	"github.com/krelinga/go-tmdb"
)

func TestGetDetails(t *testing.T) {
	ctx := context.Background()

	
	// Set up context with TMDB configuration
	tmdbCtx := tmdb.Context{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
		Client:          http.DefaultClient,
	}
	ctx = tmdb.SetContext(ctx, tmdbCtx)
options := tmdbseries.GetDetailsOptions{

	}

	httpReply, err := tmdbseries.GetDetails(ctx, 1399, options)
	if err != nil {
		t.Fatalf("GetDetails failed: %v", err)
	}
	reply, err := tmdbseries.ParseGetDetailsReply(httpReply)
	if err != nil {
		t.Fatalf("ParseGetDetailsReply failed: %v", err)
	}

	if reply.ID == nil || *reply.ID != 1399 {
		t.Errorf("unexpected ID: %v", reply.ID)
	}
	t.Log("GetDetails reply:", reply)
}
