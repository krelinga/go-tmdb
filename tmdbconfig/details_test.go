package tmdbconfig_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb/tmdbconfig"
)

func TestGetDetails(t *testing.T) {
	ctx := context.Background()

	options := tmdbconfig.GetDetailsOptions{
		ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	}

	httpReply, err := tmdbconfig.GetDetails(ctx, http.DefaultClient, options)
	if err != nil {
		t.Fatalf("GetDetails failed: %v", err)
	}
	reply, err := tmdbconfig.ParseGetDetailsReply(httpReply)
	if err != nil {
		t.Fatalf("ParseGetDetailsReply failed: %v", err)
	}

	t.Log("GetDetails reply:", reply)
}