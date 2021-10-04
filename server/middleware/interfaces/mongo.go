package interfaces

import (
	"server/middleware"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInterface interface {
	Connect() (*mongo.Client, error)
	Stop() error
	CreateNewStudent(student middleware.Student) error
	CreateNewGoogleStudent(student middleware.GoogleUser)
	FindStudent(login middleware.LoginRequest) bool
	FindUser(login middleware.LoginRequest, collection *mongo.Collection) bool
}
