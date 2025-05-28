package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
	"github.com/stretchr/testify/assert"
)

func TestGetMovie(t *testing.T) {
	config, err := tmdb.GetConfiguration(getClient(t))
	if err != nil {
		t.Fatalf("GetConfiguration failed: %v", err)
	}
	found, err := tmdb.GetMovie(getClient(t), 550,
		tmdb.WithKeywords(),
		tmdb.WithCredits(),
		tmdb.WithReleaseDates(),
	)
	if err != nil {
		t.Fatalf("GetMovie failed: %v", err)
	}

	expectedMovie := &tmdb.Movie{
		MovieSum: tmdb.MovieSum{
			Adult:            false,
			BackdropImage:    tmdb.BackdropImage("/xRyINp9KfMLVjRiO5nCsoRDdvvF.jpg"),
			MovieId:          550,
			OriginalLanguage: "en",
			OriginalTitle:    "Fight Club",
			Overview:         "A ticking-time-bomb insomniac and a slippery soap salesman channel primal male aggression into a shocking new form of therapy. Their concept catches on, with underground \"fight clubs\" forming in every town, until an eccentric gets in the way and ignites an out-of-control spiral toward oblivion.",
			Popularity:       30.8279,
			PosterImage:      "/pB8BM7pdSp6B6Ih7QZ4DrQ3PmJK.jpg",
			ReleaseDate:      "1999-10-15",
			Title:            "Fight Club",
			Video:            false,
			VoteAverage:      8.438,
			VoteCount:        30289,
		},
		BelongsToCollection: "",
		Budget:              63000000,
		Genres: []*tmdb.Genre{
			{GenreId: 18, Name: "Drama"},
		},
		Homepage: "http://www.foxmovies.com/movies/fight-club",
		ImdbId:   "tt0137523",
		ProductionCompanyShorts: []*tmdb.ProductionCompanySum{
			{
				Id:            711,
				LogoImage:     tmdb.LogoImage("/tEiIH5QesdheJmDAqQwvtN60727.png"),
				Name:          "Fox 2000 Pictures",
				OriginCountry: "US",
			},
			{
				Id:            508,
				LogoImage:     tmdb.LogoImage("/7cxRWzi4LsVm4Utfpr1hfARNurT.png"),
				Name:          "Regency Enterprises",
				OriginCountry: "US",
			},
			{
				Id:            4700,
				LogoImage:     tmdb.LogoImage("/A32wmjrs9Psf4zw0uaixF0GXfxq.png"),
				Name:          "Linson Entertainment",
				OriginCountry: "US",
			},
			{
				Id:            25,
				LogoImage:     tmdb.LogoImage("/qZCc1lty5FzX30aOCVRBLzaVmcp.png"),
				Name:          "20th Century Fox",
				OriginCountry: "US",
			},
			{
				Id:            20555,
				LogoImage:     tmdb.LogoImage("/hD8yEGUBlHOcfHYbujp71vD8gZp.png"),
				Name:          "Taurus Film",
				OriginCountry: "DE",
			},
		},
		ProductionCountrySummaries: []*tmdb.CountrySum{
			{
				CountryIso3166_1: "DE",
				EnglishName:      "Germany",
			},
			{
				CountryIso3166_1: "US",
				EnglishName:      "United States of America",
			},
		},
		Revenue: 100853753,
		Runtime: 139,
		SpokenLanguages: []*tmdb.Language{
			{
				Iso639_1:    "en",
				Name:        "English",
				EnglishName: "English",
			},
		},
		Status:  "Released",
		Tagline: "Mischief. Mayhem. Soap.",
	}
	assert.Equal(t, expectedMovie, found.Movie, "Expected movie details do not match")
	for _, pc := range found.Movie.ProductionCompanyShorts {
		checkLogoImage(t, pc.LogoImage, config)
	}
	checkDate(t, 1999, 10, 15, found.Movie.ReleaseDate)

	expectedKeywords := &tmdb.MovieKeywords{
		MovieId: 0, // TODO: looks like the API doesn't populate this field.  Should I set it anyway?
		Keywords: []*tmdb.Keyword{
			{KeywordId: 851, Name: "dual identity"},
			{KeywordId: 3927, Name: "rage and hate"},
			{KeywordId: 818, Name: "based on novel or book"},
			{KeywordId: 1541, Name: "nihilism"},
			{KeywordId: 1721, Name: "fight"},
			{KeywordId: 825, Name: "support group"},
			{KeywordId: 4565, Name: "dystopia"},
			{KeywordId: 4142, Name: "insomnia"},
			{KeywordId: 9181, Name: "alter ego"},
			{KeywordId: 11687, Name: "breaking the fourth wall"},
			{KeywordId: 156761, Name: "split personality"},
			{KeywordId: 179173, Name: "quitting a job"},
			{KeywordId: 212803, Name: "dissociative identity disorder"},
			{KeywordId: 260426, Name: "self destructiveness"},
		},
	}
	assert.Equal(t, expectedKeywords, found.Keywords, "Expected keywords do not match")

	expectedCastSubset := []*tmdb.MovieCast{
		{
			CastPerson: tmdb.CastPerson{
				CreditPerson: tmdb.CreditPerson{
					PersonSum: tmdb.PersonSum{
						Adult:              false,
						Gender:             tmdb.GenderMale,
						KnownForDepartment: "Acting",
						Name:               "Edward Norton",
						PersonId:           819,
						Popularity:         8.3963,
						ProfileImage:       tmdb.ProfileImage("/8nytsqL59SFJTVYVrN72k6qkGgJ.jpg"),
					},
					OriginalName: "Edward Norton",
					CreditId:     "52fe4250c3a36847f80149f3",
				},
				Character: "Narrator",
				Order:     0,
			},
			MovieCastId: 4,
		},
		{
			CastPerson: tmdb.CastPerson{
				CreditPerson: tmdb.CreditPerson{
					PersonSum: tmdb.PersonSum{
						Adult:              false,
						Gender:             tmdb.GenderMale,
						KnownForDepartment: "Acting",
						Name:               "Brad Pitt",
						PersonId:           287,
						Popularity:         15.325,
						ProfileImage:       tmdb.ProfileImage("/cckcYc2v0yh1tc9QjRelptcOBko.jpg"),
					},
					OriginalName: "Brad Pitt",
					CreditId:     "52fe4250c3a36847f80149f7",
				},
				Character: "Tyler Durden",
				Order:     1,
			},
			MovieCastId: 5,
		},
		{
			CastPerson: tmdb.CastPerson{
				CreditPerson: tmdb.CreditPerson{
					PersonSum: tmdb.PersonSum{
						Adult:              false,
						Gender:             tmdb.GenderFemale,
						KnownForDepartment: "Acting",
						Name:               "Helena Bonham Carter",
						PersonId:           1283,
						Popularity:         5.9771,
						ProfileImage:       tmdb.ProfileImage("/hJMbNSPJ2PCahsP3rNEU39C8GWU.jpg"),
					},
					OriginalName: "Helena Bonham Carter",
					CreditId:     "631f0de8bd32090082733691",
				},
				Character: "Marla Singer",
				Order:     2,
			},
			MovieCastId: 285,
		},
	}
	for _, ec := range expectedCastSubset {
		assert.Contains(t, found.Credits.Cast, ec, "Expected cast member not found: %v", ec)
		checkProfileImage(t, ec.ProfileImage, config)
	}

	expectedCrewSubset := []*tmdb.MovieCrew{
		{
			CreditPerson: tmdb.CreditPerson{
				PersonSum: tmdb.PersonSum{
					Adult:              false,
					Gender:             tmdb.GenderMale,
					KnownForDepartment: "Directing",
					Name:               "David Fincher",
					PersonId:           7467,
					Popularity:         9.2504,
					ProfileImage:       "/tpEczFclQZeKAiCeKZZ0adRvtfz.jpg",
				},
				OriginalName: "David Fincher",
				CreditId:     "631f0289568463007bbe28a5",
			},
			Department: "Directing",
			Job:        "Director",
		},
		{
			CreditPerson: tmdb.CreditPerson{
				PersonSum: tmdb.PersonSum{
					Adult:              false,
					Gender:             tmdb.GenderMale,
					KnownForDepartment: "Production",
					Name:               "Arnon Milchan",
					PersonId:           376,
					Popularity:         2.9801,
					ProfileImage:       "/b2hBExX4NnczNAnLuTBF4kmNhZm.jpg",
				},
				OriginalName: "Arnon Milchan",
				CreditId:     "55731b8192514111610027d7",
			},
			Department: "Production",
			Job:        "Executive Producer",
		},
	}
	for _, ec := range expectedCrewSubset {
		assert.Contains(t, found.Credits.Crew, ec, "Expected crew member not found: %v", ec)
		checkProfileImage(t, ec.ProfileImage, config)
	}

	premiere := &tmdb.MovieReleaseDate{
		Certification:    "",
		Language:         "",
		Note:             "CMJ Film Festival",
		ReleaseDate:      "1999-09-21T00:00:00.000Z",
		MovieReleaseType: tmdb.MovieReleaseTypePremiere,
	}
	checkDate(t, 1999, 9, 21, premiere.ReleaseDate)
	theatrical := &tmdb.MovieReleaseDate{
		Certification:    "R",
		Language:         "",
		Note:             "",
		ReleaseDate:      "1999-10-15T00:00:00.000Z",
		MovieReleaseType: tmdb.MovieReleaseTypeTheatrical,
	}
	checkDate(t, 1999, 10, 15, theatrical.ReleaseDate)
	vhs := &tmdb.MovieReleaseDate{
		Certification:    "R",
		Language:         "",
		Note:             "VHS",
		ReleaseDate:      "2000-04-25T00:00:00.000Z",
		MovieReleaseType: tmdb.MovieReleaseTypePhysical,
	}
	checkDate(t, 2000, 4, 25, vhs.ReleaseDate)
	for _, rc := range found.Releases.MovieReleaseCountries {
		if rc.Iso3166_1 != "US" {
			continue // We only care about the US release dates.
		}
		for _, erd := range []*tmdb.MovieReleaseDate{premiere, theatrical, vhs} {
			assert.Contains(t, rc.MovieReleaseDates, erd, "Expected US release date not found: %v", erd)
		}
	}
}

func TestHttpError(t *testing.T) {
	_, err := tmdb.GetMovie(getClient(t), 0) // Invalid movie ID
	if err == nil {
		t.Fatal("Expected error for invalid movie ID, got nil")
	}
	t.Logf("Expected error: %v", err)
}
