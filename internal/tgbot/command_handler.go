package tgbot

import (
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/hack"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/store"
	"github.com/av1ppp/dada-ptizza_tg-bot/internal/tgbot/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type pictureUploadHack struct {
	filepath string
	fileID   string
}

type hackUploadHacks struct {
	id   string
	blur []pictureUploadHack
	orig []pictureUploadHack
}

func (bot *Bot) handleCommand(update *tgbotapi.Update, p *store.Purchase) {
	command := update.Message.Command()

	switch command {
	case "start":
		// ds.Reset()
		bot.Send(message.MessageStart(p.ChatID, update.Message.From.FirstName))
		bot.Send(message.MessageSelectSocialNetwork(p.ChatID))

	case "test":
		msg := tgbotapi.NewMediaGroup(p.ChatID, []interface{}{
			tgbotapi.NewInputMediaPhoto("AgACAgIAAxkDAAIEymF0QXpkR1THarhI0d8I3sjlDfPyAAKKuDEbS6ehS25Py_wVqXrwAQADAgADcwADIQQ"),
			tgbotapi.NewInputMediaPhoto("AgACAgIAAxkDAAIEy2F0QXqnIgyPyZR8YhYZJFmSU2wtAAI2uTEbS6ehS5K8ixwtvobIAQADAgADcwADIQQ"),
			tgbotapi.NewInputMediaPhoto("AgACAgIAAxkDAAIEzGF0QXpl6ThYhWeN3fIUAAGIYafh3AACirgxG0unoUtuT8v8Fal68AEAAwIAA3MAAyEE"),
			tgbotapi.NewInputMediaPhoto("AgACAgIAAxkDAAIEzWF0QXoxkwuwA0b_xQABj6Tdz2MjFAACNrkxG0unoUuSvIscLb6GyAEAAwIAA3MAAyEE"),
		})
		bot.Send(msg)

	case "uploadHacks":
		var c tgbotapi.Chattable

		hacks, err := hack.Parse()
		if err != nil {
			bot.sendRequestError(p.ChatID, err)
			return
		}

		uploadHacks := []hackUploadHacks{}

		for _, h := range hacks {
			uploadHack := hackUploadHacks{
				id:   h.ID,
				blur: []pictureUploadHack{},
				orig: []pictureUploadHack{},
			}

			for _, picture := range h.BlurFilepaths {
				c = tgbotapi.NewPhotoUpload(p.ChatID, picture)
				msg, err := bot.Send(c)
				if err != nil {
					bot.sendRequestError(p.ChatID, err)
					return
				}
				bot.DeleteMessage(tgbotapi.NewDeleteMessage(p.ChatID, msg.MessageID))
				uploadHack.blur = append(uploadHack.blur, pictureUploadHack{
					filepath: picture,
					fileID:   (*msg.Photo)[0].FileID,
				})
			}

			for _, picture := range h.OrigFilepaths {
				c = tgbotapi.NewPhotoUpload(p.ChatID, picture)
				msg, err := bot.Send(c)
				if err != nil {
					bot.sendRequestError(p.ChatID, err)
					return
				}
				bot.DeleteMessage(tgbotapi.NewDeleteMessage(p.ChatID, msg.MessageID))
				uploadHack.orig = append(uploadHack.orig, pictureUploadHack{
					filepath: picture,
					fileID:   (*msg.Photo)[0].FileID,
				})
			}

			uploadHacks = append(uploadHacks, uploadHack)
		}

		/// send response message
		text := "```\n"
		for _, h := range uploadHacks {
			text += "hack_id: " + h.id + "\n\n"
			text += "blur_files:" + "\n"
			for _, p := range h.blur {
				text += "\tfile_id: " + p.fileID + "\n"
				text += "\tfilepath: " + p.filepath + "\n\n"
			}
			text += "orig_files:" + "\n"
			for _, p := range h.orig {
				text += "\tfile_id: " + p.fileID + "\n"
				text += "\tfilepath: " + p.filepath + "\n\n"
			}
			text += "\n\n"
		}
		text += "```"

		msg := tgbotapi.NewMessage(p.ChatID, text)
		msg.ParseMode = "MarkdownV2"
		if _, err = bot.Send(msg); err != nil {
			bot.sendRequestError(p.ChatID, err)
		}

	default:
		bot.Send(message.MessageUnknownCommand(p.ChatID))
	}
}
