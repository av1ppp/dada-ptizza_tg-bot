package instagram

import (
	"encoding/json"
	"net/http"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
)

type RawUserResp struct {
	GraphQL struct {
		User struct {
			FullName        string `json:"full_name"`
			ProfilePicUrl   string `json:"profile_pic_url"`
			ProfilePicUrlHD string `json:"profile_pic_url_hd"`
		} `json:"user"`
	} `json:"graphql"`
}

func request(url string) *http.Request {
	url = url + "?__a=1"

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

	// Загрузка JSON
	dec := json.NewDecoder(res.Body)
	r := RawUserResp{}
	if err = dec.Decode(&r); err != nil {
		return nil, err
	}

	// Парсинг
	fullName := r.GraphQL.User.FullName
	if fullName == "" {
		return nil, parser.ErrWrongServerReponse
	}

	var picture *parser.Picture
	pucUrl := r.GraphQL.User.ProfilePicUrlHD
	if pucUrl != "" {
		picture, _ = parser.GetPicture(pucUrl)
	}

	return &parser.UserInfo{
		FullName: fullName,
		Picture:  picture,
	}, nil
}
