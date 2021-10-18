package vk

import (
	"net/url"
	"strings"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/vkapi"
)

func GetUserInfo(u *url.URL, vkApi *vkapi.API) (*parser.UserInfo, error) {
	userUrlSplit := strings.Split(u.Path, "/")
	if len(userUrlSplit) < 2 {
		return nil, parser.ErrWrongUrl
	}

	userId := userUrlSplit[1]

	if userId == "" {
		return nil, parser.ErrUserNotFound
	}

	users, err := vkApi.UsersGet(vkapi.UsersGetParams{
		UserIds: userId,
		Fields:  "photo_400_orig,sex",
	})
	if err != nil {
		return nil, err
	}

	if users == nil || len(*users) == 0 {
		return nil, parser.ErrUserNotFound
	}

	user := (*users)[0]

	// Get picture
	var picture *parser.Picture

	if user.Photo400Orig != "" {
		picture, _ = parser.GetPicture(user.Photo400Orig)
	}

	return &parser.UserInfo{
		URL:       u,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Picture:   picture,
		Sex:       parser.Sex(user.Sex),
	}, nil
}
