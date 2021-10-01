package app

import (
	"log"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/config"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/vkapi"
)

type App struct {
	vkApi       *vkapi.API
	telegramBot *tgbot.Bot
	store       *store.Store
}

func New(conf *config.Config) (*App, error) {
	s, err := store.New()
	if err != nil {
		return nil, err
	}

	vkApi := vkapi.NewClient("cd5e4b3de057ae5124b4eafd730922b1481f6775cb6984cb731fe3bc1e9129ab7582e99c2994153ea8f9b")

	log.Printf("vk api - success")

	tgBot, err := tgbot.New(conf.TelegramBot.Token, vkApi, s)
	if err != nil {
		return nil, err
	}

	return &App{
		vkApi:       vkApi,
		telegramBot: tgBot,
		store:       s,
	}, nil
}

func (app *App) Do() error {
	return app.telegramBot.Start()
}
