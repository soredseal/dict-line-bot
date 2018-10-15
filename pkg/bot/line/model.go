package bot

import (
	"errors"
	"github.com/sariakra/line-dictionary-bot/internal"
	"strings"
)

type LineWebhookRequest struct {
	Events []LineEvent `json:"events"`
}

type LineEvent struct {
	Type string `json:"type"`
	ReplyToken string `json:"replyToken"`
	Source LineEventSource `json:"source"`
	Timestamp internal.Time `json:"timestamp"`
	Message LineEventMessage `json:"message"`
}

type LineEventSource struct {
	UserId string `json:"userId"`
	SourceType string `json:"type"`
}

type LineEventMessage struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Text string `json:"text"`
}

func (w LineWebhookRequest) GetWord() (string, error) {
	if w.Events[0].Message.Type != "text" {
		return "", errors.New("message should be text")
	}

	text := w.Events[0].Message.Text
	word := strings.Trim(strings.Split(strings.Split(strings.Split(text, "\n")[0], " ")[0], ",")[0], " ")
	if strings.ContainsAny(word, "!@#$%^&*()_+`~[]{}\\|/,'\":;<>?-=1234567890") {
		return "", errors.New("message should not contain special characters")
	}
	return word, nil
}

func (w LineWebhookRequest) GetReplyToken() string {
	return w.Events[0].ReplyToken
}

type LineReply struct {
	ReplyToken string `json:"replyToken"`
	Messages []LineMessage `json:"messages"`
}

type LineMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}