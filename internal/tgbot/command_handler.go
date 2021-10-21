package tgbot

import (
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleCommand(update *tgbotapi.Update, p *store.Purchase) {
	command := update.Message.Command()

	switch command {
	case "start":
		// ds.Reset()
		bot.Send(message.MessageStart(p.ChatID, update.Message.From.FirstName))
		bot.Send(message.MessageSelectSocialNetwork(p.ChatID))

	default:
		bot.Send(message.MessageUnknownCommand(p.ChatID))
	}
}
