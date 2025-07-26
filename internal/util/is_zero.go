package util

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

func setIfNotZero[T comparable](values *url.Values, key string, value T) {
	if !IsZero(value) {
		values.Set(key, fmt.Sprint(value))
	}
}

func setAuthIfNotZero(request *http.Request, value string) {
	if !IsZero(value) {
		if request.Header == nil {
			request.Header = http.Header{}
		}
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", value))
	}
}

func FmtOrNil[T any](v *T) string {
	if v == nil {
		return "<nil>"
	}
	value := reflect.ValueOf(*v)
	switch value.Kind() {
	case reflect.String:
		return fmt.Sprintf("\"%v\"", *v)
	default:
		return fmt.Sprint(*v)
	}
}

func SetIfNil[T any](target **T, value T) {
	if *target == nil {
		*target = &value
	}
}
