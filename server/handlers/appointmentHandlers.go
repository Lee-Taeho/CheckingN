package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/middleware"
	"time"
	"github.com/gorilla/mux"
)

var loc, _ = time.LoadLocation("America/Los_Angeles")

func (h *Handlers) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment middleware.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		log.Println("ERROR [handlers/appointmentHandlers.go] Couldn't get data: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	appointment.StartTime = appointment.StartTime.In(loc)
	appointment.EndTime = appointment.EndTime.In(loc)
	if (appointment.StartTime.Before(time.Now())) {
		log.Println("ERROR [handlers/appointmentHandlers.go] Cannot create appointment for date that passed")
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err := h.db.AddAppointment(appointment); err != nil {
		log.Printf("Couldn't Create Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) CancelAppointment(w http.ResponseWriter, r *http.Request) {
	appointmentId := mux.Vars(r)["id"]

	if err := h.db.DeleteAppointment(appointmentId); err != nil {
		log.Printf("Couldn't Delete Appointment Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) ViewAppointment(w http.ResponseWriter, r *http.Request) {
	appointmentId := mux.Vars(r)["id"]

	appointment, err := h.db.GetAppointment(appointmentId)
	if err != nil {
		log.Printf("Couldn't View Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(appointment)
}

func (h *Handlers) ViewAllStudentAppointment(w http.ResponseWriter, r *http.Request) {
	studentEmail := mux.Vars(r)["id"]

	appointments, err := h.db.GetAppointmentsForStudent(studentEmail)
	if err != nil {
		log.Printf("Couldn't Get Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("start")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

/*
func (h *Handlers) ViewAllTutorAppointment(w http.ResponseWriter, r *http.Request) {
	tutorId := mux.Vars(r)["id"]

	appointments, err := h.db.GetAppointmentsForTutor(tutorId)
	if err != nil {
		log.Printf("Couldn't Get Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, appointment := range appointments {
		fmt.Fprintln(w, appointment.TutorEmail)
		fmt.Fprintln(w, appointment.StudentEmail)
		fmt.Fprintln(w, appointment.CourseCode)
		fmt.Fprintln(w, appointment.MeetingLocation)
		fmt.Fprintln(w, appointment.StartTime.String())
		fmt.Fprintln(w, appointment.EndTime.String())
	}
	w.WriteHeader(http.StatusOK)
}*/

/*
func (h *Handlers) EditAppointment(w http.ResponseWriter, r *http.Request) {
	appointmentId := mux.Vars(r)["id"]
	var appointment middleware.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		log.Println("ERROR [handlers/appointmentHandlers.go] Couldn't get data: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.db.UpdateAppointment(appointmentId, appointment); err != nil {
		log.Printf("Couldn't Edit Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusConflict)
		return
	}

	log.Println("Successfully edit appointment")

}*/
