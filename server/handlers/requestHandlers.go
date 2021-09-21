package handlers

import (
	"fmt"
	"net/http"
	"server/middleware"
	"server/utils"
)

func (h *Handlers) SaveNewUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) LoginRequest(w http.ResponseWriter, r *http.Request) {
	course := new(middleware.Course)
	course.Name = "CS160"
	course.Department = "computer science"
	fmt.Fprint(w, utils.Jsonify(course))
}
