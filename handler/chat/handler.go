package chat

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/model"
)

type Handler struct {
	UserService    model.UserService
	TokenService   model.TokenService
	OrgService     model.OrgService
	ChannelService model.ChannelService
	MessageService model.MessageService
	ChatService    model.ChatService
}

type Config struct {
	R               *gin.Engine
	UserService     model.UserService
	TokenService    model.TokenService
	OrgService      model.OrgService
	ChannelService  model.ChannelService
	MessageService  model.MessageService
	ChatService     model.ChatService
	BaseURL         string
	TimeoutDuration time.Duration
}

func NewHandler(c *Config) {

	h := &Handler{
		UserService:    c.UserService,
		TokenService:   c.TokenService,
		MessageService: c.MessageService,
		OrgService:     c.OrgService,
		ChannelService: c.ChannelService,
		ChatService:    c.ChatService,
	}
	g := c.R.Group(c.BaseURL)
	g.GET("/:user_id", h.UserConnection)
	g.POST("/", h.GetChat)

}
