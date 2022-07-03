package message

type MessageCreateReq struct {
	UserID        string   `json:"user_id" validate:"required,len=8"`
	ChannelID     string   `json:"channel_id" validate:"required,len=8"`
	TextContent   string   `json:"text_content" validate:"required"`
	AttachmentURL string   `json:"attachment_url" validate:"omitempty,url"`
	SeenUserIds   []string `json:"seen_user_ids" validate:"omitempty,gt=0"`
} // @name MessageCreateReq

type MessageUpdateReq struct {
	TextContent   string `json:"text_content" validate:"omitempty"`
	AttachmentURL string `json:"attachment_url" validate:"omitempty,url"`
} // @name MessageUpdateReq
