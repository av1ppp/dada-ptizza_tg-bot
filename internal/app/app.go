package app

import (
	"log"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/config"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/vkapi"
)

type App struct {
	VkApi       *vkapi.API
	TelegramBot *tgbot.Bot
}

func New(conf *config.Config) (*App, error) {
	vkApi := vkapi.NewClient("cd5e4b3de057ae5124b4eafd730922b1481f6775cb6984cb731fe3bc1e9129ab7582e99c2994153ea8f9b")

	log.Printf("vk api - success")

	tgBot, err := tgbot.New(conf.TelegramBot.Token, vkApi)
	if err != nil {
		return nil, err
	}

	return &App{
		VkApi:       vkApi,
		TelegramBot: tgBot,
	}, nil
}

func (app *App) Do() error {
	return app.TelegramBot.Start()
}
