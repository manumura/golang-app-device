package channel

// Channel type
type Channel struct {
	ID          int    `json:"distChannelId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
