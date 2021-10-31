package interfaces

import (
	"server/middleware"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInterface interface {
	Connect() (*mongo.Client, error)
	Stop() error
	CreateNewStudent(student middleware.Student)
	CreateNewGoogleStudent(student middleware.GoogleUser)
	FindStudent(login middleware.LoginRequest) *middleware.Student
	GetUUID() int
	FindStudentUUID(uuid int) *middleware.Student
	// FindUser(login middleware.LoginRequest, collection *mongo.Collection) bool
	AddAppointment(appointment middleware.Appointment) error
	GetAppointment(id string) (*middleware.Appointment, error)
	DeleteAppointment(id string) error
	DeleteAppointmentFromTutor(app_id string, tutor_id primitive.ObjectID) error
	DeleteAppointmentFromStudent(app_id string, student_id primitive.ObjectID) error
	GetAppointmentsForTutor(tutor_id string) ([]middleware.Appointment, error)
	GetAppointmentsForStudent(student_id string) ([]middleware.Appointment, error)
}
