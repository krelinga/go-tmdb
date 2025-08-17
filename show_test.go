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
