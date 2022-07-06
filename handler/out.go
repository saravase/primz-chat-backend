package handler

type CreateResponse struct {
	ID string `json:"id"`
} // @name CreateResponse

type UpdateResponse struct {
	IsUpdated bool `json:"is_updated"`
} // @name UpdateResponse

type DeleteResponse struct {
	IsDeleted bool `json:"is_deleted"`
} // @name DeleteResponse

type ChatMessage struct {
	MsgID         string `json:"msg_id"`
	UserID        string `json:"user_id"`
	ChannelID     string `json:"channel_id"`
	TextContent   string `json:"text_content"`
	AttachmentURL string `json:"attachment_url"`
}

type ChatUser struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	AvatarURL    string `json:"avatar_url"`
	ActiveStatus bool   `json:"active_status"`
}

type Chat struct {
	ChannelId string              `json:"channel_id"`
	UserIDs   []string            `json:"user_ids"`
	Users     map[string]ChatUser `json:"users"`
	Messages  []ChatMessage       `json:"messages"`
}

type Chats struct {
	Chats []Chat `json:"chats"`
} // @name Chats
