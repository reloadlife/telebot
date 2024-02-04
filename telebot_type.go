package telebot

import "time"

type Bot interface {
	// Debug sends a debug message to the bot.
	Debug(...any)

	// OnError Debug sends a debug message to the bot.
	OnError(error, Context)

	// Close this method to close the bot instance
	// before moving it from one local server to another.
	// You need to delete the webhook before calling this
	// method to ensure that the bot isn't launched again after the server restarts.
	// The method will return error 429 in the first 10 minutes after the bot is launched.
	Close() error

	// Logout Use this method to log out from the cloud
	// Bot API server before launching the bot locally.
	// You must log out the bot before running it locally,
	// otherwise there is no guarantee that the bot will receive updates.
	// After a successful call,
	// you can immediately log in on a local server,
	// but will not be able to log in back to the cloud Bot API server for 10 minutes.
	Logout() error

	// GetUpdates returns a list of updates (Long Polling).
	GetUpdates(offset, limit int, timeout time.Duration, allowed ...UpdateType) ([]Update, error)

	// GetMe returns basic information about the bot as a User object.
	GetMe() (*User, error)

	// SendMessage sends a text message.
	SendMessage(recipient Recipient, text string, options ...Option) (*Message, error)

	GetFile(fileID string) (*File, error)

	GetUserProfilePhotos(userID int64, offset, limit int) (*UserProfilePhotos, error)

	// Ban a user from a group, a supergroup or a channel.
	Ban(chatID Recipient, userID int64, untilDate *int64, revokeMessages *bool) error

	// Unban a user from a group, a supergroup or a channel.
	Unban(chatID Recipient, userID int64, onlyIfBanned *bool) error

	// Restrict restricts a user in a supergroup.
	Restrict(chatID Recipient, userID int64, permissions ChatPermissions, useIndependentChatPermissions *bool, untilDate *time.Duration) error

	// Promote promotes or demotes a user in a supergroup or a channel.
	Promote(chatID Recipient, userID int64, roles ...ChatMemberPermission) error

	// SetChatAdministratorCustomTitle sets a custom title for an administrator in a supergroup promoted by the bot.
	SetChatAdministratorCustomTitle(chatID Recipient, userID int64, customTitle string) error

	// BanChatSenderChat restricts a user in a supergroup.
	BanChatSenderChat(chatID Recipient, userID int64) error

	// UnbanChatSenderChat promotes or demotes a user in a supergroup or a channel.
	UnbanChatSenderChat(chatID Recipient, userID int64) error

	// SetChatPermissions sets default chat permissions for all members.
	SetChatPermissions(chatID Recipient, permissions ChatPermissions, useIndependentChatPermissions *bool) error

	// ExportChatInviteLink generates a new primary invite link for a chat.
	ExportChatInviteLink(chatID Recipient) (*string, error)

	// CreateChatInviteLink creates an additional invite link for a chat.
	CreateChatInviteLink(chatID Recipient, name string, expireDate int64, memberLimit int, createsJoinRequest bool) (*ChatInviteLink, error)

	// EditChatInviteLink edits a non-primary invite link created by the bot.
	EditChatInviteLink(chatID Recipient, inviteLink, name string, expireDate int64, memberLimit int, createsJoinRequest bool) (*ChatInviteLink, error)

	// RevokeChatInviteLink revokes an invite link created by the bot.
	RevokeChatInviteLink(chatID Recipient, inviteLink string) (*ChatInviteLink, error)

	// ApproveChatJoinRequest approves a chat join request for the specified user in the chat.
	// The bot must be an administrator in the chat and have the can_invite_users administrator right.
	// Returns an error on failure.
	ApproveChatJoinRequest(chatID Recipient, userID int64) error

	// DeclineChatJoinRequest declines a chat join request for the specified user in the chat.
	// The bot must be an administrator in the chat and have the can_invite_users administrator right.
	// Returns an error on failure.
	DeclineChatJoinRequest(chatID Recipient, userID int64) error

	// SetChatPhoto sets a new profile photo for the chat.
	// Photos can't be changed for private chats.
	// The bot must be an administrator in the chat and have the appropriate administrator rights.
	// Returns an error on failure.
	SetChatPhoto(chatID Recipient, photo File) error

	// DeleteChatPhoto deletes the chat photo.
	// Photos can't be changed for private chats.
	// The bot must be an administrator in the chat and have the appropriate administrator rights.
	// Returns an error on failure.
	DeleteChatPhoto(chatID Recipient) error

	// SetChatTitle changes the title of the chat.
	// Titles can't be changed for private chats.
	// The bot must be an administrator in the chat and have the appropriate administrator rights.
	// Returns an error on failure.
	SetChatTitle(chatID Recipient, title string) error

	// SetChatDescription changes the description of a group, supergroup, or channel.
	// The bot must be an administrator in the chat and have the appropriate administrator rights.
	// Returns an error on failure.
	SetChatDescription(chatID Recipient, description string) error

	PinChatMessage(chatID Recipient, messageID int, disableNotification bool) error
	UnpinChatMessage(chatID Recipient, messageID int) error
	UnpinAllChatMessages(chatID Recipient) error
	LeaveChat(chatID Recipient) error

	GetChat(chatID Recipient) (*Chat, error)
	GetChatAdministrators(chatID Recipient) ([]ChatMember, error)
	GetChatMemberCount(chatID Recipient) (*int, error)
	GetChatMember(chatID Recipient, userID int64) (*ChatMember, error)

	SetChatStickerSet(chatID Recipient, stickerSetName string) error
	DeleteChatStickerSet(chatID Recipient) error

	GetForumTopicIconStickers() ([]Sticker, error)
	CreateForumTopic(chatID Recipient, name string, options ...any) (*ForumTopic, error)
	EditForumTopic(chatID Recipient, topicID int64, options ...any) error
	CloseForumTopic(chatID Recipient, topicID int64) error
	ReopenForumTopic(chatID Recipient, topicID int64) error
	DeleteForumTopic(chatID Recipient, topicID int64) error
	UnpinAllForumTopicMessages(chatID Recipient, topicID int64) error
	EditGeneralForumTopic(chatID Recipient, name string) error
	CloseGeneralForumTopic(chatID Recipient) error
	ReopenGeneralForumTopic(chatID Recipient) error
	HideGeneralForumTopic(chatID Recipient) error
	UnhideGeneralForumTopic(chatID Recipient) error
	UnpinAllGeneralForumTopicMessages(chatID Recipient) error

	AnswerCallbackQuery(callback *Callback, opts ...any) error

	SetMyCommands(commands []BotCommand, opts ...any) error
	GetMyCommands(opts ...any) ([]BotCommand, error)
	DeleteMyCommands(opts ...any) error

	SetMyName(name string, opts ...any) error
	GetMyName(opts ...any) (*string, error)

	SetMyDescription(description string, opts ...any) error
	GetMyDescription(opts ...any) (*string, error)

	SetMyShortDescription(shortDescription string, opts ...any) error
	GetMyShortDescription(opts ...any) (*string, error)

	SetChatMenuButton(opts ...any) error
	GetChatMenuButton(opts ...any) (*MenuButton, error)

	SetMyDefaultAdministratorRights(opts ...any) error
	GetMyDefaultAdministratorRights(opts ...any) (*Rights, error)

	// Handle Register Routes
	Handle(endpoint any, h HandlerFunc, m ...MiddlewareFunc)

	Start()
	Stop()

	StartInWebhook()
	StopInWebhook()
}
