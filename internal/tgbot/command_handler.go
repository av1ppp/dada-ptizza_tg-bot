package tgbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleCommand(command string, update *tgbotapi.Update) {
	switch command {
	case "start":
		bot.handleCommand_start(update)

	default:
		bot.handleCommand_default(update)
	}
}

func (bot *Bot) handleCommand_default(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Извини, но я не знаю эту команду")
	bot.Send(msg)
}

func (bot *Bot) handleCommand_start(update *tgbotapi.Update) {
	msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "assets/start.jpg")
	msg.Caption = fmt.Sprintf("👋 Привет, %s 😈\\!\n\n"+
		"*Этот бот может найти приватные фотографии девушек, "+
		"анализируя их профили во всех социальных сетях и в слитых базах данных 😏*\n\n"+
		"Приступим? 👇", update.Message.From.FirstName)
	msg.ParseMode = "MarkdownV2"

	bot.Send(msg)
	bot.sendSelectSocialNetworkKeyboard(update.Message.Chat.ID)
}

// Отправка сообщения для выбора соц. сети
func (bot *Bot) sendSelectSocialNetworkKeyboard(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "🔥 Выбери, где будем искать:")
	msg.ReplyMarkup = &selectSocialNetworkKeyboard
	bot.sendAndSave(msg)
}

// Клавиатура с выбором соц. сети
var selectSocialNetworkKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Instagram", "social-network__instagram"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ВКонтакте", "social-network__vkontakte"),
	),
	// tgbotapi.NewInlineKeyboardRow(
	// 	tgbotapi.NewInlineKeyboardButtonData("Telegram", "social-network__telegram"),
	// ),
	// tgbotapi.NewInlineKeyboardRow(
	// 	tgbotapi.NewInlineKeyboardButtonData("What's App", "social-network__whatsapp"),
	// ),
	// tgbotapi.NewInlineKeyboardRow(
	// 	tgbotapi.NewInlineKeyboardButtonData("Viber", "social-network__viber"),
	// ),
)
