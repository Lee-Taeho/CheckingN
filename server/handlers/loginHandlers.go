package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/middleware"
	"server/utils"
)

func (h *Handlers) SaveNewUser(w http.ResponseWriter, r *http.Request) {
	log.Println(LOGGER_INFO_LOGIN + " Request to Save New User")

	student := &middleware.Student{}
	if err := json.NewDecoder(r.Body).Decode(student); err != nil {
		log.Printf("ERROR [handlers/requestHandlers.go] Couldn't get data: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	uuid := h.db.GetUUID()
	student.Uuid = uuid

	h.db.CreateNewStudent(*student)
	log.Println(LOGGER_INFO_LOGIN + " Successfully Saved New User")
}

func (h *Handlers) LoginRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(LOGGER_INFO_LOGIN + " Request to Log In")

	login := &middleware.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(login); err != nil {
		log.Printf("ERROR [handlers/requestHandlers.go] Couldn't get data: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	student := h.db.FindStudent(*login)
	if student == nil {
		log.Println(LOGGER_INFO_LOGIN + " Failed Log In")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	header := &middleware.Header{
		Key:   "Authorization",
		Value: "Bearer " + encrypt(aes_key, fmt.Sprint(student.Uuid)),
	}

	w.Write([]byte(utils.Jsonify(header)))
	log.Println(LOGGER_INFO_LOGIN + " Log In Successful")
}

func (h *Handlers) Authorized(w http.ResponseWriter, r *http.Request) {
	student := h.authorized(r)
	if student != nil {
		log.Println(LOGGER_INFO_LOGIN, "Student Authorized")
		w.Write([]byte(utils.Jsonify(student)))
	} else {
		log.Println(LOGGER_INFO_LOGIN, "Student Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
	}
}
