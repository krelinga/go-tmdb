package tmdb_test

import (
	"context"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetShow(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	_, err := tmdb.GetShow(context.Background(), client, 1399, tmdb.WithAppendToResponse("aggregate_credits", "content_ratings", "credits", "external_ids", "keywords"))
	if err != nil {
		t.Fatalf("failed to get show: %v", err)
	}
}