package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (h *Handlers) GetTutorsByCourseAndDate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["course_code"]
	year, _ := strconv.Atoi(params["year"])
	month, _ := strconv.Atoi(params["month"])
	day, _ := strconv.Atoi(params["day"])
	date_check := time.Now()
	date := time.Date(year, time.Month(month), day, 23, 59, 59, 0, loc)
	if date.Before(date_check) {
		w.WriteHeader(http.StatusConflict)
		return
	}

	tutors := h.db.GetTutorsByCourseAndDate(code, date)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tutors)
}
