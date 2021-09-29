package main

import (
	"log"

	tgbotApp "github.com/av1ppp/dada-ptizza_tg-bot/internal/app"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/config"
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

	app, err := tgbotApp.New(conf)
	if err != nil {
		return err
	}

	return app.Do()
}
