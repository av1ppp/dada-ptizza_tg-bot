package tgbot

import (
	"fmt"
	"net/url"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser/instagram"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser/vk"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot/state"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleSelectUser_sendError(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Не удалось найти пользователя ☹️")
	bot.sendAndSave(msg)
}

// high probability of detecting intimate photos

// Клавиатура с выбором оплаты
var buyKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Оплата | 39.0₽ 💳", "buy"),
		// TODO: Добавить ссылку для оплаты
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Проверить", "check"),
	),
)

// Отправить информацию о найденом пользователе
func (bot *Bot) sendUserInfo(chatID int64, ui *parser.UserInfo) error {
	fileBytes := tgbotapi.FileBytes{Name: ui.Picture.Filename, Bytes: *ui.Picture.Data}

	msg := tgbotapi.NewPhotoUpload(chatID, fileBytes)
	msg.Caption = "**Пользователь найден:**\n\n" +
		"*Имя: " + ui.FullName + "*\n\n〰️〰️〰️〰️〰️〰️〰️\n\n" +
		"🔞 _Приватные фотографии пользователя: ?\n" +
		"⛔️ Скрытые данные пользователя: ?\n" +
		"👥 Скрытые друзья пользователя: ?\n\n" +
		"💰 Стоимость проверки:_ *39\\.0 RUB*"
	msg.ParseMode = "MarkdownV2"

	_, err := bot.Send(msg)
	return err
}

// Отправить сообщение "высокая верятность обнаружения.."
func (bot *Bot) sendHighProbDetectingPhotos(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "⚙️ Высокая вероятность обраружения интимных фотографий")
	msg.ReplyMarkup = &buyKeyboard
	bot.Send(msg)
}

func (bot *Bot) handleSelectUser(update *tgbotapi.Update, ds *state.DialogState) {
	chatID := update.Message.Chat.ID
	var ui *parser.UserInfo
	var err error

	if ds.State == state.SELECT_USER_INSTAGRAM {
		ui, err = instagram.GetUserInfo(update.Message.Text)
		if err != nil {
			bot.handleSelectUser_sendError(chatID)
			return
		}

	} else if ds.State == state.SELECT_USER_VKONTAKTE {
		url_, err := url.Parse(update.Message.Text)
		if err != nil {
			bot.handleSelectUser_sendError(chatID)
			return
		}

		ui, err = vk.GetUserInfo(url_, bot.vkApi)
		if err != nil {
			bot.handleSelectUser_sendError(chatID)
			return
		}
	}

	err = bot.sendUserInfo(chatID, ui)
	if err != nil {
		fmt.Printf("bot.sendUserInfo | Ошибка при отправке сообщения: %s\n", err)
		bot.handleSelectUser_sendError(chatID)
		return
	}

	bot.sendHighProbDetectingPhotos(chatID)
}