package handlers

import (
	"net/http"
	"server/middleware/interfaces"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	AUTO_LOGOUT_TIME    = time.Minute * 20
	TIME_BEFORE_EXPIRED = time.Minute * 30

	rANDOM_STATE = "random"

	gOOGLE_CLIENT_ID      = "533962375262-sn2l2op591vabl5i85f6vf7sptad47tt.apps.googleusercontent.com"
	gOOGLE_CLIENT_SECRET  = "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"
	GOOGLE_EMAIL_SCOPE    = "https://www.googleapis.com/auth/userinfo.email"
	GOOGLE_PROFILE_SCOPE  = "https://www.googleapis.com/auth/userinfo.profile"
	GOOGLE_CALENDAR_SCOPE = "https://www.googleapis.com/auth/calendar"

	LOGGER_ERROR_GOOGLE    = "ERROR [handlers/googleHandlers.go]"
	LOGGER_INFO_GOOGLE     = "INFO [handlers/googleHandlers.go]"
	LOGGER_ERROR_HANDLERS  = "ERROR [handlers/handlers.go]"
	LOGGER_INFO_HANDLERS   = "INFO [handlers/handlers.go]"
	LOGGER_ERROR_LOGIN     = "ERROR [handlers/loginHandlers.go]"
	LOGGER_INFO_LOGIN      = "INFO [handlers/loginHandlers.go]"
	LOGGER_ERROR_TEMPORARY = "ERROR [handlers/temporary.go]"
	LOGGER_INFO_TEMPORARY  = "INFO [handlers/temporary.go]"
	LOGGER_ERROR_HELPERS   = "ERROR [handlers/requestHelpers.go]"
	LOGGER_INFO_HELPERS    = "INFO [handlers/requestHelpers.go]"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/google_callback",
		ClientID:     gOOGLE_CLIENT_ID,
		ClientSecret: gOOGLE_CLIENT_SECRET,
		Scopes:       []string{GOOGLE_EMAIL_SCOPE, GOOGLE_PROFILE_SCOPE, GOOGLE_CALENDAR_SCOPE},
		Endpoint:     google.Endpoint,
	}
	aes_iv              = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
	aes_key             = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	aes_password_db_key = []byte{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
)

type Handlers struct {
	db            interfaces.MongoInterface
	hostIpBinding string
}

func NewHandlers(gr interfaces.Router, db interfaces.MongoInterface, ip string) {
	handlers := new(Handlers)
	handlers.db = db
	handlers.hostIpBinding = ip

	gr.AddRoute("/api/example_json_response", http.MethodGet, handlers.ExampleJsonReponse)
	gr.AddRoute("/api/save_new_user", http.MethodPost, handlers.SaveNewUser)
	gr.AddRoute("/api/login_request", http.MethodPost, handlers.LoginRequest)
	gr.AddRoute("/api/home", http.MethodPost, handlers.Home)
	gr.AddRoute("/api/google", http.MethodGet, handlers.Google)
	gr.AddRoute("/api/google_login_request", http.MethodPost, handlers.GoogleLoginInfoSaver)
	gr.AddRoute("/api/google_callback", http.MethodGet, handlers.GoogleLoginCallback)
	gr.AddRoute("/api/authorized", http.MethodGet, handlers.Authorized)
	gr.AddRoute("/api/appointment", http.MethodPost, handlers.CreateAppointment)
	gr.AddRoute("/api/appointment/{id}", http.MethodDelete, handlers.CancelAppointment)
	gr.AddRoute("/api/appointment/{id}", http.MethodGet, handlers.ViewAppointment)
	gr.AddRoute("/api/appointment/tutor/{id}", http.MethodGet, handlers.ViewAllTutorAppointment)
	gr.AddRoute("/api/appointment/student/{id}", http.MethodGet, handlers.ViewAllStudentAppointment)
	//gr.AddRoute("/api/appointment/{id}", http.MethodPatch, handlers.EditAppointment)
	gr.AddRoute("/api/courses_by_departments", http.MethodGet, handlers.GetCoursesGroupedByDepartments)
	gr.AddRoute("/api/{course_code}/tutors/{year}/{month}/{day}", http.MethodGet, handlers.GetTutorsByCourseAndDate)
}
