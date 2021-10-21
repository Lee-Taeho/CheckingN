package middleware

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
// bson tag will tell golang to extract that specific field from mongodb into the variable
// json tag will be the way it is called from front end
type Tutor struct {
	FirstName       string     `bson:"first_name" json:"first_name"`
	LastName        string     `bson:"last_name" json:"last_name"`
	Students        []string   `bson:"students" json:"students"`
	FluentLanguages []string   `bson:"fluent_languages" json:"fluent_languages"`
	Courses         []string   `bson:"courses" json:"courses"`
	Availability	[][]int    `bson:"availability" json:"availability"`
	Appointments    []string   `bson:"appointments" json:"appointments"`
}

type Student struct {
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password"`
}

type GoogleUser struct {
	Id            string `json:"id" bson:"id"`
	Email         string `json:"email" bson:"email"`
	VerifiedEmail bool   `json:"verified_email" bson:"verified_email"`
	FirstName     string `json:"given_name" bson:"first_name"`
	LastName      string `json:"family_name" bson:"last_name"`
	PictureLink   string `json:"picture" bson:"picture_link"`
	Locale        string `json:"locale" bson:"locale"`
}

type LoginRequest struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type Course struct {
	Name       string `bson:"name" json:"name"`
	Department string `bson:"department" json:"department"`
}

type Appointment struct {
	TutorID			string 		`bson:"tutor_id" json:"tutor_id"`
	StudentID		string 		`bson:"student_id" json:"student_id"`
	CourseID		string 		`bson:"course_id" json:"course_id"`
	MeetingLocation string 		`bson:"meeting_location" json:"meeting_location"`
	StartTime		primitive.DateTime 	`bson:"start_time" json:"start_time"`
	EndTime			primitive.DateTime 	`bson:"end_time" json:"end_time"`
}
