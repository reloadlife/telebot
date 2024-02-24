package telebot

import "fmt"

// Venue represents a venue.
type Venue struct {
	// Location is the Venue location. It can't be a live location.
	Location Location `json:"location"`

	// Title is the name of the venue.
	Title string `json:"title"`

	// Address is the address of the venue.
	Address string `json:"address"`

	// FoursquareID is an optional field representing the Foursquare identifier of the venue.
	FoursquareID string `json:"foursquare_id,omitempty"`

	// FoursquareType is an optional field representing the Foursquare type of the venue.
	// For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.
	FoursquareType string `json:"foursquare_type,omitempty"`

	// GooglePlaceID is an optional field representing the Google Places identifier of the venue.
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// GooglePlaceType is an optional field representing the Google Places type of the venue.
	// See supported types for more information.
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

func (v *Venue) ReflectType() string { return fmt.Sprintf("%T", v) }
func (v *Venue) Type() string        { return "Venue" }
