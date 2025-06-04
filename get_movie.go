package tmdb

import (
	"context"
	"fmt"
	"iter"
	"net/url"
	"slices"
	"time"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type GetMovieOptions struct {
	Language LanguageId
	Columns  []MovieDataCol
}

// In addition to what is listed in MovieData, the following data is also available:
// - Companies()
//   - Logo()
//   - Name()
//   - OriginCountry()
func GetMovie(ctx context.Context, c *Client, id MovieId, options *GetMovieOptions) (Movie, error) {
	if options == nil {
		options = &GetMovieOptions{}
	}
	parts, err := getMovie(ctx, c, id, options.Language, options.Columns...)
	if err != nil {
		return nil, fmt.Errorf("getting movie %d: %w", id, err)
	}

	return &movie{
		client:    c,
		id:        id,
		language:  options.Language,
		MovieData: parts,
	}, nil
}

func getMovie(ctx context.Context, c *Client, id MovieId, language LanguageId, columns ...MovieDataCol) (*getMovieData, error) {
	v := url.Values{}
	if language != "" {
		v.Set("language", string(language))
	}
	if len(columns) > 0 {
		v.Set("append_to_response", appendToResponse(columns))
	}

	raw := &raw.GetMovie{}
	if err := get(ctx, c, fmt.Sprintf("movie/%d", id), v, raw); err != nil {
		return nil, fmt.Errorf("getting movie %d: %w", id, err)
	}

	out := &getMovieData{}
	out.init(raw, c)
	return out, nil
}

type getMovieData struct {
	client         *Client
	rawDetails     *raw.GetMovieDetails
	keywords       []Keyword
	rawExternalIds *raw.GetMovieExternalIds
	companies      []Company
}

func (p *getMovieData) init(raw *raw.GetMovie, client *Client) {
	p.client = client

	p.rawDetails = raw.GetMovieDetails

	if raw.Keywords != nil {
		p.keywords = make([]Keyword, len(raw.Keywords.Keywords))
		for i, kw := range raw.Keywords.Keywords {
			p.keywords[i] = keyword{
				id:   KeywordId(kw.Id),
				name: kw.Name,
			}
		}
	}

	if raw.ExternalIds != nil {
		p.rawExternalIds = raw.ExternalIds
	}

	p.companies = make([]Company, 0, len(raw.ProductionCompanies))
	for _, rawCompany := range raw.ProductionCompanies {
		p.companies = append(p.companies, &company{
			id: CompanyId(rawCompany.Id),
			CompanyData: &getMovieCompanyData{
				client:      client,
				raw:         rawCompany,
				CompanyData: companyNoData{},
			},
		})
	}
}

func (p *getMovieData) upgrade(in *getMovieData) MovieData {
	if in == nil {
		return p
	}
	if p.rawDetails == nil {
		p.rawDetails = in.rawDetails
	}
	if p.keywords == nil {
		p.keywords = in.keywords
	}
	if p.rawExternalIds == nil {
		p.rawExternalIds = in.rawExternalIds
	}
	return p
}

func (p *getMovieData) Adult() bool {
	return *p.rawDetails.Adult
}

func (p *getMovieData) Backdrop() Image {
	return image{
		raw:    p.rawDetails.BackdropPath,
		client: p.client,
	}
}

func (p *getMovieData) BelongsToCollection() string {
	return p.rawDetails.BelongsToCollection
}

func (p *getMovieData) Budget() int {
	return p.rawDetails.Budget
}

func (p *getMovieData) GenreIds() iter.Seq[GenreId] {
	return func(yield func(GenreId) bool) {
		for _, genre := range p.rawDetails.Genres {
			if !yield(GenreId(genre.Id)) {
				return
			}
		}
	}
}

func (p *getMovieData) Genres() iter.Seq[Genre] {
	return func(yield func(Genre) bool) {
		for _, g := range p.rawDetails.Genres {
			if !yield(genre{
				id:   GenreId(g.Id),
				name: g.Name,
			}) {
				return
			}
		}
	}
}

func (p *getMovieData) Homepage() string {
	return p.rawDetails.Homepage
}

func (p *getMovieData) ImdbId() ImdbMovieId {
	return ImdbMovieId(p.rawDetails.ImdbId)
}

func (p *getMovieData) OriginalLanguage() LanguageId {
	return LanguageId(p.rawDetails.OriginalLanguage)
}

func (p *getMovieData) OriginalTitle() string {
	return p.rawDetails.OriginalTitle
}

func (p *getMovieData) Overview() string {
	return p.rawDetails.Overview
}

func (p *getMovieData) Popularity() float64 {
	return p.rawDetails.Popularity
}

func (p *getMovieData) Poster() Image {
	return image{
		raw:    p.rawDetails.PosterPath,
		client: p.client,
	}
}

func (p *getMovieData) Companies() iter.Seq[Company] {
	return func(yield func(Company) bool) {
		for _, c := range p.companies {
			if !yield(c) {
				return
			}
		}
	}
}

func (p *getMovieData) Countries() iter.Seq[Country] {
	return func(yield func(Country) bool) {
		for _, rawCountry := range p.rawDetails.ProductionCountries {
			if !yield(country{
				id:   CountryId(rawCountry.Iso3166_1),
				name: rawCountry.Name,
			}) {
				return
			}
		}
	}
}

func (p *getMovieData) ReleaseDate() Date {
	return newDateYYYYMMDD(p.rawDetails.ReleaseDate)
}

func (p *getMovieData) Revenue() int {
	return p.rawDetails.Revenue
}

func (p *getMovieData) Runtime() time.Duration {
	return time.Duration(p.rawDetails.Runtime) * time.Minute
}

func (p *getMovieData) Cast() iter.Seq[Cast] {
	return nil // TODO: implement
}

func (p *getMovieData) Crew() iter.Seq[Crew] {
	return nil // TODO: implement
}

func (p *getMovieData) WikidataId() WikidataMovieId {
	if p.rawExternalIds == nil {
		panic(ErrMovieNoDataWikidataId)
	}
	return WikidataMovieId(p.rawExternalIds.WikidataId)
}

func (p *getMovieData) Keywords() iter.Seq[Keyword] {
	if p.keywords == nil {
		panic(ErrMovieNoDataKeywords)
	}
	return slices.Values(p.keywords)
}

type getMovieCompanyData struct {
	client *Client
	raw    *raw.GetMovieProductionCompany
	CompanyData
}

func (c *getMovieCompanyData) Logo() Image {
	return image{
		raw:    c.raw.LogoPath,
		client: c.client,
	}
}

func (c *getMovieCompanyData) Name() string {
	return c.raw.Name
}

func (c *getMovieCompanyData) OriginCountry() CountryId {
	return CountryId(c.raw.OriginCountry)
}
