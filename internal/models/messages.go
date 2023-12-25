package models

type MessageSet struct {
	Key   string      `json:"key"  validate:"required"`
	Value interface{} `json:"value" validate:"required"`
}
