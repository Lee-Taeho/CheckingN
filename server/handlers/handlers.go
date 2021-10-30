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

	gOOGLE_CLIENT_ID      = "533962375262-dvbofsom1ocmc8mjb6aq4alfd7rdisi0.apps.googleusercontent.com"
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
	aes_iv  = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
	aes_key = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
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
}
