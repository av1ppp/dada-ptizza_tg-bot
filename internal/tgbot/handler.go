package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleUpdate(update *tgbotapi.Update) {
	var chatID int64
	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID
	} else if update.Message != nil {
		chatID = update.Message.Chat.ID
	}

	ds, err := bot.GetDialogState(chatID)
	if err != nil {
		bot.sendRequestError(chatID, err)
		return
	}

	// Обработка callbacks
	if update.CallbackQuery != nil {
		bot.handleCallback(update, ds)
		return
	}

	// Обработка сообщений
	if update.Message.Command() != "" {
		bot.handleCommand(update.Message.Command(), update)
		return
	}

	// Обработка сообщений
	if update.Message.Text != "" {
		bot.handleMessage(update.Message.Text, update, ds)
		return
	}

	// if update.Message != nil {
	// 	// Обработка команды /start
	// 	if update.Message.Text == "/start" {
	// 		bot.handleStartMessage(update)
	// 		return
	// 	}

	// 	ds := state.Get(update.Message.From.ID)

	// 	// Обработка "получение ссылки на пользователя"
	// 	if ds.IsSelectUser() {
	// 		bot.handleSelectUser(update, ds)
	// 	}

	// } else if update.CallbackQuery != nil {
	// 	bot.handleCallback(update)
	// }

}

// func (bot *Bot) handleCallback_instagram(update *tgbotapi.Update) {
// 	// edit := tgbotapi.EditMessageTextConfig{
// 	// 	BaseEdit: tgbotapi.BaseEdit{
// 	// 		ChatID:    update.Message.Chat.ID,
// 	// 		MessageID: lastMsg.MessageID,
// 	// 	},
// 	// 	Text: "✅️ Отправьте ссылку на девушку из Instagram!\n\n" +
// 	// 		"📝 Пример:\nhttps://instagram.com/buzova86",
// 	// }
// 	text := "✅️ Отправьте ссылку на девушку из Instagram!\n\n" +
// 		"📝 Пример:\nhttps://instagram.com/buzova86"

// 	chatID := update.CallbackQuery.Message.Chat.ID

// 	// var msg tgbotapi.Chattable

// 	// if lastMsg.MessageID != 0 {
// 	// 	msg = tgbotapi.NewEditMessageText(chatID, lastMsg.MessageID, text)

// 	// } else {
// 	// 	msg = tgbotapi.NewMessage(chatID, text)
// 	// }
// 	// bot.Send(msg)

// 	if lastMsg.MessageID != 0 {
// 		msg := tgbotapi.NewEditMessageText(chatID, lastMsg.MessageID, text)
// 		bot.Send(msg)

// 	} else {
// 		msg := tgbotapi.NewMessage(chatID, text)
// 		bot.Send(msg)
// 	}

// }

/*
✅️ Отправьте ссылку на девушку из Instagram!

📝 Пример: https://instagram.com/buzova86
*/
