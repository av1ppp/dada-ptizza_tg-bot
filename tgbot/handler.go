package tgbot

import (
	"fmt"
	"log"
	"regexp"

	"github.com/av1ppp/dada-ptizza_tg-bot/parser"
	"github.com/av1ppp/dada-ptizza_tg-bot/parser/instagram"
	"github.com/av1ppp/dada-ptizza_tg-bot/parser/vkontakte"
	"github.com/av1ppp/dada-ptizza_tg-bot/state"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleUpdate(update *tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.Text == "/start" {
			bot.handleStartMessage(update)
			return
		}

		ds := state.Get(update.Message.From.ID)
		if ds.IsSelectUser() {
			var ui *parser.UserInfo
			var err error

			if ds.State == state.SELECT_USER_INSTAGRAM {
				ui, err = instagram.GetUserInfo(update.Message.Text)
				if err != nil {
					log.Fatal(err)
				}
			} else if ds.State == state.SELECT_USER_VKONTAKTE {
				ui, err = vkontakte.GetUserInfo(update.Message.Text)
				if err != nil {
					log.Fatal(err)
				}
			}

			fmt.Println(ui)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"**Имя пользователя: "+ui.FullName+"**\n\n"+
					"➖➖➖➖➖➖➖➖➖➖➖➖➖")
			bot.Send(msg)

		}

	} else if update.CallbackQuery != nil {
		// if ds := state.Get(update.CallbackQuery.From.ID); ds.State == state.SELECT_USER {

		// }
		bot.handleCallback(update)
	}

}

var selectNetworkPattern = regexp.MustCompile(`social-network__([\w-]+)`)

func (bot *Bot) handleCallback(update *tgbotapi.Update) {
	data := update.CallbackQuery.Data

	if selectNetworkPattern.MatchString(data) {
		bot.handleSelectNetworkCallback(update, selectNetworkPattern.FindStringSubmatch(data)[1])
	}
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
