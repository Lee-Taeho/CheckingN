package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/middleware"
	"server/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (h *Handlers) tokenValid(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(sECRET_KEY), nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}

func (h *Handlers) createTokenAndSetCookie(w http.ResponseWriter, email string) {
	expirationTime := time.Now().Add(AUTO_LOGOUT_TIME)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(sECRET_KEY))

	if err != nil {
		log.Printf("%s Error signing token: %s\n", LOGGER_ERROR_HELPERS, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
}

func (h *Handlers) googleRespDecoder(resp http.Response) middleware.GoogleUser {
	contents := []byte(utils.JsonifyHttpResponse(resp))

	var user middleware.GoogleUser
	json.Unmarshal(contents, &user)

	return user
}
