package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/config"
	_ "github.com/saravase/primz-chat-backend/docs"
	"github.com/saravase/primz-chat-backend/handler/middleware"
	"github.com/saravase/primz-chat-backend/injector"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

func inject(cfg config.Config, db *mongo.Database) (*gin.Engine, error) {
	log.Println("Injecting data sources")
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	icfg := injector.InjectorConfig{
		Cfg:    cfg,
		DB:     db,
		Engine: router,
	}
	i := injector.NewInjector(&icfg)
	err := i.TokenInjector()
	if err != nil {
		return router, err
	}
	i.OrgInjector()
	i.ChannelInjector()
	i.MessageInjector()
	i.ChatInjector()
	i.UserInjector()
	return router, nil
}
