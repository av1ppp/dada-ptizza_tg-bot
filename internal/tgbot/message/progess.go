package message

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Сообщение "Идет поиск.."
func MessageSearchInProgess(chatID int64) tgbotapi.Chattable {
	return tgbotapi.NewMessage(chatID, "Идёт поиск 🔍...")
}

// Сообщение "Проверяем наши взломы.."
func EditMessageCheckOurHacks(chatID int64, messageID int) tgbotapi.Chattable {
	return tgbotapi.NewEditMessageText(chatID, messageID, "Проверяем наши взломы😈...")
}

// Сообщение "Проверяем наши сливы.."
func EditMessageCheckOurPlums(chatID int64, messageID int) tgbotapi.Chattable {
	return tgbotapi.NewEditMessageText(chatID, messageID, "Проверяем наши сливы🤯...")
}
