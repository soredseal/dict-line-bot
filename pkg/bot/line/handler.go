package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sariakra/line-dictionary-bot/pkg/dictionary"
	"io"
	"log"
	"net/http"
)

func (l LineClient) Handler(dictClient dictionary.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body WebhookRequest
		var buff bytes.Buffer
		io.Copy(&buff, request.Body)
		err := json.NewDecoder(&buff).Decode(&body)
		if err != nil || request.Method != http.MethodPost {
			log.Println(err, buff.String())
			writer.Write([]byte(`{"error":"invalid request"}`))
			return
		}

		word, err := body.GetWord()
		if err != nil {
			log.Println(err, buff.String())
			writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
			return
		}

		replyToken := body.GetReplyToken()
		dictClient.Fetch(dictionary.Request{
			Word: word,
			Token: replyToken,
		})
	}
}