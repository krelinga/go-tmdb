package tmdb

import "github.com/krelinga/go-jsonflex"

type ReleaseDates Object

func (r ReleaseDates) ID() (int32, error) {
	return jsonflex.GetField(r, "id", jsonflex.AsInt32())
}

func (r ReleaseDates) Results() ([]CountryReleaseDates, error) {
	return jsonflex.GetField(r, "results", jsonflex.AsArray(jsonflex.AsObject[CountryReleaseDates]()))
}

type CountryReleaseDates Object

func (c CountryReleaseDates) ISO31661() (string, error) {
	return jsonflex.GetField(c, "iso_3166_1", jsonflex.AsString())
}

func (c CountryReleaseDates) ReleaseDates() ([]ReleaseDate, error) {
	return jsonflex.GetField(c, "release_dates", jsonflex.AsArray(jsonflex.AsObject[ReleaseDate]()))
}

type ReleaseDate Object

func (c ReleaseDate) Certification() (string, error) {
	return jsonflex.GetField(c, "certification", jsonflex.AsString())
}

// Descriptors?

func (c ReleaseDate) ISO639_1() (string, error) {
	return jsonflex.GetField(c, "iso_639_1", jsonflex.AsString())
}

func (c ReleaseDate) Note() (string, error) {
	return jsonflex.GetField(c, "note", jsonflex.AsString())
}

func (c ReleaseDate) ReleaseDate() (string, error) {
	return jsonflex.GetField(c, "release_date", jsonflex.AsString())
}

const (
	ReleaseTypePremiere          int32 = 1
	ReleaseTypeTheatricalLimited int32 = 2
	ReleaseTypeTheatrical        int32 = 3
	ReleaseTypeDigital           int32 = 4
	ReleaseTypePhysical          int32 = 5
	ReleaseTypeTV                int32 = 6
)

func (c ReleaseDate) Type() (int32, error) {
	return jsonflex.GetField(c, "type", jsonflex.AsInt32())
}