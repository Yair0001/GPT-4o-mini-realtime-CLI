package config

import "os"

// GetAPIKey retrieves the OpenAI API Key
func GetAPIKey() string {
	return os.Getenv("OPENAI_API_KEY")
}
