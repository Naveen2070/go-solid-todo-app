package models

type Todo struct {
	Id          int    `json:"id"`
	IsCompleted bool   `json:"isCompleted"`
	Body        string `json:"body"`
}
