package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
	"github.com/stretchr/testify/assert"
)

func TestGetTvEpisode(t *testing.T) {
	const Series tmdb.TvSeriesId = 1399 // "Game of Thrones"
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
	assert.Equal(t, tmdb.TvEpisodeId(63056), episode.TvEpisodeId, "TvEpisodeId should match")
	assert.Equal(t, "101", episode.ProductionCode, "ProductionCode should match")
	assert.Equal(t, "standard", episode.EpisodeType, "EpisodeType should be 'standard'")
	assert.Equal(t, "Winter Is Coming", episode.Name, "Episode name should match")
	const expectedOverview = "Jon Arryn, the Hand of the King, is dead. King Robert Baratheon plans to ask his oldest friend, Eddard Stark, to take Jon's place. Across the sea, Viserys Targaryen plans to wed his sister to a nomadic warlord in exchange for an army."
	assert.Equal(t, expectedOverview, episode.Overview, "Episode overview should match")
	assert.Equal(t, tmdb.Minutes(62), episode.Runtime, "Runtime should be 62 minutes")
}
