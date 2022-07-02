package injector

import (
	"time"

	"github.com/saravase/primz-chat-backend/handler/org"
	"github.com/saravase/primz-chat-backend/repository"
	"github.com/saravase/primz-chat-backend/service"
)

func (i *Injector) OrgInjector() {
	baseURL := i.Cfg.OrgBaseURL()
	handlerTimeout := i.Cfg.HandlerTimeout()
	orgRepository := repository.NewOrgRepository(i.DB)
	orgService := service.NewOrgService(&service.OrgConfig{
		OrgRepository: orgRepository,
	})
	i.OrgService = orgService
	org.NewHandler(&org.Config{
		R:               i.Engine,
		UserService:     i.UserService,
		TokenService:    i.TokenService,
		OrgService:      i.OrgService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(handlerTimeout) * time.Second),
	})
}
