package tmdb

func ImageUrl(base, size, path string) string {
	return base + size + path
}
