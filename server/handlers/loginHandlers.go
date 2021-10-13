package handlers

import (
<<<<<<< HEAD:server/handlers/requestHandlers.go
	"encoding/json"
	"fmt"
=======
>>>>>>> 06a9f7b1e538b4a410fa7a16db5118d33eadc4b1:server/handlers/loginHandlers.go
	"log"
	"net/http"
	"server/middleware"
)

func (h *Handlers) SaveNewUser(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD:server/handlers/requestHandlers.go
	log.Println("INFO [handlers/requestHandlers.go] Request to Save New User")
	var student middleware.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		log.Printf("ERROR [handlers/requestHandlers.go] Couldn't get data: %s\n", err.Error())
		http.Error(w, "Couldn't get data", http.StatusInternalServerError)
		return
	}

	if err := h.db.CreateNewStudent(student); err != nil {
		log.Printf("ERROR [handlers/requestHandlers.go] Couldn't Save New User: %s\n", err.Error())
=======
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
>>>>>>> 06a9f7b1e538b4a410fa7a16db5118d33eadc4b1:server/handlers/loginHandlers.go
		http.Error(w, "Couldn't Save New User", http.StatusInternalServerError)
		return
	}
	log.Println(LOGGER_INFO_LOGIN + " Successfully Saved New User")
	// fmt.Fprint(w, "Thanks For signing up for CheckingN!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) LoginRequest(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD:server/handlers/requestHandlers.go
	log.Println("INFO [handlers/requestHandlers.go] Request to Log In")
	var login middleware.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		log.Printf("ERROR [handlers/requestHandlers.go] Couldn't get data: %s\n", err.Error())
		http.Error(w, "Couldn't get data", http.StatusInternalServerError)
		return
	}

	if found := h.db.FindStudent(login); !found {
		log.Println("INFO [handlers/requestHandlers.go] Failed Log In")
=======
	log.Println(LOGGER_INFO_LOGIN + " Request to Log In")
	r.ParseForm()
	login := &middleware.LoginRequest{
		Email:    r.PostForm.Get("Email"),
		Password: r.PostForm.Get("Password"),
	}

	if found := h.db.FindStudent(*login); !found {
		log.Println(LOGGER_INFO_LOGIN + " Failed Log In")
>>>>>>> 06a9f7b1e538b4a410fa7a16db5118d33eadc4b1:server/handlers/loginHandlers.go
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println(LOGGER_INFO_LOGIN + " Log In Successful")

<<<<<<< HEAD:server/handlers/requestHandlers.go
	expirationTime := time.Now().Add(AUTO_LOGOUT_TIME)

	claims := &Claims{
		Email: login.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.LoginRequestSuccess(w, r, tokenString)
}

func (h *Handlers) LoginRequestSuccess(w http.ResponseWriter, r *http.Request, tokenString string) {
	r.AddCookie(
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: time.Now().Add(AUTO_LOGOUT_TIME),
		})
	if h.tokenValid(w, r) {
		// w.Header().Set("Access-Control-Allow-Credentials", "true")
		fmt.Fprint(w, "<h1>Successful Login!<h1>")
	}
=======
	h.createTokenAndSetCookie(w, r, login.Email)
	h.Home(w, r)
>>>>>>> 06a9f7b1e538b4a410fa7a16db5118d33eadc4b1:server/handlers/loginHandlers.go
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
