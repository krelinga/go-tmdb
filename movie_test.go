package tmdb_test

import (
	"os"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func getClient(t *testing.T) tmdb.Client {
	key, ok := os.LookupEnv("TMDB_API_KEY")
	if !ok {
		t.Fatal("environment variable TMDB_API_KEY not set")
	}
	return tmdb.NewClient(key)
}

func TestGetMovie(t *testing.T) {
	client := getClient(t)
	movie, err := tmdb.GetMovie(client, 550, tmdb.WithDetails(), tmdb.WithKeywords())
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
}