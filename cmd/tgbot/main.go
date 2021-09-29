package main

import (
	"log"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/config"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot"
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

// func main() {
// 	vk := vkapi.NewClient("cd5e4b3de057ae5124b4eafd730922b1481f6775cb6984cb731fe3bc1e9129ab7582e99c2994153ea8f9b")

// 	// vk.Client.UserInfo
// 	users, err := vk.UsersGet(vkapi.UsersGetParams{
// 		UserIds: "durov",
// 		Fields:  "photo_400_orig",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	if len(*users) == 0 {
// 		panic("user not found")
// 	}

// 	user := (*users)[0]

// 	fmt.Println(user.Photo400Orig)

// 	// vk.Photos

// 	// resp, err := vk.PhotosGetAll(vkapi.PhotosGetAllParams{
// 	// 	OwnerID: user.ID,
// 	// })
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// if resp.Count == 0 {
// 	// 	panic("photos not found")
// 	// }

// 	// for _, photo := range resp.Items {
// 	// 	fmt.Println(photo.)
// 	// }

// 	// vk.PhotosGet(vkapi.PhotosGetParams{
// 	// 	OwnerID: ,
// 	// })
// }
