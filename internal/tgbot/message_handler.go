package tgbot

import (
	"fmt"
	"net/url"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser/instagram"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser/vk"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *Bot) handleMessage(message string, update *tgbotapi.Update, ds *DialogState) {
	// TODO if ds.SocicalNetwork == ""

	chatID := update.Message.Chat.ID

	if ds.SocicalNetwork == "" {
		return
	}

	u, err := url.Parse(message)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Неверный формат ссылки")
		bot.sendAndSave(msg)
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

	bot.sendUserInfoAndBuyKeyboard(chatID, ui, ds)

}

func (bot *Bot) sendRequestError(chatID int64, err error) {
	fmt.Println("Ошибка:", err)
	msg := tgbotapi.NewMessage(chatID, "Во время обработки запроса произошла ошибка. Попробуй повторить попытку позже.")
	bot.Send(msg)
}

func (bot *Bot) sendUserNotFound(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Не удалось найти пользователя ☹️")
	bot.sendAndSave(msg)
}

// Клавиатура с выбором оплаты
func (bot *Bot) getBuyKeyboard(ds *DialogState) (*tgbotapi.InlineKeyboardMarkup, error) {
	// TODO: add accountInfo cache
	accountInfoResp, err := bot.yoomoneyApi.CallAccountInfo()
	if err != nil {
		return nil, err
	}

	createFormResp, err := bot.yoomoneyApi.CreateFormURL(yoomoney.CreateFormOptions{
		PaymentType:  "AC",
		Receiver:     accountInfoResp.Account,
		QuickpayForm: "shop",

		FormComment: "Телеграм бот",
		ShortDest:   "Телеграм бот",

		Label:   ds.Label,
		Targets: "Оплата | Телеграм бот",
		Sum:     ds.Price,
	})
	if err != nil {
		return nil, err
	}

	buyKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(fmt.Sprintf("Оплата | %.1f₽ 💳", ds.Price), createFormResp.TempURL.String()),
			// TODO: Добавить ссылку для оплаты
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Проверить", "check_payment"),
		),
	)

	return &buyKeyboard, nil
}

// Отправить информацию о найденом пользователе
func (bot *Bot) sendUserInfoAndBuyKeyboard(chatID int64, ui *parser.UserInfo, ds *DialogState) {
	fileBytes := tgbotapi.FileBytes{Name: ui.Picture.Filename, Bytes: *ui.Picture.Data}

	buyKeyboard, err := bot.getBuyKeyboard(ds)
	if err != nil {
		bot.sendRequestError(chatID, err)
		return
	}

	msg := tgbotapi.NewPhotoUpload(chatID, fileBytes)
	msg.Caption = "**Пользователь найден:**\n\n" +
		"*Имя: " + ui.FullName + "*\n\n〰️〰️〰️〰️〰️〰️〰️\n\n" +
		"🔞 _Приватные фотографии пользователя: ?\n" +
		"⛔️ Скрытые данные пользователя: ?\n" +
		"👥 Скрытые друзья пользователя: ?_\n\n"
	msg.ParseMode = "MarkdownV2"
	msg.ReplyMarkup = buyKeyboard

	bot.Send(msg)
}
