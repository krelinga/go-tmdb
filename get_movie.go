package tmdb

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type GetMovieOptions struct {
	Language LanguageId
	Columns  []MovieDataColumn
}

// In addition to what is listed in MovieData, the following data is also available:
// - Companies()
//   - Logo()
//   - Name()
//   - OriginCountry()
func GetMovie(ctx context.Context, c *Client, id MovieId, options ...Option) (*Movie, error) {
	var callOpts allOptions
	callOpts.apply(options)
	v := url.Values{}
	if callOpts.Language != nil {
		v.Set("language", *callOpts.Language)
	}
	if len(callOpts.MovieDataColumns) > 0 {
		v.Set("append_to_response", appendToResponse(callOpts.MovieDataColumns))
	}

	raw := &raw.GetMovie{}
	if err := get(ctx, c, fmt.Sprintf("movie/%d", id), v, raw); err != nil {
		return nil, fmt.Errorf("getting movie %d: %w", id, err)
	}

	genres := make([]Genre, len(raw.Genres))
	for i, g := range raw.Genres {
		genres[i] = Genre{
			Id:   GenreId(g.Id),
			Name: &g.Name,
		}
	}
	companies := make([]*Company, len(raw.ProductionCompanies))
	for i, rawCompany := range raw.ProductionCompanies {
		companies[i] = &Company{
			Id: CompanyId(rawCompany.Id),
			Logo: NewPtr[Image](Image(rawCompany.LogoPath)),
			Name: &rawCompany.Name,
			OriginCountry: &rawCompany.OriginCountry,
		}
	}
	countries := make([]*Country, len(raw.ProductionCountries))
	for i, rawCountry := range raw.ProductionCountries {
		countries[i] = &Country{
			Code: &rawCountry.Iso3166_1,
			Name: &rawCountry.Name,
		}
	}
	spokenLanguages := make([]*Language, len(raw.SpokenLanguages))
	for i, rawLang := range raw.SpokenLanguages {
		spokenLanguages[i] = &Language{
			Code: &rawLang.Iso639_1,
			Name: &rawLang.Name,
			EnglishName: &rawLang.EnglishName,
		}
	}
	out := &Movie{
		Id: MovieId(raw.Id),
		Adult: raw.Adult,
		Backdrop: NewPtr[Image](Image(raw.BackdropPath)),
		BelongsToCollection: &raw.BelongsToCollection,
		Budget: &raw.Budget,
		Genres: genres,
		Homepage: &raw.Homepage,
		ImdbId: &raw.ImdbId,
		OriginalLanguage: &raw.OriginalLanguage,
		OriginalTitle: &raw.OriginalTitle,
		Overview: &raw.Overview,
		Popularity: &raw.Popularity,
		Poster: NewPtr[Image](Image(raw.PosterPath)),
		ProductionCompanies: companies,
		ProductionCountries: countries,
		ReleaseDate: NewPtr(DateYYYYMMDD(raw.ReleaseDate)),
		Revenue: &raw.Revenue,
		Runtime: NewPtr(time.Duration(raw.Runtime) * time.Minute),
		SpokenLanguages: spokenLanguages,
		Status: &raw.Status,
		Tagline: &raw.Tagline,
		Title: &raw.Title,
		Video: &raw.Video,
		VoteAverage: &raw.VoteAverage,
		VoteCount: &raw.VoteCount,
	}
	return out, nil
}

