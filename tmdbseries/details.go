package tmdbseries

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

func GetDetails(ctx context.Context, client *http.Client, SeriesID int32, options GetDetailsOptions) (*http.Response, error) {
	values := url.Values{}
	util.SetIfNotZero(&values, "api_key", options.Key)
	util.SetIfNotZero(&values, "language", options.Language)
	url := &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     "/3/tv/" + fmt.Sprintf("%d", SeriesID),
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

func ParseGetDetailsReply(httpReply *http.Response) (*GetDetailsReply, error) {
	defer httpReply.Body.Close()

	if httpReply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %v", httpReply.Status)
	}

	rawReply := &struct {
		ID *int32 `json:"id"`
		*Details
	}{}
	if err := json.NewDecoder(httpReply.Body).Decode(rawReply); err != nil {
		return nil, err
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
	builder := strings.Builder{}
	builder.WriteString("{")
	fmt.Fprintf(&builder, "ID: %v", util.FmtOrNil(gdr.ID))
	fmt.Fprintf(&builder, " Details: %s", gdr.Details.String())
	builder.WriteString("}")
	return builder.String()
}

type Details struct {
	Adult               *bool                `json:"adult"`
	BackdropPath        *string              `json:"backdrop_path"`
	Creators            []*Creator           `json:"created_by"`
	EpisodeRunTime      []int32              `json:"episode_run_time"`
	FirstAirDate        *string              `json:"first_air_date"`
	Genres              []*Genre             `json:"genres"`
	Homepage            *string              `json:"homepage"`
	InProduction        *bool                `json:"in_production"`
	Languages           []string             `json:"languages"`
	LastAirDate         *string              `json:"last_air_date"`
	LastEpisodeToAir    *Episode             `json:"last_episode_to_air"`
	Name                *string              `json:"name"`
	NextEpisodeToAir    *string              `json:"next_episode_to_air"`
	Networks            []*Network           `json:"networks"`
	NumberOfEpisodes    *int32               `json:"number_of_episodes"`
	NumberOfSeasons     *int32               `json:"number_of_seasons"`
	OriginCountry       []string             `json:"origin_country"`
	OriginalLanguage    *string              `json:"original_language"`
	OriginalName        *string              `json:"original_name"`
	Overview            *string              `json:"overview"`
	Popularity          *float32             `json:"popularity"`
	PosterPath          *string              `json:"poster_path"`
	ProductionCompanies []*ProductionCompany `json:"production_companies"`
	ProductionCountries []*ProductionCountry `json:"production_countries"`
	Seasons             []*Season            `json:"seasons"`
	SpokenLanguages     []*SpokenLanguage    `json:"spoken_languages"`
	Status              *string              `json:"status"`
	Tagline             *string              `json:"tagline"`
	Type                *string              `json:"type"`
	VoteAverage         *float32             `json:"vote_average"`
	VoteCount           *int32               `json:"vote_count"`
}

func (d *Details) SetDefaults() {
	if d == nil {
		return
	}
	util.SetIfNil(&d.Adult, true)
	for _, creator := range d.Creators {
		creator.SetDefaults()
	}
	for _, genre := range d.Genres {
		genre.SetDefaults()
	}
	util.SetIfNil(&d.InProduction, true)
	d.LastEpisodeToAir.SetDefaults()
	for _, network := range d.Networks {
		network.SetDefaults()
	}
	util.SetIfNil(&d.NumberOfEpisodes, 0)
	util.SetIfNil(&d.NumberOfSeasons, 0)
	util.SetIfNil(&d.Popularity, 0.0)
	for _, company := range d.ProductionCompanies {
		company.SetDefaults()
	}
	for _, country := range d.ProductionCountries {
		country.SetDefaults()
	}
	for _, season := range d.Seasons {
		season.SetDefaults()
	}
	for _, lang := range d.SpokenLanguages {
		lang.SetDefaults()
	}
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
	fmt.Fprintf(&builder, " Creators: %v", d.Creators)
	fmt.Fprintf(&builder, " EpisodeRunTime: %v", d.EpisodeRunTime)
	fmt.Fprintf(&builder, " FirstAirDate: %s", util.FmtOrNil(d.FirstAirDate))
	fmt.Fprintf(&builder, " Genres: %v", d.Genres)
	fmt.Fprintf(&builder, " Homepage: %s", util.FmtOrNil(d.Homepage))
	fmt.Fprintf(&builder, " InProduction: %s", util.FmtOrNil(d.InProduction))
	fmt.Fprintf(&builder, " Languages: %v", d.Languages)
	fmt.Fprintf(&builder, " LastAirDate: %s", util.FmtOrNil(d.LastAirDate))
	fmt.Fprintf(&builder, " LastEpisodeToAir: %s", d.LastEpisodeToAir.String())
	fmt.Fprintf(&builder, " Name: %s", util.FmtOrNil(d.Name))
	fmt.Fprintf(&builder, " NextEpisodeToAir: %s", util.FmtOrNil(d.NextEpisodeToAir))
	fmt.Fprintf(&builder, " Networks: %v", d.Networks)
	fmt.Fprintf(&builder, " NumberOfEpisodes: %s", util.FmtOrNil(d.NumberOfEpisodes))
	fmt.Fprintf(&builder, " NumberOfSeasons: %s", util.FmtOrNil(d.NumberOfSeasons))
	fmt.Fprintf(&builder, " OriginCountry: %v", d.OriginCountry)
	fmt.Fprintf(&builder, " OriginalLanguage: %s", util.FmtOrNil(d.OriginalLanguage))
	fmt.Fprintf(&builder, " OriginalName: %s", util.FmtOrNil(d.OriginalName))
	fmt.Fprintf(&builder, " Overview: %s", util.FmtOrNil(d.Overview))
	fmt.Fprintf(&builder, " Popularity: %s", util.FmtOrNil(d.Popularity))
	fmt.Fprintf(&builder, " PosterPath: %s", util.FmtOrNil(d.PosterPath))
	fmt.Fprintf(&builder, " ProductionCompanies: %v", d.ProductionCompanies)
	fmt.Fprintf(&builder, " ProductionCountries: %v", d.ProductionCountries)
	fmt.Fprintf(&builder, " Seasons: %v", d.Seasons)
	fmt.Fprintf(&builder, " SpokenLanguages: %v", d.SpokenLanguages)
	fmt.Fprintf(&builder, " Status: %s", util.FmtOrNil(d.Status))
	fmt.Fprintf(&builder, " Tagline: %s", util.FmtOrNil(d.Tagline))
	fmt.Fprintf(&builder, " Type: %s", util.FmtOrNil(d.Type))
	fmt.Fprintf(&builder, " VoteAverage: %s", util.FmtOrNil(d.VoteAverage))
	fmt.Fprintf(&builder, " VoteCount: %s", util.FmtOrNil(d.VoteCount))
	builder.WriteString("}")
	return builder.String()
}

type Creator struct {
	ID          *int32  `json:"id"`
	CreditID    *string `json:"credit_id"`
	Name        *string `json:"name"`
	Gender      *int32  `json:"gender"`
	ProfilePath *string `json:"profile_path"`
}

func (c *Creator) SetDefaults() {
	if c == nil {
		return
	}
	util.SetIfNil(&c.ID, 0)
	util.SetIfNil(&c.Gender, 0)
}

func (c *Creator) String() string {
	if c == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ID: %s CreditID: %s Name: %s Gender: %s ProfilePath: %s}", util.FmtOrNil(c.ID), util.FmtOrNil(c.CreditID), util.FmtOrNil(c.Name), util.FmtOrNil(c.Gender), util.FmtOrNil(c.ProfilePath))
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

type Episode struct {
	ID             *int32   `json:"id"`
	Name           *string  `json:"name"`
	Overview       *string  `json:"overview"`
	VoteAverage    *float32 `json:"vote_average"`
	VoteCount      *int32   `json:"vote_count"`
	AirDate        *string  `json:"air_date"`
	EpisodeNumber  *int32   `json:"episode_number"`
	ProductionCode *string  `json:"production_code"`
	Runtime        *int32   `json:"runtime"`
	SeasonNumber   *int32   `json:"season_number"`
	SeriesID       *int32   `json:"show_id"`
	StillPath      *string  `json:"still_path"`
}

func (e *Episode) SetDefaults() {
	if e == nil {
		return
	}
	util.SetIfNil(&e.ID, 0)
	util.SetIfNil(&e.VoteAverage, 0.0)
	util.SetIfNil(&e.VoteCount, 0)
	util.SetIfNil(&e.EpisodeNumber, 0)
	util.SetIfNil(&e.Runtime, 0)
	util.SetIfNil(&e.SeasonNumber, 0)
	util.SetIfNil(&e.SeriesID, 0)
}

func (e *Episode) String() string {
	if e == nil {
		return "<nil>"
	}
	var builder strings.Builder
	builder.WriteString("{")
	fmt.Fprintf(&builder, "ID: %s", util.FmtOrNil(e.ID))
	fmt.Fprintf(&builder, " Name: %s", util.FmtOrNil(e.Name))
	fmt.Fprintf(&builder, " Overview: %s", util.FmtOrNil(e.Overview))
	fmt.Fprintf(&builder, " VoteAverage: %s", util.FmtOrNil(e.VoteAverage))
	fmt.Fprintf(&builder, " VoteCount: %s", util.FmtOrNil(e.VoteCount))
	fmt.Fprintf(&builder, " AirDate: %s", util.FmtOrNil(e.AirDate))
	fmt.Fprintf(&builder, " EpisodeNumber: %s", util.FmtOrNil(e.EpisodeNumber))
	fmt.Fprintf(&builder, " ProductionCode: %s", util.FmtOrNil(e.ProductionCode))
	fmt.Fprintf(&builder, " Runtime: %s", util.FmtOrNil(e.Runtime))
	fmt.Fprintf(&builder, " SeasonNumber: %s", util.FmtOrNil(e.SeasonNumber))
	fmt.Fprintf(&builder, " SeriesID: %s", util.FmtOrNil(e.SeriesID))
	fmt.Fprintf(&builder, " StillPath: %s", util.FmtOrNil(e.StillPath))
	builder.WriteString("}")
	return builder.String()
}

type Network struct {
	ID            *int32  `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          *string `json:"name"`
	OriginCountry *string `json:"origin_country"`
}

func (n *Network) SetDefaults() {
	if n == nil {
		return
	}
	util.SetIfNil(&n.ID, 0)
}

func (n *Network) String() string {
	if n == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ID: %s LogoPath: %s Name: %s OriginCountry: %s}", util.FmtOrNil(n.ID), util.FmtOrNil(n.LogoPath), util.FmtOrNil(n.Name), util.FmtOrNil(n.OriginCountry))
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

type Season struct {
	AirDate      *string  `json:"air_date"`
	EpisodeCount *int32   `json:"episode_count"`
	ID           *int32   `json:"id"`
	Name         *string  `json:"name"`
	Overview     *string  `json:"overview"`
	PosterPath   *string  `json:"poster_path"`
	SeasonNumber *int32   `json:"season_number"`
	VoteAverage  *float32 `json:"vote_average"`
}

func (s *Season) SetDefaults() {
	if s == nil {
		return
	}
	util.SetIfNil(&s.EpisodeCount, 0)
	util.SetIfNil(&s.ID, 0)
	util.SetIfNil(&s.SeasonNumber, 0)
	util.SetIfNil(&s.VoteAverage, 0.0)
}

func (s *Season) String() string {
	if s == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{AirDate: %s EpisodeCount: %s ID: %s Name: %s Overview: %s PosterPath: %s SeasonNumber: %s VoteAverage: %s}", util.FmtOrNil(s.AirDate), util.FmtOrNil(s.EpisodeCount), util.FmtOrNil(s.ID), util.FmtOrNil(s.Name), util.FmtOrNil(s.Overview), util.FmtOrNil(s.PosterPath), util.FmtOrNil(s.SeasonNumber), util.FmtOrNil(s.VoteAverage))
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
