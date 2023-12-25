package models

type Message struct {
	Key     string `json:"key"  validate:"required"`
	Content Value  `json:"content"`
}

type Value struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}
