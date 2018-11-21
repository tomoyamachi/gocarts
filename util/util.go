package util

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

// FetchURL returns HTTP response body
func FetchURL(url string) ([]byte, error) {
	var errs []error
	httpProxy := viper.GetString("http-proxy")

	resp, body, err := gorequest.New().Proxy(httpProxy).Get(url).Type("text").EndBytes()
	if len(errs) > 0 || resp == nil {
		return nil, fmt.Errorf("HTTP error. errs: %v, url: %s", err, url)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error. errs: %v, status code: %d, url: %s", err, resp.StatusCode, url)
	}
	return body, nil
}

func RemoveDuplicateFromSlice(args []string) []string {
	results := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		dup := false
		for j := 0; j < len(results); j++ {
			if args[i] == results[j] {
				dup = true
				break
			}
		}
		if !dup {
			results = append(results, args[i])
		}
	}
	return results
}
