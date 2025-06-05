package tmdb

type Image string

func (i Image) Url(client *Client, size string) string {
	if i == "" {
		return ""
	}
	return client.configuration.Images.SecureBaseUrl + size + string(i)
}
