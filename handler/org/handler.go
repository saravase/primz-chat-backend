package org

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
	OrgService   model.OrgService
}

type Config struct {
	R               *gin.Engine
	UserService     model.UserService
	TokenService    model.TokenService
	OrgService      model.OrgService
	BaseURL         string
	TimeoutDuration time.Duration
}

func NewHandler(c *Config) {

	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
		OrgService:   c.OrgService,
	}
	g := c.R.Group(c.BaseURL)

	if gin.Mode() == gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
		g.GET("/", middleware.AuthUser(h.TokenService), h.Orgs)
		g.GET("/:org_id", middleware.AuthUser(h.TokenService), h.Org)
		g.POST("/", middleware.AuthUser(h.TokenService), h.Create)
		g.PUT("/:org_id", middleware.AuthUser(h.TokenService), h.Update)
		g.DELETE("/:org_id", middleware.AuthUser(h.TokenService), h.Delete)
	} else {
		g.GET("/", h.Orgs)
		g.GET("/:org_id", h.Org)
		g.POST("/", h.Create)
		g.PUT("/:org_id", h.Update)
		g.DELETE("/:org_id", h.Delete)
	}
}
