package handlers

import (
	"net/http"
	"server/middleware/interfaces"
)

type HandlersEngine struct {
	// add db
}

func NewEngine(gr interfaces.Router) {
	engine := new(HandlersEngine)
	gr.AddRoute("/api/save_new_user", http.MethodPost, engine.SaveNewUser)
	gr.AddRoute("/api/login_request", http.MethodGet, engine.LoginRequest)
}
