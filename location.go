package telebot

import "fmt"

// Location represents a point on the map.
type Location struct {
	// Longitude is the longitude as defined by the sender.
	Longitude float64 `json:"longitude"`

	// Latitude is the latitude as defined by the sender.
	Latitude float64 `json:"latitude"`

	// HorizontalAccuracy is an optional field representing the radius of uncertainty for the location, measured in meters; 0-1500.
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// LivePeriod is an optional field representing the time relative to the message sending date, during which the location can be updated, in seconds.
	// For active live locations only.
	LivePeriod int `json:"live_period,omitempty"`

	// Heading is an optional field representing the direction in which the user is moving, in degrees; 1-360.
	// For active live locations only.
	Heading int `json:"heading,omitempty"`

	// ProximityAlertRadius is an optional field representing the maximum distance for proximity alerts about approaching another chat member, in meters.
	// For sent live locations only.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

func (c *Location) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *Location) Type() string        { return "Location" }
