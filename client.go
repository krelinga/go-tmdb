package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strings"
	"sync"
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

type savedReply struct {
	Reply []byte         `json:"reply"`
	Code  ClientHttpCode `json:"code"`
}

func savedKey(path string, params GetParams) string {
	values := url.Values{}

	for k, v := range params {
		if k == "append_to_response" {
			parts := strings.Split(v, ",")
			slices.Sort(parts)
			v = strings.Join(parts, ",")
		}
		values[k] = []string{v}
	}

	// Create a unique key for the request based on the path and parameters.
	// This is a simple implementation; you might want to use a more robust hashing function.
	return fmt.Sprintf("%s?%s", path, url.Values(values).Encode())
}

type MemoClientReadMode int

const (
	MemoClientReadModeNone MemoClientReadMode = iota
	MemoClientReadModeRead
)

type MemoClientWriteMode int

const (
	MemoClientWriteModeNone MemoClientWriteMode = iota
	MemoClientWriteModeWrite
	MemoClientWriteModeAppend
	MemoClientWriteModeReplace
)

func invalidMemoClientWriteMode(mode MemoClientWriteMode) error {
	return fmt.Errorf("invalid client memo write mode: %d", mode)
}

func invalidMemoClientReadMode(mode MemoClientReadMode) error {
	return fmt.Errorf("invalid client memo read mode: %d", mode)
}

func NewMemoClient(upstream Client, readMode MemoClientReadMode, writeMode MemoClientWriteMode, dataDir string) (Client, error) {
	switch readMode {
	case MemoClientReadModeNone, MemoClientReadModeRead: // valid modes
	default:
		return nil, invalidMemoClientReadMode(readMode)
	}

	switch writeMode {
	case MemoClientWriteModeNone: // nothing to do.
	case MemoClientWriteModeWrite, MemoClientWriteModeAppend, MemoClientWriteModeReplace:
		if err := os.MkdirAll(dataDir, 0755); err != nil {
			return nil, fmt.Errorf("creating data directory: %w", err)
		}
	default:
		return nil, invalidMemoClientWriteMode(writeMode)
	}

	if writeMode == MemoClientWriteModeAppend && readMode == MemoClientReadModeNone {
		return nil, fmt.Errorf("append mode is not allowed with read mode none")
	}

	return &memoClient{
		upstream:  upstream,
		readMode:  readMode,
		writeMode: writeMode,
		dataDir:   dataDir,
	}, nil
}

type memoClient struct {
	upstream  Client
	readMode  MemoClientReadMode
	writeMode MemoClientWriteMode
	dataDir   string
	mu        sync.Mutex
}

func (c *memoClient) Get(ctx context.Context, path string, params GetParams) ([]byte, ClientHttpCode, error) {
	if c.writeMode == MemoClientWriteModeNone && c.readMode == MemoClientReadModeNone {
		// If both read and write modes are none, just pass through to the upstream client & avoid grabbing the mutex.
		return c.upstream.Get(ctx, path, params)
	}

	// write mode | read mode yes        | read mode no
	// -----------|----------------------|-------------
	// none       | read		         | passthrough
	// write      | read + (over)write   | (over)write
	// append     | read + create if new | INVALID
	// replace    | read + (over)write   | (over)write
	key := savedKey(path, params)
	filePath := fmt.Sprintf("%s/%s.json", c.dataDir, url.QueryEscape(key))

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.readMode == MemoClientReadModeRead {
		if data, err := os.ReadFile(filePath); err == nil {
			var saved savedReply
			if err := json.Unmarshal(data, &saved); err != nil {
				return nil, 0, fmt.Errorf("unmarshalling saved reply: %w", err)
			}
			return saved.Reply, saved.Code, nil
		} else if !os.IsNotExist(err) {
			return nil, 0, fmt.Errorf("reading saved reply file: %w", err)
		}
	}

	data, code, err := c.upstream.Get(ctx, path, params)
	if err != nil {
		return nil, code, err
	}

	if c.writeMode != MemoClientWriteModeNone {
		saved := savedReply{
			Reply: data,
			Code:  code,
		}
		outData, err := json.Marshal(saved)
		if err != nil {
			return nil, code, fmt.Errorf("marshalling reply: %w", err)
		}
		if err := os.WriteFile(filePath, outData, 0644); err != nil {
			return nil, code, fmt.Errorf("writing reply to file: %w", err)
		}
	}

	return data, code, nil
}
