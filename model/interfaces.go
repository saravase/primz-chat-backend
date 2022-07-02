package model

import (
	"context"
	"time"
)

type UserService interface {
	Get(ctx context.Context, uid string) (*User, error)
	GetUsers(ctx context.Context) ([]*User, error)
	Signup(ctx context.Context, u *User) error
	Signin(ctx context.Context, u *User) error
	Update(ctx context.Context, id string, user *User) error
	Delete(ctx context.Context, id string) error
}

type UserRepository interface {
	Find(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, uid string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, id string, user *User) error
	Delete(ctx context.Context, id string) error
}

type TokenService interface {
	NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
	Signout(ctx context.Context, uid string) error
	ValidateIDToken(tokenString string) (*User, error)
	ValidateRefreshToken(refreshTokenString string) (*RefreshToken, error)
}

type TokenRepository interface {
	SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userID string, prevTokenID string) error
	DeleteUserRefreshTokens(ctx context.Context, userID string) error
}

type ChatService interface {
	Get(ctx context.Context, uid string) (*Chat, error)
	GetUsers(ctx context.Context, info *Chat) ([]*Chat, error)
	GetChannel(ctx context.Context, channelID string) (*Channel, error)
}

type ChatRepository interface {
	FindByID(ctx context.Context, uid string) (*Chat, error)
	FindByOrgID(ctx context.Context, orgID string) ([]*Chat, error)
	FindByDeptID(ctx context.Context, deptID string) ([]*Chat, error)
	FindByGroupID(ctx context.Context, grpID string) ([]*Chat, error)
	FindByChannelID(ctx context.Context, channelID string) (*Channel, error)
}

type OrgService interface {
	Get(ctx context.Context, id string) (*Org, error)
	GetByName(ctx context.Context, id string) (*Org, error)
	GetOrgs(ctx context.Context) ([]*Org, error)
	Create(ctx context.Context, org *Org) error
	Update(ctx context.Context, id string, org *Org) error
	Delete(ctx context.Context, id string) error
}

type OrgRepository interface {
	FindByID(ctx context.Context, id string) (*Org, error)
	FindByName(ctx context.Context, name string) (*Org, error)
	Find(ctx context.Context) ([]*Org, error)
	Create(ctx context.Context, org *Org) error
	Update(ctx context.Context, id string, org *Org) error
	Delete(ctx context.Context, id string) error
}

type ChannelService interface {
	Get(ctx context.Context, id string) (*Channel, error)
	GetByUserID(ctx context.Context, id string) ([]*Channel, error)
	Create(ctx context.Context, channel *Channel) error
	Update(ctx context.Context, id string, channel *Channel) error
	Delete(ctx context.Context, id string) error
}

type ChennelRepository interface {
	FindByID(ctx context.Context, id string) (*Channel, error)
	FindByUserID(ctx context.Context, id string) ([]*Channel, error)
	Create(ctx context.Context, channel *Channel) error
	Update(ctx context.Context, id string, channel *Channel) error
	Delete(ctx context.Context, id string) error
}

type MessageService interface {
	Get(ctx context.Context, id string) (*Message, error)
	GetByChennelID(ctx context.Context, id string) ([]*Message, error)
	Create(ctx context.Context, msg *Message) error
	Update(ctx context.Context, id string, msg *Message) error
	Delete(ctx context.Context, id string) error
}

type MessageRepository interface {
	FindByID(ctx context.Context, id string) (*Message, error)
	FindByChannelID(ctx context.Context, id string) ([]*Message, error)
	Create(ctx context.Context, msg *Message) error
	Update(ctx context.Context, id string, msg *Message) error
	Delete(ctx context.Context, id string) error
}
