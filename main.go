package main

import (
	"log"

	"github.com/av1ppp/dada-ptizza_tg-bot/config"
	"github.com/av1ppp/dada-ptizza_tg-bot/tgbot"
)

func main() {
	if err := mainInner(); err != nil {
		log.Fatal(err)
	}
}

func mainInner() error {
	conf, err := config.ParseFile("./config.yaml")
	if err != nil {
		return err
	}

	bot, err := tgbot.New(conf.TelegramBot.Token)
	if err != nil {
		return err
	}

	return bot.Start()
}
