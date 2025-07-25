package tmdbmovie

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/krelinga/go-tmdb/internal/util"
)

type GetReleaseDatesOptions struct {
	Key             string
	ReadAccessToken string
}

func GetReleaseDates(ctx context.Context, client *http.Client, id int32, options GetReleaseDatesOptions) (*http.Response, error) {
	values := url.Values{}
	util.SetIfNotZero(&values, "api_key", options.Key)
	url := &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     "/3/movie/" + fmt.Sprint(id) + "/release_dates",
		RawQuery: values.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    url,
	}
	util.SetAuthIfNotZero(request, options.ReadAccessToken)
	httpReply, err := client.Do(request.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return httpReply, nil
}

func ParseGetReleaseDatesReply(httpReply *http.Response) (*GetReleaseDatesReply, error) {
	defer httpReply.Body.Close()
	if httpReply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", httpReply.StatusCode)
	}
	if httpReply.Header.Get("Content-Type") != "application/json;charset=utf-8" {
		return nil, fmt.Errorf("unexpected content type: %s", httpReply.Header.Get("Content-Type"))
	}

	reply := &GetReleaseDatesReply{}
	if err := json.NewDecoder(httpReply.Body).Decode(reply); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return reply, nil
}

type GetReleaseDatesReply struct {
	ID                  *int32                `json:"id"`
	CountryReleaseDates []*CountryReleaseDate `json:"results"`
}

func (grd *GetReleaseDatesReply) SetDefaults() {
	if grd == nil {
		return
	}
	util.SetIfNil(&grd.ID, 0)
	for _, crd := range grd.CountryReleaseDates {
		crd.SetDefaults()
	}
}

func (grd *GetReleaseDatesReply) String() string {
	if grd == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ID: %s CountryReleaseDates: %v}", util.FmtOrNil(grd.ID), grd.CountryReleaseDates)
}

type CountryReleaseDate struct {
	ISO3166_1    *string        `json:"iso_3166_1"`
	ReleaseDates []*ReleaseDate `json:"release_dates"`
}

func (crd *CountryReleaseDate) SetDefaults() {
	if crd == nil {
		return
	}
	for _, rd := range crd.ReleaseDates {
		rd.SetDefaults()
	}
}

func (crd *CountryReleaseDate) String() string {
	if crd == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ISO3166_1: %s ReleaseDates: %v}", util.FmtOrNil(crd.ISO3166_1), crd.ReleaseDates)
}

type ReleaseDate struct {
	Certification *string `json:"certification"`
	// TODO: Descriptors ... what's up with this?
	ISO639_1    *string `json:"iso_639_1"`
	Note        *string `json:"note"`
	ReleaseDate *string `json:"release_date"`
	Type        *int32  `json:"type"`
}

func (rd *ReleaseDate) SetDefaults() {
	if rd == nil {
		return
	}
	util.SetIfNil(&rd.Type, 0)
}

func (rd *ReleaseDate) String() string {
	if rd == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Certification: %s ISO639_1: %s Note: %s ReleaseDate: %s Type: %s}",
		util.FmtOrNil(rd.Certification), util.FmtOrNil(rd.ISO639_1), util.FmtOrNil(rd.Note),
		util.FmtOrNil(rd.ReleaseDate), util.FmtOrNil(rd.Type))
}
