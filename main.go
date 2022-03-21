package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	MessageModel "github.com/mrnegativetw/FacebookArchivePhotosRenamer/models/messages"
)

const baseFolderPath string = "target/"
const photosFolderPath string = "photos/"
const messageFileName string = "message_11.json"

func main() {
	jsonFile, err := os.Open(baseFolderPath + messageFileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s opened successfully! \n", messageFileName)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var messages MessageModel.Messages
	json.Unmarshal(byteValue, &messages)

	// Uncomment the code below to run the feature you need.

	// 1. Print all messages from single file.
	// Utils.Viewer{}.PrintMessage(messages)

	// 2. Print all messages from single file with timestamp and name.
	// Utils.Viewer{}.PrintMessageDetails(messages)

	// 3. Calc total messages.
	// totalMessageCount := Utils.Calculator{}.CalculateTotalMessage(baseFolderPath)
	// fmt.Printf("Total message count: %d\n", totalMessageCount)

	// 4. Rename photos from single json file.
	renamePhotosFromSingleJsonFile(messages)

	// 5. Rename all photos from all json files. (Recommend)
	// renamePhotosFromAllJsonFile()
}
