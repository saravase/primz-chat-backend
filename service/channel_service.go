package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type channelService struct {
	ChannelRepository model.ChennelRepository
}

type ChannelConfig struct {
	ChannelRepository model.ChennelRepository
}

func NewChannelService(c *ChannelConfig) model.ChannelService {
	return &channelService{
		ChannelRepository: c.ChannelRepository,
	}
}

func (s *channelService) Get(ctx context.Context, id string) (*model.Channel, error) {
	channel, err := s.ChannelRepository.FindByID(ctx, id)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, apperrors.NewNotFound("channel_id", id)
	}
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (s *channelService) GetByUsers(ctx context.Context, users *[]model.ChannelUser) (*model.Channel, error) {
	org, err := s.ChannelRepository.FindByUsers(ctx, users)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, apperrors.NewNotFound("users", fmt.Sprintf("%#v", users))
	}
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (s *channelService) GetByUserIDs(ctx context.Context, id []string) ([]*model.Channel, error) {
	channels, err := s.ChannelRepository.FindByUserIDs(ctx, id)
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func (s *channelService) Create(ctx context.Context, channel *model.Channel) error {
	channel.ID = primitive.NewObjectID()
	channel.ChannelID = RandomID()
	channel.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	channel.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	if err := s.ChannelRepository.Create(ctx, channel); err != nil {
		return err
	}
	return nil
}

func (s *channelService) Update(ctx context.Context, id string, channel *model.Channel) error {
	channel.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	if err := s.ChannelRepository.Update(ctx, id, channel); err != nil {
		return err
	}
	return nil
}

func (s *channelService) Delete(ctx context.Context, id string) error {
	if err := s.ChannelRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
