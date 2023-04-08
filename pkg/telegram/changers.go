package telegram

import (
	"MyFit/internal/database"
	"MyFit/pkg/params"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateField struct {
	Field string
	Value string
}

func updateUser(message *tgbotapi.Message, updateField UpdateField) error {
	client, ctx, cancel, err := database.Connect(params.ConnectMongoDB)
	if err != nil {
		return err
	}
	defer database.Close(client, ctx, cancel)

	filter := bson.D{
		{"Username", bson.D{{"$eq", message.Chat.UserName}}},
	}

	update := bson.D{
		{"$set", bson.D{
			{updateField.Field, updateField.Value},
		}},
	}

	_, err = database.UpdateOne(client, ctx, "Users", "data", filter, update)
	if err != nil {
		return err
	}

	return nil
}

func changeName(message *tgbotapi.Message) error {
	updateField := UpdateField{
		Field: "Name",
		Value: message.Text,
	}
	return updateUser(message, updateField)
}

func changeAge(message *tgbotapi.Message) error {
	updateField := UpdateField{
		Field: "Age",
		Value: message.Text,
	}
	return updateUser(message, updateField)
}

func changeWeight(message *tgbotapi.Message) error {
	updateField := UpdateField{
		Field: "Weight",
		Value: message.Text,
	}
	return updateUser(message, updateField)
}

func changeHeight(message *tgbotapi.Message) error {
	updateField := UpdateField{
		Field: "Height",
		Value: message.Text,
	}
	return updateUser(message, updateField)
}
