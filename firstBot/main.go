package main

import (
	"Telegram/youtube"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// запрос обнвлений
func getUpdates(botURL string, offset int) ([]Update, error) {
	resp, err := http.Get(botURL + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponce
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

// ответ на обновления
func respond(botURL string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatID = update.Message.Chat.ChatID
	videoURL, err := youtube.GetLastVideo(update.Message.Text)
	// botMessage.Text = update.Message.Text
	botMessage.Text = videoURL
	if err != nil {
		return err
	}
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

// точка входа
func main() {
	botToken := "1376500281:AAFvC7D6hKamYbfRQM67qBJbKdYZfBk8VMs"
	// https://api.telegram.org/bot<token>/METHOD_NAME
	botAPI := "https://api.telegram.org/bot"
	botURL := botAPI + botToken
	offset := 0
	for {
		updates, err := getUpdates(botURL, offset)
		if err != nil {
			log.Println("Smth went wrong: ", err.Error())
		}
		for _, update := range updates {
			err = respond(botURL, update)
			offset = update.UpdateID + 1
		}
		fmt.Println(updates)
	}
}
