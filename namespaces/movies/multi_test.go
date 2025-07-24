package movies_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/namespaces/movies"
)

func TestGetMulti(t *testing.T) {
	ctx := context.Background()

	options := movies.GetMultiOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
		WantDetails:     true,
		WantCredits:     true,
	}

	reply, err := movies.GetMulti(ctx, http.DefaultClient, 11, options)
	if err != nil {
		t.Fatalf("GetMulti failed: %v", err)
	}

	if reply.ID == nil || *reply.ID != 11 {
		t.Errorf("unexpected ID: %v", reply.ID)
	}
	t.Log("GetMulti reply:", reply)
}
