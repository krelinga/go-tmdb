package movies_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/movies"
)

func TestGetReleaseDates(t *testing.T) {
	ctx := context.Background()

	options := movies.GetReleaseDatesOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	}

	reply, err := movies.GetReleaseDates(ctx, http.DefaultClient, 11, options)
	if err != nil {
		t.Fatalf("GetReleaseDates failed: %v", err)
	}

	if reply.ID == nil || *reply.ID != 11 {
		t.Errorf("unexpected ID: %v", reply.ID)
	}
	t.Log("GetReleaseDates reply:", reply)
}
