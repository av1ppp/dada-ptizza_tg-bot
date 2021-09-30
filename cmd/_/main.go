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

	u := ymClient.CreateFormURL(yoomoney.CreateFormOptions{
		Receiver:     accountInfo.Account,
		QuickpayForm: "donate",
		FormComment:  "Проект Железный человек Long",
		ShortDest:    "Проект Железный человек Short",
		Label:        "label",
		Targets:      "На да",
		Sum:          12.0,
		Comment:      "COMMENT",
	})

	fmt.Println(u.String())
}
