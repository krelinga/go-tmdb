package tmdb

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
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

	// String returns a string representation of the client, useful for cache-busting in tests.
	String() string
}

type internalClient struct {
	apiKey     string
	createTime time.Time
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

func (c *internalClient) String() string {
	return fmt.Sprintf("live data  client (created at: %s)", c.createTime.Format(time.RFC3339))
}

func NewClient(apiKey string) Client {
	return &internalClient{
		apiKey:     apiKey,
		createTime: time.Now(),
	}
}

type savedReply struct {
	Reply []byte         `json:"reply"`
	Code  ClientHttpCode `json:"code"`
}

func savedKey(path string, params GetParams) string {
	path = strings.TrimLeft(path, "/")
	path = strings.TrimRight(path, "/")
	path = strings.ReplaceAll(path, "/", "_")

	type kv struct {
		k, v string
	}
	kvs := make([]kv, 0, len(params))
	for k, v := range params {
		if k == "append_to_response" {
			parts := strings.Split(v, ",")
			slices.Sort(parts)
			v = strings.Join(parts, ",")
		}
		kvs = append(kvs, kv{k: k, v: v})
	}
	slices.SortFunc(kvs, func(a, b kv) int {
		return strings.Compare(a.k, b.k)
	})
	values := make([]string, 0, len(kvs))
	for _, kv := range kvs {
		values = append(values, fmt.Sprintf("%s=%s", kv.k, kv.v))
	}

	prefix := "GET_"
	var suffix string
	if len(values) > 0 {
		suffix = "?" + strings.Join(values, "&")
	}
	return prefix + path + suffix
}

func NewReplayClient(upstream Client, dataDir string) (Client, error) {
	if upstream == nil {
		return newReadOnlyReplayClient(dataDir)
	}
	return newUpdatingReplayClient(upstream, dataDir)
}

func newReadOnlyReplayClient(dataDir string) (Client, error) {
	data, fp, err := readDataDir(dataDir)
	if err != nil {
		return nil, fmt.Errorf("reading data directory: %w", err)
	}

	return &readOnlyReplayClient{data: data, fp: fp}, nil
}

type readOnlyReplayClient struct {
	data map[string]*savedReply
	fp   string // fingerprint of the data directory contents
}

func (c *readOnlyReplayClient) Get(ctx context.Context, path string, params GetParams) ([]byte, ClientHttpCode, error) {
	key := savedKey(path, params)
	if reply, ok := c.data[key]; ok {
		return reply.Reply, reply.Code, nil
	}
	return nil, 0, fmt.Errorf("no saved reply for %s with params %v", path, params)
}

func (c *readOnlyReplayClient) String() string {
	return fmt.Sprintf("read-only replay client (data fingerprint: %s)", c.fp)
}

func newUpdatingReplayClient(upstream Client, dataDir string) (Client, error) {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("creating data directory: %w", err)
	}

	data, fp, err := readDataDir(dataDir)
	if err != nil {
		return nil, fmt.Errorf("reading data directory: %w", err)
	}
	return &updatingReplayClient{
		upstream: upstream,
		dataDir:  dataDir,
		data:     data,
		fp:       fp,
	}, nil
}

type updatingReplayClient struct {
	upstream Client
	dataDir  string
	mu       sync.Mutex
	data     map[string]*savedReply
	fp       string // fingerprint of the data directory contents
}

func (c *updatingReplayClient) Get(ctx context.Context, path string, params GetParams) ([]byte, ClientHttpCode, error) {
	key := savedKey(path, params)
	c.mu.Lock()
	defer c.mu.Unlock()

	if reply, ok := c.data[key]; ok {
		return reply.Reply, reply.Code, nil
	}

	data, code, err := c.upstream.Get(ctx, path, params)
	if err != nil {
		return nil, code, err
	}

	saved := &savedReply{
		Reply: data,
		Code:  code,
	}
	c.data[key] = saved

	filePath := fmt.Sprintf("%s/%s.json", c.dataDir, key)
	outData, err := json.Marshal(saved)
	if err != nil {
		return nil, code, fmt.Errorf("marshalling reply: %w", err)
	}
	if err := os.WriteFile(filePath, outData, 0644); err != nil {
		return nil, code, fmt.Errorf("writing reply to file: %w", err)
	}

	return data, code, nil
}

func (c *updatingReplayClient) String() string {
	var upstramPart string
	if c.upstream != nil {
		upstramPart = fmt.Sprintf("upstream: %s", c.upstream)
	} else {
		upstramPart = "no upstream"
	}
	return fmt.Sprintf("updating replay client (data fingerprint: %s) with %s", c.fp, upstramPart)
}

func readDataDir(dataDir string) (map[string]*savedReply, string, error) {
	files, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, "", fmt.Errorf("reading data directory: %w", err)
	}

	data := make(map[string]*savedReply)
	fp := sha256.New()
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			fp.Write([]byte(file.Name()))
			filePath := fmt.Sprintf("%s/%s", dataDir, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				return nil, "", fmt.Errorf("reading file %s: %w", filePath, err)
			}
			fp.Write(content)
			var saved savedReply
			if err := json.Unmarshal(content, &saved); err != nil {
				return nil, "", fmt.Errorf("unmarshalling file %s: %w", filePath, err)
			}
			key := strings.TrimSuffix(file.Name(), ".json")
			data[key] = &saved
		}
	}
	return data, fmt.Sprintf("%x", fp.Sum(nil)), nil
}
