package repository

import (
	"context"
	"music-player/package-by-layer/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var songDBName = "music-player"
var songCollection = "Song"

type Song struct {
	MongoClient *mongo.Client
}

func (r *Song) Create(song *domain.Song) (string, error) {
	result, err := r.MongoClient.Database(songDBName).Collection(songCollection).InsertOne(context.Background(), song)
	if err != nil {
		return "", err
	}

	ID, _ := result.InsertedID.(primitive.ObjectID)
	return ID.Hex(), nil
}

func (r *Song) Get(id string) (*domain.Song, error) {
	var song domain.Song

	objectId, _ := primitive.ObjectIDFromHex(id)
	result := r.MongoClient.Database(songDBName).Collection(songCollection).FindOne(context.Background(), bson.M{"_id": objectId})
	if result.Err() != nil {
		return &song, result.Err()
	}

	return &song, result.Decode(&song)
}

func (r *Song) GetAll() ([]*domain.Song, error) {
	var song []*domain.Song

	cursor, err := r.MongoClient.Database(songDBName).Collection(songCollection).Find(context.Background(), bson.D{})
	if err != nil {
		return song, err
	}

	return song, cursor.All(context.Background(), &song)
}
