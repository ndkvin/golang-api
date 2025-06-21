package model

type Link struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	UserID uint `json:"user_id"`
}
