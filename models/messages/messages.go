package messages

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
