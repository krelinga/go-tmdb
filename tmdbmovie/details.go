package tmdbmovie

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

	AppendCredits      bool
	AppendExternalIDs  bool
	AppendReleaseDates bool
}

func GetDetails(ctx context.Context, id int32, options GetDetailsOptions) (*http.Response, error) {
	return util.NewRequestBuilder().
		SetPath("/3/movie/"+fmt.Sprint(id)).
		SetValueString("language", options.Language).
		AppendToResponse("credits", options.AppendCredits).
		AppendToResponse("external_ids", options.AppendExternalIDs).
		AppendToResponse("release_dates", options.AppendReleaseDates).
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
		Credits      *Credits              `json:"credits"`
		ExternalIDs  *ExternalIDs          `json:"external_ids"`
		ReleaseDates *GetReleaseDatesReply `json:"release_dates"`
	}{}
	if err := json.NewDecoder(httpReply.Body).Decode(rawReply); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &GetDetailsReply{
		ID:          rawReply.ID,
		Details:     rawReply.Details,
		Credits:     rawReply.Credits,
		ExternalIDs: rawReply.ExternalIDs,
		CountryReleaseDates: func() []*CountryReleaseDate {
			if rawReply.ReleaseDates == nil {
				return nil
			}
			return rawReply.ReleaseDates.CountryReleaseDates
		}(),
	}, nil
}

type GetDetailsReply struct {
	ID      *int32
	Details *Details

	Credits             *Credits
	ExternalIDs         *ExternalIDs
	CountryReleaseDates []*CountryReleaseDate
}

func (gdr *GetDetailsReply) SetDefaults() {
	if gdr == nil {
		return
	}
	util.SetIfNil(&gdr.ID, 0)
	gdr.Details.SetDefaults()
	gdr.Credits.SetDefaults()
	gdr.ExternalIDs.SetDefaults()
	for _, countryReleaseDate := range gdr.CountryReleaseDates {
		countryReleaseDate.SetDefaults()
	}
}

func (ge *GetDetailsReply) String() string {
	if ge == nil {
		return "<nil>"
	}
	builder := strings.Builder{}
	builder.WriteString("{")
	fmt.Fprintf(&builder, "ID: %v", util.FmtOrNil(ge.ID))
	fmt.Fprintf(&builder, " Details: %s", ge.Details.String())
	fmt.Fprintf(&builder, " Credits: %v", ge.Credits)
	fmt.Fprintf(&builder, " ExternalIDs: %v", ge.ExternalIDs)
	fmt.Fprintf(&builder, " CountryReleaseDates: %v", ge.CountryReleaseDates)
	builder.WriteString("}")
	return builder.String()
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

func (d *Details) SetDefaults() {
	if d == nil {
		return
	}
	util.SetIfNil(&d.Adult, true)
	util.SetIfNil(&d.Budget, 0)
	for _, genre := range d.Genres {
		genre.SetDefaults()
	}
	util.SetIfNil(&d.Popularity, 0.0)
	for _, company := range d.ProductionCompanies {
		company.SetDefaults()
	}
	for _, country := range d.ProductionCountries {
		country.SetDefaults()
	}
	util.SetIfNil(&d.Revenue, 0)
	util.SetIfNil(&d.Runtime, 0)
	for _, lang := range d.SpokenLanguages {
		lang.SetDefaults()
	}
	util.SetIfNil(&d.Video, false)
	util.SetIfNil(&d.VoteAverage, 0.0)
	util.SetIfNil(&d.VoteCount, 0)
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

func (g *Genre) SetDefaults() {
	if g == nil {
		return
	}
	util.SetIfNil(&g.ID, 0)
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

func (pc *ProductionCompany) SetDefaults() {
	if pc == nil {
		return
	}
	util.SetIfNil(&pc.ID, 0)
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

func (pc *ProductionCountry) SetDefaults() {
	// Nothing to do here for now, we have this function for consistency with other types.
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

func (sl *SpokenLanguage) SetDefaults() {
	// Nothing to do here for now, we have this function for consistency with other types.
}

func (sl *SpokenLanguage) String() string {
	if sl == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EnglishName: %s ISO639_1: %s Name: %s}", util.FmtOrNil(sl.EnglishName), util.FmtOrNil(sl.ISO639_1), util.FmtOrNil(sl.Name))
}
