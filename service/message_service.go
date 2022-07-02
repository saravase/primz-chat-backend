package service

import (
	"context"

	"github.com/saravase/primz-chat-backend/model"
)

type messageService struct {
	MessageRepository model.MessageRepository
}

type MessageConfig struct {
	MessageRepository model.MessageRepository
}

func NewMessageService(c *MessageConfig) model.MessageService {
	return &messageService{
		MessageRepository: c.MessageRepository,
	}
}

func (s *messageService) Get(ctx context.Context, id string) (*model.Message, error) {
	return nil, nil
}

func (s *messageService) GetByChennelID(ctx context.Context, id string) ([]*model.Message, error) {
	return nil, nil
}

func (s *messageService) Create(ctx context.Context, msg *model.Message) error {
	return nil
}

func (s *messageService) Update(ctx context.Context, id string, msg *model.Message) error {
	return nil
}

func (s *messageService) Delete(ctx context.Context, id string) error {
	return nil
}
