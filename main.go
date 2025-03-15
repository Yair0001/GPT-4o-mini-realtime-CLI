package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	// "sync"
	"github.com/gorilla/websocket"
)

var (
	conn      *websocket.Conn
	sessionID string
	// wg        sync.WaitGroup
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY not set")
		return
	}

	headers := http.Header{
		"Authorization": []string{"Bearer " + apiKey},
		"OpenAI-Beta":   []string{"realtime=v1"},
	}

	var err error
	conn, _, err = websocket.DefaultDialer.Dial(
		"wss://api.openai.com/v1/realtime?model=gpt-4o-mini-realtime-preview-2024-12-17",
		headers,
	)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close()

	// Handle initial session creation
	var session SessionCreated
	err = conn.ReadJSON(&session)
	if err != nil {
		fmt.Println("Session init error:", err)
		return
	}
	sessionID = session.Session.ID
	fmt.Printf("Session created: %s (expires at %d)\n", sessionID, session.Session.ExpiresAt)

	//Create function
	var sessUp SessionUpdate
	sessUp.Type = "session.update"
	sessUp.Session.Tools = []Function{
		{
			Type: "function",
			Name: "Multiplication",
			Description: "Multiply two numbers",
			Parameters: Parameters{
				Type: "object",
				Properties: map[string]interface{}{
					"num1": map[string]interface{}{
						"type": "integer",
					},
					"num2": map[string]interface{}{
						"type": "integer",
					},
				},
				Required: []string{"num1", "num2"},
			},
		},
	}

	jsonMsg, err := json.Marshal(sessUp)
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, jsonMsg); err != nil {
		fmt.Println("Write error:", err)
		return
	}

	for {
		userInputHandler()
		messageReceiver()
	}

}

func userInputHandler() {
	// defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\n> ")
	scanner.Scan()
	text := scanner.Text()

	if text == "exit" {
		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
		return
	}

	//send Message
	msg := UserMessage{
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
			}{
				{
					ContentType: "input_text",
					Text:        text,
				},
			},
		},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, jsonMsg); err != nil {
		fmt.Println("Write error:", err)
		return
	}

	//initiate response
	resp := ModelResponseCreate{
		Type: "response.create",
		Response: Response{
			Modalities: []string{"text"},
		},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, jsonResp); err != nil {
		fmt.Println("Write error:", err)
		return
	}

}


func messageReceiver() {
	// defer wg.Done()
	for {
		if conn == nil {
			fmt.Println("WebSocket connection is nil")
			return
		}

		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				fmt.Println("Read error:", err)
			}
			return
		}

		// Print the raw JSON message
		// fmt.Printf("Raw message: %s\n", string(message))

		var baseMsg struct{ Type string }
		if err := json.Unmarshal(message, &baseMsg); err != nil {
			fmt.Println("JSON parse error:", err)
			continue
		}

		switch baseMsg.Type {

		case "response.content_part.added":
			// fmt.Println("\nresponse.content_part.added created")
			
		case "response.text.delta":
			var curText ModelResponseTextDelta

			if err := json.Unmarshal(message, &curText); err != nil {
				fmt.Println("Message error:", err)
				continue
			}

			fmt.Printf("%s", curText.Delta)

			fmt.Println("\nresponse.text.delta created")
		
		case "response.function_call_arguments.done":
			fmt.Println("\nresponse.function_call_arguments.delta created")
			var curText ModelResponseFunction

			if err := json.Unmarshal(message, &curText); err != nil {
				fmt.Println("Message error:", err)
				continue
			}

			fmt.Printf("%s", curText.Arguments)

		case "conversation.item.created":
			// fmt.Println("\nConversation item created")

		case "response.done":
			fmt.Println("")
			return

		case "session_ended":
			// fmt.Println("\nSession ended")
			return

		default:
			fmt.Printf("\nReceived system message of type: %s\n", baseMsg.Type)
		}
	}
}