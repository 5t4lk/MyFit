package telegram

import (
	"MyFit/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var switcher int

const (
	commandStart      = "start"
	commandMembership = "membership"
	commandTrainings  = "trainings"
	commandConsult    = "consult"
	commandQ          = "q"
	commandProfile    = "profile"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleCommandStart(message)
	case commandMembership:
		return b.handleMembershipCommand(message)
	case commandTrainings:
		return b.handleTrainingsCommand(message)
	case commandConsult:
		return b.handleCommandConsult(message)
	case commandQ:
		return b.handleCommandQ(message)
	case commandProfile:
		return b.handleCommandProfile(message)
	default:
		return b.handleUnknownCommand(message)
	}

	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	if switcher == 0 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "I couldn't catch your message")
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

		return nil
	}

	msgGPT, err := api.AnswerGPT(message)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, msgGPT)

	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleCommandStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, displayStart(message))

	err := InitDatabase(message)
	if err != nil {
		return err
	}

	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know such a command!")

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleMembershipCommand(message *tgbotapi.Message) error {
	msgPic := tgbotapi.NewPhotoUpload(message.Chat.ID, displayMembershipPic())
	msg := tgbotapi.NewMessage(message.Chat.ID, displayMembershipText())

	_, err := b.bot.Send(msgPic)
	if err != nil {
		return err
	}

	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleTrainingsCommand(message *tgbotapi.Message) error {
	msgPic := tgbotapi.NewPhotoUpload(message.Chat.ID, displayTrainingsPic())
	msg := tgbotapi.NewMessage(message.Chat.ID, displayTrainingsText())

	_, err := b.bot.Send(msgPic)
	if err != nil {
		return err
	}

	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleCommandConsult(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hi! How can i help you? *сonsultant connected* ")

	switcher = 1

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleCommandQ(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "*consultant disconnected*")
	switcher = 0

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleCommandProfile(message *tgbotapi.Message) error {
	res, err := RequestUserInfo(message)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, res)

	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
