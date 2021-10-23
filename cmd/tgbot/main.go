package main

import (
	"log"

	tgbotApp "github.com/av1ppp/dada-ptizza_tg-bot/internal/app"
)

func main() {
	if err := mainInner(); err != nil {
		log.Fatal(err)
	}
}

func mainInner() error {
	app, err := tgbotApp.New()
	if err != nil {
		return err
	}

	return app.Do()
}
