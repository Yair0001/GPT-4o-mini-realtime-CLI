# GPT-4o-mini-CLI


websockets - gorrila package in go
i connect with a websocket to wss://api.openai.com/v1/realtime?model=gpt-4o-mini-realtime-preview-2024-12-17
and send headers:
Authorization: Bearer + API_KEY
OpenAI-Beta: realtime=v1
and functions for on_open (connection)
and on_message (recieved event)

type AssistantMessage struct {
	Type    string `json:"type"`
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
}