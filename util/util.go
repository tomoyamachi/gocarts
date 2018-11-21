package util

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"strconv"
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

func ToIntOr0(arg string, notExist int) int {
	var i int
	var err error
	if i, err = strconv.Atoi(arg); err == nil {
		return i
	}
	return notExist
}

func transformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	} else {
		return "", err
	}
}

// Convert a string encoding from ShiftJIS to UTF-8
func FromISO2022JP(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), japanese.ISO2022JP.NewDecoder())
}
