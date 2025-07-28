package tmdb

type Paged[T Data[Object]] struct {
	data[Object]
}

func NewPaged[T Data[Object]](o Object) Paged[T] {
	return Paged[T]{
		data: rootData(o),
	}
}

func (p Paged[T]) Page() Data[int32] {
	return fieldDataInt32(p, "page")
}

func (p Paged[T]) Results() Slice[T] {
	return Slice[T]{
		data: fieldData[Array](p, "results"),
	}
}

func (p Paged[T]) TotalPages() Data[int32] {
	return fieldDataInt32(p, "total_pages")
}

func (p Paged[T]) TotalResults() Data[int32] {
	return fieldDataInt32(p, "total_results")
}