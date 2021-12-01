package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/middleware"
	"server/utils"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func (h *Handlers) GoogleCalendarEventPost(info middleware.GoogleCalendarEventInfo) error {
	tok := &oauth2.Token{
		AccessToken: info.AccessToken,
		TokenType:   info.TokenType,
		Expiry:      info.Expiry,
	}

	client := googleOauthConfig.Client(context.TODO(), tok)
	ctx := context.TODO()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Printf("ERROR [handlers/googleHandlers.go] Token Doesn't Work: %s\n", err.Error())
		return err
	}

	event := &calendar.Event{
		Id:          info.ID,
		Summary:     "Tutoring Appointment",
		Location:    info.MeetingLocation,
		Description: info.JoinLink,
		Start: &calendar.EventDateTime{
			DateTime: info.StartTime.Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: info.EndTime.Format(time.RFC3339),
		},
		Attendees: []*calendar.EventAttendee{
			&calendar.EventAttendee{Email: info.TutorEmail},
			&calendar.EventAttendee{Email: info.StudentEmail},
		},
	}

	calendarId := "primary"
	if _, err = srv.Events.Insert(calendarId, event).Do(); err != nil {
		return err
	}
	return nil
}

func (h *Handlers) GoogleCalendarEventDelete(info middleware.GoogleCalendarEventInfo) {
	tok := &oauth2.Token{
		AccessToken: info.AccessToken,
		TokenType:   info.TokenType,
		Expiry:      info.Expiry,
	}

	client := googleOauthConfig.Client(context.TODO(), tok)
	ctx := context.TODO()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Printf("ERROR [handlers/googleHandlers.go] Couldn't get data: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := srv.Events.Delete("primary", info.ID).Do(); err != nil {
		log.Printf("ERROR [handlers/googleHandlers.go] Couldn't delete event: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handlers) GoogleLoginInfoSaver(w http.ResponseWriter, r *http.Request) {
	googStudent := &middleware.GoogleUser{}
	if err := json.NewDecoder(r.Body).Decode(googStudent); err != nil {
		log.Printf("ERROR [handlers/requestHandlers.go] Couldn't get data: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uuid := h.db.GetUUID()
	user := &middleware.Student{
		Uuid:      uuid,
		FirstName: googStudent.FirstName,
		LastName:  googStudent.LastName,
		Email:     googStudent.Email,
	}
	h.db.CreateNewStudent(*user)

	header := &middleware.Header{
		Key:   "Authorization",
		Value: "Bearer " + encrypt(aes_key, fmt.Sprint(uuid)),
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(utils.Jsonify(header)))
	log.Println(LOGGER_INFO_LOGIN + " Log In Successful")
}

func (h *Handlers) GoogleLoginRequest(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(rANDOM_STATE)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handlers) GoogleLoginCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != rANDOM_STATE {
		log.Println(LOGGER_INFO_GOOGLE + " State is invalid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := googleOauthConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		log.Printf("%s Could not get token %s", LOGGER_ERROR_GOOGLE, err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Printf("%s Could not create get user info %s", LOGGER_ERROR_GOOGLE, err.Error())
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}
	defer resp.Body.Close()

	googStudent := h.googleRespDecoder(*resp)
	uuid := h.db.GetUUID()
	if uuid == 0 {
		log.Println(LOGGER_ERROR_GOOGLE + " error creating uuid")
		return
	}

	user := &middleware.Student{
		Uuid:      uuid,
		FirstName: googStudent.FirstName,
		LastName:  googStudent.LastName,
		Email:     googStudent.Email,
	}
	h.db.CreateNewStudent(*user)

	header := &middleware.Header{
		Key:   "Authorization",
		Value: "Bearer " + encrypt(aes_key, fmt.Sprint(uuid)),
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write([]byte(utils.Jsonify(header)))
	log.Println(LOGGER_INFO_LOGIN + " Log In Successful")
}
