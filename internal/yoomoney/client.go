package yoomoney

import "net/http"

type Client struct {
	httpClient *http.Client
	token      string
}

func NewClient(token string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		token:      token,
	}
}
