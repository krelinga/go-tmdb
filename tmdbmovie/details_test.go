package tmdbmovie_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb"
	"github.com/krelinga/go-tmdb/tmdbmovie"
)

func TestGetDetails(t *testing.T) {
	ctx := context.Background()
	
	// Set up context with TMDB configuration
	tmdbCtx := tmdb.Context{
		Key:             os.Getenv("TMDB_API_KEY"),
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
		Client:          http.DefaultClient,
	}
	ctx = tmdb.SetContext(ctx, tmdbCtx)

	options := tmdbmovie.GetDetailsOptions{
		AppendCredits:      true,
		AppendExternalIDs:  true,
		AppendReleaseDates: true,
	}

	httpReply, err := tmdbmovie.GetDetails(ctx, 11, options)
	if err != nil {
		t.Fatalf("GetDetails failed: %v", err)
	}
	reply, err := tmdbmovie.ParseGetDetailsReply(httpReply)
	if err != nil {
		t.Fatalf("ParseGetDetailsReply failed: %v", err)
	}

	if reply.ID == nil || *reply.ID != 11 {
		t.Errorf("unexpected ID: %v", reply.ID)
	}
	t.Log("GetDetails reply:", reply)
}
