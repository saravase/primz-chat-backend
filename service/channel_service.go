package service

import (
	"context"

	"github.com/saravase/primz-chat-backend/model"
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
	return nil, nil
}

func (s *channelService) GetByUserID(ctx context.Context, id string) ([]*model.Channel, error) {
	return nil, nil
}

func (s *channelService) Create(ctx context.Context, channel *model.Channel) error {
	return nil
}

func (s *channelService) Update(ctx context.Context, id string, channel *model.Channel) error {
	return nil
}

func (s *channelService) Delete(ctx context.Context, id string) error {
	return nil
}
