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
	config, err := tmdb.GetConfiguration(getClient(t))
	if err != nil {
		t.Fatalf("GetConfiguration failed: %v", err)
	}
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
	assert.Equal(t, tmdb.DateYYYYMMDD("2011-04-17"), episode.AirDate, "AirDate should match")
	checkDate(t, 2011, 4, 17, episode.AirDate)
	assert.Equal(t, tmdb.StillImage("/wrGWeW4WKxnaeA8sxJb2T9O6ryo.jpg"), episode.StillImage, "StillImage should match")
	checkStillImage(t, episode.StillImage, config)
	assert.Equal(t, 8.063, episode.VoteAverage, "VoteAverage should match")
	assert.Equal(t, 374, episode.VoteCount, "VoteCount should match")

	expectedCrewSubset := []*tmdb.CrewPerson{
		{
			CreditPerson: tmdb.CreditPerson{
				PersonSum: tmdb.PersonSum{
					Adult:              false,
					Gender:             tmdb.GenderMale,
					PersonId:           9813,
					KnownForDepartment: "Writing",
					Name:               "David Benioff",
					Popularity:         3.065,
					ProfileImage:       tmdb.ProfileImage("/bOlW8pymCeQLfwPIvc2D1MRcUoF.jpg"),
				},
				CreditId:     "5256c8a019c2956ff6046e2b",
				OriginalName: "David Benioff",
			},
			Job:        "Writer",
			Department: "Writing",
		},
	}
	for _, ec := range expectedCrewSubset {
		assert.Contains(t, episode.Crew, ec, "Expected crew member not found: %v", ec)
		checkProfileImage(t, ec.ProfileImage, config)
	}
}
