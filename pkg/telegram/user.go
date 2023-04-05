package telegram

import (
	"MyFit/internal/database"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
)

type User struct {
	username         string  `json:"username" bson:"username"`
	balance          float64 `json:"balance" bson:"balance"`
	firstAndLastName string  `json:"firstAndLastName" bson:"firstAndLastName"`
	age              int     `json:"age" bson:"age"`
	weight           string  `json:"weight" bson:"weight"`
	height           string  `json:"height" bson:"height"`
}

func InitDatabase(message *tgbotapi.Message) error {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
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
			username:         message.Chat.UserName,
			balance:          0.00,
			firstAndLastName: "",
			age:              0,
			weight:           "",
			height:           "",
		}

		doc = bson.D{
			{"Username", user.username},
			{"Balance", user.balance},
			{"First and last name", user.firstAndLastName},
			{"Age", user.age},
			{"Weight", user.weight},
			{"Height", user.height},
		}

		_, err := database.InsertOne(client, ctx, "Users", "data", doc)
		if err != nil {
			return err
		}

		log.Printf("User signed: %s", user.username)
	}

	return nil
}

func RequestUserInfo(message *tgbotapi.Message) (string, error) {
	client, ctx, _, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return "", err
	}

	err = database.Ping(client, ctx)
	if err != nil {
		return "", err
	}

	var filter, option interface{}

	filter = bson.D{
		{"Username", message.Chat.UserName},
	}

	cursor, err := database.Query(client, ctx, "Users", "data", filter, option)
	if err != nil {
		return "", err
	}

	var results []bson.D

	if err := cursor.All(ctx, &results); err != nil {
		return "", err
	}

	err = MarshalBSON(results)
	if err != nil {
		return "", err
	}

	var user User

	user = User{
		username:         user.username,
		balance:          user.balance,
		firstAndLastName: user.firstAndLastName,
		age:              user.age,
		weight:           user.weight,
		height:           user.height,
	}
	msg := fmt.Sprintf(`
	Username: %s,
	Balance: %.2f,
	First and last name: %s,
	Age: %d,
	Weight: %s,
	Height: %s
`, user.username, user.balance, user.firstAndLastName, user.age, user.weight, user.height)

	return msg, nil
}

func displayStart(message *tgbotapi.Message) string {
	msg := fmt.Sprintf(`
	Hello, %s!

	You can control me by sending these commands:

	/start - start a bot
	/profile - check my profile
	/membership - view membership prices
	/trainings - see training plans and prices
	/consult - talk with consultant
	/q - finish chat with consultant
`, message.Chat.FirstName)

	return msg
}

func displayMembershipPic() tgbotapi.FileBytes {
	photoBytes, err := ioutil.ReadFile("/Users/5t4lk/GolandProjects/MyFit/pics/membership.png")
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
	message := `
	To buy membership, you need to...
`

	return message
}

func displayTrainingsPic() tgbotapi.FileBytes {
	photoBytes, err := ioutil.ReadFile("/Users/5t4lk/GolandProjects/MyFit/pics/training.png")
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
	message := `
	To buy training plan, you need to...
`

	return message
}

func MarshalBSON(bsonData []bson.D) error {
	_, err := bson.Marshal(bsonData)
	if err != nil {
		return err
	}

	return err
}
