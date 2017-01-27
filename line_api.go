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

func httpRequest(method string, url string, payload []byte) (*http.Response, error) {

	var req = &http.Request{}
	var err error

	if payload != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	log.Println(os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// Throw an error if the response is over 400
	if resp.StatusCode >= 400 {
		err = errors.New("ERROR Status Code is " + strconv.Itoa(resp.StatusCode))
	}

	return resp, err
}

func replyMessage(e Event, message string) {

	log.Println("Entered reply message")

	replyMessage := ReplyMessage{
		Type: "text",
		Text: message,
	}

	reply := Reply{
		SendReplyToken: e.ReplyToken,
		Messages:       []ReplyMessage{replyMessage},
	}

	jsonPayload, err := json.Marshal(reply)

	url := apiEndpoint + "message/reply"
	resp, err := httpRequest("POST", url, jsonPayload)
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

	resp, err := httpRequest("GET", url, nil)

	if resp == nil {
		log.Println("resp is nil")
	}

	log.Println("Request URL: " + url)
	log.Println("Response Code: " + strconv.Itoa(resp.StatusCode))

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
