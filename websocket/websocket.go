package websocket

import (

	"net/http"

	"github.com/Yair0001/GPT-4o-mini-CLI/models"
	"github.com/Yair0001/GPT-4o-mini-CLI/utils"

	"github.com/gorilla/websocket"
)

// Connect establishes a WebSocket connection
func Connect(apiKey string) (*websocket.Conn, error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + apiKey},
		"OpenAI-Beta":   []string{"realtime=v1"},
	}

	conn, _, err := websocket.DefaultDialer.Dial(
		"wss://api.openai.com/v1/realtime?model=gpt-4o-mini-realtime-preview-2024-12-17",
		headers,
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// InitSession initializes the WebSocket session
func InitSession(conn *websocket.Conn) (string, error) {
	var session models.SessionCreated
	err := conn.ReadJSON(&session)
	if err != nil {
		return "", err
	}
	return session.Session.ID, nil
}

// RegisterFunctions sends function registration
func RegisterFunctions(conn *websocket.Conn, funcs []models.Function) error {
	sessUp := models.SessionUpdate{
		Type: "session.update",
		Session: struct {
			Tools []models.Function `json:"tools"`
		}{
			Tools: funcs,
		},
	}
	return utils.SendJSONMessage(conn, sessUp)
}
