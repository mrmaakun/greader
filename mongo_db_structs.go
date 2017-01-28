package main

import ()

type User struct {
	UserId        string
	ImageUploaded bool
	ImageData     ImageInformation
	EmotionData   map[int]string
}
