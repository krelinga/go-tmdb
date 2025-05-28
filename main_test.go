package tmdb_test

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/krelinga/go-tmdb"
	"github.com/stretchr/testify/assert"
)

var replayFlag = flag.String("replay", "read", "Which mode to run the replay client in.  Options are: 'direct', 'read', 'write', or 'replace'")

var clientOnce = sync.OnceValues(func() (tmdb.Client, error) {
	var direct, write, replace bool
	switch *replayFlag {
	case "direct":
		direct = true
	case "read":
		// Nothing to do.
	case "write":
		write = true
	case "replace":
		replace = true
	default:
		return nil, errors.New("Invalid replay mode. Options are: 'direct', 'read', 'write', or 'replace'")
	}

	const dataDir = "testdata"

	if replace {
		// Remove all the existing files.
		files, err := os.ReadDir(dataDir)
		if err != nil && !os.IsNotExist(err) {
			return nil, fmt.Errorf("Failed to read test data directory: %w", err)
		}
		for _, file := range files {
			if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
				continue
			}
			err := os.Remove(dataDir + "/" + file.Name())
			if err != nil {
				return nil, fmt.Errorf("Failed to remove file %s: %w", file.Name(), err)
			}
		}
	}

	var upstream tmdb.Client
	if direct || write || replace {
		key, ok := os.LookupEnv("TMDB_API_KEY")
		if !ok {
			return nil, errors.New("environment variable TMDB_API_KEY not set")
		}
		upstream = tmdb.NewClient(key)
	}
	var client tmdb.Client
	if direct {
		client = upstream
	} else {
		var err error
		client, err = tmdb.NewReplayClient(upstream, dataDir)
		if err != nil {
			return nil, fmt.Errorf("Failed to create replay client: %w", err)
		}
	}
	return client, nil
})

func getClient(t *testing.T) tmdb.Client {
	t.Helper()
	client, err := clientOnce()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	} else {
		t.Logf("Using client: %s", client)
	}
	return client
}

type imageSize interface {
	tmdb.BackdropSize | tmdb.PosterSize | tmdb.ProfileSize | tmdb.LogoSize | tmdb.StillSize
}

type image[imageSizeType imageSize] interface {
	GetSecureUrl(config *tmdb.Configuration, size imageSizeType) (string, bool)
	GetUrl(config *tmdb.Configuration, size imageSizeType) (string, bool)
	~string
}

func checkImage[imageSizeType imageSize, imageType image[imageSizeType]](t *testing.T, img imageType, typeName string, config *tmdb.Configuration, size imageSizeType) bool {
	t.Helper()

	secureUrl, ok := img.GetSecureUrl(config, size)
	if !assert.Truef(t, ok, "%s %q GetSecureUrl() should support size %q", typeName, img, size) {
		return false
	}
	if !assert.Truef(t, strings.HasSuffix(secureUrl, string(img)), "%s %q GetSecureUrl() should end with %q", typeName, secureUrl, img) {
		return false
	}
	if !assert.Truef(t, strings.HasPrefix(secureUrl, config.Images.SecureBaseUrl), "%s %q GetSecureUrl() should start with %q", typeName, secureUrl, config.Images.SecureBaseUrl) {
		return false
	}

	insecureUrl, ok := img.GetUrl(config, size)
	if !assert.Truef(t, ok, "%s %q GetUrl() should support size %q", typeName, img, size) {
		return false
	}
	if !assert.Truef(t, strings.HasSuffix(secureUrl, string(img)), "%s %q GetUrl() should end with %q", typeName, insecureUrl, img) {
		return false
	}
	if !assert.Truef(t, strings.HasPrefix(secureUrl, config.Images.SecureBaseUrl), "%s %q GetUrl() should start with %q", typeName, insecureUrl, config.Images.SecureBaseUrl) {
		return false
	}

	return true
}

func checkBackdropImage(t *testing.T, backdropImage tmdb.BackdropImage, config *tmdb.Configuration) bool {
	size := config.Images.BackdropSizes[0]
	return checkImage(t, backdropImage, "BackdropImage", config, size)
}

func checkPosterImage(t *testing.T, posterImage tmdb.PosterImage, config *tmdb.Configuration) bool {
	size := config.Images.PosterSizes[0]
	return checkImage(t, posterImage, "PosterImage", config, size)
}

func checkProfileImage(t *testing.T, profileImage tmdb.ProfileImage, config *tmdb.Configuration) bool {
	size := config.Images.ProfileSizes[0]
	return checkImage(t, profileImage, "ProfileImage", config, size)
}

func checkLogoImage(t *testing.T, logoImage tmdb.LogoImage, config *tmdb.Configuration) bool {
	size := config.Images.LogoSizes[0]
	return checkImage(t, logoImage, "LogoImage", config, size)
}

func checkDate(t *testing.T, expectedYear, expectedMonth, expectedDay int, actual tmdb.Date) bool {
	t.Helper()
	asTime, err := actual.GetTime()
	if !assert.NoErrorf(t, err, "GetTime() should not return an error for date %q", actual) {
		return false
	}
	assert.Equalf(t, expectedYear, asTime.Year(), "Expected year %d, got %d for date %q", expectedYear, asTime.Year(), actual)
	assert.Equalf(t, expectedMonth, int(asTime.Month()), "Expected month %d, got %d for date %q", expectedMonth, asTime.Month(), actual)
	assert.Equalf(t, expectedDay, asTime.Day(), "Expected day %d, got %d for date %q", expectedDay, asTime.Day(), actual)
	return true
}
