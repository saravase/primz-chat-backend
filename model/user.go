package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"-"`
	UserID    string             `bson:"user_id" json:"user_id"`
	FirstName string             `bson:"first_name" json:"first_name" validate:"required, min=2, max=100"`
	LastName  string             `bson:"last_name" json:"last_name" validate:"required, min=2, max=100"`
	Password  string             `bson:"password" json:"-" validate:"required, min=6"`
	Email     string             `bson:"email" json:"email" validate:"required, email"`
	Role      string             `bson:"role" json:"role" validate:"required, eq=ADMIN|eq=USER"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}
