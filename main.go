package main

import (
	"log"
	"time"

	"bingai-bot/config"
	"bingai-bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			go func(update tgbotapi.Update) {
				if update.Message != nil {
					handlers.HandleCommands(bot, update)
				} else if update.CallbackQuery != nil {
					handlers.HandleButtons(bot, update.CallbackQuery)
				}
			}(update)
		}
	}()

	select {} // keep alive
}
