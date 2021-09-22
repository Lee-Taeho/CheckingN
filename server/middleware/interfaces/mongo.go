package interfaces

import (
	"server/middleware"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInterface interface {
	Connect() (*mongo.Client, error)
	Stop() error
	CreateNewStudent(student middleware.Student) error
}
