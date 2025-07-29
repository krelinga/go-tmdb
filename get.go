package tmdb

func Get[T any](d Data[T]) (T, error) {
	return d.get()
}

func GetOr[T any](d Data[T], defaultValue T) T {
	val, err := d.get()
	if err != nil {
		return defaultValue
	}
	return val
}

// TODO: play around with this and decide if it's useful.
func Try[T any](sharedErr *error, d Data[T]) T {
	if *sharedErr != nil {
		var zero T
		return zero
	}
	val, err := d.get()
	if err != nil {
		*sharedErr = err
	}
	return val
}
