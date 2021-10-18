package parser

import "net/url"

type Sex int

const (
	SexNotKnown Sex = 0
	SexFemale   Sex = 1
	SexMale     Sex = 2
)

type UserInfo struct {
	URL       *url.URL
	FirstName string
	LastName  string
	Picture   *Picture
	Sex       Sex
}
