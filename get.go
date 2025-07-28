package tmdb

func Get[T any](p Data[T]) (T, error) {
	return p.do()
}

func GetDefault[T any](p Data[T], defaultValue T) T {
	val, err := p.do()
	if err != nil {
		return defaultValue
	}
	return val
}
