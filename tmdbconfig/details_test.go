package tmdbconfig_test

import (
	"context"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/internal/util"
	"github.com/krelinga/go-tmdb/tmdbconfig"
)

func TestGetDetails(t *testing.T) {
	ctx := util.ContextWithAPIReadAccessToken(context.Background(), os.Getenv("TMDB_READ_ACCESS_TOKEN"))

	options := tmdbconfig.GetDetailsOptions{}

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
