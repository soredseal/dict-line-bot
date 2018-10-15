package main

import (
	"github.com/sariakra/line-dictionary-bot/pkg/bot/line"
	"github.com/sariakra/line-dictionary-bot/pkg/dictionary/oxford"
	"net/http"
	"os"
)

func main() {
	appId := os.Getenv("OXFORD_APP_ID")
	appKey := os.Getenv("OXFORD_APP_KEY")
	lineAccessToken := os.Getenv("LINE_ACCESS_TOKEN")
	lineClient := bot.NewLineClient(bot.LineConfig{
		BaseUrl:     "https://api.line.me/v2",
		AccessToken: lineAccessToken,
	})
	oxfordClient := dictionary.NewOxfordClient(dictionary.OxfordConfig{
		BaseUrl: "https://od-api.oxforddictionaries.com/api/v1",
		AppId:   appId,
		AppKey:  appKey,
	}, lineClient)
	http.Handle("/api/webhook", lineClient.Handler(oxfordClient))
	http.ListenAndServe(":8090", nil)
}
