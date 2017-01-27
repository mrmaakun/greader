package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func downloadImage(imageId string) (string, error) {

	resp, err := contentDownload(imageId)
	if err != nil {
		return "", err
	}

	// Save image file

	imageFileName := "image_" + strconv.Itoa(rand.Intn(10000)) + ".jpg"
	newFile, err := os.Create("images/" + imageFileName)

	numBytesWritten, err := io.Copy(newFile, resp.Body)
	if err != nil {
		log.Println("Error download file")
		log.Println(err.Error())
		return "", err
	}

	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	log.Println("File name: " + imageFileName)

	return os.Getenv("BASE_HOSTNAME") + imageFileName, nil
}
