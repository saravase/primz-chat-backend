package injector

import (
	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/config"
	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Injector struct {
	Cfg            config.Config
	DB             *mongo.Database
	Engine         *gin.Engine
	UserService    model.UserService
	TokenService   model.TokenService
	OrgService     model.OrgService
	ChannelService model.ChannelService
	MessageService model.MessageService
}

type InjectorConfig struct {
	Cfg          config.Config
	DB           *mongo.Database
	Engine       *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
	OrgService   model.OrgService
}

func NewInjector(c *InjectorConfig) *Injector {
	return &Injector{
		Cfg:          c.Cfg,
		DB:           c.DB,
		Engine:       c.Engine,
		UserService:  c.UserService,
		TokenService: c.TokenService,
		OrgService:   c.OrgService,
	}
}
