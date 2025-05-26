package tmdb

import (
	"context"
	"encoding/json"
)

type BackdropSize string
type LogoSize string
type PosterSize string
type ProfileSize string
type StillSize string

type Configuration struct {
	Images struct {
		BaseUrl       string         `json:"base_url"`
		SecureBaseUrl string         `json:"secure_base_url"`
		BackdropSizes []BackdropSize `json:"backdrop_sizes"`
		LogoSizes     []LogoSize     `json:"logo_sizes"`
		PosterSizes   []PosterSize   `json:"poster_sizes"`
		ProfileSizes  []ProfileSize  `json:"profile_sizes"`
		StillSizes    []StillSize    `json:"still_sizes"`
	} `json:"images"`
	ChangeKeys []string `json:"change_keys"`
}

func GetConfiguration(client Client, options ...GetConfigurationOption) (*Configuration, error) {
	o := getConfigruationOptions{}
	for _, opt := range options {
		opt.applyToGetConfigurationOptions(&o)
	}
	ctx := context.Background()
	if o.useContext != nil {
		ctx = *o.useContext
	}
	endpoint := "/configuration"
	params := GetParams{}
	data, err := checkCode(client.Get(ctx, endpoint, params))
	if err != nil {
		return nil, err
	}
	if o.rawReply != nil {
		*o.rawReply = data
	}

	c := &Configuration{}
	if err := json.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return c, nil
}

type GetConfigurationOption interface {
	applyToGetConfigurationOptions(*getConfigruationOptions)
}

type getConfigruationOptions struct {
	baseOptions
}
