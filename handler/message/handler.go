package message

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
	MessageService model.MessageService
}

type Config struct {
	R               *gin.Engine
	UserService     model.UserService
	TokenService    model.TokenService
	MessageService  model.MessageService
	BaseURL         string
	TimeoutDuration time.Duration
}

func NewHandler(c *Config) {

	_ = &Handler{
		UserService:    c.UserService,
		TokenService:   c.TokenService,
		MessageService: c.MessageService,
	}
	g := c.R.Group(c.BaseURL)

	if gin.Mode() == gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))

	} else {

	}
}
