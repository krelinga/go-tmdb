package tmdbepisode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/krelinga/go-tmdb/internal/util"
)

type GetDetailsOptions struct {
	Key             string
	ReadAccessToken string
	Language        string
}

func GetDetails(ctx context.Context, client *http.Client, seriesID, seasonNumber, episodeNumber int32, options GetDetailsOptions) (*http.Response, error) {
	values := url.Values{}
	util.SetIfNotZero(&values, "api_key", options.Key)
	util.SetIfNotZero(&values, "language", options.Language)
	url := &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     fmt.Sprintf("/3/tv/%d/season/%d/episode/%d", seriesID, seasonNumber, episodeNumber),
		RawQuery: values.Encode(),
	}
	return util.MakeRequest(ctx, client, url, options.ReadAccessToken)
}

func ParseGetDetailsReply(httpReply *http.Response) (*GetDetailsReply, error) {
	defer httpReply.Body.Close()
	if httpReply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", httpReply.StatusCode)
	}
	if httpReply.Header.Get("Content-Type") != "application/json;charset=utf-8" {
		return nil, fmt.Errorf("unexpected content type: %s", httpReply.Header.Get("Content-Type"))
	}

	rawReply := &struct {
		ID *int32 `json:"id"`
		*Details
	}{}
	if err := json.NewDecoder(httpReply.Body).Decode(rawReply); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &GetDetailsReply{
		ID:      rawReply.ID,
		Details: rawReply.Details,
	}, nil
}

type GetDetailsReply struct {
	ID      *int32
	Details *Details
}

func (r *GetDetailsReply) SetDefaults() {
	if r == nil {
		return
	}
	util.SetIfNil(&r.ID, 0)
	if r.Details != nil {
		r.Details.SetDefaults()
	}
}

func (r *GetDetailsReply) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ID: %s Details: %s}", util.FmtOrNil(r.ID), r.Details.String())
}

type Details struct {
	AirDate        *string      `json:"air_date"`
	Crew           []*Crew      `json:"crew"`
	EpisodeNumber  *int32       `json:"episode_number"`
	GuestStars     []*GuestStar `json:"guest_stars"`
	Name           *string      `json:"name"`
	Overview       *string      `json:"overview"`
	ProductionCode *string      `json:"production_code"`
	Runtime        *int32       `json:"runtime"`
	SeasonNumber   *int32       `json:"season_number"`
	StillPath      *string      `json:"still_path"`
	VoteAverage    *float32     `json:"vote_average"`
	VoteCount      *int32       `json:"vote_count"`
}

func (d *Details) SetDefaults() {
	if d == nil {
		return
	}
	for _, crew := range d.Crew {
		crew.SetDefaults()
	}
	util.SetIfNil(&d.EpisodeNumber, 0)
	for _, guestStar := range d.GuestStars {
		guestStar.SetDefaults()
	}
	util.SetIfNil(&d.Runtime, 0)
	util.SetIfNil(&d.SeasonNumber, 0)
	util.SetIfNil(&d.VoteAverage, 0.0)
	util.SetIfNil(&d.VoteCount, 0)
}

func (d *Details) String() string {
	if d == nil {
		return "<nil>"
	}
	var b strings.Builder
	b.WriteString("{")
	fmt.Fprintf(&b, "AirDate: %s", util.FmtOrNil(d.AirDate))
	fmt.Fprintf(&b, " Crew: %v", d.Crew)
	fmt.Fprintf(&b, " EpisodeNumber: %s", util.FmtOrNil(d.EpisodeNumber))
	fmt.Fprintf(&b, " GuestStars: %v", d.GuestStars)
	fmt.Fprintf(&b, " Name: %s", util.FmtOrNil(d.Name))
	fmt.Fprintf(&b, " Overview: %s", util.FmtOrNil(d.Overview))
	fmt.Fprintf(&b, " ProductionCode: %s", util.FmtOrNil(d.ProductionCode))
	fmt.Fprintf(&b, " Runtime: %s", util.FmtOrNil(d.Runtime))
	fmt.Fprintf(&b, " SeasonNumber: %s", util.FmtOrNil(d.SeasonNumber))
	fmt.Fprintf(&b, " StillPath: %s", util.FmtOrNil(d.StillPath))
	fmt.Fprintf(&b, " VoteAverage: %s", util.FmtOrNil(d.VoteAverage))
	fmt.Fprintf(&b, " VoteCount: %s", util.FmtOrNil(d.VoteCount))
	b.WriteString("}")
	return b.String()
}

type Crew struct {
	Department         *string  `json:"department"`
	Job                *string  `json:"job"`
	CreditID           *string  `json:"credit_id"`
	Adult              *bool    `json:"adult"`
	Gender             *int32   `json:"gender"`
	ID                 *int32   `json:"id"`
	KnownForDepartment *string  `json:"known_for_department"`
	Name               *string  `json:"name"`
	OriginalName       *string  `json:"original_name"`
	Popularity         *float32 `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
}

func (c *Crew) SetDefaults() {
	if c == nil {
		return
	}
	util.SetIfNil(&c.Adult, true)
	util.SetIfNil(&c.Gender, 0)
	util.SetIfNil(&c.ID, 0)
	util.SetIfNil(&c.Popularity, 0.0)
}

func (c *Crew) String() string {
	if c == nil {
		return "<nil>"
	}
	var b strings.Builder
	b.WriteString("{")
	fmt.Fprintf(&b, "Department: %s", util.FmtOrNil(c.Department))
	fmt.Fprintf(&b, " Job: %s", util.FmtOrNil(c.Job))
	fmt.Fprintf(&b, " CreditID: %s", util.FmtOrNil(c.CreditID))
	fmt.Fprintf(&b, " Adult: %s", util.FmtOrNil(c.Adult))
	fmt.Fprintf(&b, " Gender: %s", util.FmtOrNil(c.Gender))
	fmt.Fprintf(&b, " ID: %s", util.FmtOrNil(c.ID))
	fmt.Fprintf(&b, " KnownForDepartment: %s", util.FmtOrNil(c.KnownForDepartment))
	fmt.Fprintf(&b, " Name: %s", util.FmtOrNil(c.Name))
	fmt.Fprintf(&b, " OriginalName: %s", util.FmtOrNil(c.OriginalName))
	fmt.Fprintf(&b, " Popularity: %s", util.FmtOrNil(c.Popularity))
	fmt.Fprintf(&b, " ProfilePath: %s", util.FmtOrNil(c.ProfilePath))
	b.WriteString("}")
	return b.String()
}

type GuestStar struct {
	Character          *string  `json:"character"`
	CreditID           *string  `json:"credit_id"`
	Order              *int32   `json:"order"`
	Adult              *bool    `json:"adult"`
	Gender             *int32   `json:"gender"`
	ID                 *int32   `json:"id"`
	KnownForDepartment *string  `json:"known_for_department"`
	Name               *string  `json:"name"`
	OriginalName       *string  `json:"original_name"`
	Popularity         *float32 `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
}

func (gs *GuestStar) SetDefaults() {
	if gs == nil {
		return
	}
	util.SetIfNil(&gs.Order, 0)
	util.SetIfNil(&gs.Adult, true)
	util.SetIfNil(&gs.ID, 0)
	util.SetIfNil(&gs.Popularity, 0.0)
}

func (gs *GuestStar) String() string {
	if gs == nil {
		return "<nil>"
	}
	var b strings.Builder
	b.WriteString("{")
	fmt.Fprintf(&b, "Character: %s", util.FmtOrNil(gs.Character))
	fmt.Fprintf(&b, " CreditID: %s", util.FmtOrNil(gs.CreditID))
	fmt.Fprintf(&b, " Order: %s", util.FmtOrNil(gs.Order))
	fmt.Fprintf(&b, " Adult: %s", util.FmtOrNil(gs.Adult))
	fmt.Fprintf(&b, " Gender: %s", util.FmtOrNil(gs.Gender))
	fmt.Fprintf(&b, " ID: %s", util.FmtOrNil(gs.ID))
	fmt.Fprintf(&b, " KnownForDepartment: %s", util.FmtOrNil(gs.KnownForDepartment))
	fmt.Fprintf(&b, " Name: %s", util.FmtOrNil(gs.Name))
	fmt.Fprintf(&b, " OriginalName: %s", util.FmtOrNil(gs.OriginalName))
	fmt.Fprintf(&b, " Popularity: %s", util.FmtOrNil(gs.Popularity))
	fmt.Fprintf(&b, " ProfilePath: %s", util.FmtOrNil(gs.ProfilePath))
	b.WriteString("}")
	return b.String()
}
