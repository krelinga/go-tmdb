package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestSearchMoviess(t *testing.T) {
	total := 0
	for m, err := range tmdb.SearchMovies(getClient(t), "Inception") {
		if err != nil {
			t.Fatalf("SearchMovies failed: %v", err)
		}
		if m == nil {
			t.Fatal("SearchMovies returned nil")
			continue
		}
		t.Logf("Found movie: %s (%d)", m.Title, m.MovieId)
		total++
	}
	t.Logf("Found %d movies", total)
	if total == 0 {
		t.Fatal("No movies found")
	}
}
