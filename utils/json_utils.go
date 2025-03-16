package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/Yair0001/GPT-4o-mini-CLI/models"
)

// SendJSONMessage sends a JSON message via WebSocket
func SendJSONMessage(conn *websocket.Conn, data interface{}) error {
	// Marshal the data into a JSON byte slice
	jsonMsg, err := json.Marshal(data)
	if err != nil {
		// Print and return the error if JSON marshaling fails
		fmt.Println("JSON error:", err)
		return err
	}
	// Write the JSON message to the WebSocket connection
	return conn.WriteMessage(websocket.TextMessage, jsonMsg)
}

// SendUserMessage sends a user message through the WebSocket connection
func SendUserMessage(conn *websocket.Conn, text string) error {
	// Create a user message with the provided text
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
	
	// Send the user message as a JSON message via WebSocket
	return SendJSONMessage(conn, userMsg)
}

// SendModelResponse sends a model response through the WebSocket connection
func SendModelResponse(conn *websocket.Conn) error {
	// Create a model response message
	modelMsg := models.ModelResponseCreate{
		Type: "response.create",
		Response: models.Response{
			Modalities: []string{"text"},
		},
	}
	
	// Send the model response as a JSON message via WebSocket
	return SendJSONMessage(conn, modelMsg)
}