package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetTvSeries(t *testing.T) {
	var raw []byte
	tv, err := tmdb.GetTvSeries(getClient(t), 1399,
		tmdb.WithRawReply(&raw))
	if err != nil {
		t.Fatalf("GetTVSeries failed: %v", err)
	}
	if tv == nil {
		t.Fatal("GetTVSeries returned nil")
	}
	if tv.Name == "" {
		t.Fatal("TV series name is empty")
	}
	t.Log(string(raw))
}
