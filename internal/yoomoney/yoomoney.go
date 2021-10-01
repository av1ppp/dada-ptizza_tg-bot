package yoomoney

import (
	"net/http"
	"net/url"
	"strings"
)

var (
	baseUrl = "https://yoomoney.ru"
)

func (client *Client) sendRequest(method string, body map[string]string) (*http.Response, error) {
	var err error
	var req *http.Request

	// Parse body
	values := url.Values{}
	for key, value := range body {
		values.Add(key, value)
	}

	bodyReader := strings.NewReader(values.Encode())

	req, err = http.NewRequest("POST", baseUrl+"/api/"+method, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+client.token)

	return http.DefaultClient.Do(req)
}
