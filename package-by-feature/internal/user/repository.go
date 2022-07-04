package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbName = "music-player"
var collection = "user"

type RepositoryImpl struct {
	MongoClient *mongo.Client
}

func (r *RepositoryImpl) Create(user *User) (string, error) {
	result, err := r.MongoClient.Database(dbName).Collection(collection).InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	ID, _ := result.InsertedID.(primitive.ObjectID)
	return ID.Hex(), nil
}

func (r *RepositoryImpl) Get(id string) (*User, error) {
	var user User

	objectId, _ := primitive.ObjectIDFromHex(id)
	result := r.MongoClient.Database(dbName).Collection(collection).FindOne(context.Background(), bson.M{"_id": objectId})
	if result.Err() != nil {
		return &user, result.Err()
	}

	return &user, result.Decode(&user)
}

func (r *RepositoryImpl) GetAll() ([]*User, error) {
	var users []*User

	cursor, err := r.MongoClient.Database(dbName).Collection(collection).Find(context.Background(), bson.D{})
	if err != nil {
		return users, err
	}

	return users, cursor.All(context.Background(), &users)
}
