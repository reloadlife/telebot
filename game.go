package telebot

// Game represents a game. Use BotFather to create and edit games; their short names will act as unique identifiers.
type Game struct {
	// Title is the title of the game.
	Title string `json:"title"`

	// Description is the description of the game.
	Description string `json:"description"`

	// Photo is an array of PhotoSize representing the photo that will be displayed in the game message in chats.
	Photo []PhotoSize `json:"photo"`

	// Text is an optional field providing a brief description of the game or high scores included in the game message.
	// It can be automatically edited to include current high scores for the game when the bot calls setGameScore,
	// or manually edited using editMessageText. It can have 0-4096 characters.
	Text *string `json:"text,omitempty"`

	// TextEntities is an optional field representing special entities that appear in text, such as usernames, URLs, bot commands, etc.
	TextEntities []Entity `json:"text_entities,omitempty"`

	// Animation is an optional field representing the animation that will be displayed in the game message in chats.
	// Upload via BotFather.
	Animation *Animation `json:"animation,omitempty"`
}

// CallbackGame represents a placeholder for game-related information.
// Use BotFather to set up your game.
type CallbackGame struct{}

type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}
