package main

import (
	"groupie-tracker-all/handlers_and_api"
	"log"
	"net/http"
)

func main() {
	//index handler
	http.HandleFunc("/", handlers_and_api.IndexHandler)

	//artist page handler
	http.HandleFunc("/artist/", handlers_and_api.ArtistHandler)

	//to prevent info leak
	http.HandleFunc("/static/", handlers_and_api.StaticFileHandler)

	//search handler
	http.HandleFunc("/search", handlers_and_api.SearchHandler)

	// Start the server
	log.Println("🎶 Rock on! The Groupie Tracker server is live at: http://localhost:8080 🤘🔥")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
