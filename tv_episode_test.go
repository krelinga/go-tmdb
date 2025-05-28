package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
	"github.com/stretchr/testify/assert"
)

func TestGetTvEpisode(t *testing.T) {
	const Series tmdb.TvSeriesId= 1399 // "Game of Thrones"
	const SeasonNumber tmdb.TvSeasonNumber = 1
	const EpisodeNumber tmdb.TvEpisodeNumber = 1
	episode, err := tmdb.GetTvEpisode(getClient(t), Series, SeasonNumber, EpisodeNumber, tmdb.WithExternalIds())
	if err != nil {
		t.Fatalf("GetTvEpisode failed: %v", err)
	}
	if episode == nil {
		t.Fatal("GetTvEpisode returned nil")
	}

	assert.Equal(t, Series, episode.TvSeriesId, "TvSeriesId should match")
	assert.Equal(t, SeasonNumber, episode.TvSeasonNumber, "TvSeasonNumber should match")
	assert.Equal(t, EpisodeNumber, episode.TvEpisodeNumber, "TvEpisodeNumber should match")
}
