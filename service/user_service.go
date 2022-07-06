package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	UserRepository model.UserRepository
}

type USConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &userService{
		UserRepository: c.UserRepository,
	}
}

func (s *userService) Get(ctx context.Context, uid string) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, apperrors.NewNotFound("user_id", uid)
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *userService) Signup(ctx context.Context, u *model.User) error {
	pw, err := hashPassword(u.Password)

	if err != nil {
		log.Printf("Unable to signup user for email: %v\n", u.Email)
		return apperrors.NewInternal()
	}
	u.ID = primitive.NewObjectID()
	u.UserID = UserID()
	u.Password = pw
	u.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	u.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	if err := s.UserRepository.Create(ctx, u); err != nil {
		return err
	}

	return nil
}

func (s *userService) Signin(ctx context.Context, u *model.User) error {
	uFetched, err := s.UserRepository.FindByEmail(ctx, u.Email)

	if err != nil {
		return apperrors.NewAuthorization("Invalid email and password combination")
	}

	match, err := comparePasswords(uFetched.Password, u.Password)

	if err != nil {
		return apperrors.NewInternal()
	}

	if !match {
		return apperrors.NewAuthorization("Invalid email and password combination")
	}

	*u = *uFetched
	return nil
}

func (s *userService) GetUsers(ctx context.Context) ([]*model.User, error) {
	users, err := s.UserRepository.Find(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) SearchUsers(ctx context.Context, queryMap map[string]string) ([]*model.User, error) {
	users, err := s.UserRepository.Search(ctx, queryMap)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) Update(ctx context.Context, id string, user *model.User) error {
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	if err := s.UserRepository.Update(ctx, id, user); err != nil {
		return err
	}
	return nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	if err := s.UserRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
