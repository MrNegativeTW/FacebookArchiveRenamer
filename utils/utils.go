package utils

import (
	"fmt"
	"time"

	MessageModel "github.com/mrnegativetw/FacebookArchivePhotosRenamer/models/messages"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type Viewer struct{}

func (util Viewer) PrintMessageDetails(messages MessageModel.Messages) {
	for i := 0; i < len(messages.Messages); i++ {
		senderName := Viewer{}.encodeToHumanReadable(messages.Messages[i].SenderName)
		timestamp := Viewer{}.convertTimestampMsToDateTime(messages.Messages[i].TimestampMs)
		content := Viewer{}.encodeToHumanReadable(messages.Messages[i].Content)
		fmt.Printf("%s <%s> %s\n", timestamp, senderName, content)
	}
}

func (util Viewer) PrintMessage(messages MessageModel.Messages) {
	for i := 0; i < len(messages.Messages); i++ {
		content := Viewer{}.encodeToHumanReadable(messages.Messages[i].Content)
		fmt.Printf("%s\n", content)
	}
}

func (util Viewer) encodeToHumanReadable(content string) string {
	toLatin := charmap.ISO8859_1.NewEncoder()
	inLatin, _, _ := transform.String(toLatin, content)
	return inLatin
}

func (util Viewer) convertTimestampMsToDateTime(timestampMs int) string {
	parsedTime := time.UnixMilli(int64(timestampMs))
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		parsedTime.Year(), parsedTime.Month(), parsedTime.Day(),
		parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second())
}
