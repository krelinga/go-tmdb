package tmdb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type Client struct {
	httpClient    *http.Client
	globalOpts    allOptions
	configuration *raw.Configuration
}

func NewClient(httpClient *http.Client, options ...Option) (*Client, error) {
	c := &Client{httpClient: httpClient}
	c.globalOpts.apply(options)

	c.configuration = &raw.Configuration{}
	if err := get(context.Background(), c, "configuration", nil, c.globalOpts, c.configuration); err != nil {
		return nil, fmt.Errorf("initializing TMDB client: %w", err)
	}

	return c, nil
}

func (c *Client) checkOk() {
	if c.httpClient == nil {
		panic("TMDB client is not properly initialized.  Use NewClient to create a new client.")
	}
}

var ErrApiHttpNotOk = errors.New("TMDB API returned non-OK HTTP status code")

func checkResponseCode(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	status := &raw.ApiStatus{}
	if err := json.NewDecoder(resp.Body).Decode(status); err != nil {
		return fmt.Errorf("decoding error response: %w", err)
	}
	return fmt.Errorf("%w: Code %d, Message: %s", ErrApiHttpNotOk, status.Code, status.Message)
}

func get(ctx context.Context, c *Client, endpoint string, params url.Values, callOpts allOptions, out raw.Raw) error {
	c.checkOk()
	if params == nil {
		params = url.Values{}
	}
	if callOpts.ApiKey != nil {
		params.Add("api_key", *callOpts.ApiKey)
	}
	theUrl := &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     fmt.Sprintf("/3/%s", strings.TrimLeft(endpoint, "/")),
		RawQuery: params.Encode(),
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, theUrl.String(), nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	if callOpts.BearerToken != nil {
		req.Header.Add("Authorization", "Bearer "+*callOpts.BearerToken)
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()
	if err := checkResponseCode(resp); err != nil {
		return err
	}
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}
	out.SetDefaults()
	return nil
}
