package main

import (
	"reflect"
)

func determineEmotion(emotionData EmotionScores) string {

	s := reflect.ValueOf(&emotionData.Scores).Elem()

	var maxEmotionValue float64 = 0
	var maxEmotionIndex int = 0

	for i := 0; i < s.NumField(); i++ {
		currentEmotionValue := s.Field(i).Interface().(float64)
		if currentEmotionValue > maxEmotionValue {
			maxEmotionValue = currentEmotionValue
			maxEmotionIndex = i
		}
	}

	// If none of the scores ar eover 50%, return "None"
	// so that the bot doesn't make any statements about
	// emotion

	if maxEmotionValue < .5 {
		return "None"
	}

	typeOfScore := s.Type()
	return typeOfScore.Field(maxEmotionIndex).Name
}
