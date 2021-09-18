package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONNECT_TIMEOUT = 1 * time.Minute
)

type MongoInfo struct {
	Uri string
	// ServerCertPath string
	// ClientCertPath string
	// ClientKeyPath  string
}

// this will "implement" MongoInterface in the interfaces directory
// in golang, to implement in interface just use all the methods defined by it (no need for implements keyword like in Java)
// someone told me the idea of this is "If it swims like a duck, it's a duck" (if Mongo looks like MongoInterface, it is a MongoInterface)
type Mongo struct {
	mongoInfo   *MongoInfo
	mongoHandle *mongo.Client
}

func NewMongo(mongoInfo *MongoInfo) *Mongo {
	m := new(Mongo)
	m.mongoInfo = mongoInfo
	return m
}

func (m *Mongo) Connect() (*mongo.Client, error) {
	log.Println("INFO [database/mongo.go] Connecting to Mongo")
	clientOptions := options.Client().ApplyURI(m.mongoInfo.Uri)

	connectCtx, connectCancel := context.WithTimeout(context.Background(), CONNECT_TIMEOUT)
	defer connectCancel()

	var err error
	m.mongoHandle, err = mongo.Connect(connectCtx, clientOptions)
	if err != nil {
		log.Println("ERROR [database/mongo.go] Failed to connect to mongo: %s", err.Error())
		return nil, err
	}

	log.Println("INFO [database/mongo.go] Successfully Connected To Mongo")
	return m.mongoHandle, err
}

func (m *Mongo) Stop() error {
	log.Println("INFO [database/mongo.go] Stop Mongo")
	return m.mongoHandle.Disconnect(context.TODO())
}
