package tgbot

import (
	"fmt"

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
}

var selectSocialNetworkKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Instagram", "social-network__instagram"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ВКонтакте", "social-network__vkontakte"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Telegram", "social-network__telegram"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("What's App", "social-network__whatsapp"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Viber", "social-network__viber"),
	),
)

func (bot *Bot) sendSelectSocialNetwork(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "🔥 Выбери, где будем искать:")
	msg.ReplyMarkup = selectSocialNetworkKeyboard
	bot.sendAndSave(msg)
}

func (bot *Bot) handleSelectNetworkCallback(update *tgbotapi.Update, socialNetwork string) {
	var text string

	switch socialNetwork {
	case "instagram":
		text = "✅️ Отправьте ссылку на девушку из Instagram!\n\n" +
			"📝 Пример:\nhttps://instagram.com/buzova86"
		break
	case "vkontakte":
		text = "✅️ Отправьте ссылку на девушку из ВКонтакте!\n\n" +
			"📝 Пример: https://vk.com/durov"
		break
	case "telegram":
		text = "✅ Отправьте номер девушки из Telegram!\n\n" +
			"📝 Пример: +79876543211"
		break
	case "whatsapp":
		text = "✅ Отправьте номер девушки из What’S App!\n\n" +
			"📝 Пример: +79876543211"
		break
	case "viber":
		text = "✅ Отправьте ссылку на девушку из Viber!\n\n" +
			"📝 Пример: +79876543211"
		break
	default:
		fmt.Printf("bot.handleSelectNetworkCallback | Неизвестный тип соц. сети: %s\n", socialNetwork)
		text = "Произошла ошибка при обработке запроса. Пожалуйста, повторите попытку позже"
	}

	var msg tgbotapi.Message

	if lastMsg.MessageID != 0 {
		msg = tgbotapi.NewEditMessageText(chatID, lastMsg.MessageID, text)
	} else {
		msg = tgbotapi.NewMessage(chatID, text)
	}

	bot.Send(msg)
}
