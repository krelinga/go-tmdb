package tmdbseries_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/tmdbseries"
)

func TestGetDetails(t *testing.T) {
	ctx := context.Background()

	options := tmdbseries.GetDetailsOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	}

	httpReply, err := tmdbseries.GetDetails(ctx, http.DefaultClient, 1399, options)
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
