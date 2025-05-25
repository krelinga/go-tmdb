package tmdb_test

import (
	"os"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func getClient(t *testing.T) tmdb.Client {
	t.Helper()
	key, ok := os.LookupEnv("TMDB_API_KEY")
	if !ok {
		t.Fatal("environment variable TMDB_API_KEY not set")
	}
	return tmdb.NewClient(key)
}

func TestGetMovie(t *testing.T) {
	client := getClient(t)
	var raw []byte
	movie, err := tmdb.GetMovie(client, 550,
		tmdb.WithKeywords(),
		tmdb.WithCredits(),
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
	t.Log(string(raw))
}
