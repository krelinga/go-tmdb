package tmdb

import (
	"iter"
)

type Slice[T any] struct {
	data[Array]
	members func(Data[any]) T
}

func NewSlice[T any](in Data[Array], members func(Data[any]) T) Slice[T] {
	return Slice[T]{
		data:    in,
		members: members,
	}
}

func (s Slice[T]) Get(index int) T {
	return s.members(GetIndex(s, index))
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

func (s Slice[T]) All() Data[iter.Seq2[int, T]] {
	return dataFunc[iter.Seq2[int, T]](func() (iter.Seq2[int, T], error) {
		arr, err := s.get()
		if err != nil {
			return nil, err
		}
		return func(yield func(int, T) bool) {
			for i := range arr {
				if !yield(i, s.members(GetIndex(s, i))) {
					return
				}
			}
		}, nil
	})
}
