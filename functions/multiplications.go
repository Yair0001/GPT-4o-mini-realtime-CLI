package functions

import (
	"github.com/Yair0001/GPT-4o-mini-CLI/models"
	"encoding/json"
)

// NewMultiplicationFunction registers a multiplication function
func NewMultiplicationFunction() models.Function {
	return models.Function{
		Type:        "function",
		Name:        "Multiplication",
		Description: "Multiply EXACTLY two numbers",
		Parameters: models.Parameters{
			Type: "object",
			Properties: map[string]interface{}{
				"num1": map[string]interface{}{"type": "integer"},
				"num2": map[string]interface{}{"type": "integer"},
			},
			Required: []string{"num1", "num2"},
		},
	}
}

func MultiplyNumbers(args string) int {
	var data models.MultiplicationArgs
	json.Unmarshal([]byte(args), &data)
	return data.Number1 * data.Number2
}
