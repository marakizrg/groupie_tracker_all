package handlers_and_api

import (
	"groupie-tracker-all/structs"
	"net/http"
	"strconv"
	"time"
)

func filtering(w http.ResponseWriter, r *http.Request) []structs.Artist {
	// Get the values that came from the client
	selectedMembers := r.URL.Query()["members"]
	selectedLocations := r.URL.Query()["locations"]
	minCreationStr := r.URL.Query().Get("creation_min")
	maxCreationStr := r.URL.Query().Get("creation_max")
	minAlbumStr := r.URL.Query().Get("album_min")
	maxAlbumStr := r.URL.Query().Get("album_max")

	// set the default year to be the current
	currentYear := time.Now().Year() // Get current year
	// Convert to integers
	var minCreation, maxCreation, minAlbum, maxAlbum int
	var err1 error

	if minCreationStr != "" {
		minCreation, err1 = strconv.Atoi(minCreationStr)
		if err1 != nil {
			http.Error(w, "Invalid creation_min value", http.StatusBadRequest)
			return nil
		}
	} else {
		minCreation = 1900
	}

	// If maxCreationStr is empty, default to the current year
	if maxCreationStr != "" {
		maxCreation, err1 = strconv.Atoi(maxCreationStr)
		if err1 != nil {
			http.Error(w, "Invalid creation_max value", http.StatusBadRequest)
			return nil
		}
	} else {
		maxCreation = currentYear
	}

	if minAlbumStr != "" {
		minAlbum, err1 = strconv.Atoi(minAlbumStr)
		if err1 != nil {
			http.Error(w, "Invalid album_min value", http.StatusBadRequest)
			return nil
		}
	} else {
		minAlbum = 1900
	}

	// If maxAlbumStr is empty, default to the current year
	if maxAlbumStr != "" {
		maxAlbum, err1 = strconv.Atoi(maxAlbumStr)
		if err1 != nil {
			http.Error(w, "Invalid album_max value", http.StatusBadRequest)
			return nil
		}
	} else {
		maxAlbum = currentYear
	}

	filteredArtists := artists
	if minCreation > 0 && maxCreation > 0 {
		filteredArtists = FilterByCreationDate(filteredArtists, minCreation, maxCreation)
	}
	if minAlbum > 0 && maxAlbum > 0 {
		filteredArtists = FilterByFirstAlbum(filteredArtists, minAlbum, maxAlbum)
	}
	if len(selectedMembers) > 0 {
		filteredArtists = FilterByMembers(filteredArtists, selectedMembers)
	}
	if len(selectedLocations) > 0 {
		filteredArtists = FilterByLocations(filteredArtists, selectedLocations)
	}

	return filteredArtists
}
