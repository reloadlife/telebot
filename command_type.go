package telebot

type setMyCommands struct {
	Commands []BotCommand     `json:"commands"`
	Scope    *BotCommandScope `json:"scope,omitempty"`
	Language string           `json:"language_code,omitempty"`
}

type getMyCommands struct {
	Scope    *BotCommandScope `json:"scope,omitempty"`
	Language string           `json:"language_code,omitempty"`
}

type deleteMyCommands struct {
	Scope    *BotCommandScope `json:"scope,omitempty"`
	Language string           `json:"language_code,omitempty"`
}

type setMyNameRequest struct {
	Name     string `json:"name"`
	Language string `json:"language_code,omitempty"`
}

type setMyDescriptionRequest struct {
	Description      string `json:"description,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
	Language         string `json:"language_code,omitempty"`
}

type setChatMenuButton struct {
	ChatID any         `json:"chat_id,omitempty"`
	Button *MenuButton `json:"menu_button,omitempty"`
}

type setMyDefaultAdministratorRights struct {
	Rights      *Rights `json:"rights,omitempty"`
	ForChannels bool    `json:"for_channels,omitempty"`
}
