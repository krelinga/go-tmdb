package tmdbseason_test

import (
	"context"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/internal/util"
	"github.com/krelinga/go-tmdb/tmdbseason"
)

func TestGetDetails(t *testing.T) {
	ctx := util.ContextWithAPIReadAccessToken(context.Background(), os.Getenv("TMDB_READ_ACCESS_TOKEN"))

	options := tmdbseason.GetDetailsOptions{}

	httpReply, err := tmdbseason.GetDetails(ctx, 1399, 1, options)
	if err != nil {
		t.Fatalf("GetDetails failed: %v", err)
	}
	reply, err := tmdbseason.ParseGetDetailsReply(httpReply)
	if err != nil {
		t.Fatalf("ParseGetDetailsReply failed: %v", err)
	}

	if reply.Details == nil {
		t.Fatal("expected non-nil Details in reply")
	}
	if reply.Details.SeasonNumber == nil || *reply.Details.SeasonNumber != 1 {
		t.Errorf("unexpected SeasonNumber: %v", reply.Details.SeasonNumber)
	}
	if len(reply.Details.Episodes) == 0 {
		t.Fatal("expected at least one episode in the season")
	}
	if reply.Details.Episodes[0].SeriesID == nil || *reply.Details.Episodes[0].SeriesID != 1399 {
		t.Errorf("unexpected SeriesID in first episode: %v", reply.Details.Episodes[0].SeriesID)
	}

	t.Log("GetDetails reply:", reply)
}
