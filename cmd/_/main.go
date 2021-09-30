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

	fmt.Println("Account:", accountInfo.Account)
}
