package repository

import (
	"context"

	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) model.UserRepository {
	return &mongoUserRepository{
		collection: db.Collection("primz_chat_users"),
	}
}

func (r *mongoUserRepository) Create(ctx context.Context, u *model.User) error {

	user, err := r.FindByID(ctx, u.UserID)
	if user != nil {
		return apperrors.NewNotFound("user_id", u.UserID)
	}

	_, err = r.collection.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoUserRepository) FindByID(ctx context.Context, uid string) (*model.User, error) {
	user := &model.User{}

	filter := bson.D{{"user_id", uid}}

	err := r.collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mongoUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}

	filter := bson.D{{"email", email}}

	err := r.collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mongoUserRepository) Find(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}

func (r *mongoUserRepository) Update(ctx context.Context, id string, user *model.User) error {
	return nil
}

func (r *mongoUserRepository) Delete(ctx context.Context, id string) error {
	return nil
}
