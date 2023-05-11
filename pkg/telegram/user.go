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

	if err := database.Ping(client, ctx); err != nil {
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

	if err = cursor.All(ctx, &results); err != nil {
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

func messageWithEmoji(message string, emoji string) string {
	return fmt.Sprintf("%s %s", message, emoji)
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
	/discount - take a survey and receive discount

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
	return messageWithEmoji(`
	To buy a membership in our club, you need to contact a manager at one of the clubs that is most convenient for you.
	`, "\xF0\x9F\x9A\x80")
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
	return messageWithEmoji(`
	To buy a training plan in our club, you need to contact a manager at one of the clubs that is most convenient for you.
	`, "\xF0\x9F\x9A\x80")
}

func printDiscountInfo() string {
	msg := fmt.Sprintf(`
	For us it's such a big pleasure that you use our application! %s
	To receive your gift%s, please just answer to some questions. We appreciate your time!
`, "\xF0\x9F\x8E\x81", "\xF0\x9F\x98\x8C")

	return msg
}

func surveyQuestionOne() string {
	msg := fmt.Sprintf(`
	%s1/3: What is your overall assessment of the application on a scale of 1 to 10? 
	Please give a number where 1 is very bad, 10 is excellent.
`, "\xE2\x9D\x94")

	commandSwitcher = 7

	return msg
}

func surveyQuestionTwo() string {
	msg := fmt.Sprintf(`
	%s2/3: What features or functionality of the app did you like best? 
	Please describe exactly what you liked and why.
`, "\xE2\x9D\x94")

	commandSwitcher = 8

	return msg
}

func surveyQuestionThree() string {
	msg := fmt.Sprintf(`
	%s3/3: Do you have any suggestions or suggestions for improving the app? 
	We value your opinion and are open to any suggestions to make our app even better.
`, "\xE2\x9D\x94")

	commandSwitcher = 9

	return msg
}

func endSurvey(message *tgbotapi.Message) string {
	couponCode, err := GenerateCoupon(message.Chat.ID, 20)
	if err != nil {
		log.Printf("error while generating coupone: %s", err)
	}

	msg := fmt.Sprintf(`
	%sThank you for taking the survey! 
	You have earned your 20 percents discount on all MyFit products!
	%sCoupon: %s
`, "\xE2\x9D\x95", "\xF0\x9F\x8E\x81", couponCode)

	commandSwitcher = 0

	return msg
}

func displayNameUser() string {
	return messageWithEmoji(`
	Write your new name.
	`, "\xE2\x9C\x92")
}

func displayNameUserTwo() string {
	msg := fmt.Sprintf(`
	Name was successfully updated %s
`, "\xE2\x9C\x85")

	return msg
}

func displayAgeUser() string {
	return messageWithEmoji(`
	Write your new age.
	`, "\xE2\x9C\x92")
}

func displayAgeUserTwo() string {
	msg := fmt.Sprintf(`
	Age was successfully updated %s
`, "\xE2\x9C\x85")

	return msg
}

func displayWeightUser() string {
	return messageWithEmoji(`
	Write your new weight.
	`, "\xE2\x9C\x92")
}

func displayWeightUserTwo() string {
	msg := fmt.Sprintf(`
	Weight was successfully updated %s
`, "\xE2\x9C\x85")

	return msg
}

func displayHeightUser() string {
	return messageWithEmoji(`
	Write your new height.
	`, "\xE2\x9C\x92")
}

func displayHeightUserTwo() string {
	msg := fmt.Sprintf(`
	Height was successfully updated %s
`, "\xE2\x9C\x85")

	return msg
}
