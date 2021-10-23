package tgbot

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func chatIDFromUpdate(update *tgbotapi.Update) int64 {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID
	} else if update.Message != nil {
		return update.Message.Chat.ID
	}
	return 0
}

func label(chatID int64) string {
	return fmt.Sprintf("%x_%x", chatID, time.Now().UnixNano())
}

func (bot *Bot) handleUpdate(update *tgbotapi.Update) {
	if update.Message == nil && update.CallbackQuery == nil {
		return
	}

	chatID := chatIDFromUpdate(update)
	if chatID == 0 {
		fmt.Println("Не удалось определить chatID")
		return
	}

	// Получаем purchase, если нету - создаем
	p, err := store.GetActivePurchaseByChatID(chatID)
	if err != nil {
		if err == sql.ErrNoRows {
			p = store.NewPurchase(&store.PurchaseConfig{
				ChatID: chatID,
				Label:  label(chatID),
				// CheckPrice:        49,
				// UnlimitCheckPrice: 999,
				// ArchivePrice:      450,
				CheckPrice:        3,
				UnlimitCheckPrice: 5,
				ArchivePrice:      7,
			})
			p.SetActive(true)

			if err = store.SavePurchase(p); err != nil {
				bot.sendRequestError(chatID, err)
				return
			}
		} else {
			bot.sendRequestError(chatID, err)
			return
		}
	}

	// Обработка callbacks
	if update.CallbackQuery != nil {
		bot.handleCallback(update, p)
		return
	}

	// Обработка команд
	if update.Message.Command() != "" {
		bot.handleCommand(update, p)
		return
	}

	// Обработка сообщений
	if update.Message.Text != "" {
		bot.handleMessage(update, p)
		return
	}
}

func (bot *Bot) sendRequestError(chatID int64, err error) {
	log.Printf("Error: %s", err)
	bot.Send(message.MessageRequestError(chatID))
}
