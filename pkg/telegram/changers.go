package telegram

import (
	"MyFit/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
)

func changeName(message *tgbotapi.Message) error {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}
	defer database.Close(client, ctx, cancel)

	filter := bson.D{
		{"Username", bson.D{{"$eq", message.Chat.UserName}}},
	}

	update := bson.D{
		{"$set", bson.D{
			{"Name", message.Text},
		}},
	}

	_, err = database.UpdateOne(client, ctx, "Users", "data", filter, update)
	if err != nil {
		return err
	}

	return nil
}

func changeAge(message *tgbotapi.Message) error {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}
	defer database.Close(client, ctx, cancel)

	filter := bson.D{
		{"Username", bson.D{{"$eq", message.Chat.UserName}}},
	}

	update := bson.D{
		{"$set", bson.D{
			{"Age", message.Text},
		}},
	}

	_, err = database.UpdateOne(client, ctx, "Users", "data", filter, update)
	if err != nil {
		return err
	}

	return nil
}

func changeWeight(message *tgbotapi.Message) error {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}
	defer database.Close(client, ctx, cancel)

	filter := bson.D{
		{"Username", bson.D{{"$eq", message.Chat.UserName}}},
	}

	update := bson.D{
		{"$set", bson.D{
			{"Weight", message.Text},
		}},
	}

	_, err = database.UpdateOne(client, ctx, "Users", "data", filter, update)
	if err != nil {
		return err
	}

	return nil
}

func changeHeight(message *tgbotapi.Message) error {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}
	defer database.Close(client, ctx, cancel)

	filter := bson.D{
		{"Username", bson.D{{"$eq", message.Chat.UserName}}},
	}

	update := bson.D{
		{"$set", bson.D{
			{"Height", message.Text},
		}},
	}

	_, err = database.UpdateOne(client, ctx, "Users", "data", filter, update)
	if err != nil {
		return err
	}

	return nil
}
