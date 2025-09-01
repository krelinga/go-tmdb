package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/krelinga/go-tmdb"
)

type command func(tmdb.Client, []string) error

var commands = map[string]command{
	"configdetails":   configDetails,
	"configcountries": configCountries,
	"configjobs":      configJobs,
	"configlanguages": configLanguages,
	"movie":           movie,
}

func toInt32(in string) (int32, error) {
	val, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

func toRequestOptions(args []string) ([]tmdb.RequestOption, error) {
	var options []tmdb.RequestOption
	appends := []string{}
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid argument format: %q", arg)
		}
		key := parts[0]
		value := parts[1]
		switch key {
		case "--append":
			appends = append(appends, value)
		case "--param":
			paramParts := strings.SplitN(value, "=", 2)
			if len(paramParts) != 2 {
				return nil, fmt.Errorf("invalid query param format: %q", value)
			}
			options = append(options, tmdb.WithQueryParam(paramParts[0], paramParts[1]))
		}
	}
	if len(appends) > 0 {
		options = append(options, tmdb.WithAppendToResponse(appends...))
	}
	return options, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Expected at least one argument")
		return
	}

	api_key := os.Getenv("TMDB_KEY")
	read_access_token := os.Getenv("TMDB_READ_ACCESS_TOKEN")
	if len(api_key) == 0 && len(read_access_token) == 0 {
		log.Fatal("TMDB_KEY or TMDB_READ_ACCESS_TOKEN must be set")
		return
	}
	client := tmdb.ClientOptions{
		APIKey:             api_key,
		APIReadAccessToken: read_access_token,
	}.NewClient()

	cmd, ok := commands[os.Args[1]]
	if !ok {
		log.Fatal("Unknown command:", os.Args[1])
		return
	}

	if err := cmd(client, os.Args[2:]); err != nil {
		log.Fatal("Error:", err)
	}
}
