package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"server/middleware"
	"server/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		ClientID:     "848508325356-2vdge0pmqndaibtj2ulfel50gf8tvj9v.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		//Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint: google.Endpoint,
	}
	randomState = "random"
	firstLogin  = true
	secret_key  = []byte("secret_key")
)

func (h *Handlers) ExampleJsonReponse(w http.ResponseWriter, r *http.Request) {
	var courses []*middleware.Course

	course := new(middleware.Course)
	course.Name = "CS160"
	course.Department = "computer science"
	courses = append(courses, course)

	course2 := new(middleware.Course)
	course2.Name = "PHIL134"
	course2.Department = "philosophy"
	courses = append(courses, course2)

	fmt.Fprint(w, utils.Jsonify(courses))
}

func (h *Handlers) SaveNewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO [handlers/requestHandlers.go] Request to Save New User")
	r.ParseForm()
	student := &middleware.Student{
		FirstName: r.PostForm.Get("first_name"),
		LastName:  r.PostForm.Get("last_name"),
		Email:     r.PostForm.Get("email"),
		Password:  r.PostForm.Get("password"),
	}

	if err := h.db.CreateNewStudent(*student); err != nil {
		log.Printf("ERROR [handlers/requestHandlers.go] Couldn't Save New User: %s\n", err.Error())
		http.Error(w, "Couldn't Save New User", http.StatusInternalServerError)
		return
	}
	log.Println("INFO [handlers/requestHandlers.go] Successfully Saved New User")
	// fmt.Fprint(w, "Thanks For signing up for CheckingN!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) LoginRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO [handlers/requestHandlers.go] Request to Log In")
	r.ParseForm()
	login := &middleware.LoginRequest{
		Email:    r.PostForm.Get("Email"),
		Password: r.PostForm.Get("Password"),
	}

	if found := h.db.FindStudent(*login); !found {
		log.Println("INFO [handlers/requestHandlers.go] Failed Log In")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("INFO [handlers/requestHandlers.go] Log In Successful")

	expirationTime := time.Now().Add(AUTO_LOGOUT_TIME)

	claims := &Claims{
		Email: login.Email,
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
	h.Home(w, r)
}

func (h *Handlers) Google(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html><body><a href="/api/google_login_request">Google Log In</a></body></html>`
	fmt.Fprintf(w, htmlIndex)
}

func (h *Handlers) GoogleLoginRequest(w http.ResponseWriter, r *http.Request) {
	if firstLogin {
		googleOauthConfig.RedirectURL = "http://" + h.hostIpBinding + "/api/callback"
		firstLogin = false
	}
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handlers) GoogleLoginCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		fmt.Println("State is invalid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	// if err != nil {
	// 	log.Printf("INFO [handlers/requestHandlers.go] Could not get token %s", err.Error())
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// 	return
	// }

	// resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	// defer resp.Body.Close()
	// if err != nil {
	// 	log.Printf("INFO [handlers/requestHandlers.go] Could not create get request %s", err.Error())
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// 	return
	// }
	// if resp.StatusCode != http.StatusOK {
	// 	w.WriteHeader(resp.StatusCode)
	// }
	// fmt.Println(resp)

	// h.createToken(w, r, token.AccessToken)

	// h.Home(w, r)
	token, err := googleOauthConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		log.Printf("INFO [handlers/requestHandlers.go] Could not get token %s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Printf("INFO [handlers/requestHandlers.go] Could not create get request %s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ERROR [handlers/requestHandlers.go] %s\n", err.Error())
		return
	}
	contents = contents

	h.Home(w, r)
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if h.tokenValid(w, r) {
		fmt.Fprint(w, "<h1>Successful Login!<h1>")
	} else {
		fmt.Println("hello")
	}
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
