package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
	"github.com/stretchr/testify/assert"
)

func TestSearchMovies(t *testing.T) {
	var found *tmdb.MovieSearchResult
	count := 0
	for m, err := range tmdb.SearchMovies(getClient(t), "Inception") {
		if err != nil {
			t.Fatalf("SearchMovies failed: %v", err)
		}
		count++
		if found != nil {
			continue // We only want the first result.
		}
		found = m
	}

	assert.Equal(t, 9, count, "SearchMovies did not return expected number of results")

	expected := &tmdb.MovieSearchResult{
		MovieShort: tmdb.MovieShort{
			Adult:            false,
			BackdropImage:    tmdb.BackdropImage("/8ZTVqvKDQ8emSGUEMjsS4yHAwrp.jpg"),
			MovieId:          27205,
			OriginalLanguage: "en",
			OriginalTitle:    "Inception",
			Overview:         "Cobb, a skilled thief who commits corporate espionage by infiltrating the subconscious of his targets is offered a chance to regain his old life as payment for a task considered to be impossible: \"inception\", the implantation of another person's idea into a target's subconscious.",
			Popularity:       28.2038,
			PosterImage:      tmdb.PosterImage("/oYuLEt3zVCKq57qu2F8dT7NIa6f.jpg"),
			ReleaseDate:      "2010-07-15",
			Title:            "Inception",
			Video:            false,
			VoteAverage:      8.369,
			VoteCount:        37482,
		},
		GenereIds: []tmdb.GenereId{28, 878, 12},
	}
	assert.Equal(t, expected, found, "SearchMovies did not return expected result")

	config, err := tmdb.GetConfiguration(getClient(t))
	if err != nil {
		t.Fatalf("GetConfiguration failed: %v", err)
	}

	checkBackdropImage(t, found.BackdropImage, config)
	checkPosterImage(t, found.PosterImage, config)
	checkDate(t, 2010, 7, 15, found.ReleaseDate)
}
