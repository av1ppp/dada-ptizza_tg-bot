package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleCommand(update *tgbotapi.Update, ds *dialogState) {
	command := update.Message.Command()

	switch command {
	case "start":
		ds.Reset()
		bot.Send(messageStart(ds.ChatID, update.Message.From.FirstName))
		bot.Send(messageStartSelectSocialNetwork(ds.ChatID))

	case "test":
		bot.Send(messageHackPhotos(ds.ChatID))
		msg, _ := messageHackInfo(bot.yoomoneyApi, ds)
		bot.Send(msg)

	default:
		bot.Send(messageUnknown(ds.ChatID))
	}
}
