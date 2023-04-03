package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
)

type User struct {
	age           int
	gender        string
	height        int
	weight        int
	activityLevel string
}

func displayStart() string {
	message := `
	You can control me by sending these commands:

	/start - start a bot
	/membership - view membership prices
	/trainings - see training plans and prices
`
	return message
}

func displayMembershipPic() tgbotapi.FileBytes {
	photoBytes, _ := ioutil.ReadFile("/Users/5t4lk/GolandProjects/MyFit/pics/membership.png")
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}
	return photoFileBytes
}

func displayMembershipText() string {
	message := `
	To buy membership, you need to...
`
	return message
}

func displayTrainingsPic() tgbotapi.FileBytes {
	photoBytes, _ := ioutil.ReadFile("/Users/5t4lk/GolandProjects/MyFit/pics/training.png")
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}
	return photoFileBytes
}

func displayTrainingsText() string {
	message := `
	To buy training plan, you need to...
`
	return message
}
