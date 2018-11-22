package util

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"strings"
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

func UniqueStrings(args []string) (newSlice []string) {
	uniq := map[string]struct{}{}
	for _, arg := range args {
		if _, ok := uniq[arg]; ok {
			continue
		}
		uniq[arg] = struct{}{}
		newSlice = append(newSlice, arg)
	}
	return newSlice
}

func transformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err != nil {
		return "", err
	} else {
		return string(ret), nil
	}
}

// Convert a string encoding from ISO2022JP to UTF-8
func FromISO2022JP(str string) (string, error) {
	ret, err := ioutil.ReadAll(
		transform.NewReader(
			strings.NewReader(str),
			japanese.ISO2022JP.NewDecoder(),
		),
	)
	if err != nil {
		return "", err
	} else {
		return string(ret), nil
	}

}
