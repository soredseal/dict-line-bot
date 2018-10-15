package bot

import (
	"errors"
	"github.com/sariakra/line-dictionary-bot/internal"
	"strings"
)

type WebhookRequest struct {
	Events []Event `json:"events"`
}

type Event struct {
	Type string `json:"type"`
	ReplyToken string `json:"replyToken"`
	Source EventSource `json:"source"`
	Timestamp internal.Time `json:"timestamp"`
	Message EventMessage `json:"message"`
}

type EventSource struct {
	UserId string `json:"userId"`
	SourceType string `json:"type"`
}

type EventMessage struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Text string `json:"text"`
}

func (w WebhookRequest) GetWord() (string, error) {
	if w.Events[0].Message.Type != "text" {
		return "", errors.New("message should be text")
	}

	text := w.Events[0].Message.Text
	word := strings.Split(strings.Split(text, "\n")[0], " ")[0]

	if strings.ContainsAny(word, "!@#$%^&*()_+`~[]{}\\|/,'\":;<>?-=1234567890") {
		return "", errors.New("message should not contain special characters")
	}
	return word, nil
}

func (w WebhookRequest) GetReplyToken() string {
	return w.Events[0].ReplyToken
}

type Reply struct {
	ReplyToken string `json:"replyToken"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}