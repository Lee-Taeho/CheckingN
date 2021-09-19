package handlers

import (
	"net/http"
	"server/middleware/interfaces"
)

type Handlers struct {
	// add db
}

func NewHandlers(gr interfaces.Router) {
	engine := new(Handlers)
	gr.AddRoute("/api/save_new_user", http.MethodPost, engine.SaveNewUser)
	gr.AddRoute("/api/login_request", http.MethodGet, engine.LoginRequest)
}
