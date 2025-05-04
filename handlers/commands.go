package handlers

import (
	"bingai-bot/utils"
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := update.Message
	text := msg.Text

	if strings.HasPrefix(text, "/bingai") {
		prompt := strings.TrimPrefix(text, "/bingai ")
		if len(prompt) < 5 {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Tolong masukkan deskripsi yang lebih panjang."))
			return
		}
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Sedang membuat gambar, mohon tunggu..."))
		imageURL, err := utils.GenerateImage(prompt)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Gagal generate gambar: %v", err)))
			return
		}
		photo := tgbotapi.NewPhotoUpload(msg.Chat.ID, nil)
		photo.FileID = imageURL
		photo.UseExisting = false
		photo.File = tgbotapi.FileURL(imageURL)
		photo.Caption = "Berikut hasil dari prompt: " + prompt
		bot.Send(photo)
		return
	}

	switch msg.Command() {
	case "start":
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Hai, selamat datang!"))
		time.Sleep(2 * time.Second)
		bot.Send(tgbotapi.NewChatAction(msg.Chat.ID, tgbotapi.Typing))
		time.Sleep(2 * time.Second)

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("Owner", "https://t.me/Lampurge"),
				tgbotapi.NewInlineKeyboardButtonData("Help", "help"),
			),
		)
		msg := tgbotapi.NewMessage(msg.Chat.ID, "Pilih opsi:")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)

	case "help":
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Bingai", "bingai_help"),
				tgbotapi.NewInlineKeyboardButtonData("Textgen", "textgen_help"),
			),
		)
		msg := tgbotapi.NewMessage(msg.Chat.ID, "Pilih cara penggunaan:")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)

	default:
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Perintah tidak dikenali."))
	}
}

func HandleButtons(bot *tgbotapi.BotAPI, cb *tgbotapi.CallbackQuery) {
	switch cb.Data {
	case "help":
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Bingai", "bingai_help"),
				tgbotapi.NewInlineKeyboardButtonData("Textgen", "textgen_help"),
			),
		)
		msg := tgbotapi.NewMessage(cb.Message.Chat.ID, "Pilih cara penggunaan:")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)

	case "bingai_help":
		bot.Send(tgbotapi.NewMessage(cb.Message.Chat.ID, "Cara menggunakan /bingai:\nKetik `/bingai cewek anime di taman` untuk hasil gambar.\nGunakan bahasa Inggris untuk hasil terbaik."))

	case "textgen_help":
		bot.Send(tgbotapi.NewMessage(cb.Message.Chat.ID, "Fitur /textgen belum tersedia.")) // placeholder
	}
}
