package tmdb

import (
	"fmt"
	"iter"
)

type Slice[T any] struct {
	data[Array]
}

func NewSlice[T any](arr Array) Slice[T] {
	return Slice[T]{
		data: planFunc[Array](func() (Array, error) {
			if arr == nil {
				return nil, fmt.Errorf("array cannot be nil")
			}
			return arr, nil
		}),
	}
}

func (s Slice[T]) Get(index int) Data[T] {
	var zero T
	return planFunc[T](func() (T, error) {
		arr, err := s.do()
		if err != nil {
			return zero, err
		}
		if index < 0 || index >= len(arr) {
			return zero, fmt.Errorf("index %d out of bounds for array of length %d", index, len(arr))
		}
		val, ok := arr[index].(T)
		if !ok {
			return zero, fmt.Errorf("value at index %d is not of type %T", index, zero)
		}
		return val, nil
	})
}

func (s Slice[T]) Len() Data[int] {
	return planFunc[int](func() (int, error) {
		arr, err := s.do()
		if err != nil {
			return 0, err
		}
		return len(arr), nil
	})
}

func (s Slice[T]) All() Data[iter.Seq2[int, Data[T]]] {
	return planFunc[iter.Seq2[int, Data[T]]](func() (iter.Seq2[int, Data[T]], error) {
		arr, err := s.do()
		if err != nil {
			return nil, err
		}
		return func(yield func(int, Data[T]) bool) {
			for i, item := range arr {
				if !yield(i, planFunc[T](func() (T, error) {
					val, ok := item.(T)
					if !ok {
						var zero T
						return zero, fmt.Errorf("value at index %d is not of type %T", i, zero)
					}
					return val, nil
				})) {
					return
				}
			}
		}, nil
	})
}
