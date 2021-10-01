package parser

import "net/url"

type UserInfo struct {
	URL      *url.URL
	FullName string
	Picture  *Picture
}
