package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Swapnil-Sagar/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)




const connectionString = "mongodb+srv://swapnilsagar09:swapnilSagar123@cluster0.ycfvj.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflixdb"
const colName = "watchlist"

//Important 
var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance 
	fmt.Println("Collection instance is ready")

}

//MONGODB helpers - files

// insert 1 record 

func insertOneMovie(movie model.Netflix) {
	inserted, err:= collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	 fmt.Println("Inserted one movie with id: ", inserted.InsertedID)
}


// update 1 record

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
        log.Fatal(err)
    }	
	fmt.Printf("Updated %v documents\n", result.ModifiedCount)
}


// delete 1 record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents\n", deleteCount)
}


// delete all record

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {	
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents\n", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

