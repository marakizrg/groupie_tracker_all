package handlers_and_api

import (
	"groupie-tracker-all/structs"
	"strconv"
)

// creation date filter
func FilterByCreationDate(allArtists []structs.Artist, minDate int, maxDate int) []structs.Artist {
	var filtered []structs.Artist

	for _, artist := range allArtists {
		if artist.CreationDate >= minDate && artist.CreationDate <= maxDate {
			filtered = append(filtered, artist)
		}
	}

	return filtered
}

// filter for album dates
func FilterByFirstAlbum(allArtists []structs.Artist, minDate int, maxDate int) []structs.Artist {
	var filtered []structs.Artist

	for _, artist := range allArtists {
		albumYear, _ := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
		if albumYear >= minDate && albumYear <= maxDate {
			filtered = append(filtered, artist)
		}
	}

	return filtered
}

// filterByMembers filters artists based on selected number of members
func FilterByMembers(allArtists []structs.Artist, selected []string) []structs.Artist {
	var filtered []structs.Artist

	// Convert selected values to a map for fast lookup
	selectedMap := make(map[int]bool)
	for _, val := range selected {
		num, err := strconv.Atoi(val)
		if err == nil {
			selectedMap[num] = true
		}
	}

	// Filter artists whose member count matches the selection
	for _, artist := range allArtists {
		if selectedMap[len(artist.Members)] {
			filtered = append(filtered, artist)
		}
	}

	return filtered
}

// FilterByLocations filters artists whose concerts match the selected locations
func FilterByLocations(allArtists []structs.Artist, selectedLocations []string) []structs.Artist {
	var filtered []structs.Artist
	selectedMap := make(map[string]bool)
	for _, loc := range selectedLocations {
		selectedMap[loc] = true
	}

	for _, artist := range allArtists {
		// Get cached details instead of API call
		for _, detail := range artistsDetails {
			if detail.ID == artist.ID {
				for _, loc := range detail.Locations {
					if selectedMap[loc] {
						filtered = append(filtered, artist)
						break
					}
				}
				break
			}
		}
	}
	return filtered
}
