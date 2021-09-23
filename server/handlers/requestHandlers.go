package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"server/middleware"
	"server/utils"
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
	r.ParseForm()
	student := &middleware.Student{
		FirstName: r.PostForm.Get("first_name"),
		LastName:  r.PostForm.Get("last_name"),
		Email:     r.PostForm.Get("email"),
		Password:  r.PostForm.Get("password"),
	}

	if err := h.db.CreateNewStudent(*student); err != nil {
		if err == errors.New("email already exist") {
			w.WriteHeader(http.StatusConflict)
			return
		}
		log.Printf("ERROR [handlers/requestHandlers.go] Couldn't Save New User: %s\n", err.Error())
		http.Error(w, "Couldn't Save New User", http.StatusInternalServerError)
		return
	}
	log.Println("INFO [handlers/requestHandlers.go] Successfully Saved New User")
	// fmt.Fprint(w, "Thanks For signing up for CheckingN!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) LoginRequest(w http.ResponseWriter, r *http.Request) {

}
