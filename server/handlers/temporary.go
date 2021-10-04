package handlers

import (
	"fmt"
	"net/http"
	"server/middleware"
	"server/utils"
)

// THESE ARE HANDLERS THAT ARE EITHER EXAMPLES OR WAITING FOR SOME FRONT END IMPLEMENTATION

func (h *Handlers) ExampleJsonReponse(w http.ResponseWriter, r *http.Request) {
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

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if h.tokenValid(w, r) {
		fmt.Fprint(w, "<h1>Successful Login!<h1>")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (h *Handlers) Google(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html><body><a href="/api/google_login_request">Google Log In</a></body></html>`
	fmt.Fprintf(w, htmlIndex)
}
