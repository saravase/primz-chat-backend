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

type orgService struct {
	OrgRepository model.OrgRepository
}

type OrgConfig struct {
	OrgRepository model.OrgRepository
}

func NewOrgService(c *OrgConfig) model.OrgService {
	return &orgService{
		OrgRepository: c.OrgRepository,
	}
}

func (s *orgService) Get(ctx context.Context, id string) (*model.Org, error) {
	org, err := s.OrgRepository.FindByID(ctx, id)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, apperrors.NewNotFound("org_id", id)
	}
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (s *orgService) GetByName(ctx context.Context, name string) (*model.Org, error) {
	org, err := s.OrgRepository.FindByName(ctx, name)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, apperrors.NewNotFound("org_name", name)
	}
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (s *orgService) GetOrgs(ctx context.Context) ([]*model.Org, error) {
	orgs, err := s.OrgRepository.Find(ctx)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func (s *orgService) Create(ctx context.Context, org *model.Org) error {
	org.ID = primitive.NewObjectID()
	org.OrgID = RandomID()
	org.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	org.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	for didx, dept := range org.Departments {
		for gidx, _ := range dept.Groups {
			org.Departments[didx].Groups[gidx].ID = RandomID()
		}
		org.Departments[didx].ID = RandomID()
	}

	if err := s.OrgRepository.Create(ctx, org); err != nil {
		return err
	}

	return nil
}

func (s *orgService) Update(ctx context.Context, id string, org *model.Org) error {
	org.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	if err := s.OrgRepository.Update(ctx, id, org); err != nil {
		return err
	}
	return nil
}

func (s *orgService) Delete(ctx context.Context, id string) error {
	if err := s.OrgRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
