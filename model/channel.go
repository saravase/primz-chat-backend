package model

type Channel struct {
	ID           string   `json:"id"`
	UserIds      []string `json:"user_ids"`
	Name         string   `json:"name"`
	GroupChannel bool     `json:"group_channel"`
	MessageIds   []string `json:"message_ids"`
}
