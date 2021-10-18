package store

import (
	"database/sql"
)

type PurchaseSQL struct {
	ID           int64
	ChatID       int64
	Price        float32
	TargetUserID int64
	Label        string
}

type Purchase struct {
	ID         int64
	ChatID     int64
	Price      float32
	TargetUser *User
	Label      string
}

// SavePurchase - сохранение purchase в базу данных
func (store *Store) SavePurchase(p *Purchase) error {
	if p.TargetUser != nil {
		if err := UpdateOrSaveUser(p.TargetUser); err != nil {
			return err
		}
	}

	sql := PurchaseSQL{
		ChatID: p.ChatID,
		Price:  p.Price,
		Label:  p.Label,
	}
	if p.TargetUser != nil {
		sql.TargetUserID = p.TargetUser.ID
	}

	result, err := store.db.Exec(
		`INSERT INTO purchases(chat_id, price, target_user_id, label)
			VALUES ($1, $2, $3, $4);`,
		sql.ChatID, sql.Price, sql.TargetUserID, sql.Label)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = lastID
	return nil
}

// DeletePurchase - удаление purchase
func (store *Store) DeletePurchase(id int64) error {
	_, err := store.db.Exec("DELETE FROM purchases WHERE id = $1", id)
	return err
}

// Purchase - получение одного purchase
func (store *Store) GetPurchaseByID(id int64) (*Purchase, error) {
	row := store.db.QueryRow(
		`SELECT id, chat_id, price, target_user_id, label
			FROM purchases
			WHERE id = $1`, id)

	var p Purchase

	if err := row.Scan(&p.ID, &p.ChatID, &p.Price, &p.TargetUserID, &p.Label); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPurchaseNotFound
		}
		return nil, err
	}

	return &p, nil
}

// PurchaseByChatID - получение одного purchase по chat_id
func (store *Store) GetPurchaseByChatID(chatID int64) (*Purchase, error) {
	row := store.db.QueryRow(
		`SELECT id, chat_id, price, target_user_id, label
			FROM purchases
			WHERE chat_id = $1`, chatID)

	var p Purchase

	if err := row.Scan(&p.ID, &p.ChatID, &p.Price, &p.TargetUserID, &p.Label); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPurchaseNotFound
		}
		return nil, err
	}

	return &p, nil
}

// Purchases - получение всех purchases
// func (store *Store) AllPurchases() ([]*Purchase, error) {
// 	purchases := []*purchase.Purchase{}

// 	result, err := store.db.Query("SELECT id, chat_id, price, social_network, target_user, label FROM purchases;")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer result.Close()

// 	for result.Next() {
// 		var p purchase.Purchase
// 		var socnet string

// 		err := result.Scan(&p.ID, &p.ChatID, &p.Price, &socnet, &p.TargetUser, &p.Label)
// 		if err != nil {
// 			return nil, err
// 		}

// 		p.SocicalNetwork = purchase.SocialNetwork(socnet)

// 		purchases = append(purchases, &p)
// 	}

// 	return purchases, nil
// }

// UpdatePurchase - обновление purchase
func (store *Store) UpdatePurchaseByID(p *Purchase) error {
	_, err := store.db.Exec(
		`UPDATE purchases
			SET
				chat_id = $1,
				price = $2,
				target_user_id = $3,
				label = $4
			WHERE
				id = $5`,
		p.ChatID, p.Price, p.TargetUserID, p.Label, p.ID)

	return err
}

func (store *Store) UpdateOrSavePurchase(p *Purchase) error {
	_, err := GetPurchaseByChatID(p.ChatID)
	if err == nil {
		return UpdatePurchaseByID(p)
	}
	return SavePurchase(p)
}
