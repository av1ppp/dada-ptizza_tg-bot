package tgbot

import (
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/vkapi"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	vkApi       *vkapi.API
	store       *store.Store
	yoomoneyApi *yoomoney.Client

	*tgbotapi.BotAPI
}

func New(token string, vkApi *vkapi.API, store *store.Store, yoomoneyApi *yoomoney.Client) (*Bot, error) {
	apiBot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	// apiBot.Debug = true

	return &Bot{vkApi, store, yoomoneyApi, apiBot}, nil
}

func (bot *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		bot.handleUpdate(&update)
	}

	return nil
}

var lastMsg *tgbotapi.Message

// Отправить и сохранить сообщение
func (bot *Bot) sendAndSave(c tgbotapi.Chattable) (*tgbotapi.Message, error) {
	msg, err := bot.Send(c)
	if err == nil {
		lastMsg = &msg
	}
	return &msg, err
}
