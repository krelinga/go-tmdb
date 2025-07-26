package tmdbseries_test

import (
	"context"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/internal/util"
	"github.com/krelinga/go-tmdb/tmdbseries"
)

func TestGetDetails(t *testing.T) {
	ctx := util.ContextWithAPIReadAccessToken(context.Background(), os.Getenv("TMDB_READ_ACCESS_TOKEN"))

	options := tmdbseries.GetDetailsOptions{}

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
