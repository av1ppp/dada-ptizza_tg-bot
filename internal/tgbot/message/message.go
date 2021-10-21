package message

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Сообщение "Ошибка во время обработки запроса"
func MessageRequestError(chatID int64) tgbotapi.Chattable {
	return tgbotapi.NewMessage(
		chatID,
		"Во время обработки запроса произошла ошибка. Попробуй повторить попытку позже.",
	)
}

// Сообщение "Товар не оплачен"
func MessageItemUnpaid(chatID int64) tgbotapi.Chattable {
	return tgbotapi.NewMessage(chatID, "Товар не оплачен..")
}

// Приветственное сообщение
func MessageStart(chatID int64, firstName string) tgbotapi.Chattable {
	msg := tgbotapi.NewPhotoUpload(chatID, "assets/start.jpg")
	msg.ParseMode = "MarkdownV2"
	msg.Caption = fmt.Sprintf("👋 Привет, %s😈\\!\n\n"+
		"*Этот бот может найти приватные фотографии девушек, "+
		"анализируя их профили во всех социальных сетях и в слитых базах данных 😏*\n\n"+
		"Приступим? 👇", firstName)
	return &msg
}

// Неизвестное сообщение
func MessageUnknownCommand(chatID int64) tgbotapi.Chattable {
	return tgbotapi.NewMessage(chatID, "Извини, но я не знаю эту команду")
}
