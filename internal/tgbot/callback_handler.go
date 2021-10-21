package tgbot

import (
	"strings"
	"time"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot/message"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleCallback(update *tgbotapi.Update, p *store.Purchase) {
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
			bot.Send(message.EditMessageSelectSocialNetwork(p.ChatID, messageID))
			return

		case "insta":
			bot.Send(message.EditMessageSendMeInstaUrl(p.ChatID, messageID))

		case "vk":
			bot.Send(message.EditMessageSendMeVKUrl(p.ChatID, messageID))
		}

		return

	case "check-payment":
		if len(dataItems) < 2 {
			return
		}
		subcommand := dataItems[1]

		switch subcommand {
		case "check":
			paid, err := bot.checkPayment_check(p)
			if err != nil {
				bot.sendRequestError(p.ChatID, err)
				return
			}

			// Поступи ли платеж на проверку
			if paid {
				bot.sendUserInfoArchiveFormed(p)
				return
			} else {
				bot.Send(message.MessageItemUnpaid(p.ChatID))
				return
			}
		case "archive":

			// ...
		case "check_unlimit":
			// ...
		}

	}
}

func (bot *Bot) sendUserInfoArchiveFormed(p *store.Purchase) {
	// Отправляем "процесс поиска"
	delay := time.Second + time.Millisecond*200 // 1.2sec
	tmpMsgCh := make(chan int)
	go bot.sendSearchProgess(p.ChatID, delay, tmpMsgCh) // ...wait
	tmpMsgID := <-tmpMsgCh

	msgs, err := message.MessagesUserInfoArchiveFormed(p, bot.yoomoneyApi)
	if err != nil {
		bot.sendRequestError(p.ChatID, err)
		return
	}

	// Удаляем временное сообщение
	bot.DeleteMessage(tgbotapi.NewDeleteMessage(p.ChatID, tmpMsgID))

	// Отправляем фейк инфу о юзере
	for _, msg := range msgs {
		bot.Send(msg)
	}
}

// Проверка, оплатит ли юзер
func (bot *Bot) checkPayment_check(p *store.Purchase) (bool, error) {
	req := yoomoney.OperationHistoryRequest{
		Label: p.Label + "__check",
	}

	resp, err := bot.yoomoneyApi.CallOperationHistory(&req)
	if err != nil {
		return false, err
	}

	return len(resp.Operations) > 0, nil
}

// Проверка, оплатит ли юзер
func (bot *Bot) checkPayment_archive(p *store.Purchase) (bool, error) {
	req := yoomoney.OperationHistoryRequest{
		Label: p.Label + "__archive",
	}

	resp, err := bot.yoomoneyApi.CallOperationHistory(&req)
	if err != nil {
		return false, err
	}

	return len(resp.Operations) > 0, nil
}
