package tmdb_test

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/krelinga/go-tmdb"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

// Replay modes that control the behavior of the test client.
const (
	// Only read data from the replay file.  Any requests not found in the file will fail.
	replayRead string = "read"
	// Delete the existing replay file and create a new one by reading from the live API.
	replayReplace string = "replace"
	// Append to the existing replay file.  Any requests not found in the file will be added by reading from the live API.
	replayAppend string = "append"
	// Ignore the replay file and read from the live API.
	replayDirect string = "direct"
)

var replayFlag = flag.String("replay", "read", "Replay mode: read, replace, append, direct")

func getReplayMode(t *testing.T) string {
	switch *replayFlag {
	case replayRead, replayReplace, replayAppend, replayDirect:
		return *replayFlag
	default:
		t.Fatalf("Invalid --replay mode: %s. Valid modes are: read, replace, append, direct", *replayFlag)
		return replayRead // Fallback to a default mode
	}
}

type testClientOptions struct {
	useApiKey             bool
	useApiReadAccessToken bool
}

func (tco testClientOptions) newClient(t *testing.T) tmdb.Client {
	const (
		replayApiKey             string = "fake-api-key"
		replayApiReadAccessToken string = "fake-api-read-access-token"
	)
	replayMode := getReplayMode(t)
	clientOptions := tmdb.ClientOptions{}
	switch replayMode {
	case replayAppend, replayReplace, replayDirect:
		if tco.useApiKey {
			clientOptions.APIKey = os.Getenv("TMDB_API_KEY")
			if clientOptions.APIKey == "" {
				t.Fatal("TMDB_API_KEY must be set for replace, append, or direct modes when using API key")
			}
		}
		if tco.useApiReadAccessToken {
			clientOptions.APIReadAccessToken = os.Getenv("TMDB_API_READ_ACCESS_TOKEN")
			if clientOptions.APIReadAccessToken == "" {
				t.Fatal("TMDB_API_READ_ACCESS_TOKEN must be set for replace, append, or direct modes when using read access token")
			}
		}
	case replayRead:
		if tco.useApiKey {
			clientOptions.APIKey = replayApiKey
		}
		if tco.useApiReadAccessToken {
			clientOptions.APIReadAccessToken = replayApiReadAccessToken
		}
	}

	removeAuth := func(i *cassette.Interaction) error {
		if _, hasAuthHeader := i.Request.Headers["Authorization"]; hasAuthHeader {
			i.Request.Headers["Authorization"] = []string{fmt.Sprintf("Bearer %s", replayApiReadAccessToken)}
		}

		parsedUrl, err := url.Parse(i.Request.URL)
		if err != nil {
			return fmt.Errorf("failed to parse URL: %w", err)
		}
		parsedQuery := parsedUrl.Query()
		if _, hasApiKey := parsedQuery["api_key"]; hasApiKey {
			parsedQuery.Set("api_key", replayApiKey)
			parsedUrl.RawQuery = parsedQuery.Encode()
			i.Request.URL = parsedUrl.String()
		}

		if _, hasForm := i.Request.Form["api_key"]; hasForm {
			i.Request.Form.Set("api_key", replayApiKey)
		}

		return nil
	}

	authlessMatcher := func(req *http.Request, target cassette.Request) bool {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		req = req.Clone(ctx)

		parsedQuery := req.URL.Query()
		if _, hasApiKey := parsedQuery["api_key"]; hasApiKey {
			parsedQuery.Set("api_key", replayApiKey)
			req.URL.RawQuery = parsedQuery.Encode()
		}

		if _, hasAuthHeader := req.Header["Authorization"]; hasAuthHeader {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", replayApiReadAccessToken))
		}

		if _, hasForm := req.Form["api_key"]; hasForm {
			req.Form.Set("api_key", replayApiKey)
		}

		return cassette.DefaultMatcher(req, target)
	}

	var recorderMode recorder.Mode
	switch replayMode {
	case replayRead:
		recorderMode = recorder.ModeReplayOnly
	case replayReplace:
		recorderMode = recorder.ModeRecordOnly
	case replayAppend:
		recorderMode = recorder.ModeReplayWithNewEpisodes
	case replayDirect:
		recorderMode = recorder.ModePassthrough
	}

	if replayMode == replayReplace {
		if err := os.RemoveAll("testdata"); err != nil {
			t.Fatalf("Failed to remove testdata directory: %v", err)
		}
	}

	cassettePath := filepath.Join("testdata", strings.ReplaceAll(t.Name(), "/", "_"))
	r, err := recorder.New(cassettePath,
		recorder.WithMode(recorderMode),
		recorder.WithReplayableInteractions(true),
		recorder.WithSkipRequestLatency(true),
		recorder.WithHook(removeAuth, recorder.AfterCaptureHook),
		recorder.WithMatcher(authlessMatcher))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := r.Stop(); err != nil {
			t.Errorf("Failed to stop recorder: %v", err)
		}
	})

	clientOptions.HttpClient = r.GetDefaultClient()

	return clientOptions.NewClient()
}
