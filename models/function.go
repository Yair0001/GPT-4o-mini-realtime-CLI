package models

type Function struct {
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
}

type Parameters struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required"`
}


type ModelResponseFunction struct{
	Type string `json:"type"`
	Response struct{
		Output []struct{
			Type string `json:"type"`
			Name string `json:"name"`
			CallId string `json:"call_id"`
			Arguments string `json:"arguments"`
		}
	}
}

type FunctionOutput struct{
	Type string `json:"type"`
	Item struct{
		Type string `json:"type"`
		CallId string `json:"call_id"`
		Output string `json:"output"`
	} `json:"item"`
}

type MultiplicationArgs struct{
	Number1 int `json:"num1"`
	Number2 int `json:"num2"`
}