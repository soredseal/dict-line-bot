package line

import (
	"encoding/json"
	"testing"
)

func TestWebhookRequest_GetWord(t *testing.T) {
	t.Run("from text contain 1 word", func(t *testing.T) {
		var body []byte
		body = []byte(`{
    "events": [
        {
            "type": "message",
            "replyToken": "5a36f792d61e43c9a027289feec6eeca",
            "source": {
                "userId": "U2d46bcd5a990b1ba8528829eae526238",
                "type": "user"
            },
            "timestamp": 1539531121658,
            "message": {
                "type": "text",
                "id": "8717484597020",
                "text": "line"
            }
        }
    ]
}`)
		expected := "line"
		var request WebhookRequest
		json.Unmarshal(body, &request)
		actual, _ := request.GetWord()

		if expected != actual {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("from sentence", func(t *testing.T) {
		var body []byte
		body = []byte(`{
    "events": [
        {
            "type": "message",
            "replyToken": "5a36f792d61e43c9a027289feec6eeca",
            "source": {
                "userId": "U2d46bcd5a990b1ba8528829eae526238",
                "type": "user"
            },
            "timestamp": 1539531121658,
            "message": {
                "type": "text",
                "id": "8717484597020",
                "text": "pick up something"
            }
        }
    ]
}`)
		expected := "pick"
		var request WebhookRequest
		json.Unmarshal(body, &request)
		actual, _ := request.GetWord()

		if expected != actual {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})
	
	t.Run("from multiple lines", func(t *testing.T) {
		var body []byte
		body = []byte(`{
    "events": [
        {
            "type": "message",
            "replyToken": "3e9b19ffe0d64f1fb293cfe975cf5e2d",
            "source": {
                "userId": "U2d46bcd5a990b1ba8528829eae526238",
                "type": "user"
            },
            "timestamp": 1539593240620,
            "message": {
                "type": "text",
                "id": "8720203178242",
                "text": "hello I'm me\nyes sir!"
            }
        }
    ]
}`)
		expected := "hello"
		var request WebhookRequest
		json.Unmarshal(body, &request)
		actual, _ := request.GetWord()

		if expected != actual {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("from message type image", func(t *testing.T) {
		var body []byte
		body = []byte(`{
    "events": [
        {
            "type": "message",
            "replyToken": "c8dec63907314e98b8ff1011921ca795",
            "source": {
                "userId": "U2d46bcd5a990b1ba8528829eae526238",
                "type": "user"
            },
            "timestamp": 1539591531747,
            "message": {
                "type": "image",
                "id": "8720085114115"
            }
        }
    ]
}`)
		var request WebhookRequest
		json.Unmarshal(body, &request)
		_, err := request.GetWord()

		if err == nil {
			t.Fatalf("Expected errors but got nil")
		}
	})

	t.Run("from text contains special character", func(t *testing.T) {
		var body []byte
		body = []byte(`{
    "events": [
        {
            "type": "message",
            "replyToken": "a83c7375fec64be18da1f5e9f18cdbf3",
            "source": {
                "userId": "U2d46bcd5a990b1ba8528829eae526238",
                "type": "user"
            },
            "timestamp": 1539591560403,
            "message": {
                "type": "text",
                "id": "8720087042476",
                "text": "(please!)(please!)(please!)(please!)"
            }
        }
    ]
}`)
		var request WebhookRequest
		json.Unmarshal(body, &request)
		_, err := request.GetWord()

		if err == nil {
			t.Fatalf("Expected errors but got nil")
		}
	})
}
