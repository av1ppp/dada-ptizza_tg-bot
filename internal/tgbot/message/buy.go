package message

import (
	"fmt"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Сообщение с информацией о найденом пользователе
func getBuyCheckKeyboard(yoomoneyApi *yoomoney.Client, p *store.Purchase) (*tgbotapi.InlineKeyboardMarkup, error) {
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

		Label:   p.Label + "__check",
		Targets: "Оплата | Телеграм бот",
		Sum:     p.CheckPrice,
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

		Label:   p.Label + "__check_unlimit",
		Targets: "Оплата | Телеграм бот",
		Sum:     p.UnlimitCheckPrice,
	})
	if err != nil {
		return nil, err
	}

	buyKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(
				fmt.Sprintf("💰 Оплата проверки | %.1f₽", p.CheckPrice),
				paymentForm.TempURL.String(),
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(
				fmt.Sprintf("Безлимит проверок на 48 часов | %.1f₽", p.UnlimitCheckPrice),
				paymentFormUnlimit.TempURL.String(),
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Проверить оплату", "check-payment__check"),
		),
	)

	return &buyKeyboard, nil
}

// Сообщение с информацией о найденом пользователе
func getBuyArchiveKeyboard(yoomoneyApi *yoomoney.Client, p *store.Purchase) (*tgbotapi.InlineKeyboardMarkup, error) {
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

		Label:   p.Label + "__archive",
		Targets: "Оплата | Телеграм бот",
		Sum:     p.ArchivePrice,
	})
	if err != nil {
		return nil, err
	}

	buyKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(
				fmt.Sprintf("💰 Приобрести архив | %.1f₽", p.CheckPrice),
				paymentForm.TempURL.String(),
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Проверить оплату", "check-payment__archive"),
		),
	)

	return &buyKeyboard, nil
}

// Сообщеие "✅ Оплата получена! Загружаем контент"
func MessagePaymentReceived(chatID int64) tgbotapi.Chattable {
	return tgbotapi.NewMessage(chatID, "✅ Оплата получена! Загружаем контент")
}
