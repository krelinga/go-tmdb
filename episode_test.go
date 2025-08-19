package tmdb_test

import (
	"context"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetEpisode(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	_, err := tmdb.GetEpisode(context.Background(), client, 1399, 1, 1, tmdb.WithAppendToResponse("credits", "external_ids"))
	if err != nil {
		t.Fatalf("failed to get episode: %v", err)
	}
}
