package tmdb

type fallback[T any] interface {
	setFallback(T)
}