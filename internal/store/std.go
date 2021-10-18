package store

import (
	"log"
	"sync"
)

type std struct {
	store *Store
	once  sync.Once
}

var stdImpl = std{}

func stdStore() *Store {
	stdImpl.once.Do(func() {
		store, err := New()
		if err != nil {
			log.Fatalf("creating store error: %s", err)
		}
		stdImpl.store = store
	})
	return stdImpl.store
}

func SaveUser(user *User) error {
	return stdStore().SaveUser(user)
}

func GetUserByID(id int64) (*User, error) {
	return stdStore().GetUserByID(id)
}

func GetUserByURL(url string) (*User, error) {
	return stdStore().GetUserByURL(url)
}

func DeleteUserByID(id int64) error {
	return stdStore().DeleteUserByID(id)
}

func SavePurchase(purchase *Purchase) error {
	return stdStore().SavePurchase(purchase)
}

func UpdateUserByID(u *User) error {
	return stdStore().UpdateUserByID(u)
}

func UpdateOrSaveUser(u *User) error {
	return stdStore().UpdateOrSaveUser(u)
}

func DeletePurchase(id int64) error {
	return stdStore().DeletePurchase(id)
}

func GetPurchaseByID(id int64) (*Purchase, error) {
	return stdStore().GetPurchaseByID(id)
}

func GetPurchaseByChatID(chatID int64) (*Purchase, error) {
	return stdStore().GetPurchaseByChatID(chatID)
}

func UpdatePurchaseByID(p *Purchase) error {
	return stdStore().UpdatePurchaseByID(p)
}

func UpdateOrSavePurchase(p *Purchase) error {
	return stdStore().UpdateOrSavePurchase(p)
}
