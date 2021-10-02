package tgbot

import (
	"strings"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleCallback(update *tgbotapi.Update, ds *DialogState) {
	chatID := update.CallbackQuery.Message.Chat.ID
	messageID := update.CallbackQuery.Message.MessageID
	data := update.CallbackQuery.Data
	dataItems := strings.Split(data, "__")

	command := dataItems[0]

	switch command {
	case "select-social-network":
		if len(dataItems) < 2 {
			return
		}
		subcommand := dataItems[1]

		switch subcommand {
		case "back":
			bot.Send(editMessageStartSelectSocialNetwork(
				chatID, messageID))
			return

		case "insta":
			bot.Send(editMessageSendMeInstaUrl(
				chatID, messageID))
			ds.SocicalNetwork = "instagram"

		case "vk":
			bot.Send(editMessageSendMeVKUrl(
				chatID, messageID))
			ds.SocicalNetwork = "vkontakte"
		}

		bot.SaveDialogState(ds)
		return

	case "check-payment":
		paid, err := bot.checkPayment(ds)
		if err != nil {
			bot.sendRequestError(ds.ChatID, err)
			return
		}

		if paid {
			bot.Send(messageItemPaid(ds.ChatID))
		} else {
			bot.Send(messageItemUnpaid(ds.ChatID))
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
