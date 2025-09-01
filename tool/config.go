package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/krelinga/go-jsonflex"
	"github.com/krelinga/go-tmdb"
)

func configDetails(client tmdb.Client, args []string) error {
	if len(args) != 0 {
		return errors.New("No arguments expected")
	}

	out, err := tmdb.GetConfigDetails(context.Background(), client)
	if err != nil {
		return err
	}
	fmt.Println(jsonflex.String(out))

	return nil
}