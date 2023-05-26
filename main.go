package main

import (
	"bot/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	botToken := "6201257452:AAEGCw0nev74KNk21CoVBo3JI7FZXPqKl8o"

	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken

	offset := 0

	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println("ошибка, какая-то херня", err)
		}
		for _, update := range updates {

			err = respond(botUrl, update)
			if err != nil {
				log.Println("ошибка, какая-то херня", err)
			}
			offset = update.UpdateId + 1
		}
		fmt.Println(updates)
	}
}

func getUpdates(botUrl string, offset int) ([]models.Update, error) {
	reply, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer reply.Body.Close()

	body, err := ioutil.ReadAll(reply.Body)
	if err != nil {
		return nil, err
	}

	var restResponse models.RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil

}

func respond(botUrl string, update models.Update) error {
	var botMessage models.BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text

	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))

	if err != nil {
		return err
	}
	return nil
}
