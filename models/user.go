package models

type UserMessage struct {
	Type string `json:"type"`
	Item struct {
		Type    string `json:"type"`
		Role    string `json:"role"`
		Content []struct {
			ContentType string `json:"type"`
			Text        string `json:"text"`
		} `json:"content"`
	} `json:"item"`
}