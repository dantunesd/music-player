package repository

import (
	"context"
	"music-player/package-by-layer/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userDBName = "music-player"
var userCollection = "user"

type User struct {
	MongoClient *mongo.Client
}

func (r *User) Create(user *domain.User) (string, error) {
	result, err := r.MongoClient.Database(userDBName).Collection(userCollection).InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	ID, _ := result.InsertedID.(primitive.ObjectID)
	return ID.Hex(), nil
}

func (r *User) Get(id string) (*domain.User, error) {
	var user domain.User

	objectId, _ := primitive.ObjectIDFromHex(id)
	result := r.MongoClient.Database(userDBName).Collection(userCollection).FindOne(context.Background(), bson.M{"_id": objectId})
	if result.Err() != nil {
		return &user, result.Err()
	}

	return &user, result.Decode(&user)
}

func (r *User) GetAll() ([]*domain.User, error) {
	var users []*domain.User

	cursor, err := r.MongoClient.Database(userDBName).Collection(userCollection).Find(context.Background(), bson.D{})
	if err != nil {
		return users, err
	}

	return users, cursor.All(context.Background(), &users)
}
