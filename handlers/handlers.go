package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/Yair0001/GPT-4o-mini-CLI/functions"
	"github.com/Yair0001/GPT-4o-mini-CLI/models"
	"github.com/Yair0001/GPT-4o-mini-CLI/utils"

	"github.com/gorilla/websocket"
)

// number of function calls
var funcCounter int = 0

// UserInputHandler listens for user input and sends messages
func UserInputHandler(conn *websocket.Conn, responseDone chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		<-responseDone
		fmt.Print("\n> ")
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			conn.Close()
			return
		}

		// Send user message to the server
		err := utils.SendUserMessage(conn, text)
		if err != nil {
			fmt.Println("Send error:", err)
			return
		}

		// Send model response request to the server
		err = utils.SendModelResponse(conn)
		if err != nil {
			fmt.Println("Send error:", err)
			return
		}
	}
}

// MessageReceiver processes incoming messages
func MessageReceiver(conn *websocket.Conn, responseDone chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}

		// Process the received message
		ProcessMessage(conn, message, responseDone)
	}
}

// ProcessMessage handles different message types
func ProcessMessage(conn *websocket.Conn, message []byte, responseDone chan bool) {
	var baseMsg struct{ Type string }
	if err := json.Unmarshal(message, &baseMsg); err != nil {
		fmt.Println("JSON parse error:", err)
		return
	}

	switch baseMsg.Type {
	case "response.text.delta":
		var curText models.ModelResponseTextDelta
		json.Unmarshal(message, &curText)
		fmt.Printf("%s", curText.Delta)
	case "response.function_call_arguments.done":
		funcCounter++
	case "response.done":
		if funcCounter > 0 {
			FunctionCallHandler(conn, message)
			funcCounter--
		} else {
			responseDone <- false
		}
	case "error":
		fmt.Printf("\nError: %s\n", string(message))

	default:
	}
}

// FunctionCallHandler handles function call responses
func FunctionCallHandler(conn *websocket.Conn, message []byte) {
	var functionCall models.ModelResponseFunction
	json.Unmarshal(message, &functionCall)

	switch functionCall.Response.Output[0].Name {
	case "Multiplication":
		args := functionCall.Response.Output[0].Arguments
		result := functions.MultiplyNumbers(args)
		sendFunctionOutput(conn, functionCall.Response.Output[0].CallId, result)

	default:
		fmt.Println("Function not found")
	}
}

// sendFunctionOutput sends the result of a function call back to the server
func sendFunctionOutput(conn *websocket.Conn, callId string, result int) {
	output := models.FunctionOutput{
		Type: "conversation.item.create",
		Item: struct {
			Type   string `json:"type"`
			CallId string `json:"call_id"`
			Output string `json:"output"`
		}{
			Type:   "function_call_output",
			CallId: callId,
			Output: fmt.Sprintf("result %d", result),
		},
	}

	// Send the function output to the server
	err := utils.SendJSONMessage(conn, output)
	if err != nil {
		fmt.Println("Send error:", err)
		return
	}
	utils.SendModelResponse(conn)
}
