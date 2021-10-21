package interfaces

import (
	"server/middleware"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoInterface interface {
	Connect() (*mongo.Client, error)
	Stop() error
	CreateNewStudent(student middleware.Student) error
	CreateNewGoogleStudent(student middleware.GoogleUser)
	FindStudent(login middleware.LoginRequest) bool
	FindUser(login middleware.LoginRequest, collection *mongo.Collection) bool
	AddAppointment(appointment middleware.Appointment) error
	GetAppointment(id string) (*middleware.Appointment, error)
	DeleteAppointment(id string) error
	DeleteAppointmentFromTutor(app_id primitive.ObjectID, tutor_id primitive.ObjectID) error
	DeleteAppointmentFromStudent(app_id primitive.ObjectID, student_id primitive.ObjectID) error
	GetAppointmentsForTutor(tutor_id string) ([]*middleware.Appointment, error) 
}
