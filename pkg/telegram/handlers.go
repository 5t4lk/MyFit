package telegram

import (
	"MyFit/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var commandSwitcher int

const (
	commandStart      = "start"
	commandMembership = "membership"
	commandTrainings  = "trainings"
	commandConsult    = "consult"
	commandQ          = "q"
	commandProfile    = "profile"
	commandName       = "name"
	commandAge        = "age"
	commandWeight     = "weight"
	commandHeight     = "height"
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
	case commandName:
		return b.handleCommandName(message)
	case commandAge:
		return b.handleCommandAge(message)
	case commandWeight:
		return b.handleCommandWeight(message)
	case commandHeight:
		return b.handleCommandHeight(message)
	default:
		return b.handleUnknownCommand(message)
	}

	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch commandSwitcher {
	case 1:
		msgGPT, err := api.AnswerGPT(message)
		if err != nil {
			return err
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, msgGPT)

		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}
	case 2:
		err := changeName(message)
		if err != nil {
			return err
		}

		commandSwitcher = 0

		msg := tgbotapi.NewMessage(message.Chat.ID, displayNameUserTwo())

		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}
	case 3:
		err := changeAge(message)
		if err != nil {
			return err
		}

		commandSwitcher = 0

		msg := tgbotapi.NewMessage(message.Chat.ID, displayAgeUserTwo())

		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}
	case 4:
		err := changeWeight(message)
		if err != nil {
			return err
		}

		commandSwitcher = 0

		msg := tgbotapi.NewMessage(message.Chat.ID, displayWeightUserTwo())

		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}
	case 5:
		err := changeHeight(message)
		if err != nil {
			return err
		}

		commandSwitcher = 0

		msg := tgbotapi.NewMessage(message.Chat.ID, displayHeightUserTwo())

		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) handleCommandName(message *tgbotapi.Message) error {
	commandSwitcher = 2

	msg := tgbotapi.NewMessage(message.Chat.ID, displayNameUser())

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleCommandAge(message *tgbotapi.Message) error {
	commandSwitcher = 3

	msg := tgbotapi.NewMessage(message.Chat.ID, displayAgeUser())

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleCommandWeight(message *tgbotapi.Message) error {
	commandSwitcher = 4

	msg := tgbotapi.NewMessage(message.Chat.ID, displayWeightUser())

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleCommandHeight(message *tgbotapi.Message) error {
	commandSwitcher = 5

	msg := tgbotapi.NewMessage(message.Chat.ID, displayHeightUser())

	_, err := b.bot.Send(msg)
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

func (b *Bot) handleCommandConsult(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "\xE2\x9C\x8BHi! How can i help you? *сonsultant connected* ")

	commandSwitcher = 1

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

func (b *Bot) handleCommandQ(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "*consultant disconnected*")
	commandSwitcher = 0

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

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know such a command!")

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
