package tgbot

import (
	"regexp"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
)

var instRe = regexp.MustCompile(`(?:(?:http|https):\/\/)?(?:www.)?(?:instagram.com|instagr.am)\/[A-Za-z0-9-_]+`)
var vkRe = regexp.MustCompile(`(?:(?:http|https):\/\/)?(?:www.)?(vk|vkontakte)\.com\/[A-Za-z0-9-_]+`)

func DetectSocialNetwork(url string) (store.SocialNetwork, error) {
	if instRe.MatchString(url) {
		return store.SocialNetworkInsta, nil
	}
	if vkRe.MatchString(url) {
		return store.SocialNetworkVK, nil
	}
	return "", ErrUnknownSocialNetwork
}
