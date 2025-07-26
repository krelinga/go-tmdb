package tmdb

import (
	"context"

	"github.com/krelinga/go-tmdb/internal/util"
)

type Context = util.Context

func SetContext(ctx context.Context, value Context) context.Context {
	return util.SetContext(ctx, value)
}