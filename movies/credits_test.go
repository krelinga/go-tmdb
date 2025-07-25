package movies_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/movies"
)

func TestGetCredits(t *testing.T) {
	ctx := context.Background()

	options := movies.GetCreditsOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	}

	reply, err := movies.GetCredits(ctx, http.DefaultClient, 11, options)
	if err != nil {
		t.Fatalf("GetCredits failed: %v", err)
	}

	if reply.ID == nil || *reply.ID != 11 {
		t.Errorf("unexpected ID: %v", reply.ID)
	}
	t.Log("GetCredits reply:", reply)
}