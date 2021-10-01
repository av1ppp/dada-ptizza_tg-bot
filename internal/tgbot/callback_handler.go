package tgbot

import (
	"fmt"
	"regexp"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var selectNetworkPattern = regexp.MustCompile(`social-network__([\w-]+)`)

func (bot *Bot) handleCallback(update *tgbotapi.Update, ds *DialogState) {
	cbData := update.CallbackQuery.Data
	chatID := update.CallbackQuery.Message.Chat.ID

	if cbData == "check_payment" {
		// Пользователь проверяет оплату
		paid, err := bot.checkPayment(ds)
		if err != nil {
			bot.sendRequestError(ds.ChatID, err)
			return
		}

		if paid {
			bot.Send(tgbotapi.NewMessage(ds.ChatID, "Товар оплачен!"))
		} else {
			bot.Send(tgbotapi.NewMessage(ds.ChatID, "Товар не оплачен.."))
		}
	}

	if selectNetworkPattern.MatchString(cbData) {
		// Пользователь выбрал соц. сеть
		command := selectNetworkPattern.FindStringSubmatch(cbData)[1]

		if command == "back" {
			// Пользователь нажал "Назад"
			if lastMsg != nil {
				msg := tgbotapi.NewEditMessageText(chatID, lastMsg.MessageID, "🔥 Выбери, где будем искать:")
				msg.ReplyMarkup = &selectSocialNetworkKeyboard
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(chatID, "🔥 Выбери, где будем искать:")
				msg.ReplyMarkup = &selectSocialNetworkKeyboard
				bot.Send(msg)
			}
			return

		} else {
			// Пользователь выбрал соц. сеть
			bot.handleSelectSocialNetworkCallback(update, command, ds)
			return
		}
	}
}

// Проверка, оплатит ли юзер
func (bot *Bot) checkPayment(ds *DialogState) (bool, error) {
	resp, err := bot.yoomoneyApi.CallOperationHistory(&yoomoney.OperationHistoryRequest{
		Label: ds.Label,
	})
	if err != nil {
		return false, err
	}
	return len(resp.Operations) > 0, nil
}

// Клавиатура с кнопкой "назад"
var selectSocialNetworkBackKayboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("◀️ Назад", "social-network__back"),
	),
)

// Обработать callback от выбора соц. сети
func (bot *Bot) handleSelectSocialNetworkCallback(update *tgbotapi.Update, data string, ds *DialogState) {
	chatID := update.CallbackQuery.Message.Chat.ID
	var text string

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
		ds.SocicalNetwork = "instagram"
	case "vkontakte":
		text = "✅️ Отправьте ссылку на девушку из ВКонтакте!\n\n" +
			"📝 Пример: https://vk.com/durov"
		ds.SocicalNetwork = "vkontakte"
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
	bot.SaveDialogState(ds)
}
