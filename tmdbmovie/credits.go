package tmdbmovie

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/krelinga/go-tmdb/internal/util"
)

type GetCreditsOptions struct {
	Key             string
	ReadAccessToken string
	Language        string
}

func GetCredits(ctx context.Context, client *http.Client, movieID int32, options GetCreditsOptions) (*http.Response, error) {
	values := url.Values{}
	util.SetIfNotZero(&values, "api_key", options.Key)
	util.SetIfNotZero(&values, "language", options.Language)
	url := &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     "/3/movie/" + fmt.Sprint(movieID) + "/credits",
		RawQuery: values.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    url,
	}
	util.SetAuthIfNotZero(request, options.ReadAccessToken)
	httpReply, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return httpReply, nil
}

func ParseGetCreditsReply(httpReply *http.Response) (*GetCreditsReply, error) {
	defer httpReply.Body.Close()
	if httpReply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", httpReply.StatusCode)
	}
	if httpReply.Header.Get("Content-Type") != "application/json;charset=utf-8" {
		return nil, fmt.Errorf("unexpected content type: %s", httpReply.Header.Get("Content-Type"))
	}

	rawReply := &struct {
		ID *int32 `json:"id"`
		*Credits
	}{}
	if err := json.NewDecoder(httpReply.Body).Decode(rawReply); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &GetCreditsReply{
		ID:      rawReply.ID,
		Credits: rawReply.Credits,
	}, nil
}

type GetCreditsReply struct {
	ID      *int32
	Credits *Credits
}

func (gcr *GetCreditsReply) SetDefaults() {
	if gcr == nil {
		return
	}
	util.SetIfNil(&gcr.ID, 0)
	gcr.Credits.SetDefaults()
}

func (gcr *GetCreditsReply) String() string {
	if gcr == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ID: %s Credits: %v}", util.FmtOrNil(gcr.ID), gcr.Credits)
}

type Credits struct {
	Cast []*Cast `json:"cast"`
	Crew []*Crew `json:"crew"`
}

func (c *Credits) SetDefaults() {
	if c == nil {
		return
	}
	for _, cast := range c.Cast {
		cast.SetDefaults()
	}
	for _, crew := range c.Crew {
		crew.SetDefaults()
	}
}

func (c *Credits) String() string {
	if c == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Cast: %v Crew: %v}", c.Cast, c.Crew)
}

type Cast struct {
	Adult              *bool    `json:"adult"`
	Gender             *int32   `json:"gender"`
	ID                 *int32   `json:"id"`
	KnownForDepartment *string  `json:"known_for_department"`
	Name               *string  `json:"name"`
	OriginalName       *string  `json:"original_name"`
	Popularity         *float32 `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
	CastID             *int32   `json:"cast_id"`
	Character          *string  `json:"character"`
	CreditID           *string  `json:"credit_id"`
	Order              *int32   `json:"order"`
}

func (c *Cast) SetDefaults() {
	if c == nil {
		return
	}
	util.SetIfNil(&c.Adult, true)
	util.SetIfNil(&c.Gender, 0)
	util.SetIfNil(&c.ID, 0)
	util.SetIfNil(&c.Popularity, 0.0)
	util.SetIfNil(&c.CastID, 0)
	util.SetIfNil(&c.Order, 0)
}

func (c *Cast) String() string {
	if c == nil {
		return "<nil>"
	}
	var builder strings.Builder
	builder.WriteString("{")
	fmt.Fprintf(&builder, "Adult: %s", util.FmtOrNil(c.Adult))
	fmt.Fprintf(&builder, " Gender: %s", util.FmtOrNil(c.Gender))
	fmt.Fprintf(&builder, " ID: %s", util.FmtOrNil(c.ID))
	fmt.Fprintf(&builder, " KnownForDepartment: %s", util.FmtOrNil(c.KnownForDepartment))
	fmt.Fprintf(&builder, " Name: %s", util.FmtOrNil(c.Name))
	fmt.Fprintf(&builder, " OriginalName: %s", util.FmtOrNil(c.OriginalName))
	fmt.Fprintf(&builder, " Popularity: %s", util.FmtOrNil(c.Popularity))
	fmt.Fprintf(&builder, " ProfilePath: %s", util.FmtOrNil(c.ProfilePath))
	fmt.Fprintf(&builder, " CastID: %s", util.FmtOrNil(c.CastID))
	fmt.Fprintf(&builder, " Character: %s", util.FmtOrNil(c.Character))
	fmt.Fprintf(&builder, " CreditID: %s", util.FmtOrNil(c.CreditID))
	fmt.Fprintf(&builder, " Order: %s", util.FmtOrNil(c.Order))
	builder.WriteString("}")
	return builder.String()
}

type Crew struct {
	Adult              *bool    `json:"adult"`
	Gender             *int32   `json:"gender"`
	ID                 *int32   `json:"id"`
	KnownForDepartment *string  `json:"known_for_department"`
	Name               *string  `json:"name"`
	OriginalName       *string  `json:"original_name"`
	Popularity         *float32 `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
	CreditID           *string  `json:"credit_id"`
	Department         *string  `json:"department"`
	Job                *string  `json:"job"`
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
	var builder strings.Builder
	builder.WriteString("{")
	fmt.Fprintf(&builder, "Adult: %s", util.FmtOrNil(c.Adult))
	fmt.Fprintf(&builder, " Gender: %s", util.FmtOrNil(c.Gender))
	fmt.Fprintf(&builder, " ID: %s", util.FmtOrNil(c.ID))
	fmt.Fprintf(&builder, " KnownForDepartment: %s", util.FmtOrNil(c.KnownForDepartment))
	fmt.Fprintf(&builder, " Name: %s", util.FmtOrNil(c.Name))
	fmt.Fprintf(&builder, " OriginalName: %s", util.FmtOrNil(c.OriginalName))
	fmt.Fprintf(&builder, " Popularity: %s", util.FmtOrNil(c.Popularity))
	fmt.Fprintf(&builder, " ProfilePath: %s", util.FmtOrNil(c.ProfilePath))
	fmt.Fprintf(&builder, " CreditID: %s", util.FmtOrNil(c.CreditID))
	fmt.Fprintf(&builder, " Department: %s", util.FmtOrNil(c.Department))
	fmt.Fprintf(&builder, " Job: %s", util.FmtOrNil(c.Job))
	builder.WriteString("}")
	return builder.String()
}
