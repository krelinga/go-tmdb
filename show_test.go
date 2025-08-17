package tmdb_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetShow(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	show, err := tmdb.GetShow(context.Background(), client, 1399, tmdb.WithAppendToResponse("aggregate_credits", "content_ratings", "credits", "external_ids", "keywords"))
	if err != nil {
		t.Fatalf("failed to get show: %v", err)
	}

	checkField(t, false, show, tmdb.Show.Adult)
	checkField(t, "/2OMB0ynKlyIenMJWI2Dy9IWT4c.jpg", show, tmdb.Show.BackdropPath)
	if createdBy, err := show.CreatedBy(); err != nil {
		t.Errorf("failed to get created_by: %v", err)
	} else if len(createdBy) != 2 {
		t.Errorf("expected 2 created_by, got %d", len(createdBy))
	} else {
		if c, err := findCredit(createdBy, "David Benioff"); err != nil {
			t.Error(err)
		} else {
			checkField(t, int32(9813), c, tmdb.Credit.ID)
			checkField(t, "5256c8c219c2956ff604858a", c, tmdb.Credit.CreditID)
			checkField(t, "David Benioff", c, tmdb.Credit.Name)
			checkField(t, "David Benioff", c, tmdb.Credit.OriginalName)
			checkField(t, int32(2), c, tmdb.Credit.Gender)
			checkField(t, "/xvNN5huL0X8yJ7h3IZfGG4O2zBD.jpg", c, tmdb.Credit.ProfilePath)
		}
		if c, err := findCredit(createdBy, "D. B. Weiss"); err != nil {
			t.Error(err)
		} else {
			checkField(t, int32(228068), c, tmdb.Credit.ID)
			checkField(t, "552e611e9251413fea000901", c, tmdb.Credit.CreditID)
			checkField(t, "D. B. Weiss", c, tmdb.Credit.Name)
			checkField(t, "D. B. Weiss", c, tmdb.Credit.OriginalName)
			checkField(t, int32(2), c, tmdb.Credit.Gender)
			checkField(t, "/6Wt006TIQoDSSnl0YaKihfn3w7K.jpg", c, tmdb.Credit.ProfilePath)
		}
	}
	checkField(t, "2011-04-17", show, tmdb.Show.FirstAirDate)
	if genres, err := show.Genres(); err != nil {
		t.Errorf("failed to get genres: %v", err)
	} else if len(genres) != 3 {
		t.Errorf("expected 3 genres, got %d", len(genres))
	} else {
		if g, err := findGenre(genres, "Drama"); err != nil {
			t.Error(err)
		} else {
			checkField(t, int32(18), g, tmdb.Genre.ID)
			checkField(t, "Drama", g, tmdb.Genre.Name)
		}
		if g, err := findGenre(genres, "Sci-Fi & Fantasy"); err != nil {
			t.Error(err)
		} else {
			checkField(t, int32(10765), g, tmdb.Genre.ID)
			checkField(t, "Sci-Fi & Fantasy", g, tmdb.Genre.Name)
		}
		if g, err := findGenre(genres, "Action & Adventure"); err != nil {
			t.Error(err)
		} else {
			checkField(t, int32(10759), g, tmdb.Genre.ID)
			checkField(t, "Action & Adventure", g, tmdb.Genre.Name)
		}
	}
	checkField(t, "https://www.hbo.com/game-of-thrones", show, tmdb.Show.Homepage)
	checkField(t, int32(1399), show, tmdb.Show.ID)
	checkField(t, false, show, tmdb.Show.InProduction)
	if languages, err := show.Languages(); err != nil {
		t.Errorf("failed to get languages: %v", err)
	} else if len(languages) != 1 {
		t.Errorf("expected 1 language, got %d", len(languages))
	} else if languages[0] != "en" {
		t.Errorf("expected language 'en', got '%s'", languages[0])
	}
	checkField(t, "2019-05-19", show, tmdb.Show.LastAirDate)
	checkField(t, int32(1551830), show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.ID)
	checkField(t, "The Iron Throne", show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.Name)
	checkField(t, "In the aftermath of the devastating attack on King's Landing, Daenerys must face the survivors.", show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.Overview)
	checkField(t, 4.586, show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.VoteAverage)
	checkField(t, int32(366), show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.VoteCount)
	checkField(t, "2019-05-19", show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.AirDate)
	checkField(t, int32(6), show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.EpisodeNumber)
	checkField(t, "finale", show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.EpisodeType)
	checkField(t, "806", show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.ProductionCode)
	checkField(t, int32(80), show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.Runtime)
	checkField(t, int32(8), show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.SeasonNumber)
	checkField(t, int32(1399), show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.ShowID)
	checkField(t, "/zBi2O5EJfgTS6Ae0HdAYLm9o2nf.jpg", show, tmdb.Show.LastEpisodeToAir, tmdb.Episode.StillPath)
	checkField(t, "Game of Thrones", show, tmdb.Show.Name)
	if networks, err := show.Networks(); err != nil {
		t.Errorf("failed to get networks: %v", err)
	} else if len(networks) != 1 {
		t.Errorf("expected 1 network, got %d", len(networks))
	} else {
		checkField(t, int32(49), networks[0], tmdb.Company.ID)
		checkField(t, "HBO", networks[0], tmdb.Company.Name)
		checkField(t, "/tuomPhY2UtuPTqqFnKMVHvSb724.png", networks[0], tmdb.Company.LogoPath)
		checkField(t, "US", networks[0], tmdb.Company.OriginCountry)
	}
	checkField(t, int32(73), show, tmdb.Show.NumberOfEpisodes)
	checkField(t, int32(8), show, tmdb.Show.NumberOfSeasons)
	if originCountries, err := show.OriginCountry(); err != nil {
		t.Errorf("failed to get origin_countries: %v", err)
	} else if len(originCountries) != 1 {
		t.Errorf("expected 1 origin_country, got %d", len(originCountries))
	} else if originCountries[0] != "US" {
		t.Errorf("expected origin_country 'US', got '%s'", originCountries[0])
	}
	checkField(t, "en", show, tmdb.Show.OriginalLanguage)
	checkField(t, "Game of Thrones", show, tmdb.Show.OriginalName)
	checkField(t, "Seven noble families fight for control of the mythical land of Westeros. Friction between the houses leads to full-scale war. All while a very ancient evil awakens in the farthest north. Amidst the war, a neglected military order of misfits, the Night's Watch, is all that stands between the realms of men and icy horrors beyond.", show, tmdb.Show.Overview)
	checkField(t, 123.003, show, tmdb.Show.Popularity)
	checkField(t, "/1XS1oqL89opfnbLl8WnZY1O1uJx.jpg", show, tmdb.Show.PosterPath)
	if productionCompanies, err := show.ProductionCompanies(); err != nil {
		t.Errorf("failed to get production_companies: %v", err)
	} else {
		if rss, err := findCompany(productionCompanies, "Revolution Sun Studios"); err != nil {
			t.Error(err)
		} else {
			checkField(t, int32(76043), rss, tmdb.Company.ID)
			checkField(t, "Revolution Sun Studios", rss, tmdb.Company.Name)
			checkField(t, "/9RO2vbQ67otPrBLXCaC8UMp3Qat.png", rss, tmdb.Company.LogoPath)
			checkField(t, "US", rss, tmdb.Company.OriginCountry)
		}
		if t360, err := findCompany(productionCompanies, "Television 360"); err != nil {
			t.Error(err)
		} else {
			checkField(t, int32(12525), t360, tmdb.Company.ID)
			checkField(t, "Television 360", t360, tmdb.Company.Name)
			if _, err := t360.LogoPath(); err == nil {
				t.Error("expected empty logo_path")
			}
			checkField(t, "", t360, tmdb.Company.OriginCountry)
		}
	}
	if countries, err := show.ProductionCountries(); err != nil {
		t.Errorf("failed to get production_countries: %v", err)
	} else if len(countries) != 2 {
		t.Errorf("expected 2 production_countries, got %d", len(countries))
	} else {
		if us, err := findCountry(countries, "US"); err != nil {
			t.Error(err)
		} else {
			checkField(t, "US", us, tmdb.Country.ISO3166_1)
			checkField(t, "United States of America", us, tmdb.Country.Name)
		}
		if gb, err := findCountry(countries, "GB"); err != nil {
			t.Error(err)
		} else {
			checkField(t, "GB", gb, tmdb.Country.ISO3166_1)
			checkField(t, "United Kingdom", gb, tmdb.Country.Name)
		}
	}
	if seasons, err := show.Seasons(); err != nil {
		t.Errorf("failed to get seasons: %v", err)
	} else if len(seasons)-1 != 8 { // Season 0 is specials
		t.Errorf("expected 8 seasons, got %d", len(seasons)-1)
	} else {
		if s0, err := findSeason(seasons, 0); err != nil {
			t.Error(err)
		} else {
			checkField(t, "2010-12-05", s0, tmdb.Season.AirDate)
			checkField(t, int32(286), s0, tmdb.Season.EpisodeCount)
			checkField(t, int32(3627), s0, tmdb.Season.ID)
			checkField(t, "Specials", s0, tmdb.Season.Name)
			checkField(t, "", s0, tmdb.Season.Overview)
			checkField(t, "/aos6lC1JGYt6ZRL85lgstNsfSeY.jpg", s0, tmdb.Season.PosterPath)
			checkField(t, int32(0), s0, tmdb.Season.SeasonNumber)
			checkField(t, 0.0, s0, tmdb.Season.VoteAverage)
		}
		if s1, err := findSeason(seasons, 1); err != nil {
			t.Error(err)
		} else {
			checkField(t, "2011-04-17", s1, tmdb.Season.AirDate)
			checkField(t, int32(10), s1, tmdb.Season.EpisodeCount)
			checkField(t, int32(3624), s1, tmdb.Season.ID)
			checkField(t, "Season 1", s1, tmdb.Season.Name)
			checkField(t, "Trouble is brewing in the Seven Kingdoms of Westeros. For the driven inhabitants of this visionary world, control of Westeros' Iron Throne holds the lure of great power. But in a land where the seasons can last a lifetime, winter is coming...and beyond the Great Wall that protects them, an ancient evil has returned. In Season One, the story centers on three primary areas: the Stark and the Lannister families, whose designs on controlling the throne threaten a tenuous peace; the dragon princess Daenerys, heir to the former dynasty, who waits just over the Narrow Sea with her malevolent brother Viserys; and the Great Wall--a massive barrier of ice where a forgotten danger is stirring.", s1, tmdb.Season.Overview)
			checkField(t, "/wgfKiqzuMrFIkU1M68DDDY8kGC1.jpg", s1, tmdb.Season.PosterPath)
			checkField(t, int32(1), s1, tmdb.Season.SeasonNumber)
			checkField(t, 8.4, s1, tmdb.Season.VoteAverage)
		}
	}
	if languages, err := show.SpokenLanguages(); err != nil {
		t.Errorf("failed to get spoken_languages: %v", err)
	} else if len(languages) != 1 {
		t.Errorf("expected 1 spoken_language, got %d", len(languages))
	} else if en, err := findLanguage(languages, "en"); err != nil {
		t.Error(err)
	} else {
		checkField(t, "English", en, tmdb.Language.EnglishName)
		checkField(t, "en", en, tmdb.Language.ISO639_1)
		checkField(t, "English", en, tmdb.Language.Name)
	}
	checkField(t, "Ended", show, tmdb.Show.Status)
	checkField(t, "Winter is coming.", show, tmdb.Show.Tagline)
	checkField(t, "Scripted", show, tmdb.Show.Type)
	checkField(t, 8.456, show, tmdb.Show.VoteAverage)
	checkField(t, int32(25382), show, tmdb.Show.VoteCount)
	// TODO: checks for other fields.
}

func findCredit(in []tmdb.Credit, want string) (tmdb.Credit, error) {
	for _, c := range in {
		if name, err := c.Name(); err != nil {
			return tmdb.Credit{}, fmt.Errorf("failed to get credit name: %w", err)
		} else if name == want {
			return c, nil
		}
	}
	return tmdb.Credit{}, fmt.Errorf("no credit found with name: %s", want)
}

func findGenre(in []tmdb.Genre, want string) (tmdb.Genre, error) {
	for _, g := range in {
		if name, err := g.Name(); err != nil {
			return tmdb.Genre{}, fmt.Errorf("failed to get genre name: %w", err)
		} else if name == want {
			return g, nil
		}
	}
	return tmdb.Genre{}, fmt.Errorf("no genre found with name: %s", want)
}

func findCompany(in []tmdb.Company, want string) (tmdb.Company, error) {
	for _, c := range in {
		if name, err := c.Name(); err != nil {
			return tmdb.Company{}, fmt.Errorf("failed to get company name: %w", err)
		} else if name == want {
			return c, nil
		}
	}
	return tmdb.Company{}, fmt.Errorf("no company found with name: %s", want)
}

func findCountry(in []tmdb.Country, want string) (tmdb.Country, error) {
	for _, c := range in {
		if iso, err := c.ISO3166_1(); err != nil {
			return tmdb.Country{}, fmt.Errorf("failed to get country iso_3166_1: %w", err)
		} else if iso == want {
			return c, nil
		}
	}
	return tmdb.Country{}, fmt.Errorf("no country found with iso_3166_1: %s", want)
}

func findSeason(in []tmdb.Season, want int32) (tmdb.Season, error) {
	for _, s := range in {
		if sn, err := s.SeasonNumber(); err != nil {
			return tmdb.Season{}, fmt.Errorf("failed to get season number: %w", err)
		} else if sn == want {
			return s, nil
		}
	}
	return tmdb.Season{}, fmt.Errorf("no season found with season_number: %d", want)
}

func findLanguage(in []tmdb.Language, want string) (tmdb.Language, error) {
	for _, l := range in {
		if iso, err := l.ISO639_1(); err != nil {
			return tmdb.Language{}, fmt.Errorf("failed to get language iso_639_1: %w", err)
		} else if iso == want {
			return l, nil
		}
	}
	return tmdb.Language{}, fmt.Errorf("no language found with iso_639_1: %s", want)
}