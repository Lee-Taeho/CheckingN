package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInterface interface {
	Connect() (*mongo.Client, error)
	Stop() error
}
