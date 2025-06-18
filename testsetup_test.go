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

	"github.com/krelinga/go-errgrp"
	"github.com/krelinga/go-tmdb"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

var replayFlag = flag.String("replay", "read", "Which mode to run the replay client in.  Options are: 'direct', 'read', 'write', or 'replace'")

func getClient(t *testing.T) *tmdb.Client {
	t.Helper()
	var mode recorder.Mode
	var envString, envToken string
	switch *replayFlag {
	case "direct":
		mode = recorder.ModePassthrough
		envString = os.Getenv("TMDB_API_KEY")
		envToken = os.Getenv("TMDB_BEARER_TOKEN")
	case "read":
		mode = recorder.ModeReplayOnly
		envString = "will_be_ignored"
		envToken = "will_be_ignored"
	case "write":
		mode = recorder.ModeReplayWithNewEpisodes
		envString = os.Getenv("TMDB_API_KEY")
		envToken = os.Getenv("TMDB_BEARER_TOKEN")
	case "replace":
		mode = recorder.ModeRecordOnly
		envString = os.Getenv("TMDB_API_KEY")
		envToken = os.Getenv("TMDB_BEARER_TOKEN")
	default:
		t.Fatal("Invalid replay mode. Options are: 'direct', 'read', 'write', or 'replace'")
		return nil
	}

	if envString == "" && envToken == "" && mode != recorder.ModeReplayOnly {
		t.Fatal("TMDB_API_KEY or TMDB_BEARER_TOKEN environment variable must be set for non-replay modes")
		return nil
	}

	redactSecrets := func(i *cassette.Interaction) error {
		// TOOD: clean this up.
		redactUrl := func(kind string, in *string, parseMethod func(string) (*url.URL, error)) error {
			parsed, err := parseMethod(*in)
			if err != nil {
				return fmt.Errorf("redactSecrets: Failed to parse %s %s: %w", kind, *in, err)
			}
			query := parsed.Query()
			if query.Get("api_key") != "" {
				query.Set("api_key", "REDACTED_TMDB_API_KEY")
				parsed.RawQuery = query.Encode()
				*in = parsed.String()
			}
			return nil
		}

		errs := errgrp.New()
		errs.Add(redactUrl("URL", &i.Request.URL, url.Parse))

		if _, ok := i.Request.Headers["Authorization"]; ok {
			i.Request.Headers.Set("Authorization", "Bearer REDACTED_TMDB_BEARER_TOKEN")
		}
		return errs.Join()
	}
	defaultMatcher := cassette.NewDefaultMatcher()
	ignoreSecretsMatcher := func(req *http.Request, saved cassette.Request) bool {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		req = req.Clone(ctx)
		q := req.URL.Query()
		if q.Get("api_key") != "" {
			q.Set("api_key", "REDACTED_TMDB_API_KEY")
			req.URL.RawQuery = q.Encode()
		}
		if req.Header.Get("Authorization") != "" {
			req.Header.Set("Authorization", "Bearer REDACTED_TMDB_BEARER_TOKEN")
		}
		return defaultMatcher(req, saved)
	}

	dataPath := filepath.Join("testdata", strings.ReplaceAll(t.Name(), "/", "_"))
	r, err := recorder.New(dataPath,
		recorder.WithMode(mode),
		recorder.WithReplayableInteractions(true),
		recorder.WithSkipRequestLatency(true),
		recorder.WithHook(redactSecrets, recorder.BeforeSaveHook),
		recorder.WithMatcher(ignoreSecretsMatcher),
	)
	if err != nil {
		t.Fatalf("Failed to create recorder: %v", err)
		return nil
	}
	t.Cleanup(func() {
		if err := r.Stop(); err != nil {
			t.Errorf("Failed to stop recorder: %v", err)
		}
	})
	client, err := tmdb.NewClient(r.GetDefaultClient(),
		tmdb.WithApiKey(envString),
		tmdb.WithBearerToken(envToken),
	)
	if err != nil {
		t.Fatalf("Failed to create TMDB client: %v", err)
		return nil
	}
	return client
}
