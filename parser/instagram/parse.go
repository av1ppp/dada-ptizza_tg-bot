package instagram

import (
	"encoding/json"
	"net/http"

	"github.com/av1ppp/dada-ptizza_tg-bot/parser"
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

func GetUserInfo(url string) (*parser.UserInfo, error) {
	url = url + "?__a=1"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:92.0) Gecko/20100101 Firefox/92.0")

	// resp, err := http.Get(url)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	r := RawUserResp{}
	if err = dec.Decode(&r); err != nil {
		return nil, err
	}

	if r.GraphQL.User.FullName == "" {
		return nil, parser.ErrWrongServerReponse
	}

	return &parser.UserInfo{
		FullName: r.GraphQL.User.FullName,
		// PictureUrl: r.GraphQL.User.ProfilePicUrl,
	}, nil
}
