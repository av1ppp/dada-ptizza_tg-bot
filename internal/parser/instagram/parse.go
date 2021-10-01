package instagram

import (
	"net/url"
	"strings"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/instagram"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
)

func GetUserInfo(u *url.URL, instagramApi *instagram.Client) (*parser.UserInfo, error) {
	username := strings.ReplaceAll(u.Path, "/", "")

	userInfo, err := instagramApi.GetUserInfo(username)
	if err != nil {
		return nil, err
	}

	// Парсинг
	if userInfo.FullName == "" {
		return nil, parser.ErrWrongServerReponse
	}

	var picture *parser.Picture
	pucUrl := userInfo.ProfilePicUrlHD
	if pucUrl != "" {
		picture, _ = parser.GetPicture(pucUrl)
	}

	return &parser.UserInfo{
		URL:      u,
		FullName: userInfo.FullName,
		Picture:  picture,
	}, nil
}
