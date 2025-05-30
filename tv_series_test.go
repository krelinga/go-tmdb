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
	assert.Equal(t, tmdb.DateYYYYMMDD("2019-05-19"), tv.LastAirDate, "Unexpected last air date")
	expectedLastEpisode := &tmdb.TvEpisodeSum{
		TvEpisodeId:     1551830,
		Name:            "The Iron Throne",
		Overview:        "In the aftermath of the devastating attack on King's Landing, Daenerys must face the survivors.",
		VoteAverage:     4.544,
		VoteCount:       353,
		AirDate:         tmdb.DateYYYYMMDD("2019-05-19"),
		TvEpisodeNumber: 6,
		EpisodeType:     "finale",
		ProductionCode:  "806",
		Runtime:         tmdb.Minutes(80),
		TvSeasonNumber:  8,
		TvSeriesId:      1399,
		StillImage:      tmdb.StillImage("/zBi2O5EJfgTS6Ae0HdAYLm9o2nf.jpg"),
	}
	assert.Equal(t, expectedLastEpisode, tv.LastEpisodeToAir, "Unexpected last episode to air")
	assert.Nil(t, tv.NextEpisodeToAir, "Next episode to air should be nil for completed series")
	expectedNetworks := []*tmdb.TvNetwork{
		{
			TvNetworkId:   49,
			Name:          "HBO",
			OriginCountry: "US",
			LogoImage:     tmdb.LogoImage("/tuomPhY2UtuPTqqFnKMVHvSb724.png"),
		},
	}
	assert.Equal(t, expectedNetworks, tv.TvNetworks, "Unexpected TV networks")
	assert.Equal(t, 73, tv.NumberOfEpisodes, "Unexpected number of episodes")
	assert.Equal(t, 8, tv.NumberOfSeasons, "Unexpected number of seasons")
	expectedProductionCompanies := []*tmdb.ProductionCompanySum{
		{
			ProductionCompanyId: 76043,
			LogoImage:           tmdb.LogoImage("/9RO2vbQ67otPrBLXCaC8UMp3Qat.png"),
			Name:                "Revolution Sun Studios",
			OriginCountry:       "US",
		},
		{
			ProductionCompanyId: 12525,
			LogoImage:           "",
			Name:                "Television 360",
			OriginCountry:       "",
		},
	}
	for _, epc := range expectedProductionCompanies {
		assert.Contains(t, tv.ProductionCompanySums, epc, "Production companies should contain expected company: %v", epc)
		if epc.LogoImage != "" {
			checkLogoImage(t, epc.LogoImage, config)
		}
	}
	expectedProductionCountries := []*tmdb.CountrySum{
		{
			CountryIso3166_1: "GB",
			EnglishName:      "United Kingdom",
		},
		{
			CountryIso3166_1: "US",
			EnglishName:      "United States of America",
		},
	}
	assert.Equal(t, expectedProductionCountries, tv.ProductionCountries, "Unexpected production countries")
	expectedSeasons := []*tmdb.TvSeasonSum{
		{
			AirDate:        tmdb.DateYYYYMMDD("2010-12-05"),
			EpisodeCount:   283,
			TvSeasonId:     3627,
			Name:           "Specials",
			Overview:       "",
			PosterImage:    tmdb.PosterImage("/aos6lC1JGYt6ZRL85lgstNsfSeY.jpg"),
			TvSeasonNumber: 0,
			VoteAverage:    0.0,
		},
		{
			AirDate:        tmdb.DateYYYYMMDD("2011-04-17"),
			EpisodeCount:   10,
			TvSeasonId:     3624,
			Name:           "Season 1",
			Overview:       "Trouble is brewing in the Seven Kingdoms of Westeros. For the driven inhabitants of this visionary world, control of Westeros' Iron Throne holds the lure of great power. But in a land where the seasons can last a lifetime, winter is coming...and beyond the Great Wall that protects them, an ancient evil has returned. In Season One, the story centers on three primary areas: the Stark and the Lannister families, whose designs on controlling the throne threaten a tenuous peace; the dragon princess Daenerys, heir to the former dynasty, who waits just over the Narrow Sea with her malevolent brother Viserys; and the Great Wall--a massive barrier of ice where a forgotten danger is stirring.",
			PosterImage:    tmdb.PosterImage("/wgfKiqzuMrFIkU1M68DDDY8kGC1.jpg"),
			TvSeasonNumber: 1,
			VoteAverage:    8.4,
		},
	}
	for _, es := range expectedSeasons {
		assert.Contains(t, tv.TvSeasons, es, "TV seasons should contain expected season: %v", es)
		checkPosterImage(t, es.PosterImage, config)
	}
	expectedSpokenLanguages := []*tmdb.Language{
		{
			EnglishName: "English",
			Iso639_1:    "en",
			Name:        "English",
		},
	}
	assert.Equal(t, expectedSpokenLanguages, tv.SpokenLanguages, "Unexpected spoken languages")
}
