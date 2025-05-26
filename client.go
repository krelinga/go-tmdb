package tmdb

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Parameters for GET requests to the TMDB API.  Used with the Client interface.
type GetParams map[string]string

type ClientHttpCode int

type Client interface {
	// Fetches data from the TMDB API.
	// The path should be the endpoint you want to access, e.g., "/movie/popular".
	// The params are optional query parameters to include in the request.
	// The out parameter is where the response will be unmarshalled into, usually something from the raw package.
	// The method returns an error if the request fails or if unmarshalling fails.
	//
	// Note that most users should not call this method directly, but instead pass a Client into
	// other methods in this package.
	//
	// Implementations of this method should handle parallel calls to Get().
	Get(ctx context.Context, path string, params GetParams) ([]byte, ClientHttpCode, error)
}

type internalClient struct {
	apiKey string
}

func (c *internalClient) Get(ctx context.Context, path string, params GetParams) ([]byte, ClientHttpCode, error) {
	baseURL := "https://api.themoviedb.org/3"
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(fmt.Sprintf("bad base URL: %s", baseURL))
	}
	u.Path = strings.TrimRight(u.Path, "/") + "/" + strings.TrimLeft(path, "/")

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return out, ClientHttpCode(resp.StatusCode), nil
}

func NewClient(apiKey string) Client {
	return &internalClient{
		apiKey: apiKey,
	}
}
