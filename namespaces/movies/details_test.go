package movies_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/namespaces/movies"
)

func TestGetDetails(t *testing.T) {
	ctx := context.Background()

	options := movies.GetDetailsOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	}

	reply, err := movies.GetDetails(ctx, http.DefaultClient, 11, options)
	if err != nil {
		t.Fatalf("GetDetails failed: %v", err)
	}

	if reply.ID == nil || *reply.ID != 11 {
		t.Errorf("unexpected ID: %v", reply.ID)
	}
	t.Log("GetDetails reply:", reply)
}