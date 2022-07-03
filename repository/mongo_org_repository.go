package repository

import (
	"context"

	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoOrgRepository struct {
	collection *mongo.Collection
}

func NewOrgRepository(db *mongo.Database) model.OrgRepository {
	return &mongoOrgRepository{
		collection: db.Collection("primz_chat_orgs"),
	}
}

func (r *mongoOrgRepository) FindByID(ctx context.Context, id string) (org *model.Org, err error) {
	filter := bson.D{{"org_id", id}}
	err = r.collection.FindOne(ctx, filter).Decode(&org)
	if err != nil {
		return
	}
	return
}

func (r *mongoOrgRepository) Find(ctx context.Context) (orgs []*model.Org, err error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}
	if err = cursor.All(ctx, &orgs); err != nil {
		return
	}
	return
}

func (r *mongoOrgRepository) Create(ctx context.Context, org *model.Org) error {
	_, err := r.collection.InsertOne(ctx, org)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoOrgRepository) Update(ctx context.Context, id string, org *model.Org) error {
	filter := bson.D{{"org_id", id}}
	_, err := r.collection.ReplaceOne(ctx, filter, org)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoOrgRepository) Delete(ctx context.Context, id string) error {
	filter := bson.D{{"org_id", id}}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoOrgRepository) FindByName(ctx context.Context, name string) (org *model.Org, err error) {
	filter := bson.D{{"name", name}}
	err = r.collection.FindOne(ctx, filter).Decode(&org)
	if err != nil {
		return
	}
	return
}
