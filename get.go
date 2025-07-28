package tmdb

func Get[T any](d Data[T]) (T, error) {
	return d.get()
}

func GetDefault[T any](d Data[T], defaultValue T) T {
	val, err := d.get()
	if err != nil {
		return defaultValue
	}
	return val
}
