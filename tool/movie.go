package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/krelinga/go-jsonflex"
	"github.com/krelinga/go-tmdb"
)

func movie(client tmdb.Client, args []string) error {
	if len(args) == 0 {
		return errors.New("expected movie ID as argument")
	}

	movieID, err := toInt32(args[0])
	if err != nil {
		return fmt.Errorf("invalid movie ID: %w", err)
	}
	options, err := toRequestOptions(args[1:])
	if err != nil {
		return err
	}
	out, err := tmdb.GetMovie(context.Background(), client, movieID, options...)
	if err != nil {
		return err
	}
	fmt.Println(jsonflex.String(out))

	return nil
}
