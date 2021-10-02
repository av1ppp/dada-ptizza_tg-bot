package tgbot

import (
	"fmt"
	"sync"
	"time"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/purchase"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
)

var mu sync.Mutex
var price float32 = 5.0

type DialogState = purchase.Purchase

func label(chatID int64) string {
	return fmt.Sprintf("%x_%x", chatID, time.Now().UnixNano())
}

func (bot *Bot) GetDialogState(chatID int64) (*DialogState, error) {
	mu.Lock()
	defer mu.Unlock()

	p, err := bot.store.PurchaseByChatID(chatID)
	if err != nil {
		if err == store.ErrPurchaseNotFound {
			p = &purchase.Purchase{
				ChatID: chatID,
				Price:  price,
				Label:  label(chatID),
			}
			if err = bot.store.SavePurchase(p); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return p, nil
}

func (bot *Bot) ResetDialogState(ds *DialogState) error {
	ds.Label = label(ds.ChatID)
	ds.TargetUser = ""
	ds.SocicalNetwork = ""

	return bot.SaveDialogState(ds)
}

func (bot *Bot) SaveDialogState(ds *DialogState) error {
	mu.Lock()
	defer mu.Unlock()

	if _, err := bot.store.Purchase(ds.ID); err != nil {
		if err == store.ErrPurchaseNotFound {
			if err := bot.store.SavePurchase(ds); err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}

	return bot.store.UpdatePurchase(ds)
}
