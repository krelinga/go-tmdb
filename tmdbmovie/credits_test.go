package tmdbmovie_test

import (
	"context"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/internal/util"
	"github.com/krelinga/go-tmdb/tmdbmovie"
)

func TestGetCredits(t *testing.T) {
	ctx := util.ContextWithAPIReadAccessToken(context.Background(), os.Getenv("TMDB_READ_ACCESS_TOKEN"))

	options := tmdbmovie.GetCreditsOptions{}

	httpReply, err := tmdbmovie.GetCredits(ctx, 11, options)
	if err != nil {
		t.Fatalf("GetCredits failed: %v", err)
	}
	reply, err := tmdbmovie.ParseGetCreditsReply(httpReply)
	if err != nil {
		t.Fatalf("ParseGetCreditsReply failed: %v", err)
	}

	if reply.ID == nil || *reply.ID != 11 {
		t.Errorf("unexpected ID: %v", reply.ID)
	}
	t.Log("GetCredits reply:", reply)
}
