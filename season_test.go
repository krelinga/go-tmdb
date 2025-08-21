package tmdb_test

import (
	"context"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetSeason(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	season, err := tmdb.GetSeason(context.Background(), client, 1399, 1, tmdb.WithAppendToResponse("aggregate_credits", "credits", "external_ids"))
	if err != nil {
		t.Fatalf("failed to get season: %v", err)
	}

	checkField(t, "2011-04-17", season, tmdb.Season.AirDate)
}