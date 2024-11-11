package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type JsonPayload struct {
	Identificador string `json:"identificador,omitempty"`
	Autor         string `json:"autor"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	Fecha         string `json:"fecha,omitempty"`
}

func (app Config) CreateNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload JsonPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// Crear la nota con los datos y la fecha actual
	nota := JsonPayload{
		Identificador: uuid.New().String(),
		Autor:         requestPayload.Autor,
		Title:         requestPayload.Title,
		Content:       requestPayload.Content,
		Fecha:         time.Now().Format("2006-01-02 15:04:05"),
	}

	// Asegurarse de que la carpeta "notes" existe
	_, err = os.Stat("notes")
	if os.IsNotExist(err) {
		err = os.Mkdir("notes", os.ModePerm)
		if err != nil {
			log.Fatal("No se pudo crear la carpeta")
			app.errorJSON(w, fmt.Errorf("error creando el directorio de notas: %v", err))
			return
		}
	}

	// Crear un archivo en "notes" usando el título y el identificador único como nombre
	filePath := fmt.Sprintf("notes/%s.txt", nota.Identificador)
	file, err := os.Create(filePath)
	if err != nil {
		app.errorJSON(w, fmt.Errorf("error creando el archivo de nota: %v", err))
		return
	}
	defer file.Close()

	// Escribir el contenido de la nota en el archivo
	err = app.writeFile(file, nota)
	if err != nil {
		app.errorJSON(w, fmt.Errorf("error escribiendo en el archivo de nota: %v", err))
		return
	}

	// Responder al cliente que la nota se creó correctamente
	response := jsonResponse{
		Error:   false,
		Message: "Nota creada correctamente",
		Data:    nota,
	}
	app.writeJSON(w, http.StatusCreated, response)
}

func (app Config) writeFile(w io.Writer, payload JsonPayload) error {
	_, err := w.Write([]byte("Título: " + payload.Title + "\n"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("Contenido: " + payload.Content + "\n"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("Fecha: " + payload.Fecha + "\n"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("Autor: " + payload.Autor))
	return err
}
