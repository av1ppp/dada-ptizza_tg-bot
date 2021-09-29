package parser

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrWrongUrl           = errors.New("wrong url")
	ErrWrongServerReponse = errors.New("wrong server response")
)
