package middleware

import (
	"github.com/ameena-zehra/golang-react-todo/server/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)
var collection *mongoCollection
func init (){
	loadTheEnv()
	createDBinstance()
}
func loadTheEnv(){
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading the .env file")
	}
}
// Defines a function that establishes a connection with a MongoDB server
// creates a collection instance for performing database operations
// a collection is a grouping of mongoDB documents with its own structure fields and datatypes

func createDBinstance(){
	connectionString:=os.Getenv("DB_URI")
	dbName:= os.Getenv("DB_NAME")
	collName:= os.Getenv("DB_COLLECTION_NAME")
	
	clientOptions:= options.Client().ApplyURL(connectionString)
	client, err:= mongo.Connect(context.TODO(), clientOptions)
	if err!= nil{
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to mongodb!")
	collection = client.Database(dbName).Collection(collName)
	fmt.Println("collection instance created")

}
func GetAllTasks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload:= getAllTasks()
	json.NewEncoder(w).Encode(payload)
}
// HTTP handler that is called when a POST request is made to a specific endpoint
// sets the necessary response headers for CORS
// decodes the JSON data from the request body into a task variable 
// w http.ResponseWriter: allows the handler to construct and send an http response back to the client
// r http.Request: contains info about the incoming request 
func CreateTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// specifies the allowed origins that can access the resource
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	// specifies the allowed http methods for the cross origin request
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	// decodes the request body and populates the task variable with the decoded data
	json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}
func TaskComplete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// specifies the allowed origins that can access the resource
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// specifies the allowed http methods for the cross origin request
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params:=mux.Vars(r) // mux is commonly used as a GoHTTP routing library to handle URL parameters
	TaskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func UndoTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","PUT")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
	param:=mux.Vars(r)
	UndoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])


}
func DeleteTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","PUT")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
}
func DeleteAllTasks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	count: = deleteAllTasks()
	json.NewEncoder(w).Encode(count)
}

// Helper Functions
func getAllTasks() []primitive.M{
	curr, err := collection.Find(context.Background, , bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}
	var results []primitive.M
	for cur.Next(context.Background()){
		var result bson.M
		e := cur.Decode(&result)
		if e!=nil{
			log.Fatal(e)
		}
		results =append(results, result)
	}
	if err := curr.Err(); err!=nil{
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results

}
func TaskComplete(task string){
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id":id}
	update:= bson.M{"$set":bson.M{"status":true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}
func insertOneTask(task models.ToDoList){
	insertResult, err:=collection.insertOneTask(context.Background(), task)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Insert a single result", insertResult.InsertedID)

}
func UndoTask(task string){
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id":id}
	update:= bson.M{"$set":bson.M{"status":false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)

}
func deleteOneTask(task string){
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id":id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Deleted Document ", d.DeletedCount)

}
func deleteAllTasks() int64{
	d,err:= collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println("Deleted document", d.DeletedCount)
	return d.DeletedCount
}