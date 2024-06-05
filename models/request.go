package models

type Request struct {
	Action string `json:"action"`
	User   User   `json:"body"`
}
