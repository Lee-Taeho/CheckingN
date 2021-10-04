package handlers

import (
	"net/http"
	"server/middleware/interfaces"
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
	gr.AddRoute("/api/google_login_request", http.MethodGet, handlers.GoogleLoginRequest)
	gr.AddRoute("/api/callback", http.MethodGet, handlers.GoogleLoginCallback)

}
