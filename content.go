package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	//"os/exec"
	"strconv"
	"time"
)

func downloadImage(imageId string) (string, error) {

	// Call the content download API to get the image
	resp, err := contentDownload(imageId)
	if err != nil {
		return "", err
	}

	// Save image file
	imageFileName := "image_" + strconv.Itoa(rand.Intn(10000)) + ".jpg"
	newFile, err := os.Create("images/" + imageFileName)

	numBytesWritten, err := io.Copy(newFile, resp.Body)
	if err != nil {
		log.Println("Error downloading image file")
		log.Println(err.Error())
		return "", err
	}

	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	log.Println("File name: " + imageFileName)

	// Delete the oldest
	cleanMediaDirectory("images")

	return os.Getenv("BASE_HOSTNAME") + "/images/" + imageFileName, nil

}

func saveAudio(audioData []byte) (string, error) {

	// Used this buildpack to install FFMPEG:
	// https://elements.heroku.com/buildpacks/jonathanong/heroku-buildpack-ffmpeg-latest

	// Save image file
	audioFileName := "audio_" + strconv.Itoa(rand.Intn(10000))
	newFile, err := os.Create("audio/" + audioFileName + ".m4a")

	numBytesWritten, err := io.Copy(newFile, bytes.NewReader(audioData))
	if err != nil {
		log.Println("Error downloading audio file")
		log.Println(err.Error())
		return "", err
	}

	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	log.Println("File name: " + audioFileName)

	/*
		cmd := "ffmpeg"
		args := []string{"-i", "audio/" + audioFileName + ".mp3", "-c:a", "libfdk_aac", "audio/" + audioFileName + ".m4a"}
		if err := exec.Command(cmd, args...).Run(); err != nil {
			os.Exit(1)
		}
		log.Println("converted mp3 to m4a")

	*/

	// Delete the oldest
	cleanMediaDirectory("audio")

	return os.Getenv("BASE_HOSTNAME") + "/audio/" + audioFileName + ".m4a", nil

}

func downloadAudio(audioId string) (string, error) {

	// Call the content download API to get the audio
	resp, err := contentDownload(audioId)
	if err != nil {
		return "", err
	}

	// Save image file
	audioFileName := "audio_" + strconv.Itoa(rand.Intn(10000)) + ".m4a"
	newFile, err := os.Create("audio/" + audioFileName)

	numBytesWritten, err := io.Copy(newFile, resp.Body)
	if err != nil {
		log.Println("Error downloading audio file")
		log.Println(err.Error())
		return "", err
	}

	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	log.Println("File name: " + audioFileName)

	// Delete the oldest
	cleanMediaDirectory("audio")

	return os.Getenv("BASE_HOSTNAME") + "/audio/" + audioFileName, nil

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
