package tmdb

import (
	"context"
	"fmt"
	"iter"
	"net/url"
	"slices"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type GetMovieOptions struct {
	Language Language
	Columns  []MovieDataCol
}

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

func getMovie(ctx context.Context, c *Client, id MovieId, language Language, columns ...MovieDataCol) (*getMovieData, error) {
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

func (p *getMovieData) Budget() int {
	return p.rawDetails.Budget
}

func (p *getMovieData) Backdrop() Image {
	return image{
		raw:    p.rawDetails.BackdropPath,
		client: p.client,
	}
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
