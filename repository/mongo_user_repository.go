package repository

import (
	"context"
	"fmt"

	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *mongoUserRepository) Find(ctx context.Context) (users []*model.User, err error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}
	if err = cursor.All(ctx, &users); err != nil {
		return
	}
	return
}

func (r *mongoUserRepository) Search(ctx context.Context, queryMap map[string]string) (users []*model.User, err error) {

	filter := bson.M{}
	for key, value := range queryMap {
		if key == "name" {
			pattern := fmt.Sprintf("^%s*", value)
			filter["first_name"] = primitive.Regex{Pattern: pattern}
		} else {
			filter[key] = value
		}
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return
	}
	if err = cursor.All(ctx, &users); err != nil {
		return
	}
	return
}

func (r *mongoUserRepository) Update(ctx context.Context, id string, user *model.User) error {
	filter := bson.D{{"user_id", id}}
	_, err := r.collection.ReplaceOne(ctx, filter, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoUserRepository) Delete(ctx context.Context, id string) error {
	filter := bson.D{{"user_id", id}}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
