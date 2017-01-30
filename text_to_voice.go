package main

import (
	"bytes"
	"log"
)

func convertToVoice(textSlice []string) string {

	var buffer bytes.Buffer

	for _, text := range textSlice {
		buffer.WriteString(text)
	}

	speechData, err := textToSpeechApi(buffer.String())

	if err != nil {
		log.Println("Error calling textToSpeechApi")
	}

	audioFilePath, err := saveAudio(speechData)

	if err != nil {
		log.Println("Error calling saving audio")
	}

	log.Println("Speech Text: " + buffer.String())

	return audioFilePath

}
