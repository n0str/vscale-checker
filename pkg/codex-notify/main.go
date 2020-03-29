package codex_notify

import (
	"log"
	"monitoring/pkg/http"
	"net/url"
)

type Message struct {
	Message string `json:"message"`
}

func SendMessage(chatURL, message string) {
	data := url.Values{}
	data.Set("message", message)
	data.Set("parse_mode", "HTML")
	_, err := http.PostRequest("POST", chatURL, []byte(data.Encode()), map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	})
	if err != nil {
		log.Fatalf("Invalid response: %s", err)
	}
}
