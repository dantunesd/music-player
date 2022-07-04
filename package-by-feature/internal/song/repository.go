package song

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbName = "music-player"
var collection = "Song"

type RepositoryImpl struct {
	MongoClient *mongo.Client
}

func (r *RepositoryImpl) Create(song *Song) (string, error) {
	result, err := r.MongoClient.Database(dbName).Collection(collection).InsertOne(context.Background(), song)
	if err != nil {
		return "", err
	}

	ID, _ := result.InsertedID.(primitive.ObjectID)
	return ID.Hex(), nil
}

func (r *RepositoryImpl) Get(id string) (*Song, error) {
	var song Song

	objectId, _ := primitive.ObjectIDFromHex(id)
	result := r.MongoClient.Database(dbName).Collection(collection).FindOne(context.Background(), bson.M{"_id": objectId})
	if result.Err() != nil {
		return &song, result.Err()
	}

	return &song, result.Decode(&song)
}

func (r *RepositoryImpl) GetAll() ([]*Song, error) {
	var song []*Song

	cursor, err := r.MongoClient.Database(dbName).Collection(collection).Find(context.Background(), bson.D{})
	if err != nil {
		return song, err
	}

	return song, cursor.All(context.Background(), &song)
}
