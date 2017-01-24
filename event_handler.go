package main

import (
	"log"
)

func processMessageEvent(e Event) {

	if e.Message.Type == "text" {
		log.Println(e.Message.Text)
		replyMessage(e)

	}
}

func handleEvent(e Event) {

	switch e.Type {

	case "message":
		processMessageEvent(e)
	default:
		log.Println("Invalid Event Type")
		return
	}

}
