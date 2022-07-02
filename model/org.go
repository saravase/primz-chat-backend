package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Department struct {
	ID     string  `json:"id" bson:"id"`
	Name   string  `json:"name" bson:"name"`
	Groups []Group `json:"groups" bson:"groups"`
}

type Group struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Org struct {
	ID          primitive.ObjectID `bson:"_id" json:"-"`
	OrgID       string             `bson:"org_id" json:"org_id"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	Website     string             `bson:"website" json:"website"`
	AvatarURL   string             `bson:"avatar_url" json:"avatar_url"`
	Departments []Department       `bson:"departments" json:"departments"`
	CreatedAt   primitive.DateTime `bson:"created_at" json:"-"`
	UpdatedAt   primitive.DateTime `bson:"updated_at" json:"-"`
}
