package tmdb

import (
	"context"
	"fmt"
	"slices"
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

	// Call Upgrade() with MoviePartCredits to ensure these methods will not panic.
	Cast() []Cast
	Crew() []Crew

	// Call Upgrade() with MoviePartExternalIds to ensure these methods will not panic.
	WikidataId() WikidataMovieId

	// Call Upgrade() with MoviePartKeywords to ensure this method will not panic.
	Keywords() []Keyword
}

type Cast any
type Crew any

type movie struct {
	id MovieId
	MovieParts
}

func (m movie) Id() MovieId {
	return m.id
}

func (m movie) Upgrade(ctx context.Context, data ...MoviePart) error {
	appendToResponse := make([]string, 0, len(data))
	for _, d := range data {
		appendToResponse = append(appendToResponse, d.Endpoint())
	}
	slices.Sort(appendToResponse)

	return nil // TODO: implement the upgrade logic
}

type movieNoParts struct{}

func (movieNoParts) panic(method string) {
	panic(fmt.Sprintf("method %s() is not supported on this Movie; call Upgrade() first", method))
}

func (m movieNoParts) Adult() bool {
	m.panic("Adult")
	return false // unreachable, but required by the interface
}

func (m movieNoParts) Cast() []Cast {
	m.panic("Cast")
	return nil // unreachable, but required by the interface
}

func (m movieNoParts) Crew() []Crew {
	m.panic("Crew")
	return nil // unreachable, but required by the interface
}

func (m movieNoParts) WikidataId() WikidataMovieId {
	m.panic("WikidataId")
	return "" // unreachable, but required by the interface
}

func (m movieNoParts) Keywords() []Keyword {
	m.panic("Keywords")
	return nil // unreachable, but required by the interface
}
