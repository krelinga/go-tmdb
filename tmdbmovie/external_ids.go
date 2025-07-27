package tmdbmovie

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krelinga/go-tmdb/internal/util"
)

type GetExternalIDsOptions struct {
}

func GetExternalIDs(ctx context.Context, id int32, options GetExternalIDsOptions) (*http.Response, error) {
	return util.NewRequestBuilder(ctx).
		SetPath("/3/movie/" + fmt.Sprint(id) + "/external_ids").
		Do()
}

func ParseGetExternalIDsReply(httpReply *http.Response) (*GetExternalIDsReply, error) {
	defer httpReply.Body.Close()
	if httpReply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", httpReply.StatusCode)
	}
	if httpReply.Header.Get("Content-Type") != "application/json;charset=utf-8" {
		return nil, fmt.Errorf("unexpected content type: %s", httpReply.Header.Get("Content-Type"))
	}

	rawReply := &struct {
		ID *int32 `json:"id"`
		*ExternalIDs
	}{}
	if err := json.NewDecoder(httpReply.Body).Decode(rawReply); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &GetExternalIDsReply{
		ID:          rawReply.ID,
		ExternalIDs: rawReply.ExternalIDs,
	}, nil
}

type GetExternalIDsReply struct {
	ID          *int32
	ExternalIDs *ExternalIDs
}

func (geir *GetExternalIDsReply) SetDefaults() {
	if geir == nil {
		return
	}
	util.SetIfNil(&geir.ID, 0)
	geir.ExternalIDs.SetDefaults()
}

func (geir *GetExternalIDsReply) String() string {
	if geir == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ID: %s ExternalIDs: %v}", util.FmtOrNil(geir.ID), geir.ExternalIDs.String())
}

type ExternalIDs struct {
	IMDBID      *string `json:"imdb_id"`
	WikidataID  *string `json:"wikidata_id"`
	FacebookID  *string `json:"facebook_id"`
	InstagramID *string `json:"instagram_id"`
	TwitterID   *string `json:"twitter_id"`
}

func (ei *ExternalIDs) SetDefaults() {
	// Nothing to do here for now, we have this function for consistency with other types.
}

func (ei *ExternalIDs) String() string {
	if ei == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{IMDBID: %s WikidataID: %s FacebookID: %s InstagramID: %s TwitterID: %s}",
		util.FmtOrNil(ei.IMDBID), util.FmtOrNil(ei.WikidataID), util.FmtOrNil(ei.FacebookID),
		util.FmtOrNil(ei.InstagramID), util.FmtOrNil(ei.TwitterID))
}
