package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestSearchTvSeries(t *testing.T) {
	client := getClient(t)
	const query = "Game of"
	for result, err := range tmdb.SearchTvSeries(client, query) {
		if err != nil {
			t.Fatal("SearchTvSeries returned an error:", err)
		}
		if result == nil {
			t.Fatal("SearchTvSeries returned a nil result")
		}
	}
}
