package purchase

type Purchase struct {
	ID             int64
	ChatID         int64         // чат с покупателем
	Price          float32       // цена приобретения
	SocicalNetwork SocialNetwork // соц. сеть
	TargetUser     string        // цель
	Label          string        // идентификатор
}

func NewPurchase(ChatID int64) *Purchase {
	return &Purchase{
		ChatID: ChatID,
		Price:  39.0,
	}
}
