package tmdb

import (
	"iter"
)

type Slice[T any] interface {
	Get(index int) T
	Len() Data[int]
	All() Data[iter.Seq2[int, T]]
}

func NewSlice(in Data[Array]) Slice[Data[any]] {
	return rootSlice{
		data: in,
	}
}

type rootSlice struct {
	data[Array]
}

func (s rootSlice) Get(index int) Data[any] {
	return getIndex(s, index)
}

// TODO: should I change the return type to be int32 here?
// That would make it consistent with other integer types in the package.
func (s rootSlice) Len() Data[int] {
	return dataFunc[int](func() (int, error) {
		arr, err := s.get()
		if err != nil {
			return 0, err
		}
		return len(arr), nil
	})
}

func (s rootSlice) All() Data[iter.Seq2[int, Data[any]]] {
	return dataFunc[iter.Seq2[int, Data[any]]](func() (iter.Seq2[int, Data[any]], error) {
		arr, err := s.get()
		if err != nil {
			return nil, err
		}
		return func(yield func(int, Data[any]) bool) {
			for i := range arr {
				if !yield(i, getIndex(s, i)) {
					return
				}
			}
		}, nil
	})
}

func ConvertSlice[From, To any](in Slice[From], convert func(From) To) Slice[To] {
	return converterSlice[From, To]{
		data:    in,
		convert: convert,
	}
}

type converterSlice[From, To any] struct {
	data    Slice[From]
	convert func(From) To
}

func (s converterSlice[From, To]) Get(index int) To {
	return s.convert(s.data.Get(index))
}

func (s converterSlice[From, To]) Len() Data[int] {
	return s.data.Len()
}

func (s converterSlice[From, To]) All() Data[iter.Seq2[int, To]] {
	return dataFunc[iter.Seq2[int, To]](func() (iter.Seq2[int, To], error) {
		seq, err := s.data.All().get()
		if err != nil {
			return nil, err
		}
		return func(yield func(int, To) bool) {
			for i, item := range seq {
				if !yield(i, s.convert(item)) {
					return
				}
			}
		}, nil
	})
}
