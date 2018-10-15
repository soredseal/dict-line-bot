package bot

import (
	"github.com/sariakra/line-dictionary-bot/pkg/bot"
	"log"
)

type LineConfig struct {
	BaseUrl string
	AccessToken string
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
		log.Println(data)
	}
}

