package tmdb

import (
	"fmt"
	"math"
)

type Data[T any] interface {
	get() (T, error)
}

// Alias for allowing embedding
type data[T any] = Data[T]

type field interface {
	bool | float64 | string | Array | Object
}

type dataFunc[T any] func() (T, error)

func (f dataFunc[T]) get() (T, error) {
	return f()
}

func fieldData[T field](parent Data[Object], key string) Data[T] {
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
		val, ok := anyVal.(T)
		if !ok {
			return zero, fmt.Errorf("value for key %q is not of type %T", key, zero)
		}
		return val, nil
	})
}

func int32FieldData(parent Data[Object], key string) Data[int32] {
	var zero int32
	return dataFunc[int32](func() (int32, error) {
		obj, err := parent.get()
		if err != nil {
			return zero, err
		}
		anyVal, ok := obj[key]
		if !ok {
			return zero, fmt.Errorf("key %q not found in object", key)
		}
		val, ok := anyVal.(float64)
		if !ok {
			return zero, fmt.Errorf("value for key %q is not a float64", key)
		}
		if val < float64(math.MinInt32) || val > float64(math.MaxInt32) || val != float64(int32(val)) {
			return zero, fmt.Errorf("value for key %q is out of int32 range", key)
		}
		return int32(val), nil
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
