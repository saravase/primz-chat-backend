package repository

import (
	"context"

	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoMessageRepository struct {
	collection *mongo.Collection
}

func NewMessageRepository(db *mongo.Database) model.MessageRepository {
	return &mongoMessageRepository{
		collection: db.Collection("primz_chat_messages"),
	}
}

func (r *mongoMessageRepository) FindByID(ctx context.Context, id string) (msg *model.Message, err error) {
	filter := bson.D{{"msg_id", id}}
	err = r.collection.FindOne(ctx, filter).Decode(&msg)
	if err != nil {
		return
	}
	return
}

func (r *mongoMessageRepository) FindByChannelID(ctx context.Context, id string) ([]*model.Message, error) {
	return nil, nil
}

func (r *mongoMessageRepository) Create(ctx context.Context, msg *model.Message) error {
	_, err := r.collection.InsertOne(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoMessageRepository) Update(ctx context.Context, id string, msg *model.Message) error {
	filter := bson.D{{"msg_id", id}}
	_, err := r.collection.ReplaceOne(ctx, filter, msg)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoMessageRepository) Delete(ctx context.Context, id string) error {
	filter := bson.D{{"msg_id", id}}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
