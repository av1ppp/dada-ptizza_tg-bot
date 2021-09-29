package parser

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Picture struct {
	Data     *[]byte
	MimeType string
	Filename string
}

func GetPicture(src string) (*Picture, error) {
	// Filename
	var filename string
	url_, err := url.Parse(src)
	if err != nil {
		return nil, err
	}
	splitUrl := strings.Split(url_.Path, "/")
	if len(splitUrl) > 0 {
		filename = splitUrl[len(splitUrl)-1]
	} else {
		return nil, ErrWrongServerReponse
	}

	resp, err := http.Get(src)
	if err != nil {
		return nil, err
	}

	// Mime type
	mimeType := resp.Header.Get("Content-Type")

	// Data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Picture{
		Data:     &data,
		MimeType: mimeType,
		Filename: filename,
	}, nil
}
