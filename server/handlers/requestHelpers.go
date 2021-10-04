package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	AUTO_LOGOUT_TIME    = time.Minute * 20
	TIME_BEFORE_EXPIRED = time.Minute * 30
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
			return secret_key, nil
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

func (h *Handlers) createToken(w http.ResponseWriter, r *http.Request, email string) {
	expirationTime := time.Now().Add(AUTO_LOGOUT_TIME)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret_key)

	if err != nil {
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
