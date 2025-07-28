package tmdb

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RequestOption struct {
	ChangeValues   func(*url.Values)
	ChangeHeader   func(*http.Header)
	ChangeRequest  func(*http.Request)
	ChangeResponse func(*http.Response)
}

func WithQueryParam(key string, value any) RequestOption {
	return RequestOption{
		ChangeValues: func(values *url.Values) {
			if *values == nil {
				*values = url.Values{}
			}
			values.Add(key, url.QueryEscape(fmt.Sprint(value)))
		},
	}
}

func WithRequestHeader(key, value string) RequestOption {
	return RequestOption{
		ChangeHeader: func(header *http.Header) {
			if *header == nil {
				*header = http.Header{}
			}
			header.Add(key, value)
		},
	}
}

func WithAppendToResponse(appends ...string) RequestOption {
	return RequestOption{
		ChangeValues: func(values *url.Values) {
			if *values == nil {
				*values = url.Values{}
			}
			values.Add("append_to_response", url.QueryEscape(strings.Join(appends, ",")))
		},
	}
}

func WithRequestInterceptor(interceptor func(*http.Request)) RequestOption {
	return RequestOption{
		ChangeRequest: interceptor,
	}
}

func WithResponseInterceptor(interceptor func(*http.Response)) RequestOption {
	return RequestOption{
		ChangeResponse: interceptor,
	}
}
