package main

import ()

type Message struct {
	Id        string  `json:"id,omitempty"`
	Type      string  `json:"type,omitempty"`
	Text      string  `json:"text,omitempty"`
	PackageId string  `json:"packageId,omitempty"`
	StickerId string  `json:"stickerId,omitempty"`
	Title     string  `json:"title,omitempty"`
	Address   string  `json:"address,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

type Event struct {
	ReplyToken string  `json:"replyToken,omitempty"`
	Type       string  `json:"type,omitempty"`
	Timestamp  int64   `json:"timestamp,omitempty"`
	Source     Source  `json:"source,omitempty"`
	Message    Message `json:"message,omitempty"`
}

type Source struct {
	Type    string `json:"type,omitempty"`
	UserId  string `json:"userid,omitempty"`
	GroupId string `json:"groupId,omitempty"`
	RoomId  string `json:"roomId,omitempty"`
}

type Reply struct {
	SendReplyToken string         `json:"replyToken,omitempty"`
	Messages       []ReplyMessage `json:"messages,omitempty"`
}

type ReplyMessage struct {
	Type               string  `json:"type,omitempty"`
	Text               string  `json:"text,omitempty"`
	OriginalContentUrl string  `json:"originalContentUrl,omitempty"`
	PreviewImageUrl    string  `json:"previewImageUrl,omitempty"`
	PackageId          string  `json:"packageId,omitempty"`
	StickerId          string  `json:"stickerId,omitempty"`
	Duration           string  `json:"duration,omitempty"`
	Title              string  `json:"title,omitempty"`
	Address            string  `json:"address,omitempty"`
	Latitude           float32 `json:"latitude,omitempty"`
	Longitude          float32 `json:"longitude,omitempty"`
	BaseUrl            string  `json:"baseUrl,omitempty"`
	AltText            string  `json:"altText,omitempty"`
}

type Profile struct {
	DisplayName   string `json:"displayName,omitempty"`
	UserId        string `json:"userId,omitempty"`
	PictureUrl    string `json:"pictureUrl,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
}

/*
type ReplyMessage struct {
	Type               string            `json:"type,omitempty"`
	Text               string            `json:"text,omitempty"`
	OriginalContentUrl string            `json:"originalContentUrl,omitempty"`
	PreviewImageUrl    string            `json:"previewImageUrl,omitempty"`
	PackageId          string            `json:"packageId,omitempty"`
	StickerId          string            `json:"stickerId,omitempty"`
	Duration           string            `json:"duration,omitempty"`
	Title              string            `json:"title,omitempty"`
	Address            string            `json:"address,omitempty"`
	Latitude           float32           `json:"latitude,omitempty"`
	Longitude          float32           `json:"longitude,omitempty"`
	BaseUrl            string            `json:"baseUrl,omitempty"`
	AltText            string            `json:"altText,omitempty"`
	BaseSize           ImagemapBaseSize  `json:"baseSize,omitempty"`
	Actions            []ImagemapActions `json:"actions,omitempty"`
	Template           Template          `json:"template,omitempty"`
}
*/
