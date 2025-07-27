package tmdb

type Movie Object

func (m Movie) ID() Field[int32] {
	return WrapNumberAsInt32(m, "id")
}

func (m Movie) Title() Field[string] {
	return GetField[string](m, "title")
}

func (m Movie) Adult() Field[bool] {
	return GetField[bool](m, "adult").WithDefault(true)
}

func (m Movie) Keywords() Field[[]Keyword] {
	return WrapArray[Keyword](m, "keywords")
}

type Keyword Object

func (k Keyword) ID() Field[int32] {
	return WrapNumberAsInt32(k, "id")
}

func (k Keyword) Name() Field[string] {
	return GetField[string](k, "name")
}
