package tmdb

// Returns a new pointer that points to the given value.
func NewPtr[T any](v T) *T {
	ptr := new(T)
	*ptr = v
	return ptr
}
