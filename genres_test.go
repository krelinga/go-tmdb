package tmdb_test

import (
	"context"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetMovieGenres(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	genres, err := tmdb.GetMovieGenres(context.Background(), client)
	if err != nil {
		t.Fatalf("failed to get movie genres: %v", err)
	} else if genreList, err := genres.Genres(); err != nil {
		t.Fatalf("failed to get genres: %v", err)
	} else {
		if action, err := findGenre(genreList, "Action"); err != nil {
			t.Fatalf("failed to find Action genre: %v", err)
		} else {
			checkField(t, int32(28), action, tmdb.Genre.ID)
			checkField(t, "Action", action, tmdb.Genre.Name)
		}
		if adventure, err := findGenre(genreList, "Adventure"); err != nil {
			t.Fatalf("failed to find Adventure genre: %v", err)
		} else {
			checkField(t, int32(12), adventure, tmdb.Genre.ID)
			checkField(t, "Adventure", adventure, tmdb.Genre.Name)
		}
	}
}

func TestGetTvGenres(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	genres, err := tmdb.GetTvGenres(context.Background(), client)
	if err != nil {
		t.Fatalf("failed to get tv genres: %v", err)
	} else if genreList, err := genres.Genres(); err != nil {
		t.Fatalf("failed to get genres: %v", err)
	} else {
		if action, err := findGenre(genreList, "Action & Adventure"); err != nil {
			t.Fatalf("failed to find Action & Adventure genre: %v", err)
		} else {
			checkField(t, int32(10759), action, tmdb.Genre.ID)
			checkField(t, "Action & Adventure", action, tmdb.Genre.Name)
		}
		if animation, err := findGenre(genreList, "Animation"); err != nil {
			t.Fatalf("failed to find Animation genre: %v", err)
		} else {
			checkField(t, int32(16), animation, tmdb.Genre.ID)
			checkField(t, "Animation", animation, tmdb.Genre.Name)
		}
	}
}
