package injector

import (
	"time"

	"github.com/saravase/primz-chat-backend/handler/user"
	"github.com/saravase/primz-chat-backend/repository"
	"github.com/saravase/primz-chat-backend/service"
)

func (i *Injector) UserInjector() {
	baseURL := i.Cfg.AuthBaseURL()
	handlerTimeout := i.Cfg.HandlerTimeout()
	userRepository := repository.NewUserRepository(i.DB)
	userService := service.NewUserService(&service.USConfig{
		UserRepository: userRepository,
	})
	i.UserService = userService
	user.NewHandler(&user.Config{
		R:               i.Engine,
		UserService:     i.UserService,
		TokenService:    i.TokenService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(handlerTimeout) * time.Second),
	})
}
