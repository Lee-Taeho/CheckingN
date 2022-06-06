package middleware

import (
	"time"

	"golang.org/x/oauth2"
)

// bson tag will tell golang to extract that specific field from mongodb into the variable
// json tag will be the way it is called from front end
type Tutor struct {
	Uuid            int      `bson:"uuid" json:"uuid,omitempty"`
	Email           string   `bson:"email" json:"email"`
	FirstName       string   `bson:"first_name" json:"first_name"`
	LastName        string   `bson:"last_name" json:"last_name"`
	Students        []string `bson:"students" json:"students,omitempty"`
	FluentLanguages []string `bson:"fluent_languages" json:"fluent_languages,omitempty"`
	Courses         []string `bson:"courses" json:"courses"`
	Availability    [][]int  `bson:"availability" json:"availability"`
	Appointments    []string `bson:"appointments" json:"appointments,omitempty"`
}

type Student struct {
	Uuid         int      `bson:"uuid" json:"uuid,omitempty"`
	FirstName    string   `bson:"first_name" json:"first_name"`
	LastName     string   `bson:"last_name" json:"last_name"`
	Email        string   `bson:"email" json:"email"`
	Password     string   `bson:"password" json:"password"`
	Appointments []string `bson:"appointments" json:"appointments,omitempty"`
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
	ClassCode  string `bson:"class_code" json:"class_code"`
}

type Department struct {
	Name string `bson:"name" json:"name"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Appointment struct {
	TutorEmail      string    `bson:"tutor_email" json:"tutor_email"`
	StudentEmail    string    `bson:"student_email" json:"student_email"`
	CourseCode      string    `bson:"course_code" json:"course_code"`
	MeetingLocation string    `bson:"meeting_location" json:"meeting_location"`
	StartTime       time.Time `bson:"start_time" json:"start_time"`
	EndTime         time.Time `bson:"end_time" json:"end_time"`
	JoinLink        string    `bson:"join_link" json:"join_link,omitempty"`
	StartLink       string    `bson:"start_link" json:"start_link,omitempty"`
}

type GoogleCalendarEventInfo struct {
	Appointment
	oauth2.Token
	ID string `json:"appointment_id"`
}
