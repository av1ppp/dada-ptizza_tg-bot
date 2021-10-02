package tgbot

import (
	"net/url"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser/instagram"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser/vk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleMessage(update *tgbotapi.Update, ds *DialogState) {
	message := update.Message.Text
	chatID := update.Message.Chat.ID

	if ds.SocicalNetwork == "" {
		return
	}

	u, err := url.Parse(message)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Неверный формат ссылки")
		bot.Send(msg)
		return
	}

	var ui *parser.UserInfo

	switch ds.SocicalNetwork {
	case "instagram":
		ui, err = instagram.GetUserInfo(u, bot.instagramApi)
		if err != nil {
			bot.sendUserNotFound(chatID)
			return
		}
	case "vkontakte":
		ui, err = vk.GetUserInfo(u, bot.vkApi)
		if err != nil {
			bot.sendUserNotFound(chatID)
			return
		}
	}

	ds.TargetUser = message
	bot.SaveDialogState(ds)

	msg, err := messageUserInfo(ui, ds, bot.yoomoneyApi)
	if err != nil {
		bot.sendRequestError(chatID, err)
	}
	bot.Send(msg)
}

func (bot *Bot) sendUserNotFound(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Не удалось найти пользователя ☹️")
	bot.Send(msg)
}
