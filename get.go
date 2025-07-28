package tmdb

func Get[T any](p Plan[T]) (T, error) {
	return p.do()
}

func GetDefault[T any](p Plan[T], defaultValue T) T {
	val, err := p.do()
	if err != nil {
		return defaultValue
	}
	return val
}