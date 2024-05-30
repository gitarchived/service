package util

import "github.com/go-resty/resty/v2"

func IsOk(url string) bool {
	client := resty.New()

	resp, err := client.R().
		SetHeader("User-Agent", "gitarhived/1.0").
		Get(url)

	if err != nil {
		return false
	}

	if resp.StatusCode() != 200 {
		return false
	}

	return true
}
