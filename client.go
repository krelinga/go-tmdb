package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type Client struct {
	// used to determine if the client is properly initialized.
	ok bool

	// Client to use for making requests
	httpClient *http.Client

	// Authentication details
	apiKey      string
	bearerToken string

	// Lazy initialization of shared state.
	getConfiguration func() (*raw.Configuration, error)
}

type ClientOption func(*Client)

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

func WithBearerToken(token string) ClientOption {
	return func(c *Client) {
		c.bearerToken = token
	}
}

func NewClient(options ...ClientOption) *Client {
	c := &Client{
		ok: true,
	}

	for _, opt := range options {
		opt(c)
	}

	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	c.getConfiguration = sync.OnceValues(func() (*raw.Configuration, error) {
		c.checkOk()
		theUrl := &url.URL{
			Path: "/3/configuration",
		}
		c.prepUrl(theUrl)
		req, err := http.NewRequest("GET", theUrl.String(), nil)
		if err != nil {
			return nil, err
		}
		c.prepRequest(req)
		reply, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer reply.Body.Close()
		if reply.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("TMDB API returned status code %d", reply.StatusCode)
		}
		configuration := &raw.Configuration{}
		decoder := json.NewDecoder(reply.Body)
		if err := decoder.Decode(configuration); err != nil {
			return nil, fmt.Errorf("decoding configuration: %w", err)
		}
		return configuration, nil
	})

	return c
}

func (c *Client) checkOk() {
	if !c.ok {
		panic("TMDB client is not properly initialized.  Use NewClient to create a new client.")
	}
}

func (c *Client) prepRequest(req *http.Request) {
	c.checkOk()
	if c.bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.bearerToken)
	}
}

func (c *Client) prepUrl(theUrl *url.URL) {
	c.checkOk()
	theUrl.Scheme = "https"
	theUrl.Host = "api.themoviedb.org"
	if c.apiKey != "" {
		q := theUrl.Query()
		q.Add("api_key", c.apiKey)
		theUrl.RawQuery = q.Encode()
	}
}
