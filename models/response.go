package models

type ModelResponseCreate struct {
	Type string `json:"type"`
	Response Response `json:"response"`
}

type Response struct {
	Modalities   []string `json:"modalities"`
	Instructions string   `json:"instructions"`
}

type ModelResponseTextDelta struct{
	Type string `json:"type"`
	Delta string `json:"delta"`
}