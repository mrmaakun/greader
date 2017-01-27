package main

import (
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
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

	// Delete the oldest
	cleanMediaDirectory("images")

	return os.Getenv("BASE_HOSTNAME") + "/images/" + imageFileName, nil
}

func cleanMediaDirectory(dirName string) {

	// This method deletes the oldest file if the directory has greater than 30 files

	// Get a slice of files in the images directory
	files, _ := ioutil.ReadDir(dirName)

	numberOfStoredImages := len(files)

	// TODO: Change the max number of stored images to a config item
	if numberOfStoredImages > 30 {

		var earliestModifiedTime time.Time
		var earliestModifiedFileName string

		for _, f := range files {

			// Ignore file if it is a directory
			if f.IsDir() == true {
				continue
			}

			// If this is the first element, set it as the earliest one
			if earliestModifiedFileName == "" {
				earliestModifiedTime = f.ModTime()
				earliestModifiedFileName = f.Name()
				continue
			}

			if earliestModifiedTime.Before(f.ModTime()) {
				earliestModifiedTime = f.ModTime()
				earliestModifiedFileName = f.Name()
			}
		}

		err := os.Remove(earliestModifiedFileName)
		if err != nil {
			log.Fatal(err)
		}
	}
}
