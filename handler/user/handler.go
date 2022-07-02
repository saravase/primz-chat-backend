package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler/middleware"
	"github.com/saravase/primz-chat-backend/model"
)

type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

type Config struct {
	R               *gin.Engine
	UserService     model.UserService
	TokenService    model.TokenService
	BaseURL         string
	TimeoutDuration time.Duration
}

func NewHandler(c *Config) {

	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}

	g := c.R.Group(c.BaseURL)

	if gin.Mode() != gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
		g.GET("/me", middleware.AuthUser(h.TokenService), h.Me)
		g.POST("/signout", middleware.AuthUser(h.TokenService), h.Signout)
	} else {
		g.GET("/me", h.Me)
		g.POST("/signout", h.Signout)
	}

	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/tokens", h.Tokens)

}
