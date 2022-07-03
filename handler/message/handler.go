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

	h := &Handler{
		UserService:    c.UserService,
		TokenService:   c.TokenService,
		MessageService: c.MessageService,
	}
	g := c.R.Group(c.BaseURL)

	if gin.Mode() == gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
		c.R.GET("/api/messages/:channel_id", middleware.AuthUser(h.TokenService), h.Messages)
		g.GET("/:msg_id", middleware.AuthUser(h.TokenService), h.Message)
		g.POST("/", middleware.AuthUser(h.TokenService), h.Create)
		g.PUT("/:msg_id", middleware.AuthUser(h.TokenService), h.Update)
		g.DELETE("/:msg_id", middleware.AuthUser(h.TokenService), h.Delete)
	} else {
		c.R.GET("/api/messages/:channel_id", h.Messages)
		g.GET("/:msg_id", h.Message)
		g.POST("/", h.Create)
		g.PUT("/:msg_id", h.Update)
		g.DELETE("/:msg_id", h.Delete)
	}
}
