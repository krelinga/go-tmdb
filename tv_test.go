package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetTVSeries(t *testing.T) {
	client := getClient(t)
	var raw []byte
	tv, err := tmdb.GetTvSeries(client, 1399,
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