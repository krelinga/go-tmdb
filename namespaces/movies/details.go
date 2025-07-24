package movies

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

func GetDetails(ctx context.Context, client *http.Client, id int32, options GetDetailsOptions) (*GetDetailsReply, error) {
	values := url.Values{}
	util.SetIfNotZero(&values, "api_key", options.Key)
	util.SetIfNotZero(&values, "language", options.Language)
	url := &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     "/3/movie/" + fmt.Sprint(id),
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

func (ge *GetDetailsReply) String() string {
	if ge == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ID: %v Details: %v}", ge.ID, ge.Details.String())
}

type Details struct {
	Adult        *bool   `json:"adult"`
	BackdropPath *string `json:"backdrop_path"`
	// TODO: this is passed as an object rather than a string?
	// BelongsToCollection *string              `json:"belongs_to_collection"`
	Budget              *int32               `json:"budget"`
	Genres              []*Genre             `json:"genres"`
	Homepage            *string              `json:"homepage"`
	IMDBID              *string              `json:"imdb_id"`
	OriginalLanguage    *string              `json:"original_language"`
	OriginalTitle       *string              `json:"original_title"`
	Overview            *string              `json:"overview"`
	Popularity          *float32             `json:"popularity"`
	PosterPath          *string              `json:"poster_path"`
	ProductionCompanies []*ProductionCompany `json:"production_companies"`
	ProductionCountries []*ProductionCountry `json:"production_countries"`
	ReleaseDate         *string              `json:"release_date"`
	Revenue             *int32               `json:"revenue"`
	Runtime             *int32               `json:"runtime"`
	SpokenLanguages     []*SpokenLanguage    `json:"spoken_languages"`
	Status              *string              `json:"status"`
	Tagline             *string              `json:"tagline"`
	Title               *string              `json:"title"`
	Video               *bool                `json:"video"`
	VoteAverage         *float32             `json:"vote_average"`
	VoteCount           *int32               `json:"vote_count"`
}

func (d *Details) String() string {
	if d == nil {
		return "<nil>"
	}
	var builder strings.Builder
	builder.WriteString("{")
	fmt.Fprintf(&builder, "Adult: %s", util.FmtOrNil(d.Adult))
	fmt.Fprintf(&builder, " BackdropPath: %s", util.FmtOrNil(d.BackdropPath))
	// TODO: uncomment this if the field gets fixed.
	// fmt.Fprintf(&builder, " BelongsToCollection: %s", util.FmtOrNil(d.BelongsToCollection))
	fmt.Fprintf(&builder, " Budget: %s", util.FmtOrNil(d.Budget))
	fmt.Fprintf(&builder, " Genres: %v", d.Genres)
	fmt.Fprintf(&builder, " Homepage: %s", util.FmtOrNil(d.Homepage))
	fmt.Fprintf(&builder, " IMDBID: %s", util.FmtOrNil(d.IMDBID))
	fmt.Fprintf(&builder, " OriginalLanguage: %s", util.FmtOrNil(d.OriginalLanguage))
	fmt.Fprintf(&builder, " OriginalTitle: %s", util.FmtOrNil(d.OriginalTitle))
	fmt.Fprintf(&builder, " Overview: %s", util.FmtOrNil(d.Overview))
	fmt.Fprintf(&builder, " Popularity: %s", util.FmtOrNil(d.Popularity))
	fmt.Fprintf(&builder, " PosterPath: %s", util.FmtOrNil(d.PosterPath))
	fmt.Fprintf(&builder, " ProductionCompanies: %v", d.ProductionCompanies)
	fmt.Fprintf(&builder, " ProductionCountries: %v", d.ProductionCountries)
	fmt.Fprintf(&builder, " ReleaseDate: %s", util.FmtOrNil(d.ReleaseDate))
	fmt.Fprintf(&builder, " Revenue: %s", util.FmtOrNil(d.Revenue))
	fmt.Fprintf(&builder, " Runtime: %s", util.FmtOrNil(d.Runtime))
	fmt.Fprintf(&builder, " SpokenLanguages: %v", d.SpokenLanguages)
	fmt.Fprintf(&builder, " Status: %s", util.FmtOrNil(d.Status))
	fmt.Fprintf(&builder, " Tagline: %s", util.FmtOrNil(d.Tagline))
	fmt.Fprintf(&builder, " Title: %s", util.FmtOrNil(d.Title))
	fmt.Fprintf(&builder, " Video: %s", util.FmtOrNil(d.Video))
	fmt.Fprintf(&builder, " VoteAverage: %s", util.FmtOrNil(d.VoteAverage))
	fmt.Fprintf(&builder, " VoteCount: %s", util.FmtOrNil(d.VoteCount))
	builder.WriteString("}")
	return builder.String()
}

type Genre struct {
	ID   *int32  `json:"id"`
	Name *string `json:"name"`
}

func (g *Genre) String() string {
	if g == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ID: %s Name: %s}", util.FmtOrNil(g.ID), util.FmtOrNil(g.Name))
}

type ProductionCompany struct {
	ID            *int32  `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          *string `json:"name"`
	OriginCountry *string `json:"origin_country"`
}

func (pc *ProductionCompany) String() string {
	if pc == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ID: %s LogoPath: %s Name: %s OriginCountry: %s}", util.FmtOrNil(pc.ID), util.FmtOrNil(pc.LogoPath), util.FmtOrNil(pc.Name), util.FmtOrNil(pc.OriginCountry))
}

type ProductionCountry struct {
	ISO3166_1 *string `json:"iso_3166_1"`
	Name      *string `json:"name"`
}

func (pc *ProductionCountry) String() string {
	if pc == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ISO3166_1: %s Name: %s}", util.FmtOrNil(pc.ISO3166_1), util.FmtOrNil(pc.Name))
}

type SpokenLanguage struct {
	EnglishName *string `json:"english_name"`
	ISO639_1    *string `json:"iso_639_1"`
	Name        *string `json:"name"`
}

func (sl *SpokenLanguage) String() string {
	if sl == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EnglishName: %s ISO639_1: %s Name: %s}", util.FmtOrNil(sl.EnglishName), util.FmtOrNil(sl.ISO639_1), util.FmtOrNil(sl.Name))
}
