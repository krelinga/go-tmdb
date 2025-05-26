package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetTvSeason(t *testing.T) {
	var raw []byte
	tv, err := tmdb.GetTvSeason(globalClient, 1399, 1,
		tmdb.WithRawReply(&raw),
	)
	if err != nil {
		t.Fatalf("GetTVSeason failed: %v", err)
	}
	if tv == nil {
		t.Fatal("GetTVSeason returned nil")
	}
	if tv.Name == "" {
		t.Fatal("TV series name is empty")
	}
	if tv.TvSeasonNumber == 0 {
		t.Fatal("Season number is zero")
	}
	t.Logf("TV season name: %s\n", tv.Name)
	t.Logf("TV season number: %d\n", tv.TvSeasonNumber)
	for _, e := range tv.TvEpisodes {
		if e == nil {
			t.Fatal("TvEpisode is nil")
		}
		if e.Name == "" {
			t.Fatal("TvEpisode name is empty")
		}
		t.Logf("Episode: %s (ID: %d)\n", e.Name, e.TvEpisodeId)
	}
}
