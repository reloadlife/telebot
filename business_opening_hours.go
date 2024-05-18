package telebot

type OpeningHours struct {
	// OpeningMinute The minute's sequence number in a week, starting on Monday, marking the start of the time interval during which the business is open; 0 - 7 * 24 * 60
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

type BusinessOpeningHours struct {
	TimeZoneName string         `json:"time_zone_name"`
	OpeningHours []OpeningHours `json:"opening_hours"`
}
