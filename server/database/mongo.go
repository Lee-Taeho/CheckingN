package database

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONNECT_TIMEOUT            = 1 * time.Minute
	USER_DATABASE              = "users"
	STUDENTS_COLLECTION        = "students"
	TUTORS_COLLECTION          = "tutors"
	GOOGLE_STUDENTS_COLLECTION = "google_students"
	UUID_COLLECTION            = "uuid"
)

type MongoDBLogin struct {
	Uri string
	// could add server cert path fields here after getting it from mongo atlas
}

// this will "implement" MongoInterface in the interfaces directory
// in golang, to implement in interface just use all the methods defined by it (no need for implements keyword like in Java)
// someone told me the idea of this is "If it swims like a duck, it's a duck" (if Mongo looks like MongoInterface, it is a MongoInterface)
type MongoDB struct {
	login    *MongoDBLogin
	mongo    *mongo.Client
	uuidLock sync.Mutex
}

func NewMongoDB(mongoInfo *MongoDBLogin) *MongoDB {
	m := new(MongoDB)
	m.login = mongoInfo
	return m
}

func (m *MongoDB) Connect() (*mongo.Client, error) {
	log.Println("INFO [database/mongo.go] Connecting to Mongo")
	clientOptions := options.Client().ApplyURI(m.login.Uri)

	ctx, cancel := context.WithTimeout(context.Background(), CONNECT_TIMEOUT)
	defer cancel()

	var err error
	m.mongo, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("ERROR [database/mongo.go] Coudln't Connect To MongoDB:", err.Error())
		return nil, err
	}

	log.Println("INFO [database/mongo.go] Connected To MongoDB")
	return m.mongo, err
}

func (m *MongoDB) Stop() error {
	log.Println("INFO [database/mongo.go] Stopping MongoDB")
	return m.mongo.Disconnect(context.TODO())
}
