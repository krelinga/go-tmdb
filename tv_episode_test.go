package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetTvEpisode(t *testing.T) {
	episode, err := tmdb.GetTvEpisode(getClient(t), 1399, 1, 1, tmdb.WithExternalIds())
	if err != nil {
		t.Fatalf("GetTvEpisode failed: %v", err)
	}
	if episode == nil {
		t.Fatal("GetTvEpisode returned nil")
	}
}
