package handlers

import (
	"context"
	"log"
	"net/http"
)

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

	user := h.googleRespDecoder(*resp)
	h.db.CreateNewGoogleStudent(user)

	// h.createTokenAndSetCookie(w, r, user.Email)
	h.Home(w, r)
}
