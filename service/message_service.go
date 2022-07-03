package service

import (
	"context"
	"errors"
	"time"

	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	org, err := s.MessageRepository.FindByID(ctx, id)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, apperrors.NewNotFound("msg_id", id)
	}
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (s *messageService) GetByChennelID(ctx context.Context, id string) ([]*model.Message, error) {
	msgs, err := s.MessageRepository.FindByChannelID(ctx, id)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (s *messageService) Create(ctx context.Context, msg *model.Message) error {
	msg.ID = primitive.NewObjectID()
	msg.MsgID = RandomID()
	msg.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	msg.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	if err := s.MessageRepository.Create(ctx, msg); err != nil {
		return err
	}
	return nil
}

func (s *messageService) Update(ctx context.Context, id string, msg *model.Message) error {
	msg.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	if err := s.MessageRepository.Update(ctx, id, msg); err != nil {
		return err
	}
	return nil
}

func (s *messageService) Delete(ctx context.Context, id string) error {
	if err := s.MessageRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
