package data

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (l *LogEntry) Insert(entry LogEntry) error {
	log.Println("mongodbdata insertinng")
	collection := client.Database("logs").Collection("logs")

	data, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error inserting into logs:", err)
		return err
	}
	log.Println("mongodbdata", data)
	return nil
}

func (l *LogEntry) All() ([]LogEntry, error) {
	collection := client.Database("logs").Collection("logs")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("error finding logs:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var logs []LogEntry

	for cursor.Next(context.TODO()) {
		var entry LogEntry

		err := cursor.Decode(&entry)
		if err != nil {
			log.Println("error decoding log:", err)
			return nil, err
		}

		logs = append(logs, entry)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func (l *LogEntry) GetOne(id string) (*LogEntry, error) {
	collection := client.Database("logs").Collection("logs")

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var entry LogEntry

	err = collection.FindOne(
		context.TODO(),
		bson.M{"_id": objectID},
	).Decode(&entry)

	if err != nil {
		log.Println("error finding log:", err)
		return nil, err
	}

	return &entry, nil
}

func (l *LogEntry) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("logs").Collection("logs")

	if err := collection.Drop(ctx); err != nil {
		log.Println("error dropping collection:", err)
		return err
	}

	return nil
}
func (l *LogEntry) Update() error {
	collection := client.Database("logs").Collection("logs")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": l.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"name":       l.Name,
			"data":       l.Data,
			"updated_at": time.Now(),
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("error updating log:", err)
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no log entry found")
	}

	return nil
}
