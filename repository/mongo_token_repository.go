package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoTokenRepository struct {
	collection *mongo.Collection
}

func NewTokenRepository(db *mongo.Database) model.TokenRepository {
	return &mongoTokenRepository{
		collection: db.Collection("primz_chat_tokens"),
	}
}

func (r *mongoTokenRepository) SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error {

	key := fmt.Sprintf("%s:%s", userID, tokenID)
	user, err := r.FindByTokenKey(ctx, key)
	if user != nil {
		return apperrors.NewNotFound("token_key", key)
	}

	token := &model.Token{
		ID:        primitive.NewObjectID(),
		TokenKey:  key,
		ExpiresIn: expiresIn,
	}

	_, err = r.collection.InsertOne(ctx, token)
	if err != nil {
		log.Printf("Could not SET refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, tokenID, err)
		return apperrors.NewInternal()
	}

	return nil
}

func (r *mongoTokenRepository) DeleteRefreshToken(ctx context.Context, userID string, tokenID string) error {
	key := fmt.Sprintf("%s:%s", userID, tokenID)

	filter := bson.D{{"token_key", key}}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Could not delete refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, tokenID, err)
		return apperrors.NewInternal()
	}

	if result.DeletedCount < 1 {
		log.Printf("Refresh token to redis for userID/tokenID: %s/%s does not exist\n", userID, tokenID)
		return apperrors.NewAuthorization("Invalid refresh token")
	}

	return nil
}

func (r *mongoTokenRepository) DeleteUserRefreshTokens(ctx context.Context, userID string) error {
	pattern := fmt.Sprintf("%s*", userID)

	filter := bson.D{{"token_key", primitive.Regex{Pattern: pattern}}}

	result, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Printf("Could not delete refresh token to mongodb for tokenID: %s: %v\n", userID, err)
		return apperrors.NewInternal()
	}

	if result.DeletedCount < 1 {
		log.Printf("Refresh token to redis for userID/tokenID: %s does not exist\n", userID)
		return apperrors.NewAuthorization("Invalid refresh token")
	}

	return nil
}

func (r *mongoTokenRepository) FindByTokenKey(ctx context.Context, key string) (*model.Token, error) {
	token := &model.Token{}

	filter := bson.D{{"token_key", key}}

	err := r.collection.FindOne(ctx, filter).Decode(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
