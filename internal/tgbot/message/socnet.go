package message

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Выбор соц. сети
var selectSocialNetworkKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Instagram", "select-social-network__insta"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ВКонтакте", "select-social-network__vk"),
	),
)

func MessageSelectSocialNetwork(chatID int64) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "🔥 Выбери, где будем искать:")
	msg.ReplyMarkup = &selectSocialNetworkKeyboard
	return msg
}

func EditMessageSelectSocialNetwork(chatID int64, messageID int) tgbotapi.Chattable {
	msg := tgbotapi.NewEditMessageText(chatID, messageID, "🔥 Выбери, где будем искать:")
	msg.ReplyMarkup = &selectSocialNetworkKeyboard
	return msg
}

// Сообщение "Отправьте ссылку на девушку из инстаграм"
var selectSocialNetworkBackKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("◀️ Назад", "select-social-network__back"),
	),
)

func EditMessageSendMeInstaUrl(chatID int64, messageID int) tgbotapi.Chattable {
	msg := tgbotapi.NewEditMessageText(chatID, messageID, "✅️ Отправьте ссылку на девушку из Instagram!\n\n"+
		"📝 Пример: https://instagram.com/buzova86")
	msg.ReplyMarkup = &selectSocialNetworkBackKeyboard
	return msg
}

// Сообщение "Отправьте ссылку на девушку из вконтакте"
func EditMessageSendMeVKUrl(chatID int64, messageID int) tgbotapi.Chattable {
	msg := tgbotapi.NewEditMessageText(chatID, messageID, "✅️ Отправьте ссылку на девушку из ВКонтакте!\n\n"+
		"📝 Пример: https://vk.com/olgabuzova")
	msg.ReplyMarkup = &selectSocialNetworkBackKeyboard
	return msg
}
