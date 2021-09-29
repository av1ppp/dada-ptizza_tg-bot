package parser

import "io"

type UserInfo struct {
	FullName      string
	Picture       *Picture
	PictureReader io.ReadCloser
}
