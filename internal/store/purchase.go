package store

import "time"

type Purchase struct {
	ID int64

	ChatID     int64
	TargetUser *User
	Label      string

	CheckPrice float32
	checkPaid  int // 1 or 0

	UnlimitCheckPrice float32
	unlimitCheckPaid  int
	unlimitCheckDate  string

	ArchivePrice float32
	archivePaid  int // 1 or 0
}

type PurchaseConfig struct {
	ChatID int64
	Label  string

	CheckPrice        float32
	UnlimitCheckPrice float32
	ArchivePrice      float32
}

func NewPurchase(conf *PurchaseConfig) *Purchase {
	return &Purchase{
		ChatID:            conf.ChatID,
		TargetUser:        nil,
		Label:             conf.Label,
		CheckPrice:        conf.CheckPrice,
		checkPaid:         0,
		UnlimitCheckPrice: conf.UnlimitCheckPrice,
		unlimitCheckPaid:  0,
		unlimitCheckDate:  "",
		ArchivePrice:      conf.ArchivePrice,
		archivePaid:       0,
	}
}

func (p *Purchase) CheckPaid() bool {
	return p.checkPaid != 0
}

func (p *Purchase) UnlimitCheckPaid() bool {
	return p.unlimitCheckPaid != 0
}

func (p *Purchase) UnlimitCheckDate() (time.Time, error) {
	return time.Parse(time.RFC3339, p.unlimitCheckDate)
}

func (p *Purchase) SetUnlimitCheckDate(d time.Time) {
	p.unlimitCheckDate = d.Format(time.RFC3339)
}

func (p *Purchase) ArchivePaid() bool {
	return p.archivePaid != 0
}

// SavePurchase - сохранение purchase в базу данных
func (store *Store) SavePurchase(p *Purchase) error {
	var targetUserID int64

	if p.TargetUser != nil {
		if err := UpdateOrSaveUser(p.TargetUser); err != nil {
			return err
		}
		targetUserID = p.TargetUser.ID
	}

	sql :=
		`INSERT INTO purchases(
			chat_id,
			target_user_id,
			label,
			check_price,
			check_paid,
			unlimit_check_price,
			unlimit_check_paid,
			unlimit_check_date,
			archive_price,
			archive_paid
		) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10 )`

	args := []interface{}{
		p.ChatID,
		targetUserID,
		p.Label,
		p.CheckPrice,
		p.checkPaid,
		p.UnlimitCheckPrice,
		p.unlimitCheckPaid,
		p.unlimitCheckDate,
		p.ArchivePrice,
		p.archivePaid,
	}

	result, err := store.db.Exec(sql, args...)
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
	var p Purchase

	sql :=
		`SELECT
			id,
			chat_id,
			target_user_id,
			label,
			check_price,
			check_paid,
			unlimit_check_price,
			unlimit_check_paid,
			unlimit_check_date,
			archive_price,
			archive_paid
		FROM purchases
		WHERE id = $1`

	var targetUserID int64

	dest := []interface{}{
		&p.ID,
		&p.ChatID,
		&targetUserID,
		&p.Label,
		&p.CheckPrice,
		&p.checkPaid,
		&p.UnlimitCheckPrice,
		&p.unlimitCheckPaid,
		&p.unlimitCheckDate,
		&p.ArchivePrice,
		&p.archivePaid,
	}

	if err := store.db.QueryRow(sql, id).Scan(dest...); err != nil {
		return nil, err
	}

	if targetUserID != 0 {
		u, err := GetUserByID(targetUserID)
		if err != nil {
			return nil, err
		}
		p.TargetUser = u
	}

	return &p, nil
}

// PurchaseByChatID - получение одного purchase по chat_id
func (store *Store) GetPurchaseByChatID(chatID int64) (*Purchase, error) {
	var p Purchase

	sql :=
		`SELECT
			id,
			chat_id,
			target_user_id,
			label,
			check_price,
			check_paid,
			unlimit_check_price,
			unlimit_check_paid,
			unlimit_check_date,
			archive_price,
			archive_paid
		FROM purchases
		WHERE chat_id = $1`

	var targetUserID int64

	dest := []interface{}{
		&p.ID,
		&p.ChatID,
		&targetUserID,
		&p.Label,
		&p.CheckPrice,
		&p.checkPaid,
		&p.UnlimitCheckPrice,
		&p.unlimitCheckPaid,
		&p.unlimitCheckDate,
		&p.ArchivePrice,
		&p.archivePaid,
	}

	if err := store.db.QueryRow(sql, chatID).Scan(dest...); err != nil {
		return nil, err
	}

	if targetUserID != 0 {
		u, err := GetUserByID(targetUserID)
		if err != nil {
			return nil, err
		}
		p.TargetUser = u
	}

	return &p, nil
}

// UpdatePurchase - обновление purchase
func (store *Store) UpdatePurchaseByID(p *Purchase) error {
	var targetUserID int64

	if p.TargetUser != nil {
		targetUserID = p.TargetUser.ID
	}

	sql :=
		`UPDATE purchases
			SET
				chat_id = $1,
				target_user_id = $2,
				label = $3,
				check_price = $4,
				check_paid = $5,
				unlimit_check_price = $6,
				unlimit_check_paid = $7,
				unlimit_check_date = $8,
				archive_price = $9,
				archive_paid = $10
			WHERE
				id = $11`

	args := []interface{}{
		p.ChatID,
		targetUserID,
		p.Label,
		p.CheckPrice,
		p.checkPaid,
		p.UnlimitCheckPrice,
		p.unlimitCheckPaid,
		p.unlimitCheckDate,
		p.ArchivePrice,
		p.archivePaid,
		p.ID,
	}

	_, err := store.db.Exec(sql, args...)
	return err
}

func (store *Store) UpdateOrSavePurchase(p *Purchase) error {
	_, err := GetPurchaseByChatID(p.ChatID)
	if err == nil {
		return UpdatePurchaseByID(p)
	}
	return SavePurchase(p)
}
