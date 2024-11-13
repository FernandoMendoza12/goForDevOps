package main

import "net/http"

type jsonPayload struct {
	Email string `json:"email"`
	Data  string `json:"name"`
}

func (app *Config) userAuth(w http.ResponseWriter, r *http.Request) {
	var requestPayload jsonPayload
	_ = app.readJSON(w, r, &requestPayload)

}
