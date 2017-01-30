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

	typeOfScore := s.Type()

	// Translate cateogry to proper language

	switch typeOfScore.Field(maxEmotionIndex).Name {

	case "Anger":
		return "angry"
	case "Contempt":
		return "like they have hate in their eyes"
	case "Disgust":
		return "disgusted"
	case "Fear":
		return "scared"
	case "Happiness":
		return "happy"
	case "Sadness":
		return "sad"
	case "Surprise":
		return "surprised"
	default:
		return ""
	}

}
