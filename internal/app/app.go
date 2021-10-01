package app

import (
	"log"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/config"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/instagram"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/vkapi"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
)

type App struct {
	vkApi        *vkapi.API
	telegramBot  *tgbot.Bot
	instagramApi *instagram.Client
	yoomoneyApi  *yoomoney.Client
	store        *store.Store
}

func New(conf *config.Config) (*App, error) {
	// store
	s, err := store.New()
	if err != nil {
		return nil, err
	}

	// vk
	vkApi := vkapi.NewClient(conf.VK.Token)
	log.Printf("vk api - success")

	// yoomoney
	yoomoneyApi := yoomoney.NewClient(conf.YooMoney.AccessToken)
	log.Printf("yoomoney api - success")

	// instagram
	insta := instagram.NewClient(conf.Instagram.SessionID)
	log.Printf("instagram - success")

	// telegram
	tgBot, err := tgbot.New(conf.TelegramBot.Token, vkApi, s, yoomoneyApi, insta)
	if err != nil {
		return nil, err
	}
	log.Printf("telegram bot - success")

	return &App{
		vkApi:        vkApi,
		telegramBot:  tgBot,
		yoomoneyApi:  yoomoneyApi,
		instagramApi: insta,
		store:        s,
	}, nil
}

func (app *App) Do() error {
	return app.telegramBot.Start()
}
