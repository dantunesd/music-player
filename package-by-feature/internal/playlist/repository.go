package playlist

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbName = "music-player"
var collection = "Playlist"

type RepositoryImpl struct {
	MongoClient *mongo.Client
}

func (r *RepositoryImpl) Create(playlist *Playlist) (string, error) {
	result, err := r.MongoClient.Database(dbName).Collection(collection).InsertOne(context.Background(), playlist)
	if err != nil {
		return "", err
	}

	ID, _ := result.InsertedID.(primitive.ObjectID)
	return ID.Hex(), nil
}

func (r *RepositoryImpl) Get(id string) (*Playlist, error) {
	var playlist Playlist

	objectId, _ := primitive.ObjectIDFromHex(id)
	result := r.MongoClient.Database(dbName).Collection(collection).FindOne(context.Background(), bson.M{"_id": objectId})
	if result.Err() != nil {
		return &playlist, result.Err()
	}

	return &playlist, result.Decode(&playlist)
}

func (r *RepositoryImpl) GetAll() ([]*Playlist, error) {
	var playlist []*Playlist

	cursor, err := r.MongoClient.Database(dbName).Collection(collection).Find(context.Background(), bson.D{})
	if err != nil {
		return playlist, err
	}

	return playlist, cursor.All(context.Background(), &playlist)
}
