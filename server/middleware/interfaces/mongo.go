package interfaces

import (
	"server/middleware"
	"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInterface interface {
	Connect() (*mongo.Client, error)
	Stop() error
	CreateNewStudent(student middleware.Student) error
	FindStudent(login middleware.LoginRequest) *middleware.Student
	GetUUID() int
	FindStudentUUID(uuid int) *middleware.Student
	AddAppointment(appointment middleware.Appointment) (string, string, error)
	GetAppointment(id string) (*middleware.Appointment, error)
	DeleteAppointment(id string) error
	DeleteAppointmentFromTutor(app_id string, tutor_email string) error
	DeleteAppointmentFromStudent(app_id string, student_email string) error
	GetAppointmentsForStudent(student_id string) ([]string, error)
	GetDepartments() []middleware.Department
	GetCoursesByDepartment(department_name string) []middleware.Course
	GetCoursesGroupedByDepartments() map[string][]middleware.Course
	GetTutorsByCourseAndDate(course_code string, date time.Time) []middleware.Tutor
}
