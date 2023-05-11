package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron"
	"log"
	"math/rand"
	"time"
)

func scheduleMotivationalMessages(bot *tgbotapi.BotAPI, chatID int64) {
	c := cron.New()

	c.AddFunc("0 10 * * *", func() {
		motivationalText := generateMotivationalText()

		if err := sendMotivationalMessage(bot, chatID, motivationalText); err != nil {
			log.Printf("failed to send notification: %v", err)
		}
	})

	c.Start()
}

func sendMotivationalMessage(bot *tgbotapi.BotAPI, chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(msg)
	return err
}

func generateMotivationalText() string {
	messages := []string{
		"Believe in yourself and your abilities! You are stronger than you think.",
		"Your efforts today will bring you results tomorrow. Don't give up!",
		"Нет ничего невозможного, если у тебя есть решимость и настойчивость.",
		"Nothing is impossible if you have determination and perseverance.",
		"Be a true warrior and push your limits. Nothing can stop you!",
		"Your success depends only on you. Don't stop, go ahead!",
		"Remember that every day is a new opportunity to become better, stronger, more successful.",
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(messages))
	return messages[index]
}
