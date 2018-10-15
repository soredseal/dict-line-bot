package dictionary

import (
	"encoding/json"
	"errors"
	"github.com/sariakra/line-dictionary-bot/internal"
	"github.com/sariakra/line-dictionary-bot/pkg/bot"
	"github.com/sariakra/line-dictionary-bot/pkg/dictionary"
	"net/http"
	"strings"
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
		client:    bot,
	}
	go producer.fetchFromOxford()
	return producer
}

func (o OxfordClient) Fetch(request dictionary.Request) {
	o.input <- request
}

func (o OxfordClient) fetchFromOxford() {
	for data := range o.input {
		definitions, err := o.getDefinitions(data)
		if err != nil {
			o.reply(data.Token, err.Error())
			continue
		}
		replyDefinition := definitions[0]

		synonyms, err := o.getSynonyms(data)
		if err != nil {
			o.reply(data.Token, replyDefinition, err.Error())
			continue
		}
		replySynonym := strings.Join(synonyms, ", ")
		replySynonym = internal.ReplaceRight(replySynonym, ", ", " and ", 1)

		o.reply(data.Token, replyDefinition, replySynonym)
	}
}

func (o OxfordClient) getDefinitions(request dictionary.Request) ([]string, error) {
	resp, err := o.get("/entries/en/" + request.Word)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New("no definitions for this word")
	}

	var response OxfordResponse
	json.NewDecoder(resp.Body).Decode(&response)

	definitions := response.GetDefinitions()
	if len(definitions) == 0 {
		return nil, errors.New("no definitions for this word")
	}

	return definitions, nil
}

func (o OxfordClient) getSynonyms(request dictionary.Request) ([]string, error) {
	resp, err := o.get("/entries/en/" + request.Word + "/synonyms")
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New("no synonyms for this word")
	}

	var response OxfordResponse
	json.NewDecoder(resp.Body).Decode(&response)

	synonyms := response.GetSynonyms()
	if len(synonyms) == 0 {
		return nil, errors.New("no synonyms for this word")
	}

	return synonyms, nil
}

func (o *OxfordClient) get(path string) (*http.Response, error) {
	u, err := internal.MakeUrl(o.config.BaseUrl, path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("app_id", o.config.AppId)
	req.Header.Set("app_key", o.config.AppKey)

	return http.DefaultClient.Do(req)
}

func (o OxfordClient) reply(token string, messages ...string) {
	botReq := bot.Request{
		Token:    token,
		Messages: messages,
	}
	o.client.Reply(botReq)
}
