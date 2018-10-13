package main

import (
	"github.com/sariakra/line-dictionary-bot/pkg/line"
	"net/http"
)

func main() {
	http.Handle("/api/webhook", line.WebhookHandler())
	http.ListenAndServe(":8090", nil)
}
