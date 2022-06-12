package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID        primitive.ObjectID `bson:"_id" json:"-"`
	TokenKey  string             `bson:"token_key" json:"token_key"`
	ExpiresIn time.Duration      `bson:"expires_in" json:"expires_in"`
}

type RefreshToken struct {
	ID  string `json:"-"`
	UID string `json:"-"`
	SS  string `json:"refreshToken"`
}

type IDToken struct {
	SS string `json:"idToken"`
}

type TokenPair struct {
	IDToken
	RefreshToken
}
