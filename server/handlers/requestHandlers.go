package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func (h *HandlersEngine) SaveNewUser(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlersEngine) LoginRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO [handlers/requestHandlers.go] Login Request")
	fmt.Fprint(w, "not so fast sir, log in please")
}
