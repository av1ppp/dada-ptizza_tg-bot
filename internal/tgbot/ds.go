package tgbot

import (
	"fmt"
	"sync"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/purchase"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
)

var mu sync.Mutex

type DialogState = purchase.Purchase

func (bot *Bot) GetDialogState(chatID int64) (*DialogState, error) {
	mu.Lock()
	defer mu.Unlock()

	p, err := bot.store.PurchaseByChatID(chatID)
	if err != nil {
		if err == store.ErrPurchaseNotFound {
			p = &purchase.Purchase{
				ChatID: chatID,
				Price:  5.0,
				// Label:  fmt.SprintL("%d_%d", chatID, time.Now().UnixNano()),
				Label: fmt.Sprint(chatID),
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
