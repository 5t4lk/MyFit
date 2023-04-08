package telegram

import (
	"MyFit/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	messages config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, messages config.Messages) *Bot {
	return &Bot{bot: bot, messages: messages}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		switch {
		case update.Message == nil:
			continue
		case update.Message.IsCommand():
			if err := b.handleCommand(update.Message); err != nil {
				log.Printf("error: %s", err)
			}
		default:
			if err := b.handleMessage(update.Message); err != nil {
				log.Printf("error: %s", err)
			}
		}
	} // c
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
