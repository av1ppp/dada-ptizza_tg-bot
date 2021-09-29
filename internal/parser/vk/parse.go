package vk

import (
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
)

func request(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:92.0) Gecko/20100101 Firefox/92.0")
	return req
}

func GetUserInfo(url string) (*parser.UserInfo, error) {
	req := request(url)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// data, _ := io.ReadAll(res.Body)
	// os.WriteFile("durov.html", data, 0755)
	// return nil, errors.New("ha")

	// Загрузка HTML документа
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	// Парсинг
	fullName := doc.Find(".op_header").Text()
	if fullName == "" {
		fullName = doc.Find(".page_name").Text()
		if fullName == "" {
			return nil, parser.ErrWrongServerReponse
		}
	}

	var picture *parser.Picture
	pageAvatarSrc, _ := doc.Find(".page_avatar_img").Attr("src")
	if pageAvatarSrc != "" {
		picture, _ = parser.GetPicture(pageAvatarSrc)
	}

	html, _ := doc.Html()
	os.WriteFile("durov.html", []byte(html), 0755)

	// TODO: Если кодировка уже utf-8, то ничего не делаем
	fullNameUTF8, err := windows1251ToUTF8(fullName)
	if err != nil {
		return nil, err
	}

	//asdasd
	res, err = http.Get(pageAvatarSrc)
	if err != nil {
		return nil, err
	}

	return &parser.UserInfo{
		FullName:      fullNameUTF8,
		Picture:       picture,
		PictureReader: res.Body,
	}, nil
}
