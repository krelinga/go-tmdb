package tmdbmovie_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/tmdbmovie"
	"github.com/krelinga/go-tmdb"
)

func TestGetReleaseDates(t *testing.T) {
	ctx := context.Background()

	
	// Set up context with TMDB configuration
	tmdbCtx := tmdb.Context{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
		Client:          http.DefaultClient,
	}
	ctx = tmdb.SetContext(ctx, tmdbCtx)
options := tmdbmovie.GetReleaseDatesOptions{

	}

	httpReply, err := tmdbmovie.GetReleaseDates(ctx, 11, options)
	if err != nil {
		t.Fatalf("GetReleaseDates failed: %v", err)
	}
	reply, err := tmdbmovie.ParseGetReleaseDatesReply(httpReply)
	if err != nil {
		t.Fatalf("ParseGetReleaseDatesReply failed: %v", err)
	}

	if reply.ID == nil || *reply.ID != 11 {
		t.Errorf("unexpected ID: %v", reply.ID)
	}
	t.Log("GetReleaseDates reply:", reply)
}
