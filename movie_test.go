package tmdb_test

import (
	"context"
	"os"
	"testing"

	"github.com/krelinga/go-tmdb"
)

type expected[T comparable] struct {
	Want      T
	wantError bool
}

func (e expected[T]) compare(t *testing.T, p tmdb.Data[T]) {
	t.Helper()
	got, err := tmdb.Get(p)
	if e.wantError {
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	} else {
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		} else if got != e.Want {
			t.Errorf("expected %v, got %v", e.Want, got)
		}
	}
}

func TestMovie(t *testing.T) {
	cases := []struct {
		name       string
		movie      tmdb.Movie
		wantTitle  expected[string]
		wantID     expected[int32]
		wantIMDBID expected[string]
	}{
		{
			name:       "Empty",
			movie:      tmdb.NewMovie(tmdb.Object{}),
			wantTitle:  expected[string]{wantError: true},
			wantID:     expected[int32]{wantError: true},
			wantIMDBID: expected[string]{wantError: true},
		},
		{
			name: "With Title and ID",
			movie: tmdb.NewMovie(tmdb.Object{
				"title": "Inception",
				"id":    tmdb.Number(12345),
				"external_ids": tmdb.Object{
					"imdb_id": "tt1375666",
				},
			}),
			wantTitle:  expected[string]{Want: "Inception"},
			wantID:     expected[int32]{Want: 12345},
			wantIMDBID: expected[string]{Want: "tt1375666"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.wantTitle.compare(t, tc.movie.Title())
			tc.wantID.compare(t, tc.movie.ID())
			tc.wantIMDBID.compare(t, tc.movie.ExternalIDs().IMDBID())
		})
	}
}

func TestGetMovie(t *testing.T) {
	client := tmdb.NewClient(tmdb.ClientOptions{
		APIReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
	})
	movieID := int32(550) // Example movie ID for "Fight Club"
	movie, err := tmdb.GetMovie(context.Background(), client, movieID)
	if err != nil {
		t.Fatalf("failed to get movie: %v", err)
	}
	id, err := tmdb.Get(movie.ID())
	if err != nil {
		t.Fatalf("failed to get movie ID: %v", err)
	}
	if id != movieID {
		t.Errorf("expected movie ID %d, got %d", movieID, id)
	}
	title, err := tmdb.Get(movie.Title())
	if err != nil {
		t.Fatalf("failed to get movie title: %v", err)
	}
	if title != "Fight Club" {
		t.Errorf("expected movie title 'Fight Club', got '%s'", title)
	}
	t.Logf("Movie ID: %d, Title: %s", id, title)
}
