package telebot

import "time"

// Bot
// Interface for a Bot instance.
// Bot is a high-level interface for interacting with the Telegram API.
type Bot interface {

	// Handle registers a handler function for the given endpoint.
	//
	// endpoint is the endpoint path to handle.
	//
	// h is the handler function to register.
	//
	// m specifies optional middleware functions.
	//
	// This allows routing HTTP requests to handler functions.
	Handle(endpoint any, h HandlerFunc, m ...MiddlewareFunc)

	// Start starts the router listening for requests.
	//
	// This starts the HTTP server on the configured address
	// and port and begins handling requests.
	Start()

	// Stop stops the router and shuts down the HTTP server.
	//
	// This closes the listener and stops handling requests.
	Stop()

	// StartInWebhook starts the router in webhook mode.
	//
	// This configures the router for webhook usage rather
	// than a HTTP server. The router will not listen on
	// the configured port, but will call handler functions
	// when SetWebhook is called externally.
	StartInWebhook()

	// StopInWebhook stops the router in webhook mode.
	//
	// This disables webhook and stops calling any webhook
	// handler functions.
	StopInWebhook()

	// Debug sends a debug message to the bot.
	//
	// debugMessage can be anything that can be formatted into a string using fmt.Sprint.
	// Useful for debugging purposes.
	Debug(debugMessage ...any)

	// OnError sends an error message to the bot.
	//
	// err is the error that occurred.
	// ctx provides context about where the error occurred.
	OnError(err error, ctx Context)

	// Close closes the bot instance before moving it from one local server to another.
	//
	// Deletes the webhook before closing to ensure the bot isn't launched again
	// after the server restarts. Returns error 429 if called within 10 mins of launch.
	//
	// Returns an error on failure.
	Close() error

	// Logout logs out the bot from the cloud Bot API server before running locally.
	//
	// Logs out to ensure the bot receives updates when running locally.
	// Can't log back in to cloud for 10 mins after logging out.
	//
	// Returns an error on failure.
	Logout() error

	// GetUpdates retrieves updates for the bot via long polling.
	//
	// offset is the update offset to start from.
	// limit limits number of updates returned.
	// timeout is the time to wait for updates in seconds.
	// allowed optionally filters updates by type.
	//
	// Returns list of updates and error on failure.
	GetUpdates(offset, limit int, timeout time.Duration, allowed ...UpdateType) ([]Update, error)

	// GetMe returns basic info about the bot as a User object.
	//
	// Returns User info and error on failure.
	GetMe() (*User, error)
	//
	//// SendMessage sends a text message to the provided recipient.
	////
	//// recipient is the chat to send the message to.
	//// text is the text content of the message.
	//// options contains additional send options like ReplyMarkup.
	////
	//// Returns the sent Message and error on failure.
	//SendMessage(recipient Recipient, text string, options ...any) (*Message, error)
	//
	//// SendPhoto sends a photo to the provided recipient.
	////
	//// recipient is the chat to send the photo to.
	//// photo is the photo File to send.
	//// options contains additional send options like caption.
	////
	//// Returns the sent Message and error on failure.
	//SendPhoto(recipient Recipient, photo File, options ...any) (*Message, error)
	//
	//// SendAudio sends an audio track to the provided recipient.
	////
	//// recipient is the chat to send the audio to.
	//// audio is the audio File to send.
	//// options contains additional send options like caption.
	////
	//// Returns the sent Message and error on failure.
	//SendAudio(recipient Recipient, audio File, options ...any) (*Message, error)
	//
	//// SendDocument sends a document to the provided recipient.
	////
	//// recipient is the chat to send the document to.
	//// document is the document File to send.
	//// options contains additional send options like caption.
	////
	//// Returns the sent Message and error on failure.
	//SendDocument(recipient Recipient, document File, options ...any) (*Message, error)
	//
	//// SendVideo sends a video to the provided recipient.
	////
	//// recipient is the chat to send the video to.
	//// video is the video File to send.
	//// options contains additional send options like supports streaming.
	////
	//// Returns the sent Message and error on failure.
	//SendVideo(recipient Recipient, video File, options ...any) (*Message, error)
	//
	//// SendAnimation sends an animation to the provided recipient.
	////
	//// recipient is the chat to send the animation to.
	//// animation is the animation File to send.
	//// options contains additional send options like caption.
	////
	//// Returns the sent Message and error on failure.
	//SendAnimation(recipient Recipient, animation File, options ...any) (*Message, error)
	//
	//// SendVoice sends a voice recording to the provided recipient.
	////
	//// recipient is the chat to send the voice to.
	//// voice is the voice recording File to send.
	//// options contains additional send options like caption.
	////
	//// Returns the sent Message and error on failure.
	//SendVoice(recipient Recipient, voice File, options ...any) (*Message, error)
	//
	//// SendVideoNote sends a video note to the provided recipient.
	////
	//// recipient is the chat to send the video note to.
	//// videoNote is the video note File to send.
	//// options contains additional send options like duration.
	////
	//// Returns the sent Message and error on failure.
	//SendVideoNote(recipient Recipient, videoNote File, options ...any) (*Message, error)
	//
	//// SendMediaGroup sends a group of photos or videos as an album to the recipient.
	////
	//// recipient is the chat to send the media album to.
	//// media is the list of photo and video Files to send.
	//// options contains additional send options like disable_notification.
	////
	//// Returns the sent Messages and error on failure.
	//SendMediaGroup(recipient Recipient, media []File, options ...any) ([]Message, error)
	//
	//// SendLocation sends a location to the provided recipient.
	////
	//// recipient is the chat to send the location to.
	//// location is the latitude and longitude location to send.
	//// options contains additional send options like live_period.
	////
	//// Returns the sent Message and error on failure.
	//SendLocation(recipient Recipient, location Location, options ...any) (*Message, error)
	//
	//// SendVenue sends a venue to the provided recipient.
	////
	//// recipient is the chat to send the venue to.
	//// venue is the venue information to send.
	//// options contains additional send options like foursquare_id.
	////
	//// Returns the sent Message and error on failure.
	//SendVenue(recipient Recipient, venue Venue, options ...any) (*Message, error)
	//
	//// SendContact sends a contact's info to the provided recipient.
	////
	//// recipient is the chat to send the contact to.
	//// contact is the contact info to send.
	//// options contains additional send options like last_name.
	////
	//// Returns the sent Message and error on failure.
	//SendContact(recipient Recipient, contact Contact, options ...any) (*Message, error)
	//
	//// SendPoll sends a poll to the provided recipient.
	////
	//// recipient is the chat to send the poll to.
	//// question is the poll question.
	//// options contains additional send options like is_anonymous.
	////
	//// Returns the sent Message and error on failure.
	//SendPoll(recipient Recipient, question string, options ...any) (*Message, error)
	//
	//// SendDice sends a dice with random value to the provided recipient.
	////
	//// recipient is the chat to send the dice to.
	//// options contains additional send options like emoji.
	////
	//// Returns the sent Message and error on failure.
	//SendDice(recipient Recipient, options ...any) (*Message, error)
	//
	//// SendChatAction sends a chat action to the provided recipient.
	////
	//// recipient is the chat to send the action to.
	//// action is the chat action to send.
	////
	//// Returns an error on failure.
	//SendChatAction(recipient Recipient, action ChatAction) error
	//
	//// SetMessageReaction adds a reaction to a message.
	////
	//// recipient is the chat the message is in.
	//// messageID is the ID of the message to react to.
	//// options contains the reaction emoji to use.
	////
	//// Returns an error on failure.
	//SetMessageReaction(recipient Recipient, messageID int, options ...any) error
	//
	//// GetFile retrieves information about a file from its file_id.
	////
	//// fileID is the ID of the file to get info about.
	////
	//// Returns File object and error on failure.
	//GetFile(fileID string) (*File, error)
	//
	//// GetUserProfilePhotos retrieves a user's profile photos.
	////
	//// userID is the ID of the user to get profile photos for.
	//// offset optionally offsets the list of returned photos.
	//// limit limits the number of photos returned.
	////
	//// Returns UserProfilePhotos and error on failure.
	//GetUserProfilePhotos(userID int64, offset, limit int) (*UserProfilePhotos, error)
	//
	//// Ban bans a user from a chat.
	////
	//// chatID is the ID of the chat to ban the user from.
	//// userID is the ID of the user to ban.
	//// untilDate optionally bans the user until a specific date.
	//// revokeMessages optionally deletes all messages by the user.
	////
	//// Returns an error on failure.
	//Ban(chatID Recipient, userID int64, untilDate *int64, revokeMessages *bool) error
	//
	//// Unban unbans a user from a chat.
	////
	//// chatID is the ID of the chat to unban the user from.
	//// userID is the ID of the user to unban.
	//// onlyIfBanned optionally only unbans if the user was banned.
	////
	//// Returns an error on failure.
	//Unban(chatID Recipient, userID int64, onlyIfBanned *bool) error
	//
	//// Restrict restricts a user in a supergroup.
	////
	//// chatID is the ID of the supergroup.
	//// userID is the ID of the user to restrict.
	//// permissions sets the user's new permissions.
	//// useIndependentChatPermissions optionally customizes the permissions.
	//// untilDate optionally restricts the user until a certain date.
	////
	//// Returns an error on failure.
	//Restrict(chatID Recipient, userID int64, permissions ChatPermissions, useIndependentChatPermissions *bool, untilDate *time.Duration) error
	//
	//// Promote promotes or demotes a user in a supergroup or channel.
	////
	//// chatID is the ID of the supergroup or channel.
	//// userID is the ID of the user to promote/demote.
	//// roles are the new member permissions to set.
	////
	//// Returns an error on failure.
	//Promote(chatID Recipient, userID int64, roles ...ChatMemberPermission) error
	//
	//// SetChatAdministratorCustomTitle sets a custom title for an admin promoted by the bot.
	////
	//// chatID is the ID of the supergroup.
	//// userID is the ID of the administrator.
	//// customTitle is the new custom title.
	////
	//// Returns an error on failure.
	//SetChatAdministratorCustomTitle(chatID Recipient, userID int64, customTitle string) error
	//
	//// BanChatSenderChat bans a user from sending messages in a supergroup.
	////
	//// chatID is the ID of the supergroup.
	//// userID is the ID of the user to ban.
	////
	//// Returns an error on failure.
	//BanChatSenderChat(chatID Recipient, userID int64) error
	//
	//// UnbanChatSenderChat unbans a user from sending messages in a supergroup/channel.
	////
	//// chatID is the ID of the supergroup/channel.
	//// userID is the ID of the user to unban.
	////
	//// Returns an error on failure.
	//UnbanChatSenderChat(chatID Recipient, userID int64) error
	//
	//// SetChatPermissions sets default permissions for members in a chat.
	////
	//// chatID is the ID of the chat.
	//// permissions are the new default permissions.
	//// useIndependentChatPermissions optionally customizes the permissions.
	////
	//// Returns an error on failure.
	//SetChatPermissions(chatID Recipient, permissions ChatPermissions, useIndependentChatPermissions *bool) error
	//
	//// ExportChatInviteLink generates a new primary invite link for a chat.
	////
	//// chatID is the ID of the chat.
	////
	//// Returns the new invite link and error on failure.
	//ExportChatInviteLink(chatID Recipient) (*string, error)
	//
	//// CreateChatInviteLink creates a new additional invite link for a chat.
	////
	//// chatID is the ID of the chat.
	//// name is the name of the invite link.
	//// expireDate optionally sets invite link expiration.
	//// memberLimit limits the maximum number of users that can join.
	//// createsJoinRequest enables users to request to join.
	////
	//// Returns the new ChatInviteLink and error on failure.
	//CreateChatInviteLink(chatID Recipient, name string, expireDate int64, memberLimit int, createsJoinRequest bool) (*ChatInviteLink, error)
	//
	//// EditChatInviteLink edits a non-primary invite link created by the bot.
	////
	//// chatID is the ID of the chat.
	//// inviteLink is the invite link to edit.
	//// name is the new name of the invite link.
	//// expireDate optionally sets new expiration date.
	//// memberLimit limits the maximum number of users that can join.
	//// createsJoinRequest enables users to request to join.
	////
	//// Returns the edited ChatInviteLink and error on failure.
	//EditChatInviteLink(chatID Recipient, inviteLink, name string, expireDate int64, memberLimit int, createsJoinRequest bool) (*ChatInviteLink, error)
	//
	//// RevokeChatInviteLink revokes an invite link created by the bot.
	////
	//// chatID is the ID of the chat.
	//// inviteLink is the invite link to revoke.
	////
	//// Returns the revoked ChatInviteLink and error on failure.
	//RevokeChatInviteLink(chatID Recipient, inviteLink string) (*ChatInviteLink, error)
	//
	//// ApproveChatJoinRequest approves a chat join request.
	////
	//// chatID is the ID of the chat.
	//// userID is the ID of the user requesting to join.
	////
	//// Returns an error on failure.
	//ApproveChatJoinRequest(chatID Recipient, userID int64) error
	//
	//// DeclineChatJoinRequest declines a chat join request.
	////
	//// chatID is the ID of the chat.
	//// userID is the ID of the user requesting to join.
	////
	//// Returns an error on failure.
	//DeclineChatJoinRequest(chatID Recipient, userID int64) error
	//
	//// SetChatPhoto sets a new chat photo.
	////
	//// chatID is the ID of the chat.
	//// photo is the new File to set as the photo.
	////
	//// Returns an error on failure.
	//SetChatPhoto(chatID Recipient, photo File) error
	//
	//// DeleteChatPhoto deletes the chat's current photo.
	////
	//// chatID is the ID of the chat.
	////
	//// Returns an error on failure.
	//DeleteChatPhoto(chatID Recipient) error
	//
	//// SetChatTitle changes the title of a chat.
	////
	//// chatID is the ID of the chat.
	//// title is the new chat title.
	////
	//// Returns an error on failure.
	//SetChatTitle(chatID Recipient, title string) error
	//
	//// SetChatDescription changes the description of a chat.
	////
	//// chatID is the ID of the chat.
	//// description is the new chat description.
	////
	//// Returns an error on failure.
	//SetChatDescription(chatID Recipient, description string) error
	//
	//// PinChatMessage pins a message in a chat.
	////
	//// chatID is the ID of the chat.
	//// messageID is the ID of the message to pin.
	//// disableNotification optionally mutes notifications for the pin.
	////
	//// Returns an error on failure.
	//PinChatMessage(chatID Recipient, messageID int, disableNotification bool) error
	//
	//// UnpinChatMessage unpins a pinned message in a chat.
	////
	//// chatID is the ID of the chat.
	//// messageID is the ID of the pinned message to unpin.
	////
	//// Returns an error on failure.
	//UnpinChatMessage(chatID Recipient, messageID int) error
	//
	//// UnpinAllChatMessages unpins all pinned messages in a chat.
	////
	//// chatID is the ID of the chat.
	////
	//// Returns an error on failure.
	//UnpinAllChatMessages(chatID Recipient) error
	//
	//// LeaveChat makes the bot leave a group, supergroup, or channel.
	////
	//// chatID is the ID of the chat to leave.
	////
	//// Returns an error on failure.
	//LeaveChat(chatID Recipient) error
	//
	//// GetChat gets information about a chat.
	////
	//// chatID is the ID of the chat to get info about.
	////
	//// Returns the Chat object and error on failure.
	//GetChat(chatID Recipient) (*Chat, error)
	//
	//// GetChatAdministrators gets a list of administrators in a chat.
	////
	//// chatID is the ID of the chat to get admins for.
	////
	//// Returns list of ChatMembers and error on failure.
	//GetChatAdministrators(chatID Recipient) ([]ChatMember, error)
	//
	//// GetChatMemberCount gets the number of members in a chat.
	////
	//// chatID is the ID of the chat.
	////
	//// Returns number of members and error on failure.
	//GetChatMemberCount(chatID Recipient) (*int, error)
	//
	//// GetChatMember gets information about a chat member.
	////
	//// chatID is the ID of the chat.
	//// userID is the ID of the user.
	////
	//// Returns ChatMember and error on failure.
	//GetChatMember(chatID Recipient, userID int64) (*ChatMember, error)
	//
	//// SetChatStickerSet sets the sticker set for a chat.
	////
	//// chatID is the ID of the chat.
	//// stickerSetName is the name of the sticker set.
	////
	//// Returns an error on failure.
	//SetChatStickerSet(chatID Recipient, stickerSetName string) error
	//
	//// DeleteChatStickerSet deletes the custom sticker set for a chat.
	////
	//// chatID is the ID of the chat.
	////
	//// Returns an error on failure.
	//DeleteChatStickerSet(chatID Recipient) error
	//
	//// GetForumTopicIconStickers gets the default stickers for forum topics.
	////
	//// Returns the list of Stickers and error on failure.
	//GetForumTopicIconStickers() ([]Sticker, error)
	//
	//// CreateForumTopic creates a new forum topic in the given chat.
	////
	//// chatID is the ID of the chat to create the topic in.
	////
	//// name is the name for the new topic.
	////
	//// options can contain additional settings for the topic like pinned, etc.
	////
	//// Returns the new ForumTopic and any error that occurred.
	//CreateForumTopic(chatID Recipient, name string, options ...any) (*ForumTopic, error)
	//
	//// EditForumTopic edits an existing forum topic in a chat.
	////
	//// chatID is the ID of the chat containing the topic.
	////
	//// topicID is the ID of the topic to edit.
	////
	//// options can contain the new topic name or other properties to edit.
	////
	//// Returns any error that occurred.
	//EditForumTopic(chatID Recipient, topicID int64, options ...any) error
	//
	//// CloseForumTopic closes an existing forum topic in a chat.
	////
	//// chatID is the ID of the chat containing the topic.
	////
	//// topicID is the ID of the topic to close.
	////
	//// Closed topics cannot be replied to but the existing messages remain.
	////
	//// Returns any error that occurred.
	//CloseForumTopic(chatID Recipient, topicID int64) error
	//
	//// ReopenForumTopic reopens a previously closed forum topic in a chat.
	////
	//// chatID is the ID of the chat containing the topic.
	////
	//// topicID is the ID of the closed topic to reopen.
	////
	//// Reopened topics can be replied to again.
	////
	//// Returns any error that occurred.
	//ReopenForumTopic(chatID Recipient, topicID int64) error
	//
	//// DeleteForumTopic deletes an existing forum topic from a chat.
	////
	//// chatID is the ID of the chat containing the topic.
	////
	//// topicID is the ID of the topic to delete.
	////
	//// Deleting a topic also deletes all messages posted in the topic.
	////
	//// Returns any error that occurred.
	//DeleteForumTopic(chatID Recipient, topicID int64) error
	//
	//// UnpinAllForumTopicMessages unpins all pinned messages in a forum topic.
	////
	//// chatID is the ID of the chat containing the topic.
	////
	//// topicID is the ID of the topic to unpin messages from.
	////
	//// Returns any error that occurred.
	//UnpinAllForumTopicMessages(chatID Recipient, topicID int64) error
	//
	//// EditGeneralForumTopic edits the general forum topic in a chat.
	////
	//// chatID is the ID of the chat containing the general forum topic.
	////
	//// name is the new name for the general forum topic.
	////
	//// Returns any error that occurred.
	//EditGeneralForumTopic(chatID Recipient, name string) error
	//
	//// CloseGeneralForumTopic closes the general forum topic in a chat.
	////
	//// chatID is the ID of the chat containing the general forum.
	////
	//// Closed general topics cannot be replied to but existing messages remain.
	////
	//// Returns any error that occurred.
	//CloseGeneralForumTopic(chatID Recipient) error
	//
	//// ReopenGeneralForumTopic reopens a closed general forum topic in a chat.
	////
	//// chatID is the ID of the chat containing the general forum topic.
	////
	//// Reopened general topics can be replied to again.
	////
	//// Returns any error that occurred.
	//ReopenGeneralForumTopic(chatID Recipient) error
	//
	//// HideGeneralForumTopic hides the general forum topic from a chat.
	////
	//// chatID is the ID of the chat containing the general forum topic.
	////
	//// Hidden general topics are not displayed in the chat but still exist.
	////
	//// Returns any error that occurred.
	//HideGeneralForumTopic(chatID Recipient) error
	//
	//// UnhideGeneralForumTopic unhides a previously hidden general forum topic.
	////
	//// chatID is the ID of the chat containing the hidden general topic.
	////
	//// Unhidden general topics are displayed in the chat again.
	////
	//// Returns any error that occurred.
	//UnhideGeneralForumTopic(chatID Recipient) error
	//
	//// UnpinAllGeneralForumTopicMessages unpins all pinned messages in the general forum topic.
	////
	//// chatID is the ID of the chat containing the general forum topic.
	////
	//// Returns any error that occurred.
	//UnpinAllGeneralForumTopicMessages(chatID Recipient) error
	//
	//// AnswerCallbackQuery answers an incoming callback query.
	////
	//// callback is the callback query to respond to.
	////
	//// opts can contain additional API options.
	////
	//// Returns any error that occurred.
	//AnswerCallbackQuery(callback *Callback, opts ...any) error
	//
	//// Commands gets the list of commands registered by the bot.
	////
	//// opts can contain additional API options.
	////
	//// Returns the commands and any error that occurred.
	//Commands(opts ...any) ([]BotCommand, error)
	//
	//// SetCommands changes the list of commands registered by the bot.
	////
	//// commands is the new list of commands.
	////
	//// opts can contain additional API options.
	////
	//// Returns any error that occurred.
	//SetCommands(commands []BotCommand, opts ...any) error
	//
	//// DeleteCommands removes all commands registered by the bot.
	////
	//// opts can contain additional API options.
	////
	//// Returns any error that occurred.
	//DeleteCommands(opts ...any) error
	//
	//// SetName changes the name of the bot.
	////
	//// name is the new name for the bot.
	////
	//// opts can contain additional API options.
	////
	//// Returns any error that occurred.
	//SetName(name string, opts ...any) error
	//
	//// GetName gets the current name of the bot.
	////
	//// opts can contain additional API options.
	////
	//// Returns the name and any error that occurred.
	//GetName(opts ...any) (*string, error)
	//
	//// SetDescription changes the description of the bot.
	////
	//// description is the new description.
	////
	//// opts can contain additional API options.
	////
	//// Returns any error that occurred.
	//SetDescription(description string, opts ...any) error
	//
	//// GetDescription gets the current description of the bot.
	////
	//// opts can contain additional API options.
	////
	//// Returns the description and any error that occurred.
	//GetDescription(opts ...any) (*string, error)
	//
	//// SetShortDescription changes the short description of the bot.
	////
	//// shortDescription is the new short description.
	////
	//// opts can contain additional API options.
	////
	//// Returns any error that occurred.
	//SetShortDescription(shortDescription string, opts ...any) error
	//
	//// GetShortDescription gets the current short description of the bot.
	////
	//// opts can contain additional API options.
	////
	//// Returns the short description and any error that occurred.
	//GetShortDescription(opts ...any) (*string, error)
	//
	//// SetChatMenuButton sets the menu button for the bot in chat screens.
	////
	//// opts can contain the new MenuButton to set.
	////
	//// Returns any error that occurred.
	//SetChatMenuButton(opts ...any) error
	//
	//// GetChatMenuButton gets the current menu button for the bot.
	////
	//// opts can contain additional API options.
	////
	//// Returns the MenuButton and any error that occurred.
	//GetChatMenuButton(opts ...any) (*MenuButton, error)
	//
	//// SetDefaultAdministratorRights changes default admin rights for bot admins.
	////
	//// opts can contain the new Rights.
	////
	//// Returns any error that occurred.
	//SetDefaultAdministratorRights(opts ...any) error
	//
	//// GetDefaultAdministratorRights gets the current default admin rights.
	////
	//// opts can contain additional API options.
	////
	//// Returns the Rights and any error that occurred.
	//GetDefaultAdministratorRights(opts ...any) (*Rights, error)
	//
	//// EditMessageText edits the text of a previously sent message.
	////
	//// recipient is the chat to edit the message in.
	////
	//// messageID is the ID of the message to edit.
	////
	//// text is the new text for the message.
	////
	//// options can specify additional editing options.
	////
	//// Returns the edited Message and any error.
	//EditMessageText(recipient Recipient, messageID int, text string, options ...any) (*Message, error)
	//
	//// EditMessageTextInline edits text of an inline message.
	////
	//// inlineMessageID is the ID of the inline message to edit.
	////
	//// text is the new text for the message.
	////
	//// options can specify additional editing options.
	////
	//// Returns any error that occurred.
	//EditMessageTextInline(inlineMessageID string, text string, options ...any) error
	//
	//// EditMessageCaption edits the caption of a previously sent message.
	////
	//// recipient is the chat to edit the message in.
	////
	//// messageID is the ID of the message to edit.
	////
	//// caption is the new caption text.
	////
	//// options can specify additional editing options.
	////
	//// Returns the edited Message and any error.
	//EditMessageCaption(recipient Recipient, messageID int, caption string, options ...any) (*Message, error)
	//
	//// EditMessageCaptionInline edits caption of an inline message.
	////
	//// inlineMessageID is the ID of the inline message to edit.
	////
	//// caption is the new caption text.
	////
	//// options can specify additional editing options.
	////
	//// Returns any error that occurred.
	//EditMessageCaptionInline(inlineMessageID string, caption string, options ...any) error
	//
	//// EditMessageMedia edits the media content of a previously sent message.
	////
	//// recipient is the chat to edit the message in.
	////
	//// messageID is the ID of the message to edit.
	////
	//// media is the new media content.
	////
	//// options can specify additional editing options.
	////
	//// Returns the edited Message and any error.
	//EditMessageMedia(recipient Recipient, messageID int, media InputMedia, options ...any) (*Message, error)
	//
	//// EditMessageMediaInline edits media content of an inline message.
	////
	//// inlineMessageID is the ID of the inline message to edit.
	////
	//// media is the new media content.
	////
	//// options can specify additional editing options.
	////
	//// Returns any error that occurred.
	//EditMessageMediaInline(inlineMessageID string, media InputMedia, options ...any) error
	//
	//// EditMessageLiveLocation edits live location in a previously sent message.
	////
	//// recipient is the chat to edit the message in.
	////
	//// messageID is the ID of the message to edit.
	////
	//// latitude and longitude are the new location.
	////
	//// options can specify additional editing options.
	////
	//// Returns the edited Message and any error.
	//EditMessageLiveLocation(recipient Recipient, messageID int, latitude float64, longitude float64, options ...any) (*Message, error)
	//
	//// EditMessageLiveLocationInline edits live location in an inline message.
	////
	//// inlineMessageID is the ID of the inline message to edit.
	////
	//// latitude and longitude are the new location.
	////
	//// options can specify additional editing options.
	////
	//// Returns any error that occurred.
	//EditMessageLiveLocationInline(inlineMessageID string, latitude float64, longitude float64, options ...any) error
	//
	//// StopMessageLiveLocation stops updating live location in a message.
	////
	//// recipient is the chat to edit the message in.
	////
	//// messageID is the ID of the message to edit.
	////
	//// options can specify additional editing options.
	////
	//// Returns the edited Message and any error.
	//StopMessageLiveLocation(recipient Recipient, messageID int, options ...any) (*Message, error)
	//
	//// StopMessageLiveLocationInline stops live location in an inline message.
	////
	//// inlineMessageID is the ID of the inline message to edit.
	////
	//// options can specify additional editing options.
	////
	//// Returns any error that occurred.
	//StopMessageLiveLocationInline(inlineMessageID string, options ...any) error
	//
	//// EditMessageReplyMarkup edits the inline keyboard markup of a message.
	////
	//// recipient is the chat to edit the message in.
	////
	//// messageID is the ID of the message to edit.
	////
	//// markup is the new inline keyboard markup.
	////
	//// Returns the edited Message and any error.
	//EditMessageReplyMarkup(recipient Recipient, messageID int, markup *InlineKeyboardMarkup) (*Message, error)
	//
	//// EditMessageReplyMarkupInline edits reply markup of an inline message.
	////
	//// inlineMessageID is the ID of the inline message to edit.
	////
	//// markup is the new inline keyboard markup.
	////
	//// Returns any error that occurred.
	//EditMessageReplyMarkupInline(inlineMessageID string, markup *InlineKeyboardMarkup) error
	//
	//// StopPoll stops an active poll and updates the message.
	////
	//// recipient is the chat containing the poll.
	////
	//// messageID is the ID of the poll message.
	////
	//// options can specify additional parameters.
	////
	//// Returns the stopped Poll and any error.
	//StopPoll(recipient Recipient, messageID int, options ...any) (*Poll, error)
	//
	//// DeleteMessage deletes a previously sent message.
	////
	//// recipient is the chat containing the message.
	////
	//// messageIDs is a list of message IDs to delete.
	////
	//// Returns any error that occurred.
	//DeleteMessage(recipient Recipient, messageIds ...int) error
	//
	//// SendSticker sends a sticker to a chat.
	//// recipient is the chat to send the sticker to.
	//// sticker is the sticker file to send.
	//// options can specify additional send options.
	//// Returns the sent Message and any error.
	//SendSticker(recipient Recipient, sticker File, options ...any) (*Message, error)
	//
	//// GetStickerSet gets info about a sticker set by name.
	////
	//// name is the short name of the sticker set.
	////
	//// Returns the StickerSet and any error.
	//GetStickerSet(name string) (*StickerSet, error)
	//
	//// GetCustomEmojiStickers gets stickers for the given custom emoji.
	////
	//// CustomEmojiIds is a list of custom emoji identifiers.
	////
	//// Returns the list of matching Stickers.
	//GetCustomEmojiStickers(CustomEmojiIds ...CustomEmoji) ([]Sticker, error)
	//
	//// UploadStickerFile uploads a sticker image to be used in a set.
	////
	//// Owner is the user ID of the sticker set owner.
	////
	//// sticker is the image file to upload.
	////
	//// Format is the image format like png, webp, etc.
	////
	//// Returns the uploaded File and any error.
	//UploadStickerFile(Owner int64, sticker File, Format string) (*File, error)
	//
	//// CreateNewStickerSet creates a new sticker set owned by a user.
	////
	//// userID is the owner user ID.
	////
	//// name is the short name for the set.
	////
	//// title is the sticker set title.
	////
	//// sticker is the list of stickers to add.
	////
	//// format is the upload sticker format.
	////
	//// options can specify additional parameters.
	////
	//// Returns any error that occurred.
	//CreateNewStickerSet(userID int64, name, title string, sticker []InputSticker, format string, options ...any) error
	//
	//// AddStickerToSet adds a new sticker to an existing set.
	////
	//// userID is the owner user ID.
	////
	//// name is the sticker set name.
	////
	//// sticker is the new sticker to add.
	////
	//// format is the upload sticker format.
	////
	//// options can specify additional parameters.
	////
	//// Returns any error that occurred.
	//AddStickerToSet(userID int64, name string, sticker InputSticker, format string, options ...any) error
	//
	//// SetStickerPositionInSet moves a sticker to a new position.
	////
	//// sticker is the file ID of the sticker.
	////
	//// position is the new 0-based position.
	////
	//// Returns any error that occurred.
	//SetStickerPositionInSet(sticker string, position int) error
	//
	//// DeleteStickerFromSet removes a sticker from a set.
	////
	//// sticker is the file ID of the sticker.
	////
	//// Returns any error that occurred.
	//DeleteStickerFromSet(sticker string)
	//
	//// SetStickerEmojiList updates the emoji list for a sticker.
	////
	//// sticker is the file ID of the sticker.
	////
	//// EmojiList is the new list of emoji.
	////
	//// Returns any error that occurred.
	//SetStickerEmojiList(sticker string, EmojiList []Emoji) error
	//
	//// SetStickerKeywords updates the search keywords for a sticker.
	////
	//// sticker is the file ID of the sticker.
	////
	//// Keywords is the new list of keywords.
	////
	//// Returns any error that occurred.
	//SetStickerKeywords(sticker string, Keywords []string) error
	//
	//// SetStickerMaskPosition updates the mask position for a mask sticker.
	////
	//// sticker is the file ID of the sticker.
	////
	//// maskPosition is the new mask position.
	////
	//// Returns any error that occurred.
	//SetStickerMaskPosition(sticker string, maskPosition MaskPosition) error
	//
	//// SetStickerSetTitle updates the title of a sticker.
	////
	//// sticker is the file ID of the sticker.
	////
	//// title is the new title.
	////
	//// Returns any error that occurred.
	//SetStickerSetTitle(sticker string, title string) error
	//
	//// SetStickerSetThumbnail updates the thumbnail for a set.
	////
	//// name is the sticker set name.
	////
	//// userId is the owner user ID.
	////
	//// thumbnail is the new thumbnail file.
	////
	//// Returns any error that occurred.
	//SetStickerSetThumbnail(name string, userId int64, thumbnail File) error
	//
	//// SetCustomEmojiStickerSetThumbnail updates emoji set thumbnail.
	////
	//// name is the sticker set name.
	////
	//// CustomEmojiID is the new custom emoji ID.
	////
	//// Returns any error that occurred.
	//SetCustomEmojiStickerSetThumbnail(name string, CustomEmojiID *string) error
	//
	//// DeleteStickerSet removes a sticker set.
	////
	//// name is the sticker set name.
	////
	//// Returns any error that occurred.
	//DeleteStickerSet(name string) error
	//
	//// AnswerInlineQuery sends results for an inline query.
	////
	//// queryID is the ID of the inline query.
	////
	//// results are the result rows to show.
	////
	//// options can specify additional parameters.
	////
	//// Returns any error that occurred.
	//AnswerInlineQuery(queryID string, results QueryResults, options ...any) error
	//
	//// AnswerWebAppQuery sends result for a Web App query.
	////
	//// webAppQueryID is the ID of the query.
	////
	//// result is the result to show.
	////
	//// Returns the sent message info and any error.
	//AnswerWebAppQuery(webAppQueryID string, result QueryResult) (*SentWebAppMessage, error)
	//
	//// SendInvoice sends an invoice requesting payment.
	////
	//// recipient is the chat to send the invoice to.
	////
	//// invoice contains invoice details.
	////
	//// options can specify additional parameters.
	////
	//// Returns the sent Message and any error.
	//SendInvoice(recipient Recipient, invoice Invoice, options ...any) (*Message, error)
	//
	//// CreateInvoiceLink generates a link for an invoice.
	////
	//// invoice contains the invoice details.
	////
	//// options can specify additional parameters.
	////
	//// Returns the generated URL as string and any error.
	//CreateInvoiceLink(invoice Invoice, options ...any) (*string, error)
	//
	//// AnswerShippingQuery responds to shipping query with options.
	////
	//// queryID is the ID of the shipping query.
	////
	//// ok indicates if shipping to the chosen address is possible.
	////
	//// errorMessage optionally contains an error message.
	////
	//// shippingOptions lists available shipping options.
	////
	//// Returns any error that occurred.
	//AnswerShippingQuery(queryID string, ok bool, errorMessage *string, shippingOptions ...any) error
	//
	//// AnswerPreCheckoutQuery responds to pre-checkout query.
	////
	//// queryID is the ID of the pre-checkout query.
	////
	//// ok indicates if checkout can continue.
	////
	//// errorMessage optionally contains error message.
	////
	//// Returns any error that occurred.
	//AnswerPreCheckoutQuery(queryID string, ok bool, errorMessage *string) error
	//
	//// SetPassportDataErrors informs a user of errors with Telegram Passport data.
	////
	//// userID is the ID of the user.
	////
	//// errors lists the errors that occurred.
	////
	//// Returns any error that occurred.
	//SetPassportDataErrors(userID int64, errors []PassportElementError) error
	//
	//// SendGame sends a game for the user to play.
	////
	//// recipient is the chat to send the game to.
	////
	//// gameShortName is the short name of the game.
	////
	//// options can specify additional parameters.
	////
	//// Returns the sent Message and any error.
	//SendGame(recipient Recipient, gameShortName string, options ...any) (*Message, error)
	//
	//// SetGameScore sets a new high score for a game.
	////
	//// recipient is the chat containing the message with the game.
	////
	//// userID is the ID of the user who achieved the high score.
	////
	//// score is the new high score value.
	////
	//// options can specify additional parameters.
	////
	//// Returns the updated Message and any error.
	//SetGameScore(recipient Recipient, userID int64, score int, options ...any) (*Message, error)
	//
	//// SetGameScoreInline sets high score for an inline game message.
	////
	//// inlineMessageID is the ID of the inline message.
	////
	//// userID is the ID of the user who achieved the high score.
	////
	//// score is the new high score value.
	////
	//// options can specify additional parameters.
	////
	//// Returns any error that occurred.
	//SetGameScoreInline(inlineMessageID string, userID int64, score int, options ...any) error
	//
	//// GetGameHighScores gets high score table for a game.
	////
	//// recipient is the chat containing the game.
	////
	//// userID is the ID of the target user.
	////
	//// options can specify additional parameters.
	////
	//// Returns the high scores and any error.
	//GetGameHighScores(recipient Recipient, userID int64, options ...any) ([]GameHighScore, error)

}
