package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"example.com/accounting/src/db/dao"
)

type DB struct {
	client *mongo.Client
}

type CreateUser struct {
	Email string
	User  string
	Pass  string
}

func MakeDB() DB {

	ctx := context.TODO()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	return DB{client}
}
func (db DB) GetUser(email string) (dao.UserDAO, error) {
	coll := db.client.Database("Accounting").Collection("User")

	var res dao.UserDAO
	err := coll.FindOne(context.TODO(), bson.D{{Key: "Email", Value: email}}).Decode(&res)

	return res, err
}

func (db DB) CreateUser(createUser CreateUser) (dao.UserDAO, error) {
	coll := db.client.Database("Accounting").Collection("User")

	doc := bson.D{{"User", createUser.User}, {"Pass", createUser.Pass}, {"Email", createUser.Email}}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return dao.UserDAO{}, err
	}

	var res dao.UserDAO
	err = coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&res)

	return res, err
}
