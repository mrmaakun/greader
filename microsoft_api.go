package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type FaceRectangle struct {
	Height int `json:"height,omitempty"`
	Left   int `json:"left,omitempty"`
	Top    int `json:"top,omitempty"`
	Width  int `json:"width,omitempty"`
}

type EmotionScores struct {
	FaceRectangle FaceRectangle `json:"faceRectangle"`
	Scores        struct {
		Anger     float64 `json:"anger"`
		Contempt  float64 `json:"contempt"`
		Disgust   float64 `json:"disgust"`
		Fear      float64 `json:"fear"`
		Happiness float64 `json:"happiness"`
		Neutral   float64 `json:"neutral"`
		Sadness   float64 `json:"sadness"`
		Surprise  float64 `json:"surprise"`
	}
}

type VisionApiRequest struct {
	Url string `json:"url,omitempty"`
}

type EmotionApiRequest struct {
	Url string `json:"url,omitempty"`
}

type ImageInformation struct {
	Categories []struct {
		Name  string  `json:"name,omitempty"`
		Score float64 `json:"score,omitempty"`
	} `json:"categories,omitempty"`
	Description struct {
		Tags     []string `json:"tags,omitempty"`
		Captions []struct {
			Text       string  `json:"text,omitempty"`
			Confidence float64 `json:"confidence,omitempty"`
		} `json:"captions,omitempty"`
	} `json:"description,omitempty"`
	RequestID string `json:"requestId,omitempty"`
	Metadata  struct {
		Width  int    `json:"width,omitempty"`
		Height int    `json:"height,omitempty"`
		Format string `json:"format,omitempty"`
	} `json:"metadata,omitempty"`
	Faces []struct {
		Age           int           `json:"age,omitempty"`
		Gender        string        `json:"gender,omitempty"`
		FaceRectangle FaceRectangle `json:"faceRectangle,omitempty"`
	} `json:"faces,omitempty"`
}

type EmotionInformation []struct {
	Face   FaceRectangle `json:"faceRectangle"`
	Scores struct {
		Anger     float64 `json:"anger"`
		Contempt  float64 `json:"contempt"`
		Disgust   float64 `json:"disgust"`
		Fear      float64 `json:"fear"`
		Happiness float64 `json:"happiness"`
		Neutral   float64 `json:"neutral"`
		Sadness   float64 `json:"sadness"`
		Surprise  float64 `json:"surprise"`
	}
}

func visionApi(imageUrl string) (ImageInformation, error) {

	var headers = map[string]string{
		"Ocp-Apim-Subscription-Key": os.Getenv("VISION_API_KEY"),
		"Content-Type":              "application/json",
	}

	requestParameters := VisionApiRequest{
		Url: imageUrl,
	}

	var returnImageInformation ImageInformation = ImageInformation{}

	jsonPayload, err := json.Marshal(requestParameters)
	if err != nil {
		log.Println("Error unmarshalling message: " + err.Error())
		return ImageInformation{}, err
	}

	url := "https://westus.api.cognitive.microsoft.com/vision/v1.0/analyze?visualFeatures=Categories,Faces,Description&language=en"

	resp, err := httpRequest("POST", url, headers, jsonPayload)

	if err != nil {
		log.Println("Error calling the vision api")
		return ImageInformation{}, err
	}

	// Read body into bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error getting profile")
		return ImageInformation{}, err
	}

	log.Println("Image Information JSON:")
	log.Println(string(body))

	// Close the Body after using. (Find a better way to do this later)
	defer resp.Body.Close()
	json.Unmarshal(body, &returnImageInformation)

	return returnImageInformation, nil
}

func emotionApi(imageUrl string) ([]EmotionScores, error) {

	var headers = map[string]string{
		"Ocp-Apim-Subscription-Key": os.Getenv("EMOTION_API_KEY"),
		"Content-Type":              "application/json",
	}

	requestParameters := EmotionApiRequest{
		Url: imageUrl,
	}

	var returnEmotionInformation []EmotionScores

	jsonPayload, err := json.Marshal(requestParameters)
	if err != nil {
		log.Println("Error unmarshalling message: " + err.Error())
		return []EmotionScores{}, err
	}

	url := "https://westus.api.cognitive.microsoft.com/emotion/v1.0/recognize"

	resp, err := httpRequest("POST", url, headers, jsonPayload)

	if err != nil {
		log.Println("Error calling the vision api")
		return []EmotionScores{}, err
	}

	// Read body into bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error getting profile")
		return []EmotionScores{}, err
	}

	log.Println("Image Information JSON:")
	log.Println(string(body))

	// Close the Body after using. (Find a better way to do this later)
	defer resp.Body.Close()
	json.Unmarshal(body, &returnEmotionInformation)

	return returnEmotionInformation, nil

}

func textToSpeechApi(text string) ([]byte, error) {

	// Get Access Token
	url := "https://api.cognitive.microsoft.com/sts/v1.0/issueToken"

	var headers = map[string]string{
		"Ocp-Apim-Subscription-Key": os.Getenv("SPEECH_API_KEY"),
	}

	resp, err := httpRequest("POST", url, headers, nil)

	if err != nil {
		log.Println("Error calling the speech API")
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error calling the speech API")
		return []byte{}, err
	}

	speechAccessToken := string(body)

	log.Println("Speech Access Token: " + speechAccessToken)

	url = "https://speech.platform.bing.com/synthesize?OutputFormat=Audio16khz128kbitrateMonoMp3"

	headers = map[string]string{
		"Authorization":            "Bearer " + speechAccessToken,
		"Content-Type":             "application/ssml+xml",
		"X-Microsoft-OutputFormat": "audio-16khz-16kbps-mono-siren",
	}

	payload := []byte("<speak version='1.0' xml:lang='en-US'><voice xml:lang='en-US' xml:gender='Female' name='Microsoft Server Speech Text to Speech Voice (en-US, ZiraRUS)'>" + text + "</voice></speak>")

	resp, err = httpRequest("POST", url, headers, payload)

	if err != nil {
		log.Println("Error calling the speech API")
		return []byte{}, err
	}

	returnData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error calling the speech API")
		return []byte{}, err
	}

	return returnData, nil
}
