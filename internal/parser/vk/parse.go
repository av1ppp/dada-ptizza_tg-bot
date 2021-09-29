package vk

import (
	"net/url"
	"strings"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/vkapi"
)

func GetUserInfo(userUrl *url.URL, vkApi *vkapi.API) (*parser.UserInfo, error) {
	userId := strings.Split(userUrl.Path, "/")[1]

	if userId == "" {
		return nil, parser.ErrUserNotFound
	}

	users, err := vkApi.UsersGet(vkapi.UsersGetParams{
		UserIds: userId,
		Fields:  "photo_400_orig",
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
		FullName: user.FirstName + " " + user.LastName,
		Picture:  picture,
	}, nil
}
