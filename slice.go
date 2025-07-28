package tmdb

import (
	"fmt"
	"iter"
)

type Slice[T any] struct {
	plan[Array]
}

func NewSlice[T any](arr Array) Slice[T] {
	return Slice[T]{
		plan: planFunc[Array](func() (Array, error) {
			if arr == nil {
				return nil, fmt.Errorf("array cannot be nil")
			}
			return arr, nil
		}),
	}
}

func (s Slice[T]) Get(index int) Plan[T] {
	var zero T
	return planFunc[T](func() (T, error) {
		arr, err := s.plan.do()
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

func (s Slice[T]) Len() Plan[int] {
	return planFunc[int](func() (int, error) {
		arr, err := s.plan.do()
		if err != nil {
			return 0, err
		}
		return len(arr), nil
	})
}

func (s Slice[T]) All() Plan[iter.Seq2[int, Plan[T]]] {
	return planFunc[iter.Seq2[int, Plan[T]]](func() (iter.Seq2[int, Plan[T]], error) {
		arr, err := s.plan.do()
		if err != nil {
			return nil, err
		}
		return func(yield func(int, Plan[T]) bool) {
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