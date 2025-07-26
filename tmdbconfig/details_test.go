package tmdbconfig_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/tmdbconfig"
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
options := tmdbconfig.GetDetailsOptions{

	}

	httpReply, err := tmdbconfig.GetDetails(ctx, options)
	if err != nil {
		t.Fatalf("GetDetails failed: %v", err)
	}
	reply, err := tmdbconfig.ParseGetDetailsReply(httpReply)
	if err != nil {
		t.Fatalf("ParseGetDetailsReply failed: %v", err)
	}

	t.Log("GetDetails reply:", reply)
}