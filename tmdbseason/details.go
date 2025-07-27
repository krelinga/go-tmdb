package tmdbseason

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/krelinga/go-tmdb/internal/util"
)

type GetDetailsOptions struct {
	Language string
}

func GetDetails(ctx context.Context, seriesID, seasonNumber int32, options GetDetailsOptions) (*http.Response, error) {
	return util.NewRequestBuilder(fmt.Sprintf("/3/tv/%d/season/%d", seriesID, seasonNumber)).
		SetValueString("language", options.Language).
		Do(ctx)
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

	reply := &GetDetailsReply{
		ID:      rawReply.ID,
		Details: rawReply.Details,
	}
	return reply, nil
}

type GetDetailsReply struct {
	ID      *int32
	Details *Details
}

func (gdr *GetDetailsReply) SetDefaults() {
	if gdr == nil {
		return
	}
	util.SetIfNil(&gdr.ID, 0)
	gdr.Details.SetDefaults()
}

func (gdr *GetDetailsReply) String() string {
	if gdr == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ID: %s Details: %s}", util.FmtOrNil(gdr.ID), gdr.Details.String())
}

type Details struct {
	AirDate      *string    `json:"air_date"`
	Episodes     []*Episode `json:"episodes"`
	Name         *string    `json:"name"`
	Overview     *string    `json:"overview"`
	PosterPath   *string    `json:"poster_path"`
	SeasonNumber *int32     `json:"season_number"`
	VoteAverage  *float32   `json:"vote_average"`
}

func (d *Details) SetDefaults() {
	if d == nil {
		return
	}
	for _, episode := range d.Episodes {
		episode.SetDefaults()
	}
	util.SetIfNil(&d.SeasonNumber, 0)
	util.SetIfNil(&d.VoteAverage, 0.0)
}

func (d *Details) String() string {
	if d == nil {
		return "<nil>"
	}
	var b strings.Builder
	b.WriteString("{")
	fmt.Fprintf(&b, "AirDate: %s", util.FmtOrNil(d.AirDate))
	fmt.Fprintf(&b, " Episodes: %v", d.Episodes)
	fmt.Fprintf(&b, " Name: %s", util.FmtOrNil(d.Name))
	fmt.Fprintf(&b, " Overview: %s", util.FmtOrNil(d.Overview))
	fmt.Fprintf(&b, " PosterPath: %s", util.FmtOrNil(d.PosterPath))
	fmt.Fprintf(&b, " SeasonNumber: %s", util.FmtOrNil(d.SeasonNumber))
	fmt.Fprintf(&b, " VoteAverage: %s", util.FmtOrNil(d.VoteAverage))
	b.WriteString("}")
	return b.String()
}

type Episode struct {
	AirDate        *string      `json:"air_date"`
	EpisodeNumber  *int32       `json:"episode_number"`
	ID             *int32       `json:"id"`
	Name           *string      `json:"name"`
	Overview       *string      `json:"overview"`
	ProductionCode *string      `json:"production_code"`
	Runtime        *int32       `json:"runtime"`
	SeasonNumber   *int32       `json:"season_number"`
	SeriesID       *int32       `json:"show_id"`
	StillPath      *string      `json:"still_path"`
	VoteAverage    *float32     `json:"vote_average"`
	VoteCount      *int32       `json:"vote_count"`
	Crew           []*Crew      `json:"crew"`
	GuestStars     []*GuestStar `json:"guest_stars"`
}

func (e *Episode) SetDefaults() {
	if e == nil {
		return
	}
	util.SetIfNil(&e.EpisodeNumber, 0)
	util.SetIfNil(&e.ID, 0)
	util.SetIfNil(&e.Runtime, 0)
	util.SetIfNil(&e.SeasonNumber, 0)
	util.SetIfNil(&e.SeriesID, 0)
	util.SetIfNil(&e.VoteAverage, 0.0)
	util.SetIfNil(&e.VoteCount, 0)
	for _, crew := range e.Crew {
		crew.SetDefaults()
	}
	for _, guestStar := range e.GuestStars {
		guestStar.SetDefaults()
	}
}

func (e *Episode) String() string {
	if e == nil {
		return "<nil>"
	}
	var b strings.Builder
	b.WriteString("{")
	fmt.Fprintf(&b, "AirDate: %s", util.FmtOrNil(e.AirDate))
	fmt.Fprintf(&b, " EpisodeNumber: %s", util.FmtOrNil(e.EpisodeNumber))
	fmt.Fprintf(&b, " ID: %s", util.FmtOrNil(e.ID))
	fmt.Fprintf(&b, " Name: %s", util.FmtOrNil(e.Name))
	fmt.Fprintf(&b, " Overview: %s", util.FmtOrNil(e.Overview))
	fmt.Fprintf(&b, " ProductionCode: %s", util.FmtOrNil(e.ProductionCode))
	fmt.Fprintf(&b, " Runtime: %s", util.FmtOrNil(e.Runtime))
	fmt.Fprintf(&b, " SeasonNumber: %s", util.FmtOrNil(e.SeasonNumber))
	fmt.Fprintf(&b, " SeriesID: %s", util.FmtOrNil(e.SeriesID))
	fmt.Fprintf(&b, " StillPath: %s", util.FmtOrNil(e.StillPath))
	fmt.Fprintf(&b, " VoteAverage: %s", util.FmtOrNil(e.VoteAverage))
	fmt.Fprintf(&b, " VoteCount: %s", util.FmtOrNil(e.VoteCount))
	fmt.Fprintf(&b, " Crew: %v", e.Crew)
	fmt.Fprintf(&b, " GuestStars: %v", e.GuestStars)
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
	util.SetIfNil(&gs.Gender, 0)
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
