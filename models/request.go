package models

type Request struct {
	Action string `json:"action"`
	Body   User   `json:"body"`
}
