package handlers

import (
	"net/http"
   	"encoding/json"
	"github.com/gorilla/mux"
	"time"
	"strconv"
)

func (h *Handlers) GetTutorsByCourseAndDate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["course_code"]
	year, _ := strconv.Atoi(params["year"])
	month, _ := strconv.Atoi(params["month"])
	day, _ := strconv.Atoi(params["day"])
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	tutors := h.db.GetTutorsByCourseAndDate(code, date)

  	w.WriteHeader(http.StatusOK)
   	json.NewEncoder(w).Encode(tutors)
}