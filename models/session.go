package models

type SessionCreated struct {
	Type    string `json:"type"`
	EventID string `json:"event_id"`
	Session struct {
		ID        string `json:"id"`
		ExpiresAt int64  `json:"expires_at"`
	} `json:"session"`
}

type SessionUpdate struct {
	Type    string `json:"type"`
	Session struct {
		Tools []Function `json:"tools"`
	} `json:"session"`
}
