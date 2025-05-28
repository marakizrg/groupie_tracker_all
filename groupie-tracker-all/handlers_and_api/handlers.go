package handlers_and_api

import (
	"encoding/json"
	"groupie-tracker-all/structs"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

var artists []structs.Artist

// apiKey for maps
var apiKey = "417f9804b7744c8d9f9bbda14e172326"

// IndexHandler serves the index.html page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//lathos path
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}
	//edw tou lew na kanei parse to index.html template kai na to apothikefsei sto tmpl
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		handleServerError(w, err) // Serve 500 error if the template fails to load
		log.Printf("Error loading index.html template: %v", err)
		return
	}
	//filtering
	filteredArtists := filtering(w, r)

	// Fetch unique locations from all artists with cached data:
	uniqueLocations := make(map[string]bool)
	for _, detail := range artistsDetails {
		for _, loc := range detail.Locations {
			uniqueLocations[loc] = true
		}
	}

	// Convert map keys to slice
	var locations []string
	for loc := range uniqueLocations {
		locations = append(locations, loc)
	}

	// Pass filtered artists and available locations to template
	data := struct {
		Artists   []structs.Artist
		Locations []string
	}{
		Artists:   filteredArtists,
		Locations: locations,
	}
	// //edw ekteloume to template afou tou dwsoume ta dedomena poy theloyme
	if err := tmpl.Execute(w, data); err != nil {
		handleServerError(w, err)
		log.Printf("Error rendering index.html template: %v", err)
	}

}

// ArtistHandler serves the artist.html page
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the artist ID from the URL path
	idStr := r.URL.Path[len("/artist/"):]
	// Convert the ID to an integer
	artistID, err := strconv.Atoi(idStr)

	// Handle invalid artist ID (Bad Request)
	if err != nil {
		handleBadRequest(w, r)
		log.Printf("Invalid artist ID: %s", idStr)
		return
	}

	// Find the artist by ID in the artists list
	selectedArtist, found := findArtistByID(artistID)
	if !found {
		// http.NotFound(w, r) // Serve 404 error if artist is not found
		NotFoundHandler(w, r)
		return
	}

	//geolocalization part

	//location coordinates array
	var finalLocations []structs.LocationWithCoordinates

	//edw pairnoume ta locations kai ta kanoume coordinates
	for _, loc := range selectedArtist.Locations {
		lat, lng, err := geocodeLocation(loc, apiKey)
		if err != nil {
			log.Println("Error geocoding:", loc, err)
			continue
		}

		finalLocations = append(finalLocations, structs.LocationWithCoordinates{
			Name: loc,
			Lat:  lat,
			Lng:  lng,
		})

		time.Sleep(1 * time.Second) // Respect API rate limit (1 request/sec for free plan)
	}

	// Parses the artist.html template
	tmpl, err := template.ParseFiles("./templates/artist.html")
	if err != nil {
		handleServerError(w, err) // Serve 500 error if the template fails to load
		log.Printf("Error loading artist.html template: %v", err)
		return
	}

	pinsJSON, err := json.Marshal(finalLocations)
	if err != nil {
		log.Println("Error marshaling Pins:", err)
		handleServerError(w, err)
		return
	}

	data := struct {
		Artist struct {
			structs.Artist
			Locations []string
			Dates     []string
			Relations map[string][]string
		}
		Pins        []structs.LocationWithCoordinates
		PinsJSONStr template.JS // key point here
	}{
		Artist:      selectedArtist,
		Pins:        finalLocations,
		PinsJSONStr: template.JS(pinsJSON), // <- injects raw JSON safely
	}

	// The detailed artist data is passed to artist.html template
	if err := tmpl.Execute(w, data); err != nil {
		handleServerError(w, err) // Serve 500 error if the template fails to render
		log.Printf("Error rendering artist.html template: %v", err)
	}
}

func StaticFileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, "/static/")
	if filePath == "" || strings.HasSuffix(filePath, "/") {
		NotFoundHandler(w, r) // Prevent directory listing
		return
	}
	http.ServeFile(w, r, "./static/"+filePath)
}

// SearchHandler handles search requests and returns JSON suggestions
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q") // Get the search query
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}

	// Convert query to lowercase for case-insensitive search
	query = strings.ToLower(query)

	// Collect all possible search results
	var results []structs.SearchResult

	// Search through artists, members, locations, etc.
	for _, artist := range artists {
		// Search by artist name
		if strings.Contains(strings.ToLower(artist.Name), query) {
			results = append(results, structs.SearchResult{
				Text:       artist.Name,
				Type:       "artist",
				ID:         artist.ID,
				ArtistName: artist.Name,
				Image:      artist.Image,
				Priority:   1, // Highest priority for artists
			})
		}

		// Search by members
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				results = append(results, structs.SearchResult{
					Text:       member,
					Type:       "member",
					ID:         artist.ID,
					ArtistName: artist.Name,
					Priority:   2, // Second priority for members
				})
			}
		}

		// Search by locations
		for _, detail := range artistsDetails {
			if detail.ID == artist.ID {
				for _, loc := range detail.Locations {
					if strings.Contains(strings.ToLower(loc), query) {
						results = append(results, structs.SearchResult{
							Text:       loc,
							Type:       "location",
							ID:         artist.ID,
							ArtistName: artist.Name,
							Priority:   3, // Third priority for locations
						})
					}
				}
				break
			}
		}

		// Search by first album date
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			results = append(results, structs.SearchResult{
				Text:       artist.FirstAlbum,
				Type:       "first album",
				ID:         artist.ID,
				ArtistName: artist.Name,
				Priority:   4, // Fourth priority for first album dates
			})
		}

		// Search by creation date
		if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), query) {
			results = append(results, structs.SearchResult{
				Text:       strconv.Itoa(artist.CreationDate),
				Type:       "creation date",
				ID:         artist.ID,
				ArtistName: artist.Name,
				Priority:   5, // Fifth priority for creation dates
			})
		}
	}

	// Sort results by priority
	sort.Slice(results, func(i, j int) bool {
		// First, sort by priority
		if results[i].Priority != results[j].Priority {
			return results[i].Priority < results[j].Priority
		}

		// Check if the text starts with the query
		startsWithI := strings.HasPrefix(strings.ToLower(results[i].Text), query)
		startsWithJ := strings.HasPrefix(strings.ToLower(results[j].Text), query)

		// Prioritize those that start with the query
		if startsWithI != startsWithJ {
			return startsWithI
		}

		// Otherwise, keep the original order
		return false
	})

	// Return results as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
