package tmdb

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/krelinga/go-tmdb/internal/raw"
)

// In addition to what is listed in MovieData, the following data is also available:
// - Companies()
//   - Logo()
//   - Name()
//   - OriginCountry()
func GetMovie(ctx context.Context, c *Client, id MovieId, options ...Option) (*Movie, error) {
	callOpts := c.globalOpts
	callOpts.apply(options)

	v := url.Values{}
	callOpts.applyLanguage(v)
	callOpts.applyMovieDataColumns(v)

	rawGetMovie := &raw.GetMovie{}
	if err := get(ctx, c, fmt.Sprintf("movie/%d", id), v, callOpts, rawGetMovie); err != nil {
		return nil, fmt.Errorf("getting movie %d: %w", id, err)
	}

	genres := make([]Genre, len(rawGetMovie.Genres))
	for i, g := range rawGetMovie.Genres {
		genres[i] = Genre{
			Key: GenreId(g.Id),
			Data: GenreData{
				Name: &g.Name,
			},
		}
	}
	companies := make([]*Company, len(rawGetMovie.ProductionCompanies))
	for i, rawCompany := range rawGetMovie.ProductionCompanies {
		companies[i] = &Company{
			Key: CompanyId(rawCompany.Id),
			Data: CompanyData{
				Logo:          NewPtr[Image](Image(rawCompany.LogoPath)),
				Name:          &rawCompany.Name,
				OriginCountry: &rawCompany.OriginCountry,
			},
		}
	}
	countries := make([]*Country, len(rawGetMovie.ProductionCountries))
	for i, rawCountry := range rawGetMovie.ProductionCountries {
		countries[i] = &Country{
			Code: &rawCountry.Iso3166_1,
			Name: &rawCountry.Name,
		}
	}
	spokenLanguages := make([]*Language, len(rawGetMovie.SpokenLanguages))
	for i, rawLang := range rawGetMovie.SpokenLanguages {
		spokenLanguages[i] = &Language{
			Code:        &rawLang.Iso639_1,
			Name:        &rawLang.Name,
			EnglishName: &rawLang.EnglishName,
		}
	}
	var keywords []*Keyword
	if rawGetMovie.Keywords != nil {
		keywords = make([]*Keyword, len(rawGetMovie.Keywords.Keywords))
		for i, rawKeyword := range rawGetMovie.Keywords.Keywords {
			keywords[i] = &Keyword{
				Key: KeywordId(rawKeyword.Id),
				Data: KeywordData{
					Name: &rawKeyword.Name,
				},
			}
		}
	}
	var cast, crew []*Credit
	if rawGetMovie.Credits != nil {
		toPerson := func(rawPerson *raw.GetMovieCreditsPerson) *Person {
			return &Person{
				Key: PersonId(rawPerson.Id),
				Data: PersonData{
					Adult:              &rawPerson.Adult,
					Gender:             NewPtr(Gender(rawPerson.Gender)),
					KnownForDepartment: &rawPerson.KnownForDepartment,
					Name:               &rawPerson.Name,
					Popularity:         &rawPerson.Popularity,
					Profile:            NewPtr(Image(rawPerson.ProfilePath)),
				},
			}
		}
		cast = make([]*Credit, len(rawGetMovie.Credits.Cast))
		for i, rawCast := range rawGetMovie.Credits.Cast {
			cast[i] = &Credit{
				Key: CreditId(rawCast.CreditId),
				Data: CreditData{
					OriginalName: &rawCast.OriginalName,
					CastId:       NewPtr(CastId(rawCast.CastId)),
					Character:    &rawCast.Character,
					Order:        &rawCast.Order,
				},
				Person: toPerson(&rawCast.GetMovieCreditsPerson),
			}
		}
		crew = make([]*Credit, len(rawGetMovie.Credits.Crew))
		for i, rawCrew := range rawGetMovie.Credits.Crew {
			crew[i] = &Credit{
				Key: CreditId(rawCrew.CreditId),
				Data: CreditData{
					OriginalName: &rawCrew.OriginalName,
					Department:   &rawCrew.Department,
					Job:          &rawCrew.Job,
				},
				Person: toPerson(&rawCrew.GetMovieCreditsPerson),
			}
		}
	}
	out := &Movie{
		Key: MovieId(rawGetMovie.Id),
		Data: MovieData{
			Adult:               rawGetMovie.Adult,
			Backdrop:            NewPtr(Image(rawGetMovie.BackdropPath)),
			BelongsToCollection: &rawGetMovie.BelongsToCollection,
			Budget:              &rawGetMovie.Budget,
			Genres:              genres,
			Homepage:            &rawGetMovie.Homepage,
			ImdbId:              &rawGetMovie.ImdbId,
			OriginalLanguage:    &rawGetMovie.OriginalLanguage,
			OriginalTitle:       &rawGetMovie.OriginalTitle,
			Overview:            &rawGetMovie.Overview,
			Popularity:          &rawGetMovie.Popularity,
			Poster:              NewPtr(Image(rawGetMovie.PosterPath)),
			ProductionCompanies: companies,
			ProductionCountries: countries,
			ReleaseDate:         NewPtr(DateYYYYMMDD(rawGetMovie.ReleaseDate)),
			Revenue:             &rawGetMovie.Revenue,
			Runtime:             NewPtr(time.Duration(rawGetMovie.Runtime) * time.Minute),
			SpokenLanguages:     spokenLanguages,
			Status:              &rawGetMovie.Status,
			Tagline:             &rawGetMovie.Tagline,
			Title:               &rawGetMovie.Title,
			Video:               &rawGetMovie.Video,
			VoteAverage:         &rawGetMovie.VoteAverage,
			VoteCount:           &rawGetMovie.VoteCount,
		},

		Cast: cast,
		Crew: crew,

		Keywords: keywords,
	}

	if rawGetMovie.ExternalIds != nil {
		out.Data.WikidataId = &rawGetMovie.ExternalIds.WikidataId
		out.Data.FacebookId = &rawGetMovie.ExternalIds.FacebookId
		out.Data.TwitterId = &rawGetMovie.ExternalIds.TwitterId
		out.Data.InstagramId = &rawGetMovie.ExternalIds.InstagramId
	}
	return out, nil
}
