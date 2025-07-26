package tmdbconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/krelinga/go-tmdb/internal/util"
)

type GetDetailsOptions struct {
	Key             string
	ReadAccessToken string
}

func GetDetails(ctx context.Context, client *http.Client, options GetDetailsOptions) (*http.Response, error) {
	return util.NewRequestBuilder(ctx, client).
		SetPath("/3/configuration").
		SetApiKey(options.Key).
		SetReadAccessToken(options.ReadAccessToken).
		Do()
}

func ParseGetDetailsReply(httpReply *http.Response) (*GetDetailsReply, error) {
	defer httpReply.Body.Close()
	if httpReply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", httpReply.StatusCode)
	}
	var reply GetDetailsReply
	if err := json.NewDecoder(httpReply.Body).Decode(&reply); err != nil {
		return nil, err
	}
	return &reply, nil
}

type GetDetailsReply struct {
	Images     *Images  `json:"images"`
	ChangeKeys []string `json:"change_keys"`
}

func (gdr *GetDetailsReply) String() string {
	if gdr == nil {
		return "<nil>"
	}
	var builder strings.Builder
	builder.WriteString("{")
	fmt.Fprintf(&builder, "Images: %s", gdr.Images.String())
	fmt.Fprintf(&builder, " ChangeKeys: %v", gdr.ChangeKeys)
	builder.WriteString("}")
	return builder.String()
}

type Images struct {
	BaseURL       *string  `json:"base_url"`
	SecureBaseURL *string  `json:"secure_base_url"`
	BackdropSizes []string `json:"backdrop_sizes"`
	LogoSizes     []string `json:"logo_sizes"`
	PosterSizes   []string `json:"poster_sizes"`
	ProfileSizes  []string `json:"profile_sizes"`
	StillSizes    []string `json:"still_sizes"`
}

func (i *Images) String() string {
	if i == nil {
		return "<nil>"
	}
	var builder strings.Builder
	builder.WriteString("{")
	fmt.Fprintf(&builder, "BaseURL: %s", util.FmtOrNil(i.BaseURL))
	fmt.Fprintf(&builder, " SecureBaseURL: %s", util.FmtOrNil(i.SecureBaseURL))
	fmt.Fprintf(&builder, " BackdropSizes: %v", i.BackdropSizes)
	fmt.Fprintf(&builder, " LogoSizes: %v", i.LogoSizes)
	fmt.Fprintf(&builder, " PosterSizes: %v", i.PosterSizes)
	fmt.Fprintf(&builder, " ProfileSizes: %v", i.ProfileSizes)
	fmt.Fprintf(&builder, " StillSizes: %v", i.StillSizes)
	builder.WriteString("}")
	return builder.String()
}
