package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

const baseFolderPath = "target/"
const photosFolderPath string = "photos/"
const messageFileName string = "message_1.json"

type Messages struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	SenderName  string   `json:"sender_name"`
	TimestampMs int      `json:"timestamp_ms"`
	Content     string   `json:"content"`
	Photos      []Photos `json:"photos"`
	Type        string   `json:"type"`
	IsUnsent    bool     `json:"is_unsent"`
}

type Photos struct {
	Uri               string `json:"uri"`
	CreationTimestamp int    `json:"creation_timestamp"`
}

func getOriginalPhotoName(uri string) string {
	fileName := strings.Split(uri, "/")
	return fileName[4]
}

func getFileExtensionFromFileName(fileName string) string {
	return strings.Split(fileName, ".")[1]
}

func convertUnixTimestampToIMGDateTime(photoCreationTimestamp int) string {
	parsedTime := time.Unix(int64(photoCreationTimestamp), 0)
	return fmt.Sprintf("IMG_%d%02d%02d_%02d%02d%02d",
		parsedTime.Year(), parsedTime.Month(), parsedTime.Day(),
		parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second())
}

func isFileExist(newPath string) bool {
	if _, err := os.Stat(newPath); err == nil {
		// EXIST
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		// DOES NOT EXIST
		return false
	} else {
		return true
	}
}

// [OK, but duplicated] File not foun.
func renamePhotos(originalPhotoName string, creationTimestamp int) {
	fmt.Printf("originalPhotoName: %s\n", originalPhotoName)
	fmt.Printf("with extension: %s\n", getFileExtensionFromFileName(originalPhotoName))

	newPhotoName := convertUnixTimestampToIMGDateTime(creationTimestamp)
	fmt.Printf("New name: %s\n", newPhotoName)

	originalPath := fmt.Sprintf("%s%s", photosFolderPath, originalPhotoName)
	newPath := fmt.Sprintf("%s%s.%s", photosFolderPath, newPhotoName, getFileExtensionFromFileName(originalPhotoName))

	for isFileExist(newPath) {
		creationTimestamp += 1
		newPhotoName = convertUnixTimestampToIMGDateTime(creationTimestamp)
		newPath = fmt.Sprintf("%s%s.%s", photosFolderPath, newPhotoName, getFileExtensionFromFileName(originalPhotoName))
	}

	if isFileExist(originalPath) {
		e := os.Rename(originalPath, newPath)
		if e != nil {
			log.Fatal(e)
		}
	} else {
		fmt.Printf("[Not Found]\n")
	}
}

func main() {
	jsonFile, err := os.Open(baseFolderPath + messageFileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Json file opened successfully! ")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var messages Messages
	json.Unmarshal(byteValue, &messages)

	fmt.Printf("There are %d messages in this file.\n", len(messages.Messages))

	// Loop through all the messages
	for i := 0; i < len(messages.Messages); i++ {
		toLatin := charmap.ISO8859_1.NewEncoder()
		inLatin, _, _ := transform.String(toLatin, messages.Messages[i].Content)
		fmt.Printf("%s\n", inLatin)

		// Check message type is photo
		if len(messages.Messages[i].Photos) != 0 {
			// Loop through photos, sometimes a message has more than one photo.
			for j := 0; j < len(messages.Messages[i].Photos); j++ {
				// Passing original photo name and creation timestamp to rename
				// photos.
				renamePhotos(
					getOriginalPhotoName(messages.Messages[i].Photos[j].Uri),
					messages.Messages[i].Photos[j].CreationTimestamp)
			}
		}
	}

}
