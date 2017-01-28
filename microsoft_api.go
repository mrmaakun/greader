package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type VisionApiRequest struct {
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
		Age           int    `json:"age,omitempty"`
		Gender        string `json:"gender,omitempty"`
		FaceRectangle struct {
			Left   int `json:"left,omitempty"`
			Top    int `json:"top,omitempty"`
			Width  int `json:"width,omitempty"`
			Height int `json:"height,omitempty"`
		} `json:"faceRectangle,omitempty"`
	} `json:"faces,omitempty"`
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

	log.Println(body)

	// Close the Body after using. (Find a better way to do this later)
	defer resp.Body.Close()
	json.Unmarshal(body, &returnImageInformation)

	return returnImageInformation, nil

}
