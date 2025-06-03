package tmdb

import (
	"context"
	"fmt"
	"iter"
)

type MovieId int

type WikidataMovieId string

type MovieDataCol int

const (
	movieDataNone MovieDataCol = iota
	movieDataMin
	MovieDataCredits
	MovieDataExternalIds
	MovieDataKeywords
	movieDataMax
)

func (d MovieDataCol) String() string {
	switch d {
	case movieDataNone:
		return "movieDataNone"
	case movieDataMin:
		return "movieDataMin"
	case MovieDataCredits:
		return "MovieDataCredits"
	case MovieDataExternalIds:
		return "MovieDataExternalIds"
	case MovieDataKeywords:
		return "MovieDataKeywords"
	case movieDataMax:
		return "movieDataMax"
	default:
		return fmt.Sprintf("MovieDataCol(%d)", d)
	}
}

func (d MovieDataCol) Endpoint() string {
	switch d {
	case MovieDataCredits:
		return "credits"
	case MovieDataExternalIds:
		return "external_ids"
	case MovieDataKeywords:
		return "keywords"
	default:
		panic(fmt.Sprintf("invalid MovieData value %d; must be between %d and %d", d, movieDataMin, movieDataMax))
	}
}

func movieNoDataError(field string, col MovieDataCol) error {
	var colPart string
	if col != movieDataNone {
		colPart = fmt.Sprintf(" with %s", col)
	}
	return fmt.Errorf("cannot access %s on MovieData without calling Upgrade()%s first", field, colPart)
}

// MovieData implementations may panic with these errors if the corresponding data is not available.
var (
	ErrMovieNoDataAdult      = movieNoDataError("Adult", movieDataNone)
	ErrMovieNoDataBackdrop   = movieNoDataError("Backdrop", movieDataNone)
	ErrMovieNoDataCollection = movieNoDataError("BelongsToCollection", movieDataNone)
	ErrMovieNoDataBudget     = movieNoDataError("Budget", movieDataNone)
	ErrMovieNoDataGenreIds   = movieNoDataError("GenreIds", movieDataNone)
	ErrMovieNoDataGenres     = movieNoDataError("Genres", movieDataNone)
	ErrMovieNoDataHomepage   = movieNoDataError("Homepage", movieDataNone)
	ErrMovieNoDataImdbId     = movieNoDataError("ImdbId", movieDataNone)

	ErrMovieNoDataCast = movieNoDataError("Cast", MovieDataCredits)
	ErrMovieNoDataCrew = movieNoDataError("Crew", MovieDataCredits)

	ErrMovieNoDataWikidataId = movieNoDataError("WikidataId", MovieDataExternalIds)

	ErrMovieNoDataKeywords = movieNoDataError("Keywords", MovieDataKeywords)
)

type Movie interface {
	// This method will never panic.
	// It is safe to call this method concurrently with any other methods.
	Id() MovieId

	// Fetch more data for this movie.
	// The movie will be unchanged if any error occurs (including context cancellation).
	// It is unsafe to call Upgrade() concurrently with calls to the methods contained in MovieData.
	Upgrade(context.Context, ...MovieDataCol) error

	// Calls to the methods contained in MovieData may panic if the data is not available.
	// Call Upgrade() with the appropriate MovieDataCol to ensure these methods will not panic.
	// It is safe to call any methods on MovieData concurrently with each other, but not with Upgrade().
	MovieData
}

type ImdbMovieId string

type MovieData interface {
	// Call Upgrade() (with any or no arguments) to ensure these methods will not panic.
	Adult() bool
	BelongsToCollection() string
	Budget() int
	Backdrop() Image
	GenreIds() iter.Seq[GenreId]
	Genres() iter.Seq[Genre]
	Homepage() string
	ImdbId() ImdbMovieId

	// Call Upgrade() with MovieDataCredits to ensure these methods will not panic.
	Cast() iter.Seq[Cast]
	Crew() iter.Seq[Crew]

	// Call Upgrade() with MovieDataExternalIds to ensure these methods will not panic.
	WikidataId() WikidataMovieId

	// Call Upgrade() with MovieDataKeywords to ensure this method will not panic.
	Keywords() iter.Seq[Keyword]

	// Internal methods, not safe to call together with any other method on MovieData.
	upgrade(*getMovieData) MovieData
}

type Cast any
type Crew any

type movie struct {
	client   *Client
	id       MovieId
	language Language

	MovieData
}

func (m *movie) Id() MovieId {
	return m.id
}

func (m *movie) Upgrade(ctx context.Context, data ...MovieDataCol) error {
	newParts, err := getMovie(ctx, m.client, m.id, m.language, data...)
	if err != nil {
		return fmt.Errorf("upgrading movie %d: %w", m.id, err)
	}
	m.MovieData = m.MovieData.upgrade(newParts)
	return nil
}

type movieNoData struct{}

func (movieNoData) upgrade(in *getMovieData) MovieData {
	return in
}

func (movieNoData) Adult() bool {
	panic(ErrMovieNoDataAdult)
}

func (movieNoData) Backdrop() Image {
	panic(ErrMovieNoDataBackdrop)
}

func (movieNoData) BelongsToCollection() string {
	panic(ErrMovieNoDataCollection)
}

func (movieNoData) Budget() int {
	panic(ErrMovieNoDataBudget)
}

func (movieNoData) GenreIds() iter.Seq[GenreId] {
	panic(ErrMovieNoDataGenreIds)
}

func (movieNoData) Genres() iter.Seq[Genre] {
	panic(ErrMovieNoDataGenres)
}

func (movieNoData) Homepage() string {
	panic(ErrMovieNoDataHomepage)
}

func (movieNoData) ImdbId() ImdbMovieId {
	panic(ErrMovieNoDataImdbId)
}

func (movieNoData) Cast() iter.Seq[Cast] {
	panic(ErrMovieNoDataCast)
}

func (movieNoData) Crew() iter.Seq[Crew] {
	panic(ErrMovieNoDataCrew)
}

func (movieNoData) WikidataId() WikidataMovieId {
	panic(ErrMovieNoDataWikidataId)
}

func (movieNoData) Keywords() iter.Seq[Keyword] {
	panic(ErrMovieNoDataKeywords)
}
