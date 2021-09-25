package handlers

import (
	"net/http"
	"server/middleware/interfaces"
)

type Handlers struct {
	// add db
	db interfaces.MongoInterface
}

func NewHandlers(gr interfaces.Router, db interfaces.MongoInterface) {
	handlers := new(Handlers)
	handlers.db = db

	gr.AddRoute("/api/example_json_response", http.MethodGet, handlers.ExampleJsonReponse)
	gr.AddRoute("/api/save_new_user", http.MethodPost, handlers.SaveNewUser)
	gr.AddRoute("/api/login_request", http.MethodPost, handlers.LoginRequest)
}
