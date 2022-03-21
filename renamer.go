package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	MessagesModel "github.com/mrnegativetw/FacebookArchivePhotosRenamer/models/messages"
	Utils "github.com/mrnegativetw/FacebookArchivePhotosRenamer/utils"
)

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

// [OK, but duplicated] File not foun.
func renamePhotos(originalPhotoName string, creationTimestamp int) {
	// fmt.Printf("originalPhotoName: %s\n", originalPhotoName)
	// fmt.Printf("with extension: %s\n", getFileExtensionFromFileName(originalPhotoName))

	originalPath := fmt.Sprintf("%s%s%s",
		baseFolderPath,
		photosFolderPath,
		originalPhotoName)

	newPhotoName := convertUnixTimestampToIMGDateTime(creationTimestamp)
	newPath := fmt.Sprintf("%s%s%s.%s",
		baseFolderPath,
		photosFolderPath,
		newPhotoName,
		getFileExtensionFromFileName(originalPhotoName))

	// Check is file name duplicated, if so add timestamp by 1 sec.
	for Utils.IsFileExist(newPath) {
		creationTimestamp += 1
		newPhotoName = convertUnixTimestampToIMGDateTime(creationTimestamp)
		newPath = fmt.Sprintf("%s%s%s.%s",
			baseFolderPath,
			photosFolderPath,
			newPhotoName,
			getFileExtensionFromFileName(originalPhotoName))
	}

	// Rename photo.
	if Utils.IsFileExist(originalPath) {
		e := os.Rename(originalPath, newPath)
		fmt.Printf("[OK] %s\n", newPhotoName)
		if e != nil {
			log.Fatal(e)
		}
	} else {
		fmt.Printf("[Not Found] %s\n", originalPath)
	}
}

func renamePhotosFromSingleJsonFile(messages MessagesModel.Messages) {
	// Loop through all messages.
	for i := 0; i < len(messages.Messages); i++ {

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

func renamePhotosFromAllJsonFile() {
	jsonFileCount := 1

	filePath := fmt.Sprintf("%smessage_%d.json", baseFolderPath, jsonFileCount)

	// Loop through all json files.
	for Utils.IsFileExist(filePath) {
		jsonFile, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s has opend successfully!\n", filePath)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var messages MessagesModel.Messages
		json.Unmarshal(byteValue, &messages)

		renamePhotosFromSingleJsonFile(messages)

		jsonFileCount++
		filePath = fmt.Sprintf("%smessage_%d.json", baseFolderPath, jsonFileCount)
	}
}
