package tmdb

// This package makes extensive use of pointers to represent optional values.  This helper function provides a convenient way to create a pointer to a given value.
func NewPtr[T any](v T) *T {
	ptr := new(T)
	*ptr = v
	return ptr
}
