package tmdb

import (
	"context"
	"fmt"
	"iter"
)

type MovieId int

type WikidataMovieId string

type MoviePart int

const (
	moviePartMin MoviePart = iota
	MoviePartCredits
	MoviePartExternalIds
	MoviePartKeywords
	moviePartMax
)

func (d MoviePart) Endpoint() string {
	switch d {
	case MoviePartCredits:
		return "credits"
	case MoviePartExternalIds:
		return "external_ids"
	case MoviePartKeywords:
		return "keywords"
	default:
		panic(fmt.Sprintf("invalid MovieData value %d; must be between %d and %d", d, moviePartMin, moviePartMax))
	}
}

type Movie interface {
	// This method will never panic.
	// It is safe to call this method concurrently with any other methods.
	Id() MovieId

	// Fetch more data for this movie.
	// The movie will be unchanged if any error occurs (including context cancellation).
	// It is unsafe to call Upgrade() concurrently with calls to the methods contained in MovieParts.
	Upgrade(context.Context, ...MoviePart) error

	// Calls to the methods contained in MovieData may panic if the data is not available.
	// Call Upgrade() with the appropriate MoviePart to ensure these methods will not panic.
	// It is safe to call any methods on MovieParts concurrently with each other, but not with Upgrade().
	MovieParts
}

type MovieParts interface {
	// Call Upgrade() (with any or no arguments) to ensure these methods will not panic.
	Adult() bool
	Budget() int

	// Call Upgrade() with MoviePartCredits to ensure these methods will not panic.
	Cast() iter.Seq[Cast]
	Crew() iter.Seq[Crew]

	// Call Upgrade() with MoviePartExternalIds to ensure these methods will not panic.
	WikidataId() WikidataMovieId

	// Call Upgrade() with MoviePartKeywords to ensure this method will not panic.
	Keywords() iter.Seq[Keyword]
}

type Cast any
type Crew any

type movie struct {
	client   *Client
	id       MovieId
	language Language

	MovieParts
}

func (m *movie) Id() MovieId {
	return m.id
}

func (m *movie) Upgrade(ctx context.Context, data ...MoviePart) error {
	newParts, err := getMovie(ctx, m.client, m.id, m.language, data...)
	if err != nil {
		return fmt.Errorf("upgrading movie %d: %w", m.id, err)
	}
	if oldParts, ok := m.MovieParts.(*getMovieParts); ok {
		oldParts.upgradeFrom(newParts)
	} else {
		m.MovieParts = newParts
	}
	return nil
}

func movieUnsupportedPanic(method string) {
	panic(fmt.Sprintf("method %s() is not supported on this Movie; call Upgrade() first", method))
}

type movieNoParts struct{}

func (movieNoParts) Adult() bool {
	movieUnsupportedPanic("Adult")
	return false // unreachable, but required by the interface
}

func (movieNoParts) Budget() int {
	movieUnsupportedPanic("Budget")
	return 0 // unreachable, but required by the interface
}

func (movieNoParts) Cast() iter.Seq[Cast] {
	movieUnsupportedPanic("Cast")
	return nil // unreachable, but required by the interface
}

func (movieNoParts) Crew() iter.Seq[Crew] {
	movieUnsupportedPanic("Crew")
	return nil // unreachable, but required by the interface
}

func (movieNoParts) WikidataId() WikidataMovieId {
	movieUnsupportedPanic("WikidataId")
	return "" // unreachable, but required by the interface
}

func (movieNoParts) Keywords() iter.Seq[Keyword] {
	movieUnsupportedPanic("Keywords")
	return nil // unreachable, but required by the interface
}
