package api

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"net/http"
	"strings"
)

func AnswerGPT(message *tgbotapi.Message) (string, error) {
	url := "https://chatgpt-ai-chat-bot.p.rapidapi.com/ask"

	userReq := fmt.Sprintf("{\n    \"query\": \"" + message.Text + "\"\n}")
	payload := strings.NewReader(userReq)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "85e69f9465msh9b8963f3d02fc11p1fbbc9jsn64207893b164")
	req.Header.Add("X-RapidAPI-Host", "chatgpt-ai-chat-bot.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	uBody, err := unmarshalGPT(body)
	if err != nil {
		return "", err
	}

	return uBody.Response, nil
}

func unmarshalGPT(data []byte) (ChatGPT, error) {
	var c ChatGPT
	err := json.Unmarshal(data, &c)
	if err != nil {
		return ChatGPT{
			ConversationID: "",
			Response:       "",
		}, err
	}

	return c, nil
}

// c
type ChatGPT struct {
	ConversationID string `json:"conversationId"`
	Response       string `json:"response"`
}
