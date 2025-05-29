package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
	"github.com/stretchr/testify/assert"
)

func TestSearchTvSeries(t *testing.T) {
	client := getClient(t)
	const query = "Game of"
	results := []*tmdb.TvSeriesSearchResult{}
	for result, err := range tmdb.SearchTvSeries(client, query) {
		if err != nil {
			t.Fatal("SearchTvSeries returned an error:", err)
		}
		if result == nil {
			t.Fatal("SearchTvSeries returned a nil result")
		}
		results = append(results, result)
	}
	gameOfThrones := &tmdb.TvSeriesSearchResult{
		TvSeriesSum: tmdb.TvSeriesSum{
			Adult:            false,
			BackdropImage:    "/zZqpAXxVSBtxV9qPBcscfXBcL2w.jpg",
			TvSeriesId:       1399,
			OriginCountries:  []tmdb.CountryIso3166_1{"US"},
			OriginalLanguage: "en",
			OriginalName:     "Game of Thrones",
			Overview:         "Seven noble families fight for control of the mythical land of Westeros. Friction between the houses leads to full-scale war. All while a very ancient evil awakens in the farthest north. Amidst the war, a neglected military order of misfits, the Night's Watch, is all that stands between the realms of men and icy horrors beyond.",
			Popularity:       178.9799,
			PosterImage:      "/1XS1oqL89opfnbLl8WnZY1O1uJx.jpg",
			FirstAirDate:     "2011-04-17",
			Name:             "Game of Thrones",
			VoteAverage:      8.456,
			VoteCount:        25046,
		},
		GenreIds: []tmdb.GenreId{10765, 18, 10759},
	}
	assert.Contains(t, results, gameOfThrones, "SearchTvSeries did not return expected result")
}
