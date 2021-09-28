package vkontakte

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/av1ppp/dada-ptizza_tg-bot/parser"
)

func GetUserInfo(url string) (*parser.UserInfo, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:92.0) Gecko/20100101 Firefox/92.0")

	resp, err := http.DefaultClient.Do(req)
	// resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// data, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(data))
	// os.WriteFile("durov.html", data, 0755)
	// return nil, errors.New("ha")

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	fullName := doc.Find(".op_header").Text()
	// pageAvatarSrc, _ := doc.Find("#page_avatar .page_avatar_img").Attr("src")

	if fullName == "" {
		fullName = doc.Find(".page_name").Text()
		if fullName == "" {
			return nil, parser.ErrWrongServerReponse
		}
	}

	return &parser.UserInfo{
		FullName: fullName,
		// PictureUrl: pageAvatarSrc,
	}, nil
}
