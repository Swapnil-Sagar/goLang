package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Swapnil-Sagar/mongoapi/model"
	"github.com/gorilla/mux"
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

//get all movies from db 

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	
}

defer cursor.Close(context.Background())

return movies
}


// Actual controller - file 

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	json.NewEncoder(w).Encode(getAllMovies())

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
    w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

    count := deleteAllMovies()
    json.NewEncoder(w).Encode(map[string]string{"message": "All movies deleted successfully", "count": fmt.Sprintf("%v", count)})
}