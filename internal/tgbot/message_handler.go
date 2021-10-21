package tgbot

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser/instagram"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser/vk"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot/message"
	"github.com/av1ppp/dada-ptizza_tg-bot/pkg/rand"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleMessage(update *tgbotapi.Update, p *store.Purchase) {
	bot.searchUser(update, p)
}

func (bot *Bot) searchUser(update *tgbotapi.Update, p *store.Purchase) {
	msgText := update.Message.Text

	if msgText == "" {
		return
	}

	// Ищем, есть ли запрашиваемый пользователь в БД.
	// Если нету - создаем.
	// Вынес в самовызывающуюся функцию потому что так
	// легче выходить из блока кода
	var (
		err    error
		unf    bool // user not found
		u      *url.URL
		socnet store.SocialNetwork
		ui     *parser.UserInfo
	)

	p.TargetUser, err = store.GetUserByURL(msgText)
	if err != nil && err == sql.ErrNoRows {
		func() {
			socnet, err = DetectSocialNetwork(msgText)
			if err != nil {
				return
			}

			u, err = url.Parse(msgText)
			if err != nil {
				return
			}

			switch socnet {
			case store.SocialNetworkInsta:
				if ui, err = instagram.GetUserInfo(u, bot.instagramApi); err != nil {
					unf = true
					return
				}
			case store.SocialNetworkVK:
				if ui, err = vk.GetUserInfo(u, bot.vkApi); err != nil {
					unf = true
					return
				}
			default:
				err = message.ErrUnknownSocialNetwork
				return
			}

			p.TargetUser = &store.User{
				FirstName:          ui.FirstName,
				LastName:           ui.LastName,
				Picture:            *ui.Picture.Data,
				URL:                msgText,
				SocialNetwork:      socnet,
				CountPrivatePhotos: rand.IntMinMax(39, 58),
				CountPrivateVideos: rand.IntMinMax(12, 21),
				CountHiddenData:    rand.IntMinMax(10, 18),
				CountHiddenFriends: rand.IntMinMax(1, 3),
			}

			if ui.Sex == parser.SexFemale {
				p.TargetUser.Sex = store.SexFemale
			} else if ui.Sex == parser.SexMale {
				p.TargetUser.Sex = store.SexMale
			}

			if err = store.SaveUser(p.TargetUser); err != nil {
				return
			}
		}()
	} else {
		unf = false
	}

	if err != nil {
		bot.sendRequestError(p.ChatID, err)
		return
	}
	if unf {
		bot.Send(message.MessageUserNotFound(p.ChatID))
		return
	}

	// Если мужской пол - отвечаем мол пользователя нету
	if p.TargetUser.Sex == store.SexMale {
		bot.Send(message.MessageUserNotFound(p.ChatID))
		return
	}

	// Сохраняем информацию в таблицу с платежами
	err = store.UpdatePurchaseByID(p)
	if err != nil {
		bot.sendRequestError(p.ChatID, err)
		return
	}

	msg_, err := message.MessageUserInfoHiddenCounters(p, bot.yoomoneyApi)
	if err != nil {
		bot.sendRequestError(p.ChatID, err)
	}

	if _, err := bot.Send(msg_); err != nil {
		bot.sendRequestError(p.ChatID, err)
	}
}

// Отправляет статус поиска. Возвращает ID сообщения
func (bot *Bot) sendSearchProgess(chatID int64, delay time.Duration, tmpMsgCh chan int) {
	// Идёт поиск 🔍...
	tmpMsg, _ := bot.Send(message.MessageSearchInProgess(chatID))
	time.Sleep(delay)

	// Проверяем наши взломы 😈...
	bot.Send(message.EditMessageCheckOurHacks(chatID, tmpMsg.MessageID))
	time.Sleep(delay)

	// Проверяем наши сливы 🤯...
	bot.Send(message.EditMessageCheckOurPlums(chatID, tmpMsg.MessageID))
	time.Sleep(delay)

	tmpMsgCh <- tmpMsg.MessageID
}
