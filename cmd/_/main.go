package main

import (
	"fmt"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/config"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/yoomoney"
)

func main() {
	conf, err := config.ParseFile("config.yaml")
	if err != nil {
		panic(err)
	}

	ymClient := yoomoney.NewClient(conf.YooMoney.AccessToken)

	accountInfo, err := ymClient.CallAccountInfo()
	if err != nil {
		panic(err)
	}

	// Получить ссылку на форму
	var price float32 = 39.0
	u, err := ymClient.CreateFormURL(yoomoney.CreateFormOptions{
		PaymentType:  "AC",
		Receiver:     accountInfo.Account,
		QuickpayForm: "shop",

		FormComment: "Телеграм бот",
		ShortDest:   "Телеграм бот",

		Label:   "q1w2e3r4t5",
		Targets: "Оплата | Телеграм бот",
		Sum:     price,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Придет:", calculateAmount(price, "AC"))
	fmt.Println("Ссылка для транзакции:", u.OrigURL.String())
	fmt.Println("Временная ссылка для транзакции:", u.TempURL.String())
	fmt.Println()

	// Проверка платежа
	resp, err := ymClient.CallOperationHistory(&yoomoney.OperationHistoryRequest{})
	if err != nil {
		panic(err)
	}

	for _, operation := range resp.Operations {
		fmt.Println("Title:     ", operation.Title)
		fmt.Println("Amount:    ", operation.Amount)
		fmt.Println("Datetime:  ", operation.Datetime)
		fmt.Println("Label:     ", operation.Label)
		fmt.Println("Status:    ", operation.Status)
		fmt.Println("Direction: ", operation.Direction)
		fmt.Println()
	}
}

func calculateAmount(price float32, paymentMethod string) float32 {
	var a float32

	if paymentMethod == "PC" {
		a = 0.0005
		return price - price*(a/(1+a))

	} else if paymentMethod == "AC" {
		a = 0.05
		return price * (1 - a)
	}

	panic("Неизвестный способ оплаты")
}
