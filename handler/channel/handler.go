package channel

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler/middleware"
	"github.com/saravase/primz-chat-backend/model"
)

type Handler struct {
	UserService    model.UserService
	TokenService   model.TokenService
	ChannelService model.ChannelService
}

type Config struct {
	R               *gin.Engine
	UserService     model.UserService
	TokenService    model.TokenService
	ChannelService  model.ChannelService
	BaseURL         string
	TimeoutDuration time.Duration
}

func NewHandler(c *Config) {

	_ = &Handler{
		UserService:    c.UserService,
		TokenService:   c.TokenService,
		ChannelService: c.ChannelService,
	}
	g := c.R.Group(c.BaseURL)

	if gin.Mode() == gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))

	} else {

	}
}
