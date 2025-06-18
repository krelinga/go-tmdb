package tmdb

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/krelinga/go-tmdb/internal/raw"
)

func GetTv(ctx context.Context, client *Client, id TvId, options ...Option) (*Tv, error) {
	callOpts := client.globalOpts
	callOpts.apply(options)

	// TODO: support AppendToResponse here.
	params := url.Values{}
	callOpts.applyLanguage(params)

	rawTv := &raw.Tv{}
	if err := get(ctx, client, fmt.Sprintf("tv/%d", id), params, callOpts, rawTv); err != nil {
		return nil, err
	}

	createdBy := make([]*Credit, len(rawTv.CreatedBy))
	for i, rawCredit := range rawTv.CreatedBy {
		createdBy[i] = &Credit{
			Key: CreditId(rawCredit.CreditId),
			Person: &Person{
				Key: PersonId(rawCredit.Id),
				Name:    &rawCredit.Name,
				Profile: NewPtr(Image(rawCredit.ProfilePath)),
			},
		}
	}
	episodeRunTime := make([]time.Duration, len(rawTv.EpisodeRunTime))
	for i, dur := range rawTv.EpisodeRunTime {
		episodeRunTime[i] = time.Duration(dur) * time.Minute
	}
	genres := make([]Genre, len(rawTv.Genres))
	for i, g := range rawTv.Genres {
		genres[i] = Genre{
			Key: GenreId(g.Id),
			Name: &g.Name,
		}
	}
	lastEpisodeToAir := &Episode{
		Key: EpisodeKey{
			SeasonNumber:  rawTv.LastEpisodeToAir.SeasonNumber,
			EpisodeNumber: rawTv.LastEpisodeToAir.EpisodeNumber,
			TvId:          TvId(rawTv.Id),
		},
		Id:          NewPtr(EpisodeId(rawTv.LastEpisodeToAir.Id)),
		Name:        &rawTv.LastEpisodeToAir.Name,
		Overview:    &rawTv.LastEpisodeToAir.Overview,
		VoteAverage: &rawTv.LastEpisodeToAir.VoteAverage,
		VoteCount:   &rawTv.LastEpisodeToAir.VoteCount,
		AirDate:     NewPtr(DateYYYYMMDD(rawTv.LastEpisodeToAir.AirDate)),

		ProductionCode: &rawTv.LastEpisodeToAir.ProductionCode,
		Runtime:        NewPtr(time.Duration(rawTv.LastEpisodeToAir.Runtime) * time.Minute),
		Still:          NewPtr(Image(rawTv.LastEpisodeToAir.StillPath)),
	}
	networks := make([]*Network, len(rawTv.Networks))
	for i, rawNetwork := range rawTv.Networks {
		networks[i] = &Network{
			Key: NetworkId(rawNetwork.Id),
			Name:          &rawNetwork.Name,
			Logo:          NewPtr(Image(rawNetwork.LogoPath)),
			OriginCountry: &rawNetwork.OriginCountry,
		}
	}
	companies := make([]*Company, len(rawTv.ProductionCompanies))
	for i, rawCompany := range rawTv.ProductionCompanies {
		companies[i] = &Company{
			Key: CompanyId(rawCompany.Id),
			Logo:          NewPtr[Image](Image(rawCompany.LogoPath)),
			Name:          &rawCompany.Name,
			OriginCountry: &rawCompany.OriginCountry,
		}
	}
	countries := make([]*Country, len(rawTv.ProductionCountries))
	for i, rawCountry := range rawTv.ProductionCountries {
		countries[i] = &Country{
			Code: &rawCountry.Iso3166_1,
			Name: &rawCountry.Name,
		}
	}
	seasons := make([]*Season, len(rawTv.Seasons))
	for i, rawSeason := range rawTv.Seasons {
		seasons[i] = &Season{
			Key: SeasonKey{
				TvId:         TvId(rawTv.Id),
				SeasonNumber: rawSeason.SeasonNumber,
			},
			AirDate:      NewPtr(DateYYYYMMDD(rawSeason.AirDate)),
			EpisodeCount: &rawSeason.EpisodeCount,
			Id:           NewPtr(SeasonId(rawSeason.Id)),
			Name:         &rawSeason.Name,
			Overview:     &rawSeason.Overview,
			Poster:       NewPtr(Image(rawSeason.PosterPath)),
		}
	}
	spokenLanguages := make([]*Language, len(rawTv.SpokenLanguages))
	for i, rawLang := range rawTv.SpokenLanguages {
		spokenLanguages[i] = &Language{
			Code:        &rawLang.Iso639_1,
			Name:        &rawLang.Name,
			EnglishName: &rawLang.EnglishName,
		}
	}
	out := &Tv{
		Key: id,

		Adult:               rawTv.Adult,
		Backdrop:            NewPtr(Image(rawTv.BackdropPath)),
		CreatedBy:           createdBy,
		EpisodeRunTime:      episodeRunTime,
		FirstAirDate:        NewPtr(DateYYYYMMDD(rawTv.FirstAirDate)),
		Genres:              genres,
		Homepage:            &rawTv.Homepage,
		InProduction:        rawTv.InProduction,
		Languages:           rawTv.Languages,
		LastAirDate:         NewPtr(DateYYYYMMDD(rawTv.LastAirDate)),
		LastEpisodeToAir:    lastEpisodeToAir,
		Name:                &rawTv.Name,
		NextEpisodeToAir:    &rawTv.NextEpisodeToAir,
		Networks:            networks,
		NumberOfEpisodes:    &rawTv.NumberOfEpisodes,
		NumberOfSeasons:     &rawTv.NumberOfSeasons,
		OriginCountry:       rawTv.OriginCountry,
		OriginalLanguage:    &rawTv.OriginalLanguage,
		OriginalName:        &rawTv.OriginalName,
		Overview:            &rawTv.Overview,
		Popularity:          &rawTv.Popularity,
		Poster:              NewPtr(Image(rawTv.PosterPath)),
		ProductionCompanies: companies,
		ProductionCountries: countries,
		Seasons:             seasons,
		SpokenLanguages:     spokenLanguages,
		Status:              &rawTv.Status,
		Tagline:             &rawTv.Tagline,
		Type:                &rawTv.Type,
		VoteAverage:         &rawTv.VoteAverage,
		VoteCount:           &rawTv.VoteCount,
	}
	return out, nil
}
