package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/tmdbepisode"
)

type GetEpisodeDetailsOptions = tmdbepisode.GetDetailsOptions
type GetEpisodeDetailsReply = tmdbepisode.GetDetailsReply

func GetEpisodeDetails(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, options GetEpisodeDetailsOptions) (*http.Response, error) {
	return tmdbepisode.GetDetails(ctx, seriesID, seasonNumber, episodeNumber, options)
}
func ParseGetEpisodeDetailsReply(httpReply *http.Response) (*GetEpisodeDetailsReply, error) {
	return tmdbepisode.ParseGetDetailsReply(httpReply)
}