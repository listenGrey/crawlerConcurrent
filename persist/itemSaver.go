package persist

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func ItemSaver(uri, database, collection string) (chan interface{}, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	//defer client.Disconnect(context.TODO())
	userCollection := client.Database(database).Collection(collection)

	out := make(chan interface{})
	go func() {
		itemCount := 1
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			err, where := save(item, userCollection)
			if err != nil {
				log.Printf("Item Saver : error, " + where)
			}
			itemCount++
		}
	}()
	return out, nil
}

func save(item interface{}, userCollection *mongo.Collection) (err error, where string) {
	user, err := bson.Marshal(item)
	if err != nil {
		return err, "failure to convert !"
	}

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err, "id is exist !"
	}

	return nil, ""
}
