package vk

import "golang.org/x/text/encoding/charmap"

func windows1251ToUTF8(src string) (string, error) {
	decoder := charmap.Windows1251.NewDecoder()
	return decoder.String(src)
}
