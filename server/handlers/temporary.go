package handlers

import (
	"fmt"
	"log"
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
	log.Println(LOGGER_INFO_TEMPORARY + " Redirecting to Home Page")
	if h.tokenValid(w, r) {
		log.Println(LOGGER_INFO_TEMPORARY + " Login Token Valid")
		fmt.Fprint(w, "<h1>Successful Login!<h1>")
	} else {
		log.Println(LOGGER_ERROR_TEMPORARY + " Login Token Invalid")
	}
}

func (h *Handlers) Google(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html><body><a href="/api/google_login_request">Google Log In</a></body></html>`
	fmt.Fprintf(w, htmlIndex)
}
