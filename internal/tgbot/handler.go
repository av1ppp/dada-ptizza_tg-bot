package tgbot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func chatIDFromUpdate(update *tgbotapi.Update) int64 {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID
	} else if update.Message != nil {
		return update.Message.Chat.ID
	}
	return 0
}

func (bot *Bot) handleUpdate(update *tgbotapi.Update) {
	if update.Message == nil && update.CallbackQuery == nil {
		return
	}

	chatID := chatIDFromUpdate(update)
	if chatID == 0 {
		fmt.Println("Не удалось определить chatID")
		return
	}

	ds := getDialogState(chatID)

	// Обработка callbacks
	if update.CallbackQuery != nil {
		bot.handleCallback(update, ds)
		return
	}

	// Обработка команд
	if update.Message.Command() != "" {
		bot.handleCommand(update, ds)
		return
	}

	// Обработка сообщений
	if update.Message.Text != "" {
		bot.handleMessage(update, ds)
		return
	}
}

func (bot *Bot) sendRequestError(chatID int64, err error) {
	log.Printf("Error: %s", err)
	bot.Send(messageRequestError(chatID))
}
