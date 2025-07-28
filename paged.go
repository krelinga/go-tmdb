package tmdb

type PageOf[T Data[Object]] struct {
	data[Object]
	resultTrans transformer[T]
}

func newPaged[T Data[Object]](in Data[Object], resultTrans transformer[T]) PageOf[T] {
	return PageOf[T]{
		data:        in,
		resultTrans: resultTrans,
	}
}

func (p PageOf[T]) Page() Data[int32] {
	return fieldData(p, "page", asInt32)
}

func (p PageOf[T]) Results() Slice[T] {
	return newSlice(fieldData(p, "results", asArray), p.resultTrans)
}

func (p PageOf[T]) TotalPages() Data[int32] {
	return fieldData(p, "total_pages", asInt32)
}

func (p PageOf[T]) TotalResults() Data[int32] {
	return fieldData(p, "total_results", asInt32)
}
