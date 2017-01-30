package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/textproto"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func CreateAudioFormFile(w *multipart.Writer, filename string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
	h.Set("Content-Type", "audio/x-m4a")
	return w.CreatePart(h)
}

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

	buf := new(bytes.Buffer)

	newFile, err := os.Create("audio/" + audioFileName + ".mp3")

	numBytesWritten, err := io.Copy(newFile, bytes.NewReader(audioData))
	if err != nil {
		log.Println("Error downloading audio file")
		log.Println(err.Error())
		return "", err
	}

	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	log.Println("File name: " + audioFileName)

	file, _ := os.Open(audioFileName + ".mp3")
	writer := multipart.NewWriter(buf)
	audioFile, _ := CreateAudioFormFile(writer, "audio/"+audioFileName+".mp3")
	io.Copy(audioFile, file)
	writer.Close()

	cmd1 := "ffmpeg"
	args1 := []string{"-i", "audio/" + audioFileName + ".mp3", "-c", "copy", "audio/output.mp3"}
	if err := exec.Command(cmd1, args1...).Run(); err != nil {
		log.Println("Error downloading audio file")
		log.Println(err.Error())
		return "", err
	}

	cmd := "ffmpeg"
	args := []string{"-i", "audio/output.mp3", "-c:a", "aac -strict experimental", "audio/" + audioFileName + ".m4a"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Println("Error downloading audio file")
		log.Println(err.Error())
		return "", err
	}
	log.Println("converted mp3 to m4a")

	// Delete the oldest
	cleanMediaDirectory("audio")

	return os.Getenv("BASE_HOSTNAME") + "/audio/" + audioFileName + ".mp3", nil

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
