package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Note   NotePayload `json:"note,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
}

type NotePayload struct {
	Autor   string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit the broker!",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		log.Printf("error al leer el json %v", err)
		return
	}

	switch requestPayload.Action {
	case "note":
		app.requestNote(w, requestPayload.Note)
	case "log":
		app.logItem(w, requestPayload.Log)
	default:
		app.errorJSON(w, errors.New("accion no soportada"))
	}
}

func (app Config) requestNote(w http.ResponseWriter, note NotePayload) {
	jsonData, _ := json.Marshal(note)

	noteServiceURL := "http://notes-service/note"
	request, err := http.NewRequest("POST", noteServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("estatus code es distinto no es accepted"))
		return
	}

	var payload JsonResponse
	payload.Error = false
	payload.Message = "Nota creada"

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app Config) logItem(w http.ResponseWriter, log LogPayload) {
	jsonData, _ := json.Marshal(log)

	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POS", logServiceURL, bytes.NewBuffer(jsonData))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, err)
		return
	}

	var payload JsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.writeJSON(w, http.StatusAccepted, payload)

}
