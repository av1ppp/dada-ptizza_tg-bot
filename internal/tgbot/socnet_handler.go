package tgbot

import (
	"fmt"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot/state"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Клавиатура с кнопкой "назад"
var selectSocialNetworkBackKayboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("◀️ Назад", "social-network__back"),
	),
)

// Обработать callback от выбора соц. сети
func (bot *Bot) handleSelectSocialNetworkCallback(update *tgbotapi.Update, data string) {
	chatID := update.CallbackQuery.Message.Chat.ID
	var text string
	var state_ state.State

	switch data {
	case "back":
		text = "🔥 Выбери, где будем искать:"
		if lastMsg != nil {
			msg := tgbotapi.NewEditMessageText(chatID, lastMsg.MessageID, text)
			msg.ReplyMarkup = &selectSocialNetworkKeyboard
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(chatID, text)
			msg.ReplyMarkup = &selectSocialNetworkKeyboard
			bot.sendAndSave(msg)
		}
		return

	case "instagram":
		text = "✅️ Отправьте ссылку на девушку из Instagram!\n\n" +
			"📝 Пример:\nhttps://instagram.com/buzova86"
		state_ = state.SELECT_USER_INSTAGRAM
	case "vkontakte":
		text = "✅️ Отправьте ссылку на девушку из ВКонтакте!\n\n" +
			"📝 Пример: https://vk.com/durov"
		state_ = state.SELECT_USER_VKONTAKTE
	// case "telegram":
	// 	text = "✅ Отправьте номер девушки из Telegram!\n\n" +
	// 		"📝 Пример: +79876543211"
	// 	state_ = state.SELECT_USER_TELEGRAM
	// case "whatsapp":
	// 	text = "✅ Отправьте номер девушки из What’S App!\n\n" +
	// 		"📝 Пример: +79876543211"
	// 	state_ = state.SELECT_USER_WHATSAPP
	// case "viber":
	// 	text = "✅ Отправьте ссылку на девушку из Viber!\n\n" +
	// 		"📝 Пример: +79876543211"
	// 	state_ = state.SELECT_USER_VIBER
	default:
		fmt.Printf("bot.handleSelectNetworkCallback | Неизвестный тип соц. сети: %s\n", data)
		text = "Произошла ошибка при обработке запроса. Пожалуйста, повторите попытку позже"
	}

	if lastMsg != nil {
		msg := tgbotapi.NewEditMessageText(chatID, lastMsg.MessageID, text)
		msg.ReplyMarkup = &selectSocialNetworkBackKayboard
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = selectSocialNetworkBackKayboard
		bot.sendAndSave(msg)
	}

	// Сохраняем состояние
	state.Save(update.CallbackQuery.From.ID, state_)
}
