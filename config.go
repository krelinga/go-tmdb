package tmdb

import (
	"context"

	"github.com/krelinga/go-jsonflex"
)

type ConfigDetails Object

func (c ConfigDetails) Images() (ConfigImages, error) {
	return jsonflex.GetField(c, "images", jsonflex.AsObject[ConfigImages]())
}

func (c ConfigDetails) ChangeKeys() ([]string, error) {
	return jsonflex.GetField(c, "change_keys", jsonflex.AsArray(jsonflex.AsString()))
}

type ConfigImages Object

func (c ConfigImages) BaseURL() (string, error) {
	return jsonflex.GetField(c, "base_url", jsonflex.AsString())
}

func (c ConfigImages) SecureBaseURL() (string, error) {
	return jsonflex.GetField(c, "secure_base_url", jsonflex.AsString())
}

func (c ConfigImages) BackdropSizes() ([]string, error) {
	return jsonflex.GetField(c, "backdrop_sizes", jsonflex.AsArray(jsonflex.AsString()))
}

func (c ConfigImages) LogoSizes() ([]string, error) {
	return jsonflex.GetField(c, "logo_sizes", jsonflex.AsArray(jsonflex.AsString()))
}

func (c ConfigImages) PosterSizes() ([]string, error) {
	return jsonflex.GetField(c, "poster_sizes", jsonflex.AsArray(jsonflex.AsString()))
}

func (c ConfigImages) ProfileSizes() ([]string, error) {
	return jsonflex.GetField(c, "profile_sizes", jsonflex.AsArray(jsonflex.AsString()))
}

func (c ConfigImages) StillSizes() ([]string, error) {
	return jsonflex.GetField(c, "still_sizes", jsonflex.AsArray(jsonflex.AsString()))
}

func GetConfigDetails(ctx context.Context, client Client, opts ...RequestOption) (ConfigDetails, error) {
	return client.Get(ctx, "/3/configuration", opts...)
}

func GetConfigCountries(ctx context.Context, client Client, opts ...RequestOption) ([]Country, error) {
	if countries, err := client.GetArray(ctx, "/3/configuration/countries", opts...); err != nil {
		return nil, err
	} else {
		return jsonflex.FromArray(countries, jsonflex.AsObject[Country]())
	}
}