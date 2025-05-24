package tmdb

import (
	"context"
	"encoding/json"
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
	ctx := context.Background()
	if o.useContext != nil {
		ctx = *o.useContext
	}
	endpoint := fmt.Sprintf("/movie/%d", movieId)
	data, err := client.Get(ctx, endpoint, params)
	if err != nil {
		return nil, err
	}
	if o.rawReply != nil {
		*o.rawReply = data
	}
	m := &Movie{}
	if err := json.Unmarshal(data, m); err != nil {
		return nil, fmt.Errorf("unmarshalling movie: %w", err)
	}

	return m, nil
}

type getMovieOptions struct {
	baseOptions
	wantDetails  bool
	wantKeywords bool
}

type GetMovieOption interface {
	applyToGetMovieOptions(o *getMovieOptions)
}
