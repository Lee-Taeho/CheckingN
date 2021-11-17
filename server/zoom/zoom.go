package zoom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"server/middleware"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MeetingOptions struct {
	Topic       string         `json:"topic"`        //Topic
	Type        int            `json:"type"`         // 2 for scheduled meeting
	PreSchedule bool           `json:"pre_schedule"` //True
	StartTime   string         `json:"start_time" `
	Duration    int            `json:"duration"`     //120 minutes
	Email       string         `json:"schedule_for"` //email
	TimeZone    string         `json:"timezone"`
	Password    string         `json:"password"` // Max 10 characters. [a-z A-Z 0-9 @ - _ *]
	Agenda      string         `json:"agenda"`
	Settings    MeetingSetting `json:"settings"`
}
type MeetingSetting struct {
	HostVideo        bool   `json:"host_video"`
	ParticipantVideo bool   `json:"participant_video"`
	ChinaMeeting     bool   `json:"cn_meeting"`
	IndiaMeeting     bool   `json:"in_meeting"`
	JoinBeforeHost   bool   `json:"join_before_host"`
	JBHTime          int    `json:"jbh_time"`
	MuteUponEntry    bool   `json:"mute_upon_entry"`
	Watermark        bool   `json:"watermark"`
	UsePMI           bool   `json:"use_pmi"`        //false
	ApprovalType     int    `json:"approval_type"`  // 2
	Audio            string `json:"audio"`          //both
	AutoRecording    string `json:"auto_recording"` //local
	//EnforceLogin                 bool   `json:"enforce_login"`  //true
	WaitingRoom bool `json:"waiting_room"`
	//RegistrantsConfirmationEmail bool   `json:"registrants_confirmation_email"`
}

type ZoomLink struct {
	JoinLink  string `json:"join_url"`
	StartLink string `json:"start_url"`
}

func CreateZoomLink(appointment middleware.Appointment) (string, string, error) {
	jwtToken, _ := jwtToken(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	meeting := createMeetingOption(appointment)
	uri := fmt.Sprintf("https://api.zoom.us/v2/users/me/meetings")

	var bearer = "Bearer " + jwtToken
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(meeting)
	req, err := http.NewRequest("POST", uri, payloadBuf)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
	fmt.Println("response Status:", resp.Status)
	io.Copy(os.Stdout, resp.Body)

	zoomLink := ZoomLink{}
	json.Unmarshal([]byte(body), &zoomLink)

	joinLink := zoomLink.JoinLink
	startLink := zoomLink.StartLink
	return joinLink, startLink, nil

}

func jwtToken(key string, secret string) (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:    key,
		ExpiresAt: jwt.TimeFunc().Local().Add(time.Second * time.Duration(5000)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["alg"] = "HS256"
	token.Header["typ"] = "JWT"
	return token.SignedString([]byte(secret))
}

func createMeetingOption(appointment middleware.Appointment) *MeetingOptions {
	startTime := appointment.StartTime

	timeToString := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		startTime.Year(), startTime.Month(), startTime.Day(),
		startTime.Hour(), startTime.Minute(), startTime.Second())

	defaultMeetingSetting := &MeetingSetting{
		HostVideo:        false,
		ParticipantVideo: false,
		ChinaMeeting:     false,
		IndiaMeeting:     false,
		JoinBeforeHost:   true,
		JBHTime:          5,
		MuteUponEntry:    true,
		Watermark:        true,
		UsePMI:           true,
		ApprovalType:     2,
		Audio:            "both",
		AutoRecording:    "local",
		//EnforceLogin:                 true,
		WaitingRoom: true,
		//RegistrantsConfirmationEmail: true,
	}

	meeting := &MeetingOptions{
		Topic:       appointment.CourseCode,
		Type:        2,
		PreSchedule: true,
		StartTime:   timeToString,
		Duration:    120,
		//Email:    ignore for now, need Zoom license to appoint user
		//TimeZone: "Pacific/Midway",
		//Password: "TestPw@12",
		Agenda:   "Agenda",
		Settings: *defaultMeetingSetting,
	}
	return meeting
}
