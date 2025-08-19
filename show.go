package tmdb

import (
	"context"
	"fmt"

	"github.com/krelinga/go-jsonflex"
)

type Show Object

func (s Show) Adult() (bool, error) {
	return jsonflex.GetField(s, "adult", jsonflex.AsBool())
}

func (s Show) BackdropPath() (string, error) {
	return jsonflex.GetField(s, "backdrop_path", jsonflex.AsString())
}

func (s Show) CreatedBy() ([]Credit, error) {
	return jsonflex.GetField(s, "created_by", jsonflex.AsArray(jsonflex.AsObject[Credit]()))
}

func (s Show) FirstAirDate() (string, error) {
	return jsonflex.GetField(s, "first_air_date", jsonflex.AsString())
}

func (s Show) Genres() ([]Genre, error) {
	return jsonflex.GetField(s, "genres", jsonflex.AsArray(jsonflex.AsObject[Genre]()))
}

func (s Show) Homepage() (string, error) {
	return jsonflex.GetField(s, "homepage", jsonflex.AsString())
}

func (s Show) ID() (int32, error) {
	return jsonflex.GetField(s, "id", jsonflex.AsInt32())
}

func (s Show) InProduction() (bool, error) {
	return jsonflex.GetField(s, "in_production", jsonflex.AsBool())
}

func (s Show) Languages() ([]string, error) {
	return jsonflex.GetField(s, "languages", jsonflex.AsArray(jsonflex.AsString()))
}

func (s Show) LastAirDate() (string, error) {
	return jsonflex.GetField(s, "last_air_date", jsonflex.AsString())
}

func (s Show) LastEpisodeToAir() (Episode, error) {
	return jsonflex.GetField(s, "last_episode_to_air", jsonflex.AsObject[Episode]())
}

func (s Show) Name() (string, error) {
	return jsonflex.GetField(s, "name", jsonflex.AsString())
}

func (s Show) NextEpisodeToAir() (string, error) {
	return jsonflex.GetField(s, "next_episode_to_air", jsonflex.AsString())
}

func (s Show) Networks() ([]Company, error) {
	return jsonflex.GetField(s, "networks", jsonflex.AsArray(jsonflex.AsObject[Company]()))
}

func (s Show) NumberOfEpisodes() (int32, error) {
	return jsonflex.GetField(s, "number_of_episodes", jsonflex.AsInt32())
}

func (s Show) NumberOfSeasons() (int32, error) {
	return jsonflex.GetField(s, "number_of_seasons", jsonflex.AsInt32())
}

func (s Show) OriginCountry() ([]string, error) {
	return jsonflex.GetField(s, "origin_country", jsonflex.AsArray(jsonflex.AsString()))
}

func (s Show) OriginalLanguage() (string, error) {
	return jsonflex.GetField(s, "original_language", jsonflex.AsString())
}

func (s Show) OriginalName() (string, error) {
	return jsonflex.GetField(s, "original_name", jsonflex.AsString())
}

func (s Show) Overview() (string, error) {
	return jsonflex.GetField(s, "overview", jsonflex.AsString())
}

func (s Show) Popularity() (float64, error) {
	return jsonflex.GetField(s, "popularity", jsonflex.AsFloat64())
}

func (s Show) PosterPath() (string, error) {
	return jsonflex.GetField(s, "poster_path", jsonflex.AsString())
}

func (s Show) ProductionCompanies() ([]Company, error) {
	return jsonflex.GetField(s, "production_companies", jsonflex.AsArray(jsonflex.AsObject[Company]()))
}

func (s Show) ProductionCountries() ([]Country, error) {
	return jsonflex.GetField(s, "production_countries", jsonflex.AsArray(jsonflex.AsObject[Country]()))
}

func (s Show) Seasons() ([]Season, error) {
	return jsonflex.GetField(s, "seasons", jsonflex.AsArray(jsonflex.AsObject[Season]()))
}

func (s Show) SpokenLanguages() ([]Language, error) {
	return jsonflex.GetField(s, "spoken_languages", jsonflex.AsArray(jsonflex.AsObject[Language]()))
}

func (s Show) Status() (string, error) {
	return jsonflex.GetField(s, "status", jsonflex.AsString())
}

func (s Show) Tagline() (string, error) {
	return jsonflex.GetField(s, "tagline", jsonflex.AsString())
}

func (s Show) Type() (string, error) {
	return jsonflex.GetField(s, "type", jsonflex.AsString())
}

func (s Show) VoteAverage() (float64, error) {
	return jsonflex.GetField(s, "vote_average", jsonflex.AsFloat64())
}

func (s Show) VoteCount() (int32, error) {
	return jsonflex.GetField(s, "vote_count", jsonflex.AsInt32())
}

func (s Show) AggregateCredits() (Credits, error) {
	return jsonflex.GetField(s, "aggregate_credits", jsonflex.AsObject[Credits]())
}

func (s Show) ContentRatings() (ContentRatings, error) {
	return jsonflex.GetField(s, "content_ratings", jsonflex.AsObject[ContentRatings]())
}

func (s Show) Credits() (Credits, error) {
	return jsonflex.GetField(s, "credits", jsonflex.AsObject[Credits]())
}

func (s Show) ExternalIDs() (ExternalIDs, error) {
	return jsonflex.GetField(s, "external_ids", jsonflex.AsObject[ExternalIDs]())
}

func GetShow(ctx context.Context, client Client, showId int32, opts ...RequestOption) (Show, error) {
	return client.Get(ctx, fmt.Sprintf("/3/tv/%d", showId), opts...)
}