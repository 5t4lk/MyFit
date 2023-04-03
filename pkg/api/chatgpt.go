package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://chatgpt-ai-chat-bot.p.rapidapi.com/ask"

	payload := strings.NewReader("{\n    \"query\": \"What is google?\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "85e69f9465msh9b8963f3d02fc11p1fbbc9jsn64207893b164")
	req.Header.Add("X-RapidAPI-Host", "chatgpt-ai-chat-bot.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
