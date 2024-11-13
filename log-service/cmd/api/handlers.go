package main

import (
	"logger/data"
	"net/http"
)

type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) writteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload jsonPayload
	_ = app.readJSON(w, r, &requestPayload)

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errJson(w, err)
		return
	}
	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}
	app.writeJson(w, http.StatusAccepted, resp)
}
