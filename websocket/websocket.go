package websocket

import (
	"net/http"

	"github.com/Yair0001/GPT-4o-mini-CLI/models"
	"github.com/Yair0001/GPT-4o-mini-CLI/utils"

	"github.com/gorilla/websocket"
)

// Connect establishes a WebSocket connection using the provided API key
func Connect(apiKey string) (*websocket.Conn, error) {
	// Set up headers with authorization and OpenAI beta information
	headers := http.Header{
		"Authorization": []string{"Bearer " + apiKey},
		"OpenAI-Beta":   []string{"realtime=v1"},
	}

	// Dial the WebSocket server with the specified URL and headers
	conn, _, err := websocket.DefaultDialer.Dial(
		"wss://api.openai.com/v1/realtime?model=gpt-4o-mini-realtime-preview-2024-12-17",
		headers,
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// InitSession initializes the WebSocket session and returns the session ID
func InitSession(conn *websocket.Conn) (string, error) {
	var session models.SessionCreated
	// Read the JSON response from the WebSocket connection into the session struct
	err := conn.ReadJSON(&session)
	if err != nil {
		return "", err
	}
	return session.Session.ID, nil
}

// RegisterFunctions sends a function registration message over the WebSocket connection
func RegisterFunctions(conn *websocket.Conn, funcs []models.Function) error {
	// Create a session update message with the provided functions
	sessUp := models.SessionUpdate{
		Type: "session.update",
		Session: struct {
			Tools []models.Function `json:"tools"`
		}{
			Tools: funcs,
		},
	}
	// Send the session update message as JSON over the WebSocket connection
	return utils.SendJSONMessage(conn, sessUp)
}
