package main

type SessionCreated struct {
	Type    string `json:"type"`
	EventID string `json:"event_id"`
	Session struct {
		ID        string `json:"id"`
		ExpiresAt int64  `json:"expires_at"`
	} `json:"session"`
}

type SessionUpdate struct{
	Type string `json:"type"`
	Session struct{
		Tools []Function `json:"tools"`
	}`json:"session"`
}

type Function struct{
	Type string `json:"type"`
	Name string `json:"name"`
	Description string `json:"description"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct{
	Type string `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required []string `json:"required"`
}

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

type ModelResponseCreate struct {
	Type string `json:"type"`
	Response Response `json:"response"`
}

type Response struct {
	Modalities   []string `json:"modalities"`
	Instructions string   `json:"instructions"`
}

type ModelResponseRecieve struct{
	Type string `json:"type"`
	Response Response
}

type ModelResponseTextDelta struct{
	Type string `json:"type"`
	Delta string `json:"delta"`
}

type ModelResponseFunction struct{
	Type string `json:"type"`
	Arguments string `json:"arguments"`
}