package tmdbmovie_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/tmdbmovie"
)

func TestGetExternalIDs(t *testing.T) {
	ctx := context.Background()

	options := tmdbmovie.GetExternalIDsOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	}

	httpReply, err := tmdbmovie.GetExternalIDs(ctx, http.DefaultClient, 11, options)
	if err != nil {
		t.Fatalf("GetExternalIDs failed: %v", err)
	}
	reply, err := tmdbmovie.ParseGetExternalIDsReply(httpReply)
	if err != nil {
		t.Fatalf("ParseGetExternalIDsReply failed: %v", err)
	}

	if reply.ID == nil || *reply.ID != 11 {
		t.Errorf("unexpected ID: %v", reply.ID)
	}
	t.Log("GetExternalIDs reply:", reply)
}
