package tmdb

import "fmt"

func checkCode(data []byte, code ClientHttpCode, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", code)
	}

	return data, nil
}