package tgbot

import (
	"fmt"
	"sync"
	"time"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
)

type dialogState struct {
	ChatID        int64
	Label         string
	TargetUserURL string
	SocialNetwork store.SocialNetwork
}

var mu sync.Mutex
var dialogStates []*dialogState

func label(chatID int64) string {
	return fmt.Sprintf("%x_%x", chatID, time.Now().UnixNano())
}

func getDialogState(chatID int64) *dialogState {
	mu.Lock()
	defer mu.Unlock()

	var ds *dialogState

	for _, ds_ := range dialogStates {
		if ds_.ChatID == chatID {
			ds = ds_
		}
	}

	if ds == nil {
		ds = &dialogState{
			ChatID:        chatID,
			Label:         label(chatID),
			TargetUserURL: "",
			SocialNetwork: "",
		}

		dialogStates = append(dialogStates, ds)
	}

	return ds
}

func (ds *dialogState) Reset() {
	ds.Label = label(ds.ChatID)
	ds.TargetUserURL = ""
	ds.SocialNetwork = ""
}

// func (ds *DialogState) SaveDialogState() {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	if _, err := store.GetPurchaseByID(ds.ID); err != nil {
// 		if err == store.ErrPurchaseNotFound {
// 			if err := store.SavePurchase(ds); err != nil {
// 				return err
// 			}
// 			return nil
// 		} else {
// 			return err
// 		}
// 	}

// 	return store.UpdatePurchaseByID(ds)
// }
