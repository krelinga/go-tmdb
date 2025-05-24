package tmdb

import (
	"context"
	"fmt"
	"strings"
)

type Movie struct {
	Adult bool `json:"adult"`

	// Additional bits that can be fetched at the same time.
	Keywords *MovieKeywords `json:"keywords,omitempty"`
}

type MovieKeywords struct {
	Keywords []*Keyword `json:"keywords"`
}

func GetMovie(client Client, movieId int, options ...GetMovieOption) (*Movie, error) {
	o := getMovieOptions{}
	for _, opt := range options {
		opt.applyToGetMovieOptions(&o)
	}

	appends := []string{}
	if o.wantKeywords {
		appends = append(appends, "keywords")
	}

	params := GetParams{}
	if len(appends) > 0 {
		params["append_to_response"] = strings.Join(appends, ",")
	}
	m := &Movie{}
	ctx := context.Background()
	if o.useContext != nil {
		ctx = *o.useContext
	}
	endpoint := fmt.Sprintf("/movie/%d", movieId)
	if err := client.Get(ctx, endpoint, params, m); err != nil {
		return nil, err
	}
	return m, nil
}

type getMovieOptions struct {
	wantDetails  bool
	wantKeywords bool
	useContext   *context.Context
}

type GetMovieOption interface {
	applyToGetMovieOptions(o *getMovieOptions)
}
