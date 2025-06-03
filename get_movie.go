package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"net/http"
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
	theUrl := &url.URL{
		Path:     fmt.Sprintf("/3/movie/%d", id),
		RawQuery: v.Encode(),
	}
	c.prepUrl(theUrl)
	req, err := http.NewRequestWithContext(ctx, "GET", theUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	c.prepRequest(req)
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	reply, err := c.httpClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer reply.Body.Close()
	if reply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API returned status code %d", reply.StatusCode)
	}
	decoder := json.NewDecoder(reply.Body)
	raw := &raw.GetMovie{}
	if err := decoder.Decode(&raw); err != nil {
		return nil, fmt.Errorf("decoding movie %d: %w", id, err)
	}
	out := &getMovieData{}
	out.init(raw)
	return out, nil
}

type getMovieData struct {
	rawDetails     *raw.GetMovieDetails
	keywords       []Keyword
	rawExternalIds *raw.GetMovieExternalIds
}

func (p *getMovieData) init(raw *raw.GetMovie) {
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
