package injector

import (
	"time"

	"github.com/saravase/primz-chat-backend/handler/channel"
	"github.com/saravase/primz-chat-backend/repository"
	"github.com/saravase/primz-chat-backend/service"
)

func (i *Injector) ChannelInjector() {
	baseURL := i.Cfg.ChannelBaseURL()
	handlerTimeout := i.Cfg.HandlerTimeout()
	channelRepository := repository.NewChannelRepository(i.DB)
	i.ChennelRepository = channelRepository
	channelService := service.NewChannelService(&service.ChannelConfig{
		ChannelRepository: channelRepository,
	})
	i.ChannelService = channelService
	channel.NewHandler(&channel.Config{
		R:               i.Engine,
		UserService:     i.UserService,
		TokenService:    i.TokenService,
		ChannelService:  i.ChannelService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(handlerTimeout) * time.Second),
	})
}
