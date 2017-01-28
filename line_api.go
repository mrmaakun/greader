package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func httpRequest(method string, url string, headers map[string]string, payload []byte) (*http.Response, error) {

	var req = &http.Request{}
	var err error

	log.Println("Request URL: " + url)

	if payload != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	log.Println("Response Code: " + strconv.Itoa(resp.StatusCode))

	// Throw an error if the response is over 400
	if resp.StatusCode >= 400 {
		err = errors.New("ERROR Status Code is " + strconv.Itoa(resp.StatusCode))
	}

	return resp, err
}

func replyMessage(e Event, messages []string) {

	log.Println("Entered reply message")

	outgoingMessageSlice := []ReplyMessage{}

	for _, message := range messages {
		outgoingMessageSlice = append(outgoingMessageSlice, ReplyMessage{
			Type: "text",
			Text: message,
		})
	}

	reply := Reply{
		SendReplyToken: e.ReplyToken,
		Messages:       outgoingMessageSlice,
	}

	jsonPayload, err := json.Marshal(reply)

	url := apiEndpoint + "message/reply"

	var headers = map[string]string{
		"Authorization": "Bearer " + os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
		"Content-Type":  "application/json",
	}
	resp, err := httpRequest("POST", url, headers, jsonPayload)
	if err != nil {
		log.Println("Error sending reply" + err.Error())

		return
	}
	// Close the Body after using. (Find a better way to do this later. It's kind of weird doing it in a different method)
	defer resp.Body.Close()

}

func getProfile(userId string) (Profile, error) {

	log.Println("Getting user profile")

	url := apiEndpoint + "profile/" + userId

	var headers = map[string]string{
		"Authorization": "Bearer " + os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
		"Content-Type":  "application/json",
	}

	resp, err := httpRequest("GET", url, headers, nil)

	if resp == nil {
		log.Println("resp is nil")
	}

	if err != nil {
		log.Println("Error getting profile")
		return Profile{}, err
	}

	var returnProfile Profile = Profile{}

	// Read body into bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error getting profile")
		return Profile{}, err
	}

	// Close the Body after using. (Find a better way to do this later)
	defer resp.Body.Close()
	json.Unmarshal(body, &returnProfile)

	return returnProfile, nil

}

func contentDownload(contentId string) (*http.Response, error) {

	var headers = map[string]string{
		"Authorization": "Bearer " + os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
		"Content-Type":  "application/json",
	}

	url := apiEndpoint + "message/" + contentId + "/content"
	resp, err := httpRequest("GET", url, headers, nil)
	return resp, err

}
