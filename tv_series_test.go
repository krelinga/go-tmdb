package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
	"github.com/stretchr/testify/assert"
)

func TestGetTvSeries(t *testing.T) {
	config, err := tmdb.GetConfiguration(getClient(t))
	if err != nil {
		t.Fatalf("GetConfiguration failed: %v", err)
	}
	tv, err := tmdb.GetTvSeries(getClient(t), 1399)
	if err != nil {
		t.Fatalf("GetTVSeries failed: %v", err)
	}
	if tv == nil {
		t.Fatal("GetTVSeries returned nil")
	}

	assert.False(t, tv.Adult, "TV series should not be marked as adult")
	assert.Equal(t, tmdb.BackdropImage("/zZqpAXxVSBtxV9qPBcscfXBcL2w.jpg"), tv.BackdropImage, "Unexpected backdrop image")
	assert.Equal(t, tmdb.TvSeriesId(1399), tv.TvSeriesId, "Unexpected TV series ID")
	assert.Equal(t, []tmdb.CountryIso3166_1{"US"}, tv.OriginCountries, "Unexpected origin countries")
	assert.Equal(t, "en", tv.OriginalLanguage, "Unexpected original language")
	assert.Equal(t, "Game of Thrones", tv.OriginalName, "Unexpected original name")
	overview := "Seven noble families fight for control of the mythical land of Westeros. Friction between the houses leads to full-scale war. All while a very ancient evil awakens in the farthest north. Amidst the war, a neglected military order of misfits, the Night's Watch, is all that stands between the realms of men and icy horrors beyond."
	assert.Equal(t, overview, tv.Overview, "Unexpected overview")
	assert.Equal(t, 186.7642, tv.Popularity, "Unexpected popularity")
	assert.Equal(t, tmdb.PosterImage("/1XS1oqL89opfnbLl8WnZY1O1uJx.jpg"), tv.PosterImage, "Unexpected poster image")
	checkPosterImage(t, tv.PosterImage, config)
	assert.Equal(t, tmdb.DateYYYYMMDD("2011-04-17"), tv.FirstAirDate, "Unexpected first air date")
	assert.Equal(t, "Game of Thrones", tv.Name, "Unexpected name")
	assert.Equal(t, 8.456, tv.VoteAverage, "Unexpected vote average")
	assert.Equal(t, 25031, tv.VoteCount, "Unexpected vote count")

	expectedCreatorSubset := []*tmdb.TvSeriesCreator{
		{
			PersonCore: tmdb.PersonCore{
				PersonId:     9813,
				Name:         "David Benioff",
				ProfileImage: "/bOlW8pymCeQLfwPIvc2D1MRcUoF.jpg",
				Gender:       tmdb.GenderMale,
			},
			CreditId: tmdb.CreditId("5256c8c219c2956ff604858a"),
		},
	}
	for _, ec := range expectedCreatorSubset {
		assert.Contains(t, tv.CreatedBy, ec, "Creators should contain expected creator: %v", ec)
		checkProfileImage(t, ec.ProfileImage, config)
	}

	assert.Empty(t, tv.EpisodeRunTimes, "Episode run times should be empty for TV series")
	expectedGenres := []*tmdb.Genre{
		{
			GenreId: 10765,
			Name:    "Sci-Fi & Fantasy",
		},
		{
			GenreId: 18,
			Name:    "Drama",
		},
		{
			GenreId: 10759,
			Name:    "Action & Adventure",
		},
	}
	assert.Equal(t, expectedGenres, tv.Genres, "Unexpected genres")
	assert.Equal(t, "https://www.hbo.com/game-of-thrones", tv.Homepage, "Unexpected homepage")
	assert.False(t, tv.InProduction, "TV series should not be in production")
	assert.Equal(t, []string{"en"}, tv.Languages, "Unexpected languages")
}
