package repository

import (
	"context"

	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoChannelRepository struct {
	collection *mongo.Collection
}

func NewChannelRepository(db *mongo.Database) model.ChennelRepository {
	return &mongoChannelRepository{
		collection: db.Collection("primz_chat_channels"),
	}
}

func (r *mongoChannelRepository) FindByID(ctx context.Context, id string) (channel *model.Channel, err error) {
	filter := bson.D{{"channel_id", id}}
	err = r.collection.FindOne(ctx, filter).Decode(&channel)
	if err != nil {
		return
	}
	return
}

func (r *mongoChannelRepository) FindByUsers(ctx context.Context, users *[]model.ChannelUser) (channel *model.Channel, err error) {
	filter := bson.D{{"users", users}}
	err = r.collection.FindOne(ctx, filter).Decode(&channel)
	if err != nil {
		return
	}
	return
}

func (r *mongoChannelRepository) FindByUserIDs(ctx context.Context, ids []string) (channels []*model.Channel, err error) {
	filter := bson.D{{"channel_id", bson.D{{"$in", ids}}}}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return
	}
	if err = cursor.All(ctx, &channels); err != nil {
		return
	}
	return
}

func (r *mongoChannelRepository) Create(ctx context.Context, channel *model.Channel) error {
	_, err := r.collection.InsertOne(ctx, channel)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoChannelRepository) Update(ctx context.Context, id string, channel *model.Channel) error {
	filter := bson.D{{"channel_id", id}}
	_, err := r.collection.ReplaceOne(ctx, filter, channel)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoChannelRepository) Delete(ctx context.Context, id string) error {
	filter := bson.D{{"channel_id", id}}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
