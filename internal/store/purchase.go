package store

import "github.com/av1ppp/dada-ptizza_tg-bot/internal/purchase"

/*
	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
	chat_id INTEGER UNIQUE,
	price REAL,
	social_network TEXT,
	target_user TEXT,
	label TEXT
*/

// SavePurchase - сохранение purchase в базу данных
func (store *Store) SavePurchase(p *purchase.Purchase) error {
	_, err := store.db.Exec(
		`INSERT INTO purchases
			(id, chat_id, price, social_network, target_user, label)
			VALUSE ($1, $2, $3, $4, $5, $6);`,
		p.ID, p.ChatID, p.Price, string(p.SocicalNetwork), p.TargetUser, p.Label)
	return err
}

// DeletePurchase - удаление purchase
func (store *Store) DeletePurchase(id int64) error {
	_, err := store.db.Exec("DELETE FROM purchases WHERE id = $1;", id)
	return err
}

// Purchase - получение одного purchase
func (store *Store) Purchase(id int64) (*purchase.Purchase, error) {
	row := store.db.QueryRow(
		`SELECT id, chat_id, price, social_network, target_user, label
			FROM purchases
			WHERE id = $1`, id)

	var p purchase.Purchase

	if err := row.Scan(&p.ID, &p.ChatID, &p.Price, &p.SocicalNetwork, &p.TargetUser, &p.Label); err != nil {
		return nil, err
	}

	return &p, nil
}

// PurchaseByChatID - получение одного purchase по chat_id
func (store *Store) PurchaseByChatID(chatID int64) (*purchase.Purchase, error) {
	row := store.db.QueryRow(
		`SELECT id, chat_id, price, social_network, target_user, label
			FROM purchases
			WHERE chat_id = $1`, chatID)

	var p purchase.Purchase

	if err := row.Scan(&p.ID, &p.ChatID, &p.Price, &p.SocicalNetwork, &p.TargetUser, &p.Label); err != nil {
		return nil, err
	}

	return &p, nil
}

// Purchases - получение всех purchases
func (store *Store) Purchases() ([]*purchase.Purchase, error) {
	purchases := []*purchase.Purchase{}

	result, err := store.db.Query("SELECT id, chat_id, price, social_network, target_user, label FROM purchases;")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var p purchase.Purchase
		var socnet string

		err := result.Scan(&p.ID, &p.ChatID, &p.Price, &socnet, &p.TargetUser, &p.Label)
		if err != nil {
			return nil, err
		}

		p.SocicalNetwork = purchase.SocialNetwork(socnet)

		purchases = append(purchases, &p)
	}

	return purchases, nil
}

// UpdatePurchase - обновление purchase
func (store *Store) UpdatePurchase(p *purchase.Purchase) error {
	_, err := store.db.Exec(
		`UPDATE purchases
			SET chat_id = $2,
				price = $3,
				social_network = $4,
				target_user = $5,
				label = $6
			WHERE
				id = $1;`,
		p.ID, p.ChatID, p.Price, p.SocicalNetwork, p.TargetUser, p.Label)

	return err
}
