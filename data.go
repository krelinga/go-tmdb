package tmdb

import (
	"fmt"
	"math"
)

type Data[T any] interface {
	get() (T, error)
}

func NewData[T any](t T) Data[T] {
	return dataFunc[T](func() (T, error) {
		return t, nil
	})
}

// Alias for allowing unexported embedding
type data[T any] = Data[T]

type dataFunc[T any] func() (T, error)

func (f dataFunc[T]) get() (T, error) {
	return f()
}

func GetField(parent Data[Object], key string) Data[any] {
	return dataFunc[any](func() (any, error) {
		obj, err := parent.get()
		if err != nil {
			return nil, err
		}
		if val, ok := obj[key]; !ok {
			return nil, fmt.Errorf("key %q not found in object", key)
		} else {
			return val, nil
		}
	})
}

func GetIndex(parent Data[Array], index int) Data[any] {
	return dataFunc[any](func() (any, error) {
		arr, err := parent.get()
		if err != nil {
			return nil, err
		}
		if arr == nil {
			return nil, fmt.Errorf("array is nil")
		}
		if index < 0 || index >= len(arr) {
			return nil, fmt.Errorf("index %d out of bounds for array of length %d", index, len(arr))
		}
		return arr[index], nil
	})
}

func AsString(in Data[any]) Data[string] {
	return dataFunc[string](func() (string, error) {
		val, err := in.get()
		if err != nil {
			return "", err
		}
		if val == nil {
			return "", fmt.Errorf("input is nil")
		}
		if str, ok := val.(string); !ok {
			return "", fmt.Errorf("input is not a string")
		} else {
			return str, nil
		}
	})
}

func AsNumber(in Data[any]) Data[Number] {
	return dataFunc[Number](func() (Number, error) {
		val, err := in.get()
		if err != nil {
			return 0, err
		}
		if val == nil {
			return 0, fmt.Errorf("input is nil")
		}
		if num, ok := val.(Number); !ok {
			return 0, fmt.Errorf("input is not a Number")
		} else {
			return num, nil
		}
	})
}

func AsInt32(in Data[any]) Data[int32] {
	asNumberData := AsNumber(in)
	return dataFunc[int32](func() (int32, error) {
		numberVal, err := asNumberData.get()
		if err != nil {
			return 0, err
		}
		if numberVal < float64(math.MinInt32) || numberVal > float64(math.MaxInt32) || numberVal != float64(int32(numberVal)) {
			return 0, fmt.Errorf("input is out of int32 range")
		}
		return int32(numberVal), nil
	})
}

func AsObject(in Data[any]) Data[Object] {
	return dataFunc[Object](func() (Object, error) {
		val, err := in.get()
		if err != nil {
			return nil, err
		}
		if val == nil {
			return nil, fmt.Errorf("input is nil")
		}
		if obj, ok := val.(Object); !ok {
			return nil, fmt.Errorf("input is not an Object")
		} else {
			return obj, nil
		}
	})
}

func AsArray(in Data[any]) Data[Array] {
	return dataFunc[Array](func() (Array, error) {
		val, err := in.get()
		if err != nil {
			return nil, err
		}
		if val == nil {
			return nil, fmt.Errorf("input is nil")
		}
		if arr, ok := val.([]any); !ok {
			return nil, fmt.Errorf("input is not an array")
		} else {
			return arr, nil
		}
	})
}

func AsFunc[T Data[Object]](newFunc func(Data[Object]) T) func(Data[any]) T {
	return func(obj Data[any]) T {
		return newFunc(AsObject(obj))
	}
}

func Compose[In Data[X], Middle Data[Y], Out Data[Z], X , Y, Z any](f1 func(In) Middle, f2 func(Middle) Out) func(In) Out {
	return func(input In) Out {
		middle := f1(input)
		return f2(middle)
	}
}
