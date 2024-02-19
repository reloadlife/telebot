package telebot

// VideoChatScheduled represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	// StartDate is the point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator.
	StartDate int `json:"start_date"`
}

// VideoChatStarted represents a service message about a video chat started in the chat.
// Currently holds no information.
type VideoChatStarted struct{}

// VideoChatEnded represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	// Duration is the video chat duration in seconds.
	Duration int `json:"duration"`
}

// VideoChatParticipantsInvited represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	// Users is the list of new members that were invited to the video chat.
	Users []User `json:"users"`
}
