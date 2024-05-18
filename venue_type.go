package telebot

type sendVenueRequest struct {
	ChatID    any     `json:"chat_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
	Address   string  `json:"address"`

	Foursquare      *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`

	BusinessID *BusinessID      `json:"business_connection_id,omitempty"`
	ThreadID   *MessageThreadID `json:"message_thread_id,omitempty"`

	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup      `json:"reply_markup,omitempty"`
}
