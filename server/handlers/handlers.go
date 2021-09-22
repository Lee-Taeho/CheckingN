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
	// TODO: make the method of the next api POST
	gr.AddRoute("/api/save_new_user", http.MethodPost, handlers.SaveNewUser)
	gr.AddRoute("/api/login_request", http.MethodGet, handlers.LoginRequest)
}
