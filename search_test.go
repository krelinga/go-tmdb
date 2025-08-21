package tmdb_test

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestSearchMovie(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	results, err := tmdb.SearchMovie(context.Background(), client, "Inception")
	if err != nil {
		t.Fatalf("Failed to search movie: %v", err)
	}
	checkField(t, int32(1), results, tmdb.SearchResults[tmdb.Movie].Page)
	checkField(t, int32(1), results, tmdb.SearchResults[tmdb.Movie].TotalPages)
	checkField(t, int32(10), results, tmdb.SearchResults[tmdb.Movie].TotalResults)
	if results, err := results.Results(); err != nil {
		t.Fatalf("Failed to get movie results: %v", err)
	} else if inception, err := findMovie(results, "Inception"); err != nil {
		t.Fatalf("Failed to find movie: %v", err)
	} else {
		checkField(t, false, inception, tmdb.Movie.Adult)
		checkField(t, "/gqby0RhyehP3uRrzmdyUZ0CgPPe.jpg", inception, tmdb.Movie.BackdropPath)
		if genreIds, err := inception.GenreIDs(); err != nil {
			t.Fatalf("Failed to get genre IDs: %v", err)
		} else {
			if !slices.Contains(genreIds, int32(28)) {
				t.Errorf("Expected genre ID 28 not found")
			}
			if !slices.Contains(genreIds, int32(878)) {
				t.Errorf("Expected genre ID 878 not found")
			}
			if !slices.Contains(genreIds, int32(12)) {
				t.Errorf("Expected genre ID 12 not found")
			}
		}
		checkField(t, int32(27205), inception, tmdb.Movie.ID)
		checkField(t, "en", inception, tmdb.Movie.OriginalLanguage)
		checkField(t, "Inception", inception, tmdb.Movie.OriginalTitle)
		checkField(t, "Cobb, a skilled thief who commits corporate espionage by infiltrating the subconscious of his targets is offered a chance to regain his old life as payment for a task considered to be impossible: \"inception\", the implantation of another person's idea into a target's subconscious.", inception, tmdb.Movie.Overview)
		checkField(t, 25.0601, inception, tmdb.Movie.Popularity)
		checkField(t, "/ljsZTbVsrQSqZgWeep2B1QiDKuh.jpg", inception, tmdb.Movie.PosterPath)
		checkField(t, "2010-07-15", inception, tmdb.Movie.ReleaseDate)
		checkField(t, "Inception", inception, tmdb.Movie.Title)
		checkField(t, false, inception, tmdb.Movie.Video)
		checkField(t, 8.369, inception, tmdb.Movie.VoteAverage)
		checkField(t, int32(37812), inception, tmdb.Movie.VoteCount)
	}
}

func findMovie(movies []tmdb.Movie, title string) (tmdb.Movie, error) {
	for _, movie := range movies {
		if name, err := movie.Title(); err != nil {
			return nil, fmt.Errorf("failed to get movie title: %w", err)
		} else if name == title {
			return movie, nil
		}
	}
	return nil, fmt.Errorf("movie not found: %s", title)
}
