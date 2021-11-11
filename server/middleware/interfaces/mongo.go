package interfaces

import (
	"server/middleware"

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
	AddAppointment(appointment middleware.Appointment) error
	GetAppointment(id string) (*middleware.Appointment, error)
	DeleteAppointment(id string) error
	DeleteAppointmentFromTutor(app_id string, tutor_email string) error
	DeleteAppointmentFromStudent(app_id string, student_email string) error
	GetAppointmentsForStudent(student_id string) ([]string, error)
	GetDepartments() []middleware.Department
	GetCoursesByDepartment(department_name string) []middleware.Course
	GetCoursesGroupedByDepartments() map[string][]middleware.Course

	//FindTutorId(tutor middleware.Tutor) int
	// FindUser(login middleware.LoginRequest, collection *mongo.Collection) bool
	//UpdateAppointment(app_id string, newAppointment middleware.Appointment) error
	//GetAppointmentsForTutor(tutor_id string) ([]middleware.Appointment, error)
}
