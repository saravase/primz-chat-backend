package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID            primitive.ObjectID `bson:"_id" json:"-"`
	MsgID         string             `bson:"msg_id" json:"msg_id"`
	UserID        string             `bson:"user_id" json:"user_id"`
	ChannelID     string             `bson:"channel_id" json:"channel_id"`
	TextContent   string             `bson:"text_content" json:"text_content"`
	AttachmentURL string             `bson:"attachment_url" json:"attachment_url"`
	SeenUserIds   []string           `bson:"seen_user_ids" json:"seen_user_ids"`
	CreatedAt     primitive.DateTime `bson:"created_at" json:"-"`
	UpdatedAt     primitive.DateTime `bson:"updated_at" json:"-"`
}
