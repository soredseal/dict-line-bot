package bot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/sariakra/line-dictionary-bot/pkg/bot"
	"github.com/sariakra/line-dictionary-bot/pkg/dictionary"
	"io/ioutil"
	"log"
	"net/http"
)

func (l LineClient) Handler(dictClient dictionary.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body LineWebhookRequest
		var buff *bytes.Buffer

		if request.Body == nil {
			log.Println("invalid request")
			return
		}

		defer request.Body.Close()
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println("invalid request")
			return
		}
		decoded, err := base64.StdEncoding.DecodeString(request.Header.Get("X-Line-Signature"))
		if err != nil {
			log.Println("invalid request")
			return
		}
		hash := hmac.New(sha256.New, []byte(l.config.Secret))
		hash.Write(data)
		if !hmac.Equal(decoded, hash.Sum(nil)) {
			log.Println("invalid request")
			return
		}

		buff = bytes.NewBuffer(data)
		err = json.NewDecoder(buff).Decode(&body)
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
