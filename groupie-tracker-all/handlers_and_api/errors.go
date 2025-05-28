package handlers_and_api

import (
	"groupie-tracker-all/structs"
	"html/template"
	"log"
	"net/http"
)

// handleServerError serves a custom 500 error page and logs the error
func handleServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	tmpl, templateErr := template.ParseFiles("./templates/error.html")
	if templateErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Internal Server Error - Failed to load error.html template: %v", templateErr)
		return
	}

	data := structs.ErrorData{
		StatusCode: http.StatusInternalServerError,
		Message:    "It seems there was a problem on our end",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Internal Server Error - Failed to render error.html template: %v", err)
	}

	log.Printf("500 Internal Server Error: %v", err)
}

// NotFoundHandler serves a custom 404 page and logs the error
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	tmpl, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		log.Printf("Internal Server Error - Failed to load error.html template: %v", err)
		return
	}

	data := structs.ErrorData{
		StatusCode: http.StatusNotFound,
		Message:    "The page you are looking for does not exist.",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		log.Printf("Internal Server Error - Failed to render error.html template: %v", err)
	}

	log.Printf("404 Not Found: %s", r.URL.Path)
}

// handleBadRequest serves a custom 400 page and logs the error
func handleBadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)

	tmpl, templateErr := template.ParseFiles("./templates/error.html")
	if templateErr != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Internal Server Error - Failed to load error.html template: %v", templateErr)
		return
	}

	data := structs.ErrorData{
		StatusCode: http.StatusBadRequest,
		Message:    "Your request was not valid. Please check and try again.",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Internal Server Error - Failed to render error.html template: %v", err)
	}

	log.Printf("400 Bad Request: %s", r.URL.Path)
}
