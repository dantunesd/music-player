package infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDBAdapter struct {
	client     *mongo.Client
	dbName     string
	collection string
}

func NewMongoDBAdapter(client *mongo.Client, dbName, collection string) *mongoDBAdapter {
	return &mongoDBAdapter{
		client:     client,
		dbName:     dbName,
		collection: collection,
	}
}

func (d *mongoDBAdapter) Get(fieldName, fieldValue string, output interface{}) error {
	return d.findOne(fieldName, fieldValue, output)
}

func (d *mongoDBAdapter) Create(content interface{}) (string, error) {
	return d.insertOne(content)
}

func (d *mongoDBAdapter) GetAll(output interface{}) error {
	return d.find(output)
}

func (d *mongoDBAdapter) findOne(fieldName, fieldValue string, output interface{}) error {
	objectId, _ := primitive.ObjectIDFromHex(fieldValue)
	return d.getCollection().FindOne(context.TODO(), bson.M{fieldName: objectId}).Decode(output)
}

func (d *mongoDBAdapter) insertOne(output interface{}) (string, error) {
	result, err := d.getCollection().InsertOne(context.TODO(), output)
	return d.getIDFromResult(result, err)
}

func (d *mongoDBAdapter) find(output interface{}) error {
	cursor, err := d.getCollection().Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}
	return cursor.All(context.TODO(), output)
}

func (d *mongoDBAdapter) getCollection() *mongo.Collection {
	return d.client.Database(d.dbName).Collection(d.collection)
}

func (d *mongoDBAdapter) getIDFromResult(result *mongo.InsertOneResult, err error) (string, error) {
	if err != nil {
		return "", err
	}

	ID, _ := result.InsertedID.(primitive.ObjectID)
	return ID.Hex(), err
}
