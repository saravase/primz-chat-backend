package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChannelUser struct {
	ID   string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}

type Channel struct {
	ID           primitive.ObjectID `bson:"_id" json:"-"`
	ChannelID    string             `bson:"channel_id" json:"channel_id"`
	Users        []ChannelUser      `bson:"users" json:"users"`
	Name         string             `bson:"name" json:"name"`
	GroupChannel bool               `bson:"group_channel" json:"group_channel"`
	ChannelOwner ChannelUser        `bson:"channel_owner" json:"channel_owner"`
	CreatedAt    primitive.DateTime `bson:"created_at" json:"-"`
	UpdatedAt    primitive.DateTime `bson:"updated_at" json:"-"`
}
