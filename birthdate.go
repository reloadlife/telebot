package telebot

type BirthDate struct {
	// Day is the day of the month.
	Day int `json:"day"`
	// Month is the month of the year.
	Month int `json:"month"`
	// Year is the year.
	Year *int `json:"year"`
}
