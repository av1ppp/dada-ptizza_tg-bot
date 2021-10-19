package tgbot

import "errors"

var (
	ErrUnknownSocialNetwork = errors.New("unknown social network")
	ErrUnknownUser          = errors.New("unknown user")
)
