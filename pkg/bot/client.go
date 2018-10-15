package bot

import (
	"github.com/sariakra/line-dictionary-bot/pkg/dictionary"
	"net/http"
)

type Request struct {
	Messages []string
	Token string
}

type Client interface {
	Reply(request Request)
	Handler(dictClient dictionary.Client) http.HandlerFunc
}

