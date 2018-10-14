package line

import "github.com/sariakra/line-dictionary-bot/internal"

type WebhookRequest struct {
	Events []Event `json:"events"`
}

type Event struct {
	EventType string `json:"type"`
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
	MessageType string `json:"type"`
	Id string `json:"id"`
	Text string `json:"text"`
}