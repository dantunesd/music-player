package infra

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBAdapter struct {
	Client     *mongo.Client
	DBName     string
	Collection string
}

func (d *MongoDBAdapter) Get(fieldName, fieldValue string, output interface{}) error {
	return d.findOne(fieldName, fieldValue, output)
}

func (d *MongoDBAdapter) Create(content interface{}) (string, error) {
	return d.insertOne(content)
}

func (d *MongoDBAdapter) GetAll(output interface{}) error {
	return d.find(output)
}

func (d *MongoDBAdapter) findOne(fieldName, fieldValue string, output interface{}) error {
	objectId, _ := primitive.ObjectIDFromHex(fieldValue)
	return d.getCollection().FindOne(context.TODO(), bson.M{fieldName: objectId}).Decode(output)
}

func (d *MongoDBAdapter) insertOne(output interface{}) (string, error) {
	result, err := d.getCollection().InsertOne(context.TODO(), output)
	return d.getIDFromResult(result, err)
}

func (d *MongoDBAdapter) find(output interface{}) error {
	cursor, err := d.getCollection().Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}
	return cursor.All(context.TODO(), output)
}

func (d *MongoDBAdapter) getCollection() *mongo.Collection {
	return d.Client.Database(d.DBName).Collection(d.Collection)
}

func (d *MongoDBAdapter) getIDFromResult(result *mongo.InsertOneResult, err error) (string, error) {
	if err != nil {
		return "", err
	}

	ID, _ := result.InsertedID.(primitive.ObjectID)
	return ID.Hex(), err
}
