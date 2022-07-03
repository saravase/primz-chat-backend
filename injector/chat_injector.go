package injector

import (
	"time"

	"github.com/saravase/primz-chat-backend/handler/chat"
	"github.com/saravase/primz-chat-backend/service"
)

func (i *Injector) ChatInjector() {
	baseURL := i.Cfg.ChatBaseURL()
	handlerTimeout := i.Cfg.HandlerTimeout()
	chatService := service.NewChatService(&service.ChatConfig{
		UserRepository:    i.UserRepository,
		OrgRepository:     i.OrgRepository,
		ChennelRepository: i.ChennelRepository,
		MessageRepository: i.MessageRepository,
	})
	i.ChatService = chatService
	chat.NewHandler(&chat.Config{
		R:               i.Engine,
		UserService:     i.UserService,
		TokenService:    i.TokenService,
		ChannelService:  i.ChannelService,
		MessageService:  i.MessageService,
		OrgService:      i.OrgService,
		ChatService:     i.ChatService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(handlerTimeout) * time.Second),
	})

	go chatService.UsersChatManager()
}
