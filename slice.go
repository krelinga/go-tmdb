package tmdb

import (
	"fmt"
	"iter"
)

type Slice[T any] struct {
	data[Array]
	trans transformer[T]
}

func newSlice[T any](in Data[Array], trans transformer[T]) Slice[T] {
	return Slice[T]{
		data: in,
		trans: trans,
	}
}

func (s Slice[T]) Get(index int) Data[T] {
	var zero T
	return dataFunc[T](func() (T, error) {
		arr, err := s.get()
		if err != nil {
			return zero, err
		}
		if index < 0 || index >= len(arr) {
			return zero, fmt.Errorf("index %d out of bounds for array of length %d", index, len(arr))
		}
		return s.trans(arr[index])
	})
}

// TODO: should I change the return type to be int32 here?
// That would make it consistent with other integer types in the package.
func (s Slice[T]) Len() Data[int] {
	return dataFunc[int](func() (int, error) {
		arr, err := s.get()
		if err != nil {
			return 0, err
		}
		return len(arr), nil
	})
}

func (s Slice[T]) All() Data[iter.Seq2[int, Data[T]]] {
	return dataFunc[iter.Seq2[int, Data[T]]](func() (iter.Seq2[int, Data[T]], error) {
		arr, err := s.get()
		if err != nil {
			return nil, err
		}
		return func(yield func(int, Data[T]) bool) {
			for i, item := range arr {
				if !yield(i, dataFunc[T](func() (T, error) {
					return s.trans(item)
				})) {
					return
				}
			}
		}, nil
	})
}
