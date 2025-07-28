package tmdb

import (
	"fmt"
	"math"
)

type Data[T any] interface {
	get() (T, error)
}

// Alias for allowing unexported embedding
type data[T any] = Data[T]

type dataFunc[T any] func() (T, error)

func (f dataFunc[T]) get() (T, error) {
	return f()
}

type transformer[T any] func(any) (T, error)

func asString(in any) (string, error) {
	if in == nil {
		return "", fmt.Errorf("input is nil")
	}
	if str, ok := in.(string); ok {
		return str, nil
	}
	return "", fmt.Errorf("input is not a string")
}

func asFloat64(in any) (float64, error) {
	if in == nil {
		return 0, fmt.Errorf("input is nil")
	}
	if num, ok := in.(float64); ok {
		return num, nil
	}
	return 0, fmt.Errorf("input is not a float64")
}

func asInt32(in any) (int32, error) {
	f64, err := asFloat64(in)
	if err != nil {
		return 0, err
	}
	if f64 < float64(math.MinInt32) || f64 > float64(math.MaxInt32) || f64 != float64(int32(f64)) {
		return 0, fmt.Errorf("input is out of int32 range")
	}
	return int32(f64), nil
}

func asObject(in any) (Object, error) {
	if in == nil {
		return nil, fmt.Errorf("input is nil")
	}
	if obj, ok := in.(Object); ok {
		return obj, nil
	}
	return nil, fmt.Errorf("input is not an Object")
}

func asArray(in any) (Array, error) {
	if in == nil {
		return nil, fmt.Errorf("input is nil")
	}
	if arr, ok := in.([]any); ok {
		return arr, nil
	}
	return nil, fmt.Errorf("input is not an array")
}

func asTypedObject[T any](in func(Object) T) transformer[T] {
	return func(anyVal any) (T, error) {
		obj, err := asObject(anyVal)
		if err != nil {
			var zero T
			return zero, err
		}
		return in(obj), nil
	}
}

func fieldData[T any](parent Data[Object], key string, trans transformer[T]) Data[T] {
	var zero T
	return dataFunc[T](func() (T, error) {
		obj, err := parent.get()
		if err != nil {
			return zero, err
		}
		anyVal, ok := obj[key]
		if !ok {
			return zero, fmt.Errorf("key %q not found in object", key)
		}
		return trans(anyVal)
	})
}

func rootData(o Object) Data[Object] {
	return dataFunc[Object](func() (Object, error) {
		if o == nil {
			return nil, fmt.Errorf("object is nil")
		}
		return o, nil
	})
}
