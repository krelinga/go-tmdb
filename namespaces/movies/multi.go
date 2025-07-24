package movies

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/krelinga/go-tmdb/internal/util"
)

type GetMultiOptions struct {
	Key             string
	ReadAccessToken string
	Language        string

	WantDetails bool
	WantCredits bool
}

func GetMulti(ctx context.Context, client *http.Client, id int32, options GetMultiOptions) (*GetMultiReply, error) {
	var appends []string
	if options.WantCredits {
		appends = append(appends, "credits")
	}
	values := url.Values{}
	util.SetIfNotZero(&values, "api_key", options.Key)
	util.SetIfNotZero(&values, "language", options.Language)
	util.SetIfNotZero(&values, "append_to_response", strings.Join(appends, ","))
	url := &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     "/3/movie/" + fmt.Sprint(id),
		RawQuery: values.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    url,
	}
	util.SetAuthIfNotZero(request, options.ReadAccessToken)
	httpReply, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer httpReply.Body.Close()
	if httpReply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", httpReply.StatusCode)
	}
	if httpReply.Header.Get("Content-Type") != "application/json;charset=utf-8" {
		return nil, fmt.Errorf("unexpected content type: %s", httpReply.Header.Get("Content-Type"))
	}

	rawReply := &struct {
		ID *int32 `json:"id"`
		*Details
		Credits *Credits `json:"credits"`
	}{}
	if err := json.NewDecoder(httpReply.Body).Decode(rawReply); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	reply := &GetMultiReply{
		ID: rawReply.ID,
		Details: func() *Details {
			if options.WantDetails {
				return rawReply.Details
			}
			return nil
		}(),
		Credits: rawReply.Credits,
	}

	return reply, nil
}

type GetMultiReply struct {
	ID *int32

	Details *Details
	Credits *Credits
}

func (gmr *GetMultiReply) SetDefaults() {
	if gmr == nil {
		return
	}
	util.SetIfNil(&gmr.ID, 0)
	gmr.Details.SetDefaults()
	gmr.Credits.SetDefaults()
}

func (gmr *GetMultiReply) String() string {
	if gmr == nil {
		return "<nil>"
	}

	builder := strings.Builder{}
	builder.WriteString("{")
	fmt.Fprintf(&builder, "ID: %s", util.FmtOrNil(gmr.ID))
	fmt.Fprintf(&builder, " Details: %v", gmr.Details)
	fmt.Fprintf(&builder, " Credits: %v", gmr.Credits)
	builder.WriteString("}")
	return builder.String()
}
