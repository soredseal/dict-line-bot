package bot

import (
	"bytes"
	"encoding/json"
	"github.com/sariakra/line-dictionary-bot/pkg/bot"
	"github.com/sariakra/line-dictionary-bot/pkg/dictionary"
	"io"
	"log"
	"net/http"
)

func (l LineClient) Handler(dictClient dictionary.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body LineWebhookRequest
		var buff bytes.Buffer
		io.Copy(&buff, request.Body)
		err := json.NewDecoder(&buff).Decode(&body)
		replyToken := body.GetReplyToken()
		if err != nil || request.Method != http.MethodPost {
			log.Println(err)
			l.Reply(bot.Request{
				Token:    replyToken,
				Messages: []string{"invalid request"},
			})
			return
		}

		word, err := body.GetWord()
		if err != nil {
			log.Println(err)
			l.Reply(bot.Request{
				Token:    replyToken,
				Messages: []string{err.Error()},
			})
			return
		}

		dictClient.Fetch(dictionary.Request{
			Word:  word,
			Token: replyToken,
		})
	}
}
