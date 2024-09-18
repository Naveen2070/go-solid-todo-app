package models

type Todo struct {
	Id          int    `json:"id" bson:"_id"`
	IsCompleted bool   `json:"isCompleted"`
	Body        string `json:"body"`
}
