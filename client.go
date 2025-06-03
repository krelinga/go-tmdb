package tmdb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type NewClientOptions struct {
	HttpClient  *http.Client
	ApiKey      string
	BearerToken string
}

type Client struct {
	options *NewClientOptions

	// Lazy initialization of shared state.
	getConfiguration func() (*raw.Configuration, error)
}

func NewClient(options *NewClientOptions) *Client {
	if options == nil {
		options = &NewClientOptions{}
	}

	c := &Client{
		options: options,
	}

	c.getConfiguration = sync.OnceValues(func() (*raw.Configuration, error) {
		configuration := &raw.Configuration{}
		err := get(context.Background(), c, "configuration", nil, configuration)
		if err != nil {
			return nil, fmt.Errorf("getting configuration: %w", err)
		}
		return configuration, nil
	})

	return c
}

func (c *Client) checkOk() {
	if c.options == nil {
		panic("TMDB client is not properly initialized.  Use NewClient to create a new client.")
	}
}

func (c *Client) httpClient() *http.Client {
	c.checkOk()
	if c.options.HttpClient == nil {
		return http.DefaultClient
	}
	return c.options.HttpClient
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

func get(ctx context.Context, c *Client, endpoint string, params url.Values, out raw.Raw) error {
	c.checkOk()
	if c.options.ApiKey != "" {
		params.Add("api_key", c.options.ApiKey)
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
	if c.options.BearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.options.BearerToken)
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}
	resp, err := c.httpClient().Do(req)
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
