package dictionary

import (
	"github.com/sariakra/line-dictionary-bot/pkg/bot"
	"github.com/sariakra/line-dictionary-bot/pkg/dictionary"
	"log"
)

type OxfordConfig struct {
	BaseUrl string
	AppId   string
	AppKey  string
}

type OxfordClient struct {
	config OxfordConfig
	input  chan dictionary.Request
	client    bot.Client
}

func NewOxfordClient(config OxfordConfig, bot bot.Client) *OxfordClient {
	input := make(chan dictionary.Request, 1000)
	producer := &OxfordClient{
		config: config,
		input:  input,
		client: bot,
	}
	go producer.fetchFromOxford()
	return producer
}

func (o OxfordClient) Fetch(request dictionary.Request) {
	o.input <- request
}

func (o OxfordClient) fetchFromOxford() {
	for data := range o.input {
		log.Println(data)
	}
}

func (o OxfordClient) getDefinitions(request dictionary.Request) ([]string, error) {
	return []string{}, nil
}

func (o OxfordClient) getSynonyms(request dictionary.Request) ([]string, error) {
	return []string{}, nil
}
