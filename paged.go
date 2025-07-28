package tmdb

type PageOf[T Data[Object]] struct {
	data[Object]
	members func(Data[Object]) T
}

func NewPaged[T Data[Object]](in Data[Object], members func(Data[Object]) T) PageOf[T] {
	return PageOf[T]{
		data:   in,
		members: members,
	}
}

func (p PageOf[T]) Page() Data[int32] {
	return AsInt32(GetField(p, "page"))
}

func (p PageOf[T]) Results() Slice[T] {
	return NewSlice(AsArray(GetField(p, "results")), Compose(AsObject, p.members))
}

func (p PageOf[T]) TotalPages() Data[int32] {
	return AsInt32(GetField(p, "total_pages"))
}

func (p PageOf[T]) TotalResults() Data[int32] {
	return AsInt32(GetField(p, "total_results"))
}
