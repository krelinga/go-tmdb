package tmdbepisode_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/tmdbepisode"
)

func TestGetDetails(t *testing.T) {
	ctx := context.Background()

	options := tmdbepisode.GetDetailsOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	}

	httpReply, err := tmdbepisode.GetDetails(ctx, http.DefaultClient, 1399, 1, 1, options)
	if err != nil {
		t.Fatalf("GetDetails failed: %v", err)
	}
	reply, err := tmdbepisode.ParseGetDetailsReply(httpReply)
	if err != nil {
		t.Fatalf("ParseGetDetailsReply failed: %v", err)
	}

	if reply.Details == nil {
		t.Fatal("expected non-nil Details in reply")
	}
	if reply.Details.SeasonNumber == nil || *reply.Details.SeasonNumber != 1 {
		t.Errorf("unexpected SeasonNumber: %v", reply.Details.SeasonNumber)
	}
	if reply.Details.EpisodeNumber == nil || *reply.Details.EpisodeNumber != 1 {
		t.Errorf("unexpected EpisodeNumber: %v", reply.Details.EpisodeNumber)
	}

	t.Log("GetDetails reply:", reply)
}
