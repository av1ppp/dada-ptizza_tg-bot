package yoomoney

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
)

var (
	baseUrl = "https://yoomoney.ru"
	decoder = schema.NewDecoder()
)

func (client *Client) sendRequest(method string, body interface{}) (*http.Response, error) {
	var err error
	var req *http.Request

	// Get request
	if body != nil {
		// Parse body
		src := make(map[string][]string)

		if err := decoder.Decode(body, src); err != nil {
			return nil, err
		}

		values := url.Values{}
		for key, value := range src {
			values.Add(key, strings.Join(value, ","))
		}
		bodyReader := strings.NewReader(values.Encode())

		req, err = http.NewRequest("POST", baseUrl+"/api/"+method, bodyReader)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest("POST", baseUrl+"/api/"+method, nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+client.token)

	return client.httpClient.Do(req)
}
