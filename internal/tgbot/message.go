package tgbot

import (
	"fmt"
	"strings"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
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

// Сообщение "Идет поиск.."
var _messageSearchInProgress = func() tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(0, "Идёт поиск 🔍...")
}()

func messageSearchInProgess(chatID int64) tgbotapi.Chattable {
	_messageSearchInProgress.ChatID = chatID
	return &_messageSearchInProgress
}

// Сообщение "Проверяем наши взломы.."
var _editMessageCheckOurHacks = func() tgbotapi.EditMessageTextConfig {
	return tgbotapi.NewEditMessageText(0, 0, "Проверяем наши взломы😈...")
}()

func editMessageCheckOurHacks(chatID int64, messageID int) tgbotapi.Chattable {
	_editMessageCheckOurHacks.ChatID = chatID
	_editMessageCheckOurHacks.MessageID = messageID
	return &_editMessageCheckOurHacks
}

// Сообщение "Проверяем наши сливы.."
var _editMessageCheckOurPlums = func() tgbotapi.EditMessageTextConfig {
	return tgbotapi.NewEditMessageText(0, 0, "Проверяем наши сливы🤯...")
}()

func editMessageCheckOurPlums(chatID int64, messageID int) tgbotapi.Chattable {
	_editMessageCheckOurPlums.ChatID = chatID
	_editMessageCheckOurPlums.MessageID = messageID
	return &_editMessageCheckOurPlums
}

// Сообщение "Пользователь не найден"
var _messageUserNotFound = func() tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(0, "Пользователь не найден ❌")
}()

func messageUserNotFound(chatID int64) tgbotapi.Chattable {
	_messageUserNotFound.ChatID = chatID
	return &_messageUserNotFound
}

var _editMessageUserNotFound = func() tgbotapi.EditMessageTextConfig {
	return tgbotapi.NewEditMessageText(0, 0, "Пользователь не найден ❌")
}()

func editMessageUserNotFound(chatID int64, messageID int) tgbotapi.Chattable {
	_editMessageUserNotFound.ChatID = chatID
	_editMessageUserNotFound.MessageID = messageID
	return &_editMessageUserNotFound
}

// Сообщение с информацией о найденом пользователе
func getBuyKeyboard(yoomoneyApi *yoomoney.Client, ds *dialogState) (*tgbotapi.InlineKeyboardMarkup, error) {
	accountInfoResp, err := yoomoneyApi.CallAccountInfo()
	if err != nil {
		return nil, err
	}

	paymentForm, err := yoomoneyApi.CreateFormURL(yoomoney.CreateFormOptions{
		PaymentType:  "AC",
		Receiver:     accountInfoResp.Account,
		QuickpayForm: "shop",

		FormComment: "Телеграм бот",
		ShortDest:   "Телеграм бот",

		Label:   ds.Label,
		Targets: "Оплата | Телеграм бот",
		Sum:     defaultPrice,
	})
	if err != nil {
		return nil, err
	}

	paymentFormUnlimit, err := yoomoneyApi.CreateFormURL(yoomoney.CreateFormOptions{
		PaymentType:  "AC",
		Receiver:     accountInfoResp.Account,
		QuickpayForm: "shop",

		FormComment: "Телеграм бот",
		ShortDest:   "Телеграм бот",

		Label:   ds.Label,
		Targets: "Оплата | Телеграм бот",
		Sum:     999,
	})
	if err != nil {
		return nil, err
	}

	buyKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(
				fmt.Sprintf("💰 Оплата проверки | %.1f₽", defaultPrice),
				paymentForm.TempURL.String(),
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(
				fmt.Sprintf("Безлимит проверок на 48 часов | %.1f₽", defaultPriceUnlimint),
				paymentFormUnlimit.TempURL.String(),
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Проверить оплату", "check-payment"),
		),
	)

	return &buyKeyboard, nil
}

func messageUserInfo(user *store.User, ds *dialogState, yoomoneyApi *yoomoney.Client) (tgbotapi.Chattable, error) {
	file := tgbotapi.FileBytes{
		Bytes: user.Picture,
		Name:  "picture",
	}

	buyKeyboard, err := getBuyKeyboard(yoomoneyApi, ds)
	if err != nil {
		return nil, err
	}

	text := fmt.Sprintf("**Пользователь найден ✅**\n\n"+
		"*Имя: %s %s*\n\n〰️〰️〰️〰️〰️〰️〰️\n\n"+
		"🔞 Приватные фотографии пользователя: %d\n"+
		"🔞 Приватные ВИДЕО пользователя: %d\n"+
		"⛔️ Скрытые данные пользователя: %d\n"+
		"👥 Скрытые друзья пользователя: %d\n\n"+
		"💰 Стоимость проверки пользователя: %.1f₽",
		user.FirstName,
		user.LastName,
		user.CountPrivatePhotos,
		user.CountPrivateVideos,
		user.CountHiddenData,
		user.CountHiddenFriends,
		defaultPrice,
	)

	msg := tgbotapi.NewPhotoUpload(ds.ChatID, file)
	msg.Caption = strings.ReplaceAll(text, ".", "\\.")
	msg.ParseMode = "MarkdownV2"
	msg.ReplyMarkup = buyKeyboard

	return &msg, nil
}

// Сообщеие "✅ Оплата получена! Загружаем контент"
var _messagePaymentReceived = func() tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(0, "✅ Оплата получена! Загружаем контент")
}()

func messagePaymentReceived(chatID int64) tgbotapi.Chattable {
	_messagePaymentReceived.ChatID = chatID
	return &_messagePaymentReceived
}

// Сообщение "Взлом найден"
var _messageHackPhotos = func() tgbotapi.MediaGroupConfig {
	return tgbotapi.NewMediaGroup(0, []interface{}{
		tgbotapi.NewInputMediaPhoto("https://raw.githubusercontent.com/php-telegram-bot/assets/master/logo/512px/logo_plain.png"),
		tgbotapi.NewInputMediaPhoto("https://blog.pythonanywhere.com/images/bot-chat-session.png"),
	})
}()

func messageHackPhotos(chatID int64) tgbotapi.Chattable {
	_messageHackPhotos.BaseChat.ChatID = chatID
	return &_messageHackPhotos
}

// Сообщение с информацией о найденом пользователе
func getBuyArchiveKeyboard(yoomoneyApi *yoomoney.Client, ds *dialogState) (*tgbotapi.InlineKeyboardMarkup, error) {
	accountInfoResp, err := yoomoneyApi.CallAccountInfo()
	if err != nil {
		return nil, err
	}

	paymentForm, err := yoomoneyApi.CreateFormURL(yoomoney.CreateFormOptions{
		PaymentType:  "AC",
		Receiver:     accountInfoResp.Account,
		QuickpayForm: "shop",

		FormComment: "Телеграм бот",
		ShortDest:   "Телеграм бот",

		Label:   ds.Label,
		Targets: "Оплата | Телеграм бот",
		Sum:     defaultPrice,
	})
	if err != nil {
		return nil, err
	}

	buyKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(
				fmt.Sprintf("💰 Приобрести архив | %.1f₽", defaultPriceArchive),
				paymentForm.TempURL.String(),
			),
		),
	)

	return &buyKeyboard, nil
}

func messageHackInfo(yoomoneyApi *yoomoney.Client, ds *dialogState) (tgbotapi.Chattable, error) {
	keyboard, err := getBuyArchiveKeyboard(yoomoneyApi, ds)
	if err != nil {
		return nil, err
	}

	msg := tgbotapi.NewMessage(ds.ChatID, fmt.Sprintf(
		"Взлом найден ✅\n\n"+
			"Имя пользователя: %s\n"+
			"ID: 363123452\n"+
			"Дата взлома: 27.07.2021\n\n"+
			"Скачано диалогов: 36\n"+
			"Интим фото: В наличии ✅\n"+
			"Интим видео: В наличии ✅\n\n"+
			"Архив взломанной страницы уже сформирован. Все диалоги и вложения страницы готовы к отправке.",
		ds.TargetUserURL),
	)
	msg.ReplyMarkup = keyboard

	return msg, nil
}
