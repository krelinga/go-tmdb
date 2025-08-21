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

func TestSeachTv(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	results, err := tmdb.SearchTv(context.Background(), client, "Breaking Bad")
	if err != nil {
		t.Fatalf("Failed to search TV show: %v", err)
	}
	checkField(t, int32(1), results, tmdb.SearchResults[tmdb.Show].Page)
	checkField(t, int32(1), results, tmdb.SearchResults[tmdb.Show].TotalPages)
	checkField(t, int32(3), results, tmdb.SearchResults[tmdb.Show].TotalResults)
	if results, err := results.Results(); err != nil {
		t.Fatalf("Failed to get TV show results: %v", err)
	} else if bb, err := findShow(results, "Breaking Bad"); err != nil {
		t.Fatalf("Failed to find TV show: %v", err)
	} else {
		checkField(t, false, bb, tmdb.Show.Adult)
		checkField(t, "/tsRy63Mu5cu8etL1X7ZLyf7UP1M.jpg", bb, tmdb.Show.BackdropPath)
		if genreIds, err := bb.GenreIDs(); err != nil {
			t.Fatalf("Failed to get genre IDs: %v", err)
		} else {
			if !slices.Contains(genreIds, int32(80)) {
				t.Errorf("Expected genre ID 80 not found")
			}
			if !slices.Contains(genreIds, int32(18)) {
				t.Errorf("Expected genre ID 18 not found")
			}
		}
		checkField(t, int32(1396), bb, tmdb.Show.ID)
		if originCountry, err := bb.OriginCountry(); err != nil {
			t.Fatalf("Failed to get origin country: %v", err)
		} else if !slices.Contains(originCountry, "US") {
			t.Errorf("Expected origin country 'US' not found")
		}
		checkField(t, "en", bb, tmdb.Show.OriginalLanguage)
		checkField(t, "Breaking Bad", bb, tmdb.Show.OriginalName)
		checkField(t, "Walter White, a New Mexico chemistry teacher, is diagnosed with Stage III cancer and given a prognosis of only two years left to live. He becomes filled with a sense of fearlessness and an unrelenting desire to secure his family's financial future at any cost as he enters the dangerous world of drugs and crime.", bb, tmdb.Show.Overview)
		checkField(t, 110.6989, bb, tmdb.Show.Popularity)
		checkField(t, "/ztkUQFLlC19CCMYHW9o1zWhJRNq.jpg", bb, tmdb.Show.PosterPath)
		checkField(t, "2008-01-20", bb, tmdb.Show.FirstAirDate)
		checkField(t, "Breaking Bad", bb, tmdb.Show.Name)
		checkField(t, 8.92, bb, tmdb.Show.VoteAverage)
		checkField(t, int32(15966), bb, tmdb.Show.VoteCount)
	}
}

func findShow(shows []tmdb.Show, name string) (tmdb.Show, error) {
	for _, show := range shows {
		if n, err := show.Name(); err != nil {
			return nil, fmt.Errorf("failed to get show name: %w", err)
		} else if n == name {
			return show, nil
		}
	}
	return nil, fmt.Errorf("show not found: %s", name)
}
