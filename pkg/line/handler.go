package line

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func WebhookHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body WebhookRequest
		var buff bytes.Buffer
		io.Copy(&buff, request.Body)
		err := json.NewDecoder(&buff).Decode(&body)
		if err != nil || request.Method != http.MethodPost {
			log.Println(err, buff.String())
			writer.Write([]byte(`{"error":"invalid request"}`))
			return
		}

		log.Println(body)
	}
}