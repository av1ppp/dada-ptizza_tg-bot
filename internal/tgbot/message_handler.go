package tgbot

import (
	"net/url"
	"time"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser/instagram"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser/vk"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/pkg/rand"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleMessage(update *tgbotapi.Update, ds *dialogState) {
	bot.searchUser(update, ds)
}

func (bot *Bot) searchUser(update *tgbotapi.Update, ds *dialogState) {
	message := update.Message.Text
	delay := time.Second + time.Millisecond*700 // 1.7sec

	var targetUser *store.User
	var err error

	if message == "" {
		return
	}

	// Отправляем "процесс поиска"
	tmpMsgID := bot.sendSearchProgess(ds.ChatID, delay)

	targetUser, err = store.GetUserByURL(message)

	// Если юзера нету в БД - ищем и заносим туда
	if err != nil {
		u, err := url.Parse(message)
		if err != nil {
			msg := tgbotapi.NewMessage(ds.ChatID, "Неверный формат ссылки")
			bot.Send(msg)
			return
		}

		var socnet store.SocialNetwork
		var ui *parser.UserInfo

		switch ds.SocialNetwork {
		case store.SocialNetworkInsta:
			ui, err = instagram.GetUserInfo(u, bot.instagramApi)
			if err != nil {
				bot.Send(messageUserNotFound(ds.ChatID))
				return
			}
			socnet = store.SocialNetworkInsta
		case store.SocialNetworkVK:
			ui, err = vk.GetUserInfo(u, bot.vkApi)
			if err != nil {
				bot.Send(messageUserNotFound(ds.ChatID))
				return
			}
			socnet = store.SocialNetworkVK
		default:
			bot.sendRequestError(ds.ChatID, ErrUnknownSocialNetwork)
			return
		}

		var sex string
		if ui.Sex == parser.SexFemale {
			sex = "female"
		} else if ui.Sex == parser.SexMale {
			sex = "male"
		}

		targetUser = &store.User{
			FirstName:          ui.FirstName,
			LastName:           ui.LastName,
			Sex:                store.Sex(sex),
			Picture:            *ui.Picture.Data,
			URL:                message,
			SocialNetwork:      socnet,
			CountPrivatePhotos: rand.IntMinMax(39, 58),
			CountPrivateVideos: rand.IntMinMax(12, 21),
			CountHiddenData:    rand.IntMinMax(10, 18),
			CountHiddenFriends: rand.IntMinMax(1, 3),
		}

		if err := store.SaveUser(targetUser); err != nil {
			bot.sendRequestError(ds.ChatID, err)
			return
		}
	}

	// Если мужской пол - отвечаем мол пользователя нету
	if targetUser.Sex == store.SexMale {
		bot.Send(editMessageUserNotFound(ds.ChatID, tmpMsgID))
		return
	}

	// Сохраняем информацию в таблицу с платежами
	err = store.UpdateOrSavePurchase(&store.Purchase{
		ChatID:       ds.ChatID,
		Price:        defaultPrice,
		TargetUserID: targetUser.ID,
		Label:        ds.Label,
	})
	if err != nil {
		bot.sendRequestError(ds.ChatID, err)
		return
	}

	ds.TargetUserURL = message

	msg_, err := messageUserInfo(targetUser, ds, bot.yoomoneyApi)
	if err != nil {
		bot.sendRequestError(ds.ChatID, err)
	}

	if _, err := bot.Send(msg_); err != nil {
		bot.sendRequestError(ds.ChatID, err)
	}

	// Удаляем временное сообщение
	bot.DeleteMessage(tgbotapi.NewDeleteMessage(ds.ChatID, tmpMsgID))

}

// Отправляет статус поиска. Возвращает ID сообщения
func (bot *Bot) sendSearchProgess(chatID int64, delay time.Duration) int {
	// Идёт поиск 🔍...
	tmpMsg, _ := bot.Send(messageSearchInProgess(chatID))
	time.Sleep(delay)

	// Проверяем наши взломы 😈...
	bot.Send(editMessageCheckOurHacks(chatID, tmpMsg.MessageID))
	time.Sleep(delay)

	// Проверяем наши сливы 🤯...
	bot.Send(editMessageCheckOurPlums(chatID, tmpMsg.MessageID))
	time.Sleep(delay)

	return tmpMsg.MessageID
}
