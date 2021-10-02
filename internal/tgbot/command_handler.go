package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleCommand(update *tgbotapi.Update, ds *DialogState) {
	command := update.Message.Command()
	chatID := update.Message.Chat.ID

	switch command {
	case "start":
		if err := bot.ResetDialogState(ds); err != nil {
			bot.sendRequestError(chatID, err)
			return
		}
		bot.Send(messageStart(chatID, update.Message.From.FirstName))
		bot.Send(messageStartSelectSocialNetwork(chatID))

	default:
		bot.Send(messageUnknown(chatID))
	}
}
