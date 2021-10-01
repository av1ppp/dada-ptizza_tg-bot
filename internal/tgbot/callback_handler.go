package tgbot

import (
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var selectNetworkPattern = regexp.MustCompile(`social-network__([\w-]+)`)

func (bot *Bot) handleCallback(update *tgbotapi.Update) {
	cbData := update.CallbackQuery.Data

	if selectNetworkPattern.MatchString(cbData) {
		command := selectNetworkPattern.FindStringSubmatch(cbData)[1]
		chatID := update.CallbackQuery.Message.Chat.ID

		if command == "back" {
			// Пользователь нажал "Назад"
			msg := tgbotapi.NewEditMessageText(chatID, lastMsg.MessageID, "🔥 Выбери, где будем искать:")
			msg.ReplyMarkup = &selectSocialNetworkKeyboard
			bot.Send(msg)
			return

		} else {
			// Пользователь выбрал соц. сеть
			bot.handleSelectSocialNetworkCallback(update, command)
			return
		}

	}
}
