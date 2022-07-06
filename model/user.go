package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id" json:"-"`
	UserID         string             `bson:"user_id" json:"user_id"`
	FirstName      string             `bson:"first_name" json:"first_name" validate:"required, min=2, max=100"`
	LastName       string             `bson:"last_name" json:"last_name" validate:"required, min=2, max=100"`
	Password       string             `bson:"password" json:"-" validate:"required, min=6"`
	Email          string             `bson:"email" json:"email" validate:"required, email"`
	Role           string             `bson:"role" json:"role" `
	AvatarURL      string             `bson:"avatar_url" json:"avatar_url" validate:"url"`
	OrgID          string             `bson:"org_id" json:"org_id"`
	DeptID         string             `bson:"dept_id" json:"dept_id"`
	GroupID        string             `bson:"group_id" json:"group_id"`
	PrivChannelIds []string           `bson:"priv_channel_ids" json:"priv_channel_ids"`
	PubChannelIds  []string           `bson:"pub_channel_ids" json:"pub_channel_ids"`
	ActiveStatus   bool               `bson:"active_status" json:"active_status"`
	CreatedAt      primitive.DateTime `bson:"created_at" json:"-"`
	UpdatedAt      primitive.DateTime `bson:"updated_at" json:"-"`
} // @name User
