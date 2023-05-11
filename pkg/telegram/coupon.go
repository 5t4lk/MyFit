package telegram

import (
	"MyFit/internal/database"
	"MyFit/pkg/params"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GenerateCoupon(userID int64, discount int) (string, error) {
	client, ctx, cancel, err := database.Connect(params.ConnectMongoDB)
	if err != nil {
		return "", err
	}
	defer database.Close(client, ctx, cancel)

	code := RandStringBytes(10)
	coupon := bson.M{
		"code":     code,
		"discount": discount,
		"used":     false,
		"user_id":  userID,
	}

	_, err = database.InsertOne(client, ctx, "Users", "coupons", coupon)
	if err != nil {
		return "", err
	}

	return code, nil
}
