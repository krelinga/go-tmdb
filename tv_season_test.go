package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetTvSeason(t *testing.T) {
	client := getClient(t)
	var raw []byte
	tv, err := tmdb.GetTvSeason(client, 1399, 1,
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
}
