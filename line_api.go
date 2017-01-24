package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func httpRequest(method string, url string, payload []byte) (*http.Response, error) {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	req.Header.Set("Authorization", "Bearer "+os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	log.Println(os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	return resp, err
}

func replyMessage(e Event) {

	log.Println("Entered reply message")

	replyMessage := ReplyMessage{
		Type: "text",
		Text: e.Message.Text,
	}

	reply := Reply{
		SendReplyToken: e.ReplyToken,
		Messages:       []ReplyMessage{replyMessage},
	}

	jsonPayload, err := json.Marshal(reply)

	url := apiEndpoint + "message/reply"
	resp, err := httpRequest("POST", url, jsonPayload)

	log.Println("Request URL: " + url)
	log.Println("Response Code: " + resp.Status)

	if err != nil {
		log.Println("Error sending reply")
		return
	}

}
