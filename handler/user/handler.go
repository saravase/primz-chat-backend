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

	if gin.Mode() == gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
		g.GET("/me", middleware.AuthUser(h.TokenService), h.Me)
		g.POST("/signout", middleware.AuthUser(h.TokenService), h.Signout)
		g.PUT("/user/:user_id", middleware.AuthUser(h.TokenService), h.Update)
		g.DELETE("/user/:user_id", middleware.AuthUser(h.TokenService), h.Delete)
		g.GET("/users", middleware.AuthUser(h.TokenService), h.Users)
		g.GET("/users/search", middleware.AuthUser(h.TokenService), h.Search)
	} else {
		g.GET("/me", h.Me)
		g.POST("/signout", h.Signout)
		g.PUT("/user/:user_id", h.Update)
		g.DELETE("/user/:user_id", h.Delete)
		g.GET("/users", h.Users)
		g.GET("/users/search", h.Search)
	}

	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/tokens", h.Tokens)

}
