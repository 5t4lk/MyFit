package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	commandStart      = "start"
	commandMembership = "membership"
	commandTrainings  = "trainings"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleCommandStart(message)
	case commandMembership:
		return b.handleMembershipCommand(message)
	case commandTrainings:
		return b.handleTrainingsCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}

	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.bot.Send(msg)
}

func (b *Bot) handleCommandStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, displayStart())

	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know such a command!")

	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleMembershipCommand(message *tgbotapi.Message) error {
	msgPic := tgbotapi.NewPhotoUpload(message.Chat.ID, displayMembershipPic())
	msg := tgbotapi.NewMessage(message.Chat.ID, displayMembershipText())

	_, err := b.bot.Send(msgPic)
	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) handleTrainingsCommand(message *tgbotapi.Message) error {
	msgPic := tgbotapi.NewPhotoUpload(message.Chat.ID, displayTrainingsPic())
	msg := tgbotapi.NewMessage(message.Chat.ID, displayTrainingsText())

	_, err := b.bot.Send(msgPic)
	_, err = b.bot.Send(msg)

	return err
}
