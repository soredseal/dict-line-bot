package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sariakra/line-dictionary-bot/internal"
	"github.com/sariakra/line-dictionary-bot/pkg/bot"
	"log"
	"net/http"
)

type LineConfig struct {
	BaseUrl     string
	AccessToken string
	Secret      string
}

type LineClient struct {
	config LineConfig
	input  chan bot.Request
}

func NewLineClient(config LineConfig) *LineClient {
	input := make(chan bot.Request, 1000)
	producer := &LineClient{
		config: config,
		input:  input,
	}
	go producer.replyToLine()
	return producer
}

func (l LineClient) Reply(request bot.Request) {
	l.input <- request
}

func (l LineClient) replyToLine() {
	for data := range l.input {
		var messages []LineMessage
		for _, message := range data.Messages {
			messages = append(messages, LineMessage{
				Type: "text",
				Text: message,
			})
		}
		reply := LineReply{
			ReplyToken: data.Token,
			Messages:   messages,
		}

		err := l.postReply(reply)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (l LineClient) postReply(body LineReply) error {
	resp, err := l.post("/bot/message/reply", body)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("cannot reply to line: %+v", body))
	}
	return nil
}

func (l LineClient) post(path string, body interface{}) (*http.Response, error) {
	u, err := internal.MakeUrl(l.config.BaseUrl, path)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+l.config.AccessToken)
	return http.DefaultClient.Do(req)
}
