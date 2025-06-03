package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type ClientOptions struct {
	HttpClient *http.Client
	ApiKey      string
	BearerToken string
}

type Client struct {
	options *ClientOptions

	// Lazy initialization of shared state.
	getConfiguration func() (*raw.Configuration, error)
}

func NewClient(options *ClientOptions) *Client {
	if options == nil {
		options = &ClientOptions{}
	}

	c := &Client{
		options: options,
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
		reply, err := c.httpClient().Do(req)
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

func (c *Client) prepRequest(req *http.Request) {
	c.checkOk()
	if c.options.BearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.options.BearerToken)
	}
}

func (c *Client) prepUrl(theUrl *url.URL) {
	c.checkOk()
	theUrl.Scheme = "https"
	theUrl.Host = "api.themoviedb.org"
	if c.options.ApiKey != "" {
		q := theUrl.Query()
		q.Add("api_key", c.options.ApiKey)
		theUrl.RawQuery = q.Encode()
	}
}
