package tgbot

import (
	"fmt"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/parser"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Сообщение "Ошибка во время обработки запроса"
var _messageRequestError = func() tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(0, "Во время обработки запроса произошла ошибка. Попробуй повторить попытку позже.")
}()

func messageRequestError(chatID int64) tgbotapi.Chattable {
	_messageRequestError.ChatID = chatID
	return &_messageRequestError
}

// Сообщение "Товар оплачен"
var _messageItemPaid = func() tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(0, "Товар оплачен!")
}()

func messageItemPaid(chatID int64) tgbotapi.Chattable {
	_messageItemPaid.ChatID = chatID
	return &_messageItemPaid
}

// Сообщение "Товар не оплачен"
var _messageItemUnpaid = func() tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(0, "Товар не оплачен..")
}()

func messageItemUnpaid(chatID int64) tgbotapi.Chattable {
	_messageItemUnpaid.ChatID = chatID
	return &_messageItemUnpaid
}

// Приветственное сообщение
var _messageStart = func() tgbotapi.PhotoConfig {
	msg := tgbotapi.NewPhotoUpload(0, "assets/start.jpg")
	msg.ParseMode = "MarkdownV2"
	return msg
}()

func messageStart(chatID int64, firstName string) tgbotapi.Chattable {
	_messageStart.ChatID = chatID
	_messageStart.Caption = fmt.Sprintf("👋 Привет, %s😈\\!\n\n"+
		"*Этот бот может найти приватные фотографии девушек, "+
		"анализируя их профили во всех социальных сетях и в слитых базах данных 😏*\n\n"+
		"Приступим? 👇", firstName)
	return &_messageStart
}

// Приветственное сообщение -> выбор где будем искать
var selectSocialNetworkKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Instagram", "select-social-network__insta"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ВКонтакте", "select-social-network__vk"),
	),
)

var _messageStartSelectSocialNetwork = func() tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(0, "🔥 Выбери, где будем искать:")
	msg.ReplyMarkup = &selectSocialNetworkKeyboard
	return msg
}()

func messageStartSelectSocialNetwork(chatID int64) tgbotapi.Chattable {
	_messageStartSelectSocialNetwork.ChatID = chatID
	return &_messageStartSelectSocialNetwork
}

var _editMessageStartSelectSocialNetwork = func() tgbotapi.EditMessageTextConfig {
	msg := tgbotapi.NewEditMessageText(0, 0, "🔥 Выбери, где будем искать:")
	msg.ReplyMarkup = &selectSocialNetworkKeyboard
	return msg
}()

func editMessageStartSelectSocialNetwork(chatID int64, messageID int) tgbotapi.Chattable {
	_editMessageStartSelectSocialNetwork.ChatID = chatID
	_editMessageStartSelectSocialNetwork.MessageID = messageID
	return &_editMessageStartSelectSocialNetwork
}

// Неизвестное сообщение
var _messageUnknown = func() tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(0, "Извини, но я не знаю эту команду")
}()

func messageUnknown(chatID int64) tgbotapi.Chattable {
	_messageUnknown.ChatID = chatID
	return &_messageUnknown
}

// Сообщение "Отправьте ссылку на девушку из инстаграм"
var selectSocialNetworkBackKayboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("◀️ Назад", "select-social-network__back"),
	),
)

var _editMessageSendMeInstaUrl = func() tgbotapi.EditMessageTextConfig {
	msg := tgbotapi.NewEditMessageText(0, 0, "✅️ Отправьте ссылку на девушку из Instagram!\n\n"+
		"📝 Пример: https://instagram.com/buzova86")
	msg.ReplyMarkup = &selectSocialNetworkBackKayboard
	return msg
}()

func editMessageSendMeInstaUrl(chatID int64, messageID int) tgbotapi.Chattable {
	_editMessageSendMeInstaUrl.ChatID = chatID
	_editMessageSendMeInstaUrl.MessageID = messageID
	return &_editMessageSendMeInstaUrl
}

// Сообщение "Отправьте ссылку на девушку из вконтакте"
var _editMessageSendMeVKUrl = func() tgbotapi.EditMessageTextConfig {
	msg := tgbotapi.NewEditMessageText(0, 0, "✅️ Отправьте ссылку на девушку из ВКонтакте!\n\n"+
		"📝 Пример: https://vk.com/olgabuzova")
	msg.ReplyMarkup = &selectSocialNetworkBackKayboard
	return msg
}()

func editMessageSendMeVKUrl(chatID int64, messageID int) tgbotapi.Chattable {
	_editMessageSendMeVKUrl.ChatID = chatID
	_editMessageSendMeVKUrl.MessageID = messageID
	return &_editMessageSendMeVKUrl
}

// Сообщение с информацией о найденом пользователе
func getBuyKeyboard(yoomoneyApi *yoomoney.Client, ds *DialogState) (*tgbotapi.InlineKeyboardMarkup, error) {
	accountInfoResp, err := yoomoneyApi.CallAccountInfo()
	if err != nil {
		return nil, err
	}

	createFormResp, err := yoomoneyApi.CreateFormURL(yoomoney.CreateFormOptions{
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
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Проверить", "check-payment"),
		),
	)

	return &buyKeyboard, nil
}

func messageUserInfo(userInfo *parser.UserInfo, ds *DialogState, yoomoneyApi *yoomoney.Client) (tgbotapi.Chattable, error) {
	file := tgbotapi.FileBytes{
		Name:  userInfo.Picture.Filename,
		Bytes: *userInfo.Picture.Data,
	}

	buyKeyboard, err := getBuyKeyboard(yoomoneyApi, ds)
	if err != nil {
		return nil, err
	}

	msg := tgbotapi.NewPhotoUpload(ds.ChatID, file)
	msg.Caption = fmt.Sprintf("**Пользователь найден:**\n\n"+
		"*Имя: %s*\n\n〰️〰️〰️〰️〰️〰️〰️\n\n"+
		"🔞 _Приватные фотографии пользователя: ?\n"+
		"⛔️ Скрытые данные пользователя: ?\n"+
		"👥 Скрытые друзья пользователя: ?_\n\n", userInfo.FullName)
	msg.ParseMode = "MarkdownV2"
	msg.ReplyMarkup = buyKeyboard

	return &msg, nil
}
