package main

import (
	"MyFit/pkg/config"
	"MyFit/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot, cfg.Messages)
	err = telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
