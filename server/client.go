package server

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Make http client send Get request
// request query string like e.g: https://localhost:443/hello?q=golang,php,python
// response like e.g: "true,true,false"
func simpleClient(wordsStr string) (string, error) {
	// init client and config tls setting
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	rawUrl := "https://localhost:443/hello"

	params := url.Values{}
	params.Set("q", wordsStr)
	reqUrl, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return "", err
	}

	// encode url format example("bar=baz&foo=fooz")
	reqUrl.RawQuery = params.Encode()

	// send get request
	resp, err := client.Get(reqUrl.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

// Test client
func BcjClient(inputWords []string) ([]bool, error) {
	boolResult := make([]bool, len(inputWords))
	wordsStr := strings.Join(inputWords, ",")
	responseStr, err := simpleClient(wordsStr)
	if err != nil {
		return boolResult, err
	}
	wordsResult := strings.Split(responseStr, ",")
	for i, v := range wordsResult{
		boolResult[i], err = strconv.ParseBool(v)
		if err != nil {
			return boolResult, err
		}
	}
	return boolResult, nil
}



