package tgbot

import (
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleCommand(update *tgbotapi.Update, p *store.Purchase) {
	command := update.Message.Command()

	switch command {
	case "start":
		// ds.Reset()
		bot.Send(messageStart(p.ChatID, update.Message.From.FirstName))
		bot.Send(messageStartSelectSocialNetwork(p.ChatID))

	case "test":
		bot.Send(messageHackPhotos(p.ChatID))
		msg, err := messageHackInfo(bot.yoomoneyApi, p)
		if err != nil {
			bot.sendRequestError(p.ChatID, err)
			return
		}
		bot.Send(msg)

	default:
		bot.Send(messageUnknown(p.ChatID))
	}
}
