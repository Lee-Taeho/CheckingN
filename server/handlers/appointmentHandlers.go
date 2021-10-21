package handlers

import (
	"net/http"
	"server/middleware"
	"log"
	"encoding/json"
)

func (h *Handlers) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	appointment := &middleware.Appointment{}
	if err := json.NewDecoder(r.Body).Decode(appointment); err != nil {
		log.Println("ERROR [handlers/appointmentHandlers.go] Couldn't get data: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.db.AddAppointment(*appointment); err != nil {
		log.Printf("Couldn't Create Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusConflict)
		return
	}

	log.Println("Successfully booked new appointment")
}


/* Testing
	appointment := &middleware.Appointment{
		TutorID: "616f427a3b3c421b64576b51",
		StudentID: "0000000",
		CourseID: "CS146",
		MeetingLocation: "Zoom",
		StartTime: time.Date(
			2021, 10, 18, 9, 0, 0, 0, loc),
		EndTime: time.Date(
			2021, 10, 18, 10, 0, 0, 0, loc),
	}

*/