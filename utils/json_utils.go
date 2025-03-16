package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/Yair0001/GPT-4o-mini-CLI/models"
)

// SendJSONMessage sends a JSON message via WebSocket
func SendJSONMessage(conn *websocket.Conn, data interface{}) error {
	jsonMsg, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON error:", err)
		return err
	}
	return conn.WriteMessage(websocket.TextMessage, jsonMsg)
}

// SendUserMessage sends a user message through the websocket connection

func SendUserMessage(conn *websocket.Conn, text string) error {
	userMsg := models.UserMessage{
		Type: "conversation.item.create",
		Item: struct {
			Type    string `json:"type"`
			Role    string `json:"role"`
			Content []struct {
				ContentType string `json:"type"`
				Text        string `json:"text"`
			} `json:"content"`
		}{
			Type: "message",
			Role: "user",
			Content: []struct {
				ContentType string `json:"type"`
				Text        string `json:"text"`
			}{{ContentType: "input_text", Text: text}},
		},
	}
	
	return SendJSONMessage(conn, userMsg)
}

func SendModelResponse(conn *websocket.Conn) error {
	modelMsg := models.ModelResponseCreate{
		Type: "response.create",
		Response: models.Response {
			Modalities: []string{"text"},
		},
	}
	return SendJSONMessage(conn, modelMsg)
}