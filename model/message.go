package model

type Message struct {
	ID            string   `json:"id"`
	UserID        string   `json:"user_id"`
	ChannelID     string   `json:"channel_id"`
	TextContent   string   `json:"text_content"`
	AttachmentURL string   `json:"attachment_url"`
	SeenUserIds   []string `json:"seen_user_ids"`
}
