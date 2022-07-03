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

	h := &Handler{
		UserService:    c.UserService,
		TokenService:   c.TokenService,
		ChannelService: c.ChannelService,
	}
	g := c.R.Group(c.BaseURL)

	if gin.Mode() == gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
		c.R.POST("/api/channels", middleware.AuthUser(h.TokenService), h.Channels)
		g.GET("/:channel_id", middleware.AuthUser(h.TokenService), h.Channel)
		g.POST("/", middleware.AuthUser(h.TokenService), h.Create)
		g.PUT("/:channel_id", middleware.AuthUser(h.TokenService), h.Update)
		g.DELETE("/:channel_id", middleware.AuthUser(h.TokenService), h.Delete)
	} else {
		c.R.POST("/api/channels", h.Channels)
		g.GET("/:channel_id", h.Channel)
		g.POST("/", h.Create)
		g.PUT("/:channel_id", h.Update)
		g.DELETE("/:channel_id", h.Delete)
	}
}
