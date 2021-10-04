package handlers

import (
	"log"
	"net/http"
	"server/middleware"
)

func (h *Handlers) SaveNewUser(w http.ResponseWriter, r *http.Request) {
	log.Println(LOGGER_INFO_LOGIN + " Request to Save New User")
	r.ParseForm()
	student := &middleware.Student{
		FirstName: r.PostForm.Get("first_name"),
		LastName:  r.PostForm.Get("last_name"),
		Email:     r.PostForm.Get("email"),
		Password:  r.PostForm.Get("password"),
	}

	if err := h.db.CreateNewStudent(*student); err != nil {
		log.Printf("%s Couldn't Save New User: %s\n", LOGGER_ERROR_LOGIN, err.Error())
		http.Error(w, "Couldn't Save New User", http.StatusInternalServerError)
		return
	}
	log.Println(LOGGER_INFO_LOGIN + " Successfully Saved New User")
	// fmt.Fprint(w, "Thanks For signing up for CheckingN!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) LoginRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(LOGGER_INFO_LOGIN + " Request to Log In")
	r.ParseForm()
	login := &middleware.LoginRequest{
		Email:    r.PostForm.Get("Email"),
		Password: r.PostForm.Get("Password"),
	}

	if found := h.db.FindStudent(*login); !found {
		log.Println(LOGGER_INFO_LOGIN + " Failed Log In")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println(LOGGER_INFO_LOGIN + " Log In Successful")

	h.createTokenAndSetCookie(w, login.Email)
	h.Home(w, r)
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
