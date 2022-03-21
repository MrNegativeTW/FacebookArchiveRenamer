package main

import (
	"fmt"
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
	fmt.Printf("originalPhotoName: %s\n", originalPhotoName)
	fmt.Printf("with extension: %s\n", getFileExtensionFromFileName(originalPhotoName))

	newPhotoName := convertUnixTimestampToIMGDateTime(creationTimestamp)
	fmt.Printf("New name: %s\n", newPhotoName)

	originalPath := fmt.Sprintf("%s%s", photosFolderPath, originalPhotoName)
	newPath := fmt.Sprintf("%s%s.%s", photosFolderPath, newPhotoName, getFileExtensionFromFileName(originalPhotoName))

	for Utils.IsFileExist(newPath) {
		creationTimestamp += 1
		newPhotoName = convertUnixTimestampToIMGDateTime(creationTimestamp)
		newPath = fmt.Sprintf("%s%s.%s", photosFolderPath, newPhotoName, getFileExtensionFromFileName(originalPhotoName))
	}

	if Utils.IsFileExist(originalPath) {
		e := os.Rename(originalPath, newPath)
		if e != nil {
			log.Fatal(e)
		}
	} else {
		fmt.Printf("[Not Found]\n")
	}
}

func renamePhotosFromSingleJsonFile(messages MessagesModel.Messages) {
	// Loop through all messages
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

}
