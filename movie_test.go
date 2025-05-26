package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetMovie(t *testing.T) {
	var raw []byte
	movie, err := tmdb.GetMovie(getClient(t), 550,
		tmdb.WithKeywords(),
		tmdb.WithCredits(),
		tmdb.WithReleaseDates(),
		tmdb.WithRawReply(&raw))
	if err != nil {
		t.Fatalf("GetMovie failed: %v", err)
	}
	if movie == nil {
		t.Fatal("GetMovie returned nil")
	}
	if movie.Keywords == nil {
		t.Fatal("MovieKeywords is nil")
	}
	if len(movie.Keywords.Keywords) == 0 {
		t.Fatal("No keywords found")
	}
	if movie.Credits == nil {
		t.Fatal("MovieCredits is nil")
	}
	if len(movie.Credits.Cast) == 0 {
		t.Fatal("No cast found")
	}
	if len(movie.Credits.Crew) == 0 {
		t.Fatal("No crew found")
	}
	if movie.Releases == nil {
		t.Fatal("MovieReleaseDates is nil")
	}
	if len(movie.Releases.MovieReleaseCountries) == 0 {
		t.Fatal("No release dates found")
	}
	t.Log(string(raw))
}

func TestHttpError(t *testing.T) {
	_, err := tmdb.GetMovie(getClient(t), 0) // Invalid movie ID
	if err == nil {
		t.Fatal("Expected error for invalid movie ID, got nil")
	}
	t.Logf("Expected error: %v", err)
}
