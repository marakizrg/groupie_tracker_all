package handlers_and_api

//my api key for the map = 417f9804b7744c8d9f9bbda14e172326
import (
	"encoding/json"
	"fmt"
	"groupie-tracker-all/structs"

	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

var artistsDetails []struct {
	structs.Artist
	Locations []string
	Dates     []string
	Relations map[string][]string
}

// Kanei fetch tous artists mono kai meta ta details tous kai ta apothikevei sta katalila structs  (init function runs automatically when this package is imported)
func init() {
	// // Fetch artists
	if err := fetchData("/artists", &artists); err != nil {
		log.Fatalf("Failed to fetch artists: %v", err)
	}

	// Pre-fetch detailed data for all artists CONCURRENTLY
	var wg sync.WaitGroup
	detailsChan := make(chan struct {
		structs.Artist
		Locations []string
		Dates     []string
		Relations map[string][]string
	}, len(artists))

	for _, artist := range artists {
		wg.Add(1)
		go func(a structs.Artist) {
			defer wg.Done()
			detailed, err := FetchDetailedArtistData(a.ID, a)
			if err == nil {
				detailsChan <- detailed
			}
		}(artist)
	}

	// Close channel after all goroutines finish
	go func() {
		wg.Wait()
		close(detailsChan)
	}()

	// after the goroutine finishes we take the results and put them in our struct
	for detail := range detailsChan {
		artistsDetails = append(artistsDetails, detail)
	}
	// Add a semaphore to limit concurrency (e.g., 10 goroutines max)
	sem := make(chan struct{}, 10)
	for _, artist := range artists {
		wg.Add(1)
		go func(a structs.Artist) {
			sem <- struct{}{}        // Acquire slot
			defer func() { <-sem }() // Release slot
		}(artist)
	}
}

// format the location
func FormatLocation(location string) string {
	// Replace underscores and hyphens with spaces
	location = strings.ReplaceAll(location, "_", " ")
	location = strings.ReplaceAll(location, "-", ", ")
	// Capitalize the first letter of each word manually
	words := strings.Fields(location)
	for i, word := range words {
		if len(word) > 0 && (strings.Contains(word, "uk") || strings.Contains(word, "usa")) {
			words[i] = strings.ToUpper(string(word))
		} else if len(word) > 0 {
			// Convert the first character to uppercase and the rest to lowercase
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	// Join the words back together
	return strings.Join(words, " ")
}

// clean up dates
func cleanDate(date string) string {
	return strings.TrimPrefix(date, "*")
}

// searches for an artist by their ID
func findArtistByID(id int) (struct {
	structs.Artist
	Locations []string
	Dates     []string
	Relations map[string][]string
}, bool) {
	for _, artist := range artistsDetails {
		if artist.ID == id {
			return artist, true
		}
	}
	// Return zero value and false if not found
	return struct {
		structs.Artist
		Locations []string
		Dates     []string
		Relations map[string][]string
	}{}, false
}

// fetches data from the API endpoind i will choose and stores it in the target struct
func fetchData(endpoint string, target interface{}) error {

	// Create an HTTP client with a timeout to avoid waiting indefinitely
	client := &http.Client{Timeout: 10 * time.Second}

	// Perform a GET request to the API
	resp, err := client.Get(baseURL + endpoint)
	if err != nil {
		return fmt.Errorf("failed to fetch %s: %w", endpoint, err)
	}
	defer resp.Body.Close()

	// Check if the server responded with a status other than 200 (OK)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decode the JSON response and store it in the target struct
	decodeError := json.NewDecoder(resp.Body).Decode(target)
	if decodeError != nil {
		return fmt.Errorf("failed to decode %s: %w", endpoint, decodeError)
	}

	return nil
}

// fetches all the details for the artist
func FetchDetailedArtistData(artistID int, artist structs.Artist) (detailedData struct {
	structs.Artist
	Locations []string
	Dates     []string
	Relations map[string][]string
}, err error) {
	// Check if data is already cached
	for _, d := range artistsDetails {
		if d.ID == artistID {
			return d, nil
		}
	}
	// Declare variables to store API responses
	var locations structs.Locations
	var dates structs.Dates
	var relations structs.Relations

	// Fetch the locations for the artist
	if err := fetchData(fmt.Sprintf("/locations/%d", artistID), &locations); err != nil {
		return struct {
			structs.Artist
			Locations []string
			Dates     []string
			Relations map[string][]string
		}{Artist: artist}, fmt.Errorf("failed to fetch locations: %w", err)
	}

	// Apply FormatLocation to each location
	for i, loc := range locations.Locations {
		locations.Locations[i] = FormatLocation(loc)
	}

	// Fetch the dates for the artist
	if err := fetchData(fmt.Sprintf("/dates/%d", artistID), &dates); err != nil {
		return struct {
			structs.Artist
			Locations []string
			Dates     []string
			Relations map[string][]string
		}{
			Artist:    artist,
			Locations: locations.Locations,
		}, fmt.Errorf("failed to fetch dates: %w", err)
	}

	// Clean up the dates (remove leading asterisks)
	for i, date := range dates.Dates {
		dates.Dates[i] = cleanDate(date)
	}

	// Fetch the relations for the artist
	if err := fetchData(fmt.Sprintf("/relation/%d", artistID), &relations); err != nil {
		return struct {
			structs.Artist
			Locations []string
			Dates     []string
			Relations map[string][]string
		}{
			Artist:    artist,
			Locations: locations.Locations,
			Dates:     dates.Dates,
		}, fmt.Errorf("failed to fetch relations: %w", err)
	}

	// Format the keys (locations) in relations
	formattedRelations := make(map[string][]string)
	for loc, dateList := range relations.DatesLocations {
		formattedRelations[FormatLocation(loc)] = dateList
	}

	// Return a complete struct with all formatted data
	return struct {
		structs.Artist
		Locations []string
		Dates     []string
		Relations map[string][]string
	}{
		Artist:    artist,
		Locations: locations.Locations,
		Dates:     dates.Dates,
		Relations: formattedRelations, // Use formatted keys
	}, nil
}
