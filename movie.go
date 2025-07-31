package tmdb

import (
	"context"
	"fmt"

	"github.com/krelinga/go-jsonflex"
)

func GetMovie(ctx context.Context, client Client, movieID int32, opts ...RequestOption) (Movie, error) {
	return client.Get(ctx, fmt.Sprintf("/3/movie/%d", movieID), opts...)
}

type Movie Object

func (m Movie) Adult() (bool, error) {
	return jsonflex.GetField(m, "adult", jsonflex.AsBool())
}

func (m Movie) BackdropPath() (string, error) {
	return jsonflex.GetField(m, "backdrop_path", jsonflex.AsString())
}

// TODO: BelongsToCollection()

func (m Movie) Budget() (int32, error) {
	return jsonflex.GetField(m, "budget", jsonflex.AsInt32())
}

func (m Movie) Genres() ([]Genre, error) {
	return jsonflex.GetField(m, "genres", jsonflex.AsArray(jsonflex.AsObject[Genre]()))
}

func (m Movie) Homepage() (string, error) {
	return jsonflex.GetField(m, "homepage", jsonflex.AsString())
}

func (m Movie) ID() (int32, error) {
	return jsonflex.GetField(m, "id", jsonflex.AsInt32())
}

func (m Movie) IMDBID() (string, error) {
	return jsonflex.GetField(m, "imdb_id", jsonflex.AsString())
}

func (m Movie) OriginalLanguage() (string, error) {
	return jsonflex.GetField(m, "original_language", jsonflex.AsString())
}

func (m Movie) OriginalTitle() (string, error) {
	return jsonflex.GetField(m, "original_title", jsonflex.AsString())
}

func (m Movie) Overview() (string, error) {
	return jsonflex.GetField(m, "overview", jsonflex.AsString())
}

func (m Movie) Popularity() (float64, error) {
	return jsonflex.GetField(m, "popularity", jsonflex.AsFloat64())
}

func (m Movie) PosterPath() (string, error) {
	return jsonflex.GetField(m, "poster_path", jsonflex.AsString())
}

func (m Movie) ProductionCompanies() ([]Company, error) {
	return jsonflex.GetField(m, "production_companies", jsonflex.AsArray(jsonflex.AsObject[Company]()))
}

func (m Movie) ProductionCountries() ([]Country, error) {
	return jsonflex.GetField(m, "production_countries", jsonflex.AsArray(jsonflex.AsObject[Country]()))
}

func (m Movie) ReleaseDate() (string, error) {
	return jsonflex.GetField(m, "release_date", jsonflex.AsString())
}

func (m Movie) Revenue() (int32, error) {
	return jsonflex.GetField(m, "revenue", jsonflex.AsInt32())
}

func (m Movie) Runtime() (int32, error) {
	return jsonflex.GetField(m, "runtime", jsonflex.AsInt32())
}

func (m Movie) SpokenLanguages() ([]Language, error) {
	return jsonflex.GetField(m, "spoken_languages", jsonflex.AsArray(jsonflex.AsObject[Language]()))
}

func (m Movie) Status() (string, error) {
	return jsonflex.GetField(m, "status", jsonflex.AsString())
}

func (m Movie) Tagline() (string, error) {
	return jsonflex.GetField(m, "tagline", jsonflex.AsString())
}

func (m Movie) Title() (string, error) {
	return jsonflex.GetField(m, "title", jsonflex.AsString())
}

func (m Movie) Video() (bool, error) {
	return jsonflex.GetField(m, "video", jsonflex.AsBool())
}

func (m Movie) VoteAverage() (float64, error) {
	return jsonflex.GetField(m, "vote_average", jsonflex.AsFloat64())
}

func (m Movie) VoteCount() (int32, error) {
	return jsonflex.GetField(m, "vote_count", jsonflex.AsInt32())
}

func (m Movie) OriginCountry() ([]string, error) {
	return jsonflex.GetField(m, "origin_country", jsonflex.AsArray(jsonflex.AsString()))
}