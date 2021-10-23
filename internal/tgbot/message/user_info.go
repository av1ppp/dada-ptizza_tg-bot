package message

import (
	"fmt"
	"strings"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/config"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Сообщение "Пользователь не найден"
func MessageUserNotFound(chatID int64) tgbotapi.Chattable {
	return tgbotapi.NewMessage(chatID, "Пользователь не найден ❌")
}

func EditMessageUserNotFound(chatID int64, messageID int) tgbotapi.Chattable {
	return tgbotapi.NewEditMessageText(chatID, messageID, "Пользователь не найден ❌")
}

// func MessageUserInfoWithCounter(user *store.User, p *store.Purchase, yoomoneyApi *yoomoney.Client) (tgbotapi.Chattable, error) {
// 	file := tgbotapi.FileBytes{
// 		Bytes: user.Picture,
// 		Name:  "picture",
// 	}

// 	buyKeyboard, err := getBuyKeyboard(yoomoneyApi, p)
// 	if err != nil {
// 		return nil, err
// 	}

// 	text := fmt.Sprintf("**Пользователь найден ✅**\n\n"+
// 		"*Имя: %s %s*\n\n〰️〰️〰️〰️〰️〰️〰️\n\n"+
// 		"🔞 Приватные фотографии пользователя: %d\n"+
// 		"🔞 Приватные ВИДЕО пользователя: %d\n"+
// 		"⛔️ Скрытые данные пользователя: %d\n"+
// 		"👥 Скрытые друзья пользователя: %d\n\n"+
// 		"💰 Стоимость проверки пользователя: %.1f₽",
// 		user.FirstName,
// 		user.LastName,
// 		user.CountPrivatePhotos,
// 		user.CountPrivateVideos,
// 		user.CountHiddenData,
// 		user.CountHiddenFriends,
// 		p.CheckPrice,
// 	)

// 	msg := tgbotapi.NewPhotoUpload(p.ChatID, file)
// 	msg.Caption = strings.ReplaceAll(text, ".", "\\.")
// 	msg.ParseMode = "MarkdownV2"
// 	msg.ReplyMarkup = buyKeyboard

// 	return &msg, nil
// }

func MessagesUserInfoArchiveFormed(p *store.Purchase, yoomoneyApi *yoomoney.Client) ([]tgbotapi.Chattable, error) {
	var err error

	files := []interface{}{}
	for _, p := range config.Global().Hacks[0].Blur {
		files = append(files, tgbotapi.NewInputMediaPhoto(p))
	}
	msgPhotos := tgbotapi.NewMediaGroup(p.ChatID, files)

	text := fmt.Sprintf("**Взлом найден ✅**\n\n"+
		"*Имя пользователя: %s %s\n"+
		"ID: 363123452\n"+
		"Дата взлома: 27.07.2021*\n\n"+
		"〰️〰️〰️〰️〰️〰️〰️\n\n"+
		"🔞 Приватные фотографии пользователя: %d\n"+
		"🔞 Приватные ВИДЕО пользователя: %d\n"+
		"⛔️ Скрытые данные пользователя: %d\n"+
		"👥 Скрытые друзья пользователя: %d\n\n"+
		"Интим фото: В наличии ✅\n"+
		"Интим видео: В наличии ✅\n\n"+
		"Архив взломанной страницы уже сформирован. Все диалоги и вложения страницы готовы к отправке.",
		p.TargetUser.FirstName,
		p.TargetUser.LastName,
		p.TargetUser.CountPrivatePhotos,
		p.TargetUser.CountPrivateVideos,
		p.TargetUser.CountHiddenData,
		p.TargetUser.CountHiddenFriends,
	)

	msgText := tgbotapi.NewMessage(p.ChatID, strings.ReplaceAll(text, ".", "\\."))
	msgText.ParseMode = "MarkdownV2"

	if msgText.ReplyMarkup, err = getBuyArchiveKeyboard(yoomoneyApi, p); err != nil {
		return nil, err
	}

	return []tgbotapi.Chattable{msgPhotos, msgText}, nil
}

func MessagesUserInfoArchivePictures(p *store.Purchase) tgbotapi.Chattable {
	files := []interface{}{}
	for _, p := range config.Global().Hacks[0].Orig {
		files = append(files, tgbotapi.NewInputMediaPhoto(p))
	}
	msgPhotos := tgbotapi.NewMediaGroup(p.ChatID, files)

	return msgPhotos
}

func MessageUserInfoHiddenCounters(p *store.Purchase, yoomoneyApi *yoomoney.Client) (tgbotapi.Chattable, error) {
	file := tgbotapi.FileBytes{
		Bytes: p.TargetUser.Picture,
		Name:  "picture",
	}

	buyKeyboard, err := getBuyCheckKeyboard(yoomoneyApi, p)
	if err != nil {
		return nil, err
	}

	text := fmt.Sprintf("**Пользователь найден ✅**\n\n"+
		"*Имя: %s %s*\n\n〰️〰️〰️〰️〰️〰️〰️\n\n"+
		"🔞 Приватные фотографии пользователя: ?\n"+
		"🔞 Приватные ВИДЕО пользователя: ?\n"+
		"⛔️ Скрытые данные пользователя: ?\n"+
		"👥 Скрытые друзья пользователя: ?\n\n"+
		"💰 Стоимость проверки пользователя: %.1f₽",
		p.TargetUser.FirstName,
		p.TargetUser.LastName,
		p.CheckPrice,
	)

	msg := tgbotapi.NewPhotoUpload(p.ChatID, file)
	msg.Caption = strings.ReplaceAll(text, ".", "\\.")
	msg.ParseMode = "MarkdownV2"
	msg.ReplyMarkup = buyKeyboard

	return &msg, nil
}
