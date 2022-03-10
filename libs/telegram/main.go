package telegram

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Telegram struct {
	Api    string
	ChatId string
}

func NewTelegram(api, chatId string) *Telegram {
	return &Telegram{
		Api:    api,
		ChatId: chatId,
	}
}

func (t *Telegram) SendMessage(message interface{}) {

	var telegramApi string = fmt.Sprintf("https://api.telegram.org/bot%s", t.Api+"/sendMessage")

	_, telegramErr := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {t.ChatId},
			"text":    {fmt.Sprintf("%v", message)},
		})

	if telegramErr != nil {
		log.Printf(telegramErr.Error())
		return
	}
}
