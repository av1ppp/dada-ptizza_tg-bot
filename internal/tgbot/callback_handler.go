package tgbot

import (
	"strings"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleCallback(update *tgbotapi.Update, ds *dialogState) {
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
				ds.ChatID, messageID))
			return

		case "insta":
			bot.Send(editMessageSendMeInstaUrl(
				ds.ChatID, messageID))
			ds.SocialNetwork = store.SocialNetworkInsta

		case "vk":
			bot.Send(editMessageSendMeVKUrl(
				ds.ChatID, messageID))
			ds.SocialNetwork = store.SocialNetworkVK
		}

		return

	case "check-payment":
		paid, err := bot.checkPayment(ds)
		if err != nil {
			bot.sendRequestError(ds.ChatID, err)
			return
		}

		if paid {
			bot.Send(messagePaymentReceived(ds.ChatID))

		} else {
			bot.Send(messageItemUnpaid(ds.ChatID))
		}
	}
}

// Проверка, оплатит ли юзер
func (bot *Bot) checkPayment(ds *dialogState) (bool, error) {
	resp, err := bot.yoomoneyApi.CallOperationHistory(&yoomoney.OperationHistoryRequest{
		Label: ds.Label,
	})
	if err != nil {
		return false, err
	}
	// TODO: Проверять еще и сумму
	return len(resp.Operations) > 0, nil
}

// Проверка, оплатит ли юзер
// TODO: Проверять не по ds, а по store.purchase
// func (bot *Bot) checkPayment2() (bool, error) {

// 	resp, err := bot.yoomoneyApi.CallOperationHistory(&yoomoney.OperationHistoryRequest{
// 		Label: ds.Label,
// 	})
// 	if err != nil {
// 		return false, err
// 	}
// 	// TODO: Проверять еще и сумму
// 	return len(resp.Operations) > 0, nil
// }
