package tmdbmovie_test

import (
	"context"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/internal/util"
	"github.com/krelinga/go-tmdb/tmdbmovie"
)

func TestGetReleaseDates(t *testing.T) {
	ctx := util.ContextWithAPIReadAccessToken(context.Background(), os.Getenv("TMDB_READ_ACCESS_TOKEN"))

	options := tmdbmovie.GetReleaseDatesOptions{}

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
