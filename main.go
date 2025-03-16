package main

import (
	"fmt"
	"sync"

	"github.com/Yair0001/GPT-4o-mini-CLI/config"
	"github.com/Yair0001/GPT-4o-mini-CLI/functions"
	"github.com/Yair0001/GPT-4o-mini-CLI/handlers"
	"github.com/Yair0001/GPT-4o-mini-CLI/models"
	// "github.com/gorilla/websocket"
	// "github.com/Yair0001/GPT-4o-mini-CLI/utils"
	"github.com/Yair0001/GPT-4o-mini-CLI/websocket"

	
)

var (
	sessionID   string
	wg          sync.WaitGroup
	responseDone chan bool
)

func main() {
	apiKey := config.GetAPIKey()
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY not set")
		return
	}

	var err error
	conn, err := websocket.Connect(apiKey)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close()

	sessionID, err = websocket.InitSession(conn)
	if err != nil {
		fmt.Println("Session init error:", err)
		return
	}

	fmt.Printf("Session created: %s\n", sessionID)

	// Register functions
	err = websocket.RegisterFunctions(conn, []models.Function{
		functions.NewMultiplicationFunction(),
	})
	if err != nil {
		fmt.Println("Failed to register functions:", err)
		return
	}

	responseDone = make(chan bool, 1)
	responseDone <- true

	wg.Add(2)
	go handlers.UserInputHandler(conn, responseDone, &wg)
	go handlers.MessageReceiver(conn, responseDone, &wg)
	wg.Wait()
}
