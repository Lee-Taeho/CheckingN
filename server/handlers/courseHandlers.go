package handlers

import (
 	"net/http"
	"encoding/json"
)

func (h *Handlers) GetCoursesGroupedByDepartments(w http.ResponseWriter, r *http.Request) {
	mapOfCourses := h.db.GetCoursesGroupedByDepartments()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapOfCourses)
}

/*
func (h *Handlers) GetCourses(w http.ResponseWriter, r *http.Request) {
 	params := mux.Vars(r)
 	dep := params["department"]

 	courses := h.db.GetCoursesByDepartment(dep)

 	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(courses)
} 

func (h *Handlers) GetDepartments(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting deps!")
 	deps := h.db.GetDepartments()

 	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(deps)
}
*/