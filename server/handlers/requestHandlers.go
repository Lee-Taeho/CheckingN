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
	var courses []*middleware.Course

	course := new(middleware.Course)
	course.Name = "CS160"
	course.Department = "computer science"
	courses = append(courses, course)

	course2 := new(middleware.Course)
	course2.Name = "PHIL134"
	course2.Department = "philosophy"
	courses = append(courses, course2)

	fmt.Fprint(w, utils.Jsonify(courses))
}
