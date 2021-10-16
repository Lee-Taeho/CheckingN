package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/middleware"
	"server/utils"
	"time"
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

	if err := h.db.CreateNewStudent(*student); err != nil {
		log.Printf("%s Couldn't Save New User: %s\n", LOGGER_ERROR_LOGIN, err.Error())
		w.WriteHeader(http.StatusSeeOther)
	}

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

	cookie := &middleware.Cookie{
		Name:      "Bearer",
		Value:     Encrypt(aes_key, fmt.Sprint(student.Uuid)),
		ExpiresAt: time.Now().Add(AUTO_LOGOUT_TIME).Unix(),
	}
	w.Write([]byte(utils.Jsonify(cookie)))

	log.Println(LOGGER_INFO_LOGIN + " Log In Successful")
}

// func (h *Handlers) Refresh(w http.ResponseWriter, r *http.Request) {
// 	cookie, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	tokenStr := cookie.Value

// 	claims := &Claims{}

// 	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
// 		func(t *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})

// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > TIME_BEFORE_EXPIRED {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	expirationTime := time.Now().Add(AUTO_LOGOUT_TIME)

// 	claims.ExpiresAt = expirationTime.Unix()

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	http.SetCookie(w,
// 		&http.Cookie{
// 			Name:    "token",
// 			Value:   tokenString,
// 			Expires: expirationTime,
// 		})

// }
