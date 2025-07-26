package testutil

import (
	"testing"

	"github.com/krelinga/go-tmdb"
)

// Test our function with actual TMDB structures from the codebase
func TestAssertAllFieldsSetWithTMDBStructures(t *testing.T) {
	// Test with a Movie struct that has all slices initialized (even if empty)
	movie := &tmdb.Movie{
		MovieSum: tmdb.MovieSum{
			Adult:            false,
			BackdropImage:    tmdb.BackdropImage("/test.jpg"),
			MovieId:          550,
			OriginalLanguage: "en",
			OriginalTitle:    "Test Movie",
			Overview:         "A test movie",
			Popularity:       10.0,
			PosterImage:      tmdb.PosterImage("/poster.jpg"),
			ReleaseDate:      "2023-01-01",
			Title:            "Test Movie",
			Video:            false,
			VoteAverage:      8.0,
			VoteCount:        1000,
		},
		BelongsToCollection: "test collection",
		Budget:              1000000,
		Genres: []*tmdb.Genre{
			{GenreId: 18, Name: "Drama"},
		},
		Homepage:                   "http://example.com",
		ImdbId:                     "tt1234567",
		ProductionCompanyShorts:    []*tmdb.ProductionCompanySum{}, // Empty but not nil
		ProductionCountrySummaries: []*tmdb.CountrySum{},           // Empty but not nil
		Revenue:                    2000000,
		Runtime:                    120,
		SpokenLanguages:            []*tmdb.Language{}, // Empty but not nil
		Status:                     "Released",
		Tagline:                    "A great movie",
	}
	
	// This should pass since all slices are initialized (even if empty)
	result := AssertAllFieldsSet(t, movie)
	if !result {
		t.Error("Expected movie struct with initialized slices to pass")
	}
}

func TestAssertAllFieldsSetWithTMDBStructures_NilSlices(t *testing.T) {
	// Test with a struct that has nil slices - should fail
	movieWithNilSlices := &tmdb.Movie{
		MovieSum: tmdb.MovieSum{
			Adult:            false,
			BackdropImage:    tmdb.BackdropImage("/test.jpg"),
			MovieId:          550,
			OriginalLanguage: "en",
			OriginalTitle:    "Test Movie",
			Overview:         "A test movie",
			Popularity:       10.0,
			PosterImage:      tmdb.PosterImage("/poster.jpg"),
			ReleaseDate:      "2023-01-01",
			Title:            "Test Movie",
			Video:            false,
			VoteAverage:      8.0,
			VoteCount:        1000,
		},
		BelongsToCollection:        "test collection",
		Budget:                     1000000,
		Genres:                     nil, // This should cause failure
		Homepage:                   "http://example.com",
		ImdbId:                     "tt1234567",
		ProductionCompanyShorts:    nil, // This should cause failure
		ProductionCountrySummaries: nil, // This should cause failure
		Revenue:                    2000000,
		Runtime:                    120,
		SpokenLanguages:            nil, // This should cause failure
		Status:                     "Released",
		Tagline:                    "A great movie",
	}
	
	// Create a separate test to capture errors without affecting this test
	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, movieWithNilSlices)
	if result {
		t.Error("Expected AssertAllFieldsSet to return false for movie with nil slices")
	}
}