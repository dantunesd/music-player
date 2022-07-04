package repository

import (
	"context"
	"music-player/package-by-layer/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var playlistDbName = "music-player"
var playlistCollection = "Playlist"

type Playlist struct {
	MongoClient *mongo.Client
}

func (r *Playlist) Create(playlist *domain.Playlist) (string, error) {
	result, err := r.MongoClient.Database(playlistDbName).Collection(playlistCollection).InsertOne(context.Background(), playlist)
	if err != nil {
		return "", err
	}

	ID, _ := result.InsertedID.(primitive.ObjectID)
	return ID.Hex(), nil
}

func (r *Playlist) Get(id string) (*domain.Playlist, error) {
	var playlist domain.Playlist

	objectId, _ := primitive.ObjectIDFromHex(id)
	result := r.MongoClient.Database(playlistDbName).Collection(playlistCollection).FindOne(context.Background(), bson.M{"_id": objectId})
	if result.Err() != nil {
		return &playlist, result.Err()
	}

	return &playlist, result.Decode(&playlist)
}

func (r *Playlist) GetAll() ([]*domain.Playlist, error) {
	var playlist []*domain.Playlist

	cursor, err := r.MongoClient.Database(playlistDbName).Collection(playlistCollection).Find(context.Background(), bson.D{})
	if err != nil {
		return playlist, err
	}

	return playlist, cursor.All(context.Background(), &playlist)
}
