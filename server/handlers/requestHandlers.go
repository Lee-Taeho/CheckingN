package handlers

import (
	"fmt"
	"log"
	"net/http"
	"server/middleware"
	"server/utils"

	"github.com/gorilla/mux"
)

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

func (h *Handlers) SaveNewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO [handlers/requestHandlers.go] Request to Save New User")
	student := new(middleware.Student)
	student.FirstName = mux.Vars(r)["firstName"]
	student.LastName = mux.Vars(r)["lastName"]
	student.LoginInfo = &middleware.LoginInfo{
		Email:    mux.Vars(r)["email"],
		Password: mux.Vars(r)["password"],
	}
	if err := h.db.CreateNewStudent(*student); err != nil {
		log.Printf("ERROR [handlers/requestHandlers.go] Couldn't Save New User: %s\n", err.Error())
		return
	}
	log.Println("INFO [handlers/requestHandlers.go] Successfully Saved New User")
	fmt.Fprint(w, "hello")
}

func (h *Handlers) LoginRequest(w http.ResponseWriter, r *http.Request) {

}
