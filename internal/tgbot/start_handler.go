package tgbot

import (
	"fmt"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot/state"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleStartMessage(update *tgbotapi.Update) {
	msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "assets/start.jpg")
	msg.Caption = fmt.Sprintf("👋 Привет, %s 😈\\!\n\n"+
		"*Этот бот может найти приватные фотографии девушек, "+
		"анализируя их профили во всех социальных сетях и в слитых базах данных 😏*\n\n"+
		"Приступим? 👇", update.Message.From.FirstName)
	msg.ParseMode = "MarkdownV2"

	bot.Send(msg)

	bot.sendSelectSocialNetwork(update.Message.Chat.ID)

	// Сохраняем статус
	state.Save(update.Message.From.ID, state.SELECT_SOCIAL_NETWORK)
}
