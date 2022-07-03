package injector

import (
	"time"

	"github.com/saravase/primz-chat-backend/handler/message"
	"github.com/saravase/primz-chat-backend/repository"
	"github.com/saravase/primz-chat-backend/service"
)

func (i *Injector) MessageInjector() {
	baseURL := i.Cfg.MessageBaseURL()
	handlerTimeout := i.Cfg.HandlerTimeout()
	messageRepository := repository.NewMessageRepository(i.DB)
	i.MessageRepository = messageRepository
	messageService := service.NewMessageService(&service.MessageConfig{
		MessageRepository: messageRepository,
	})
	i.MessageService = messageService
	message.NewHandler(&message.Config{
		R:               i.Engine,
		UserService:     i.UserService,
		TokenService:    i.TokenService,
		MessageService:  i.MessageService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(handlerTimeout) * time.Second),
	})
}
