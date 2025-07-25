package tmdbmovie_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	tmdbmovie "github.com/krelinga/go-tmdb/movie"
)

func TestGetDetails(t *testing.T) {
	ctx := context.Background()

	options := tmdbmovie.GetDetailsOptions{
		ReadAccessToken:    os.Getenv("TMDB_READ_ACCESS_TOKEN"),
		AppendCredits:      true,
		AppendExternalIDs:  true,
		AppendReleaseDates: true,
	}

	httpReply, err := tmdbmovie.GetDetails(ctx, http.DefaultClient, 11, options)
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
