package structs

// SearchResult represents a single search suggestion
type SearchResult struct {
	Text       string `json:"text"`       // The suggestion text (e.g., "London, UK")
	Type       string `json:"type"`       // The type of suggestion (e.g., "location")
	ID         int    `json:"id"`         // The artist ID (for redirects)
	ArtistName string `json:"artistName"` // The artist's name (e.g., "Queen")
	Image      string `json:"image"`      //artist's image
	Priority   int    `json:"-"`          // Priority for sorting (not sent to frontend)
}
