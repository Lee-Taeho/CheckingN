package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func (h *Handlers) SaveNewUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) LoginRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO [handlers/requestHandlers.go] Login Request")
	fmt.Fprint(w, "<h1>Not so fast sir, log in please</h1>")
}
