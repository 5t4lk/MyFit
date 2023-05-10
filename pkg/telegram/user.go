package telegram

import (
	"MyFit/internal/database"
	"MyFit/pkg/params"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
)

type User struct {
	Username string  `json:"username" bson:"username"`
	Balance  float64 `json:"balance" bson:"balance"`
	Name     string  `json:"name" bson:"name"`
	Age      string  `json:"age" bson:"age"`
	Weight   string  `json:"weight" bson:"weight"`
	Height   string  `json:"height" bson:"height"`
}

func InitDatabase(message *tgbotapi.Message) error {
	client, ctx, cancel, err := database.Connect(params.ConnectMongoDB)
	if err != nil {
		return err
	}
	defer database.Close(client, ctx, cancel)

	err = database.Ping(client, ctx)
	if err != nil {
		return err
	}

	var filter, option interface{}

	filter = bson.D{
		{"Username", message.Chat.UserName},
	}

	cursor, err := database.Query(client, ctx, "Users", "data", filter, option)
	if err != nil {
		return err
	}

	var results []bson.D

	err = cursor.All(ctx, &results)
	if err != nil {
		return err
	}

	if results == nil {
		var doc interface{}

		user := User{
			Username: message.Chat.UserName,
			Name:     "",
			Age:      "",
			Weight:   "",
			Height:   "",
		}

		doc = bson.D{
			{"Username", user.Username},
			{"Name", user.Name},
			{"Age", user.Age},
			{"Weight", user.Weight},
			{"Height", user.Height},
		}

		_, err := database.InsertOne(client, ctx, "Users", "data", doc)
		if err != nil {
			return err
		}

		log.Printf("User signed: %s", user.Username)
	}

	return nil
}

func RequestUserInfo(message *tgbotapi.Message) (string, error) {
	client, ctx, _, err := database.Connect(params.ConnectMongoDB)
	if err != nil {
		return "", err
	}

	if err = database.Ping(client, ctx); err != nil {
		return "", err
	}

	col := client.Database("Users").Collection("data")

	var result User
	if err = col.FindOne(context.TODO(), bson.M{"Username": message.Chat.UserName}).Decode(&result); err != nil {
		return "", err
	}

	msgF := fmt.Sprintf(`
	%sYour profile @%s:

	[- Username: %s
	[- Name: %s
	[- Age: %s
	[- Weight: %s
	[- Height: %s

	%sTo make changes in your profile write one of these commands below:
	/name - change your first and last names
	/age - change your age
	/weight - change your weight
	/height - change your height
`, "\xF0\x9F\x91\xA4", result.Username, result.Username, result.Name, result.Age, result.Weight, result.Height, "\xF0\x9F\x9A\xA8")

	return msgF, err
}

func displayStart(message *tgbotapi.Message) string {
	msg := fmt.Sprintf(`
	Hello, %s! %s

	%s You can control me by sending commands below:

	/start - start a bot
	/profile - check your profile
	/membership - view membership prices
	/trainings - see training plans and prices
	/consult - talk with consultant
	/q - finish chat with consultant

`, message.Chat.FirstName, "\xE2\x9C\x8B", "\xE2\x84\xB9")

	return msg
}

func displayMembershipPic() tgbotapi.FileBytes {
	photoBytes, err := ioutil.ReadFile("./pics/membership.png")
	if err != nil {
		log.Printf("picture not found/impossible to read file: %s", err)
	}

	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}

	return photoFileBytes
}

func displayMembershipText() string {
	message := fmt.Sprintf(`
	%s To buy a membership in our club, you need to contact a manager at one of the clubs that is most convenient for you.
`, "\xF0\x9F\x9A\x80")

	return message
}

func displayTrainingsPic() tgbotapi.FileBytes {
	photoBytes, err := ioutil.ReadFile("./pics/training.png")
	if err != nil {
		log.Printf("picture not found/impossible to read file: %s", err)
	}

	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}

	return photoFileBytes
}

func displayTrainingsText() string {
	message := fmt.Sprintf(`
	%s To buy a training plan in our club, you need to contact a manager at one of the clubs that is most convenient for you.
`, "\xF0\x9F\x9A\x80")

	return message
}

func displayNameUser() string {
	msg := fmt.Sprintf(`
	Write your new name %s
`, "\xE2\x9C\x92")

	return msg
}

func displayNameUserTwo() string {
	msg := fmt.Sprintf(`
	Name was successfully updated %s
`, "\xE2\x9C\x85")

	return msg
}

func displayAgeUser() string {
	msg := fmt.Sprintf(`
	Write your new age %s
`, "\xE2\x9C\x92")

	return msg
}

func displayAgeUserTwo() string {
	msg := fmt.Sprintf(`
	Age was successfully updated %s
`, "\xE2\x9C\x85")

	return msg
}

func displayWeightUser() string {
	msg := fmt.Sprintf(`
	Write your new weight %s
`, "\xE2\x9C\x92")

	return msg
}

func displayWeightUserTwo() string {
	msg := fmt.Sprintf(`
	Weight was successfully updated %s
`, "\xE2\x9C\x85")

	return msg
}

func displayHeightUser() string {
	msg := fmt.Sprintf(`
	Write your new height %s
`, "\xE2\x9C\x92")

	return msg
}

func displayHeightUserTwo() string {
	msg := fmt.Sprintf(`
	Height was successfully updated %s
`, "\xE2\x9C\x85")

	return msg
}
