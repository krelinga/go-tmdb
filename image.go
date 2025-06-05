package tmdb

type Image string

func (i Image) Url(client *Client, size string) string {
	if i == "" {
		return ""
	}
	baseUrl, _ := client.getSecureImageBaseUrl()
	return baseUrl + size + string(i)
}
