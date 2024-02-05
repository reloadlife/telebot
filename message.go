package telebot

import (
	"encoding/json"
	"fmt"
)

type MaybeInaccessibleMessage struct {
	*Message
	*InaccessibleMessage
}

func (m *MaybeInaccessibleMessage) IsAccessible() bool {
	return m.InaccessibleMessage.Date != 0 && m.Message.Date != 0
}

type InaccessibleMessage struct {
	MessageID int64 `json:"message_id"`
	Date      int64 `json:"date"`
	Chat      *Chat `json:"chat"`
}

// Message
// This struct represents a Telegram message and includes various fields to capture different aspects of the message.
//
// Fields:
// — ID (int64): Unique message identifier inside this chat.
// — ThreadID (int64): Optional. A unique identifier of a message thread to which the message belongs; for supergroups only.
// — From (*User): Optional. Sender of the message; empty for messages sent to channels. For backward compatibility,
// the field contains a fake sender user in non-channel chats if the message was sent on behalf of a chat.
// — SenderChat (*Chat): Optional. Sender of the message, sent on behalf of a chat. For example, the channel itself for channel posts,
// the supergroup itself for messages from anonymous group administrators,
// the linked channel for messages automatically forwarded to the discussion group. For backward compatibility, the field contains a fake sender user in non-channel chats if the message was sent on behalf of a chat.
// — Date (int64): Date the message was sent in Unix time. It is always a positive number, representing a valid date.
// — Chat (*Chat): Chat the message belongs to.
// — ForwardOrigin (*MessageOrigin): Optional. Information about the original message for forwarded messages.
// — IsTopic (*bool): Optional. True, if the message is sent to a forum topic.
// — IsAutomaticForward (*bool): Optional. True,
// if the message is a channel post that was automatically forwarded to the connected discussion group.
// — ReplyTo (*Message): Optional. For replies in the same chat and message thread,
// the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is
// a reply.
// — ExternalReply (*ExternalReplyInfo): Optional. Information about the message that is being replied to,
// which may come from another chat or forum topic.
// — Quote (*TextQuote): Optional. For replies that quote part of the original message, the quoted part of the message.
// — ViaBot (*User): Optional. Bot through which the message was sent.
// — EditDate (*int64): Optional. Date the message was last edited in Unix time.
// — HasProtectedContent (*bool): Optional. True, if the message can't be forwarded.
// — MediaGroupID (*string): Optional. The unique identifier of a media message group this message belongs to.
// — AuthorSignature (*string): Optional. Signature of the post-author for messages in channels,
// or the custom title of an anonymous group administrator.
// — Text (*string): Optional. For text messages, the actual UTF-8 text of the message.
// — Entities ([]Entity): Optional. For text messages, special entities like usernames, URLs, bot commands,
// etc. that appear in the text.
// — LinkPreviewOptions (*LinkPreviewOptions): Optional. Options used for link preview generation for the message,
// if it is a text message and link preview options were changed.
// — Animation (*Animation): Optional. Message is an animation, information about the animation.
// — Audio (*Audio): Optional. Message is an audio file, information about the file.
// — Document (*Document): Optional. Message is a general file, information about the file.
// — Photo ([]PhotoSize): Optional. Message is a photo, available sizes of the photo.
// — Sticker (*Sticker): Optional. Message is a sticker, information about the sticker.
// — Story (*Story): Optional. Message is a forwarded story.
// — Video (*Video): Optional. Message is a video, information about the video.
// — VideoNote (*VideoNote): Optional. Message is a video note, information about the video message.
// — Voice (*Voice): Optional. Message is a voice message, information about the file.
// — Caption (*string): Optional. Caption for the animation, audio, document, photo, video, or voice.
// — CaptionEntities ([]Entity): Optional. For messages with a caption, special entities like usernames, URLs, bot commands,
// etc. that appear in the caption.
// — HasMediaSpoiler (*bool): Optional. True, if the message media is covered by a spoiler animation.
// — Contact (*Contact): Optional. Message is a shared contact, information about the contact.
// — Dice (*Dice): Optional. Message is a 'dice' with a random value.
// — Game (*Game): Optional. Message is a game, information about the game. More about games »
// — Poll (*Poll): Optional. Message is a native poll, information about the poll.
// — Venue (*Venue): Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set,
// the location field will also be set.
// — Location (*Location): Optional. Message is a shared location, information about the location.
// — NewChatMembers ([]User): Optional. New members that were added to the group or supergroup and information about them
// (the bot itself may be one of these members).
// — LeftChatMember (*User): Optional. A member was removed from the group, information about them (this member may be the bot itself).
// — NewChatTitle (*string): Optional. A chat title was changed to this value.
// — NewChatPhoto ([]PhotoSize): Optional. A chat photo was changed to this value.
// — DeleteChatPhoto (*bool): Optional. Service message: the chat photo was deleted.
// — GroupChatCreated (*bool): Optional. Service message: the group has been created.
// — SupergroupChatCreated (*bool): Optional. Service message: the supergroup has been created.
// This field can't be received in a message coming through updates because the bot can't be a member of a supergroup when it is created.
// It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
// — ChannelChatCreated (*bool): Optional. Service message: the channel has been created.
// This field can't be received in a message coming through updates because the bot can't be a member of a channel when it is created.
// It can only be found in reply_to_message if someone replies to a very first message in a channel.
// — AutoDeleteTimerChanged (*AutoDeleteTimerChanged): Optional. Service message: auto-delete timer settings changed in the chat.
// — MigrateToChatID (*int64): Optional. The group has been migrated to a supergroup with the specified identifier.
// This number may have more than 32 significant bits,
// and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits,
// so a signed 64-bit integer or double-precision float type is safe for storing this identifier.
// — MigrateFromChatID (*int64): Optional. The supergroup has been migrated from a group with the specified identifier.
// This number may have more than 32 significant bits,
// and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits,
// so a signed 64-bit integer or double-precision float type is safe for storing this identifier.
// — PinnedMessage (*MaybeInaccessibleMessage): Optional. A specified message was pinned.
// Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
// — Invoice (*Invoice): Optional. Message is an invoice for a payment, information about the invoice.
// — SuccessfulPayment (*SuccessfulPayment): Optional. Message is a service message about a successful payment,
// information about the payment.
// — UsersShared (*UsersShared): Optional. Service message: users were shared with the bot.
// — ChatShared (*ChatShared): Optional. Service message: a chat was shared with the bot.
// — ConnectedWebsite (*string): Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
// — WriteAccessAllowed (*WriteAccessAllowed): Optional. Service message: the user allowed the bot to write messages after adding it to the
// attachment or side menu, launching a Web App from a link,
// or accepting an explicit request from a Web App sent by the method requestWriteAccess.
// — PassportData (*PassportData): Optional. Telegram Passport data.
// — ProximityAlertTriggered (*ProximityAlertTriggered): Optional. Service message.
// A user in the chat triggered another user's proximity alert while sharing Live Location.
// — ForumTopicCreated (*ForumTopicCreated): Optional. Service message: forum topic created.
// — ForumTopicEdited (*ForumTopicEdited): Optional. Service message: forum topic edited.
// — ForumTopicClosed (*ForumTopicClosed): Optional. Service message: forum topic closed.
// — ForumTopicReopened (*ForumTopicReopened): Optional. Service message: a forum topic reopened.
// — GeneralForumTopicHidden (*GeneralForumTopicHidden): Optional. Service message: the 'General' forum topic hidden.
// — GeneralForumTopicUnhidden (*GeneralForumTopicUnhidden): Optional. Service message: the 'General' forum topic unhidden.
// — GiveawayCreated (*GiveawayCreated): Optional. Service message: a scheduled giveaway was created.
// — Giveaway (*Giveaway): Optional. The message is a scheduled giveaway message.
// — GiveawayWinners (*GiveawayWinners): Optional. A giveaway with public winners was completed.
// — GiveawayCompleted (*GiveawayCompleted): Optional. Service message: a giveaway without public winners was completed.
// — VideoChatScheduled (*VideoChatScheduled): Optional. Service message: a video chat scheduled.
// — VideoChatStarted (*VideoChatStarted): Optional. Service message: a video chat started.
// — VideoChatEnded (*VideoChatEnded): Optional. Service message: a video chat ended.
// — VideoChatParticipantsInvited (*VideoChatParticipantsInvited): Optional. Service message: new participants invited to a video chat.
// — WebAppData (*WebAppData): Optional. Service message: data sent by a Web App.
// — ReplyMarkup (*InlineKeyboardMarkup): Optional. Inline keyboard attached to the message.
// login_url buttons are represented as ordinary URL buttons.
type Message struct {
	ID                           int64                         `json:"message_id"`
	ThreadID                     *int64                        `json:"message_thread_id,omitempty"`
	From                         *User                         `json:"from,omitempty"`
	SenderChat                   *Chat                         `json:"sender_chat,omitempty"`
	Date                         int64                         `json:"date"`
	Chat                         *Chat                         `json:"chat"`
	ForwardOrigin                *MessageOrigin                `json:"forward_origin,omitempty"`
	IsTopic                      *bool                         `json:"is_topic_message,omitempty"`
	IsAutomaticForward           *bool                         `json:"is_automatic_forward,omitempty"`
	ReplyTo                      *Message                      `json:"reply_to_message,omitempty"`
	ExternalReply                *ExternalReplyInfo            `json:"external_reply,omitempty"`
	Quote                        *TextQuote                    `json:"quote,omitempty"`
	ViaBot                       *User                         `json:"via_bot,omitempty"`
	EditDate                     *int64                        `json:"edit_date,omitempty"`
	HasProtectedContent          *bool                         `json:"has_protected_content,omitempty"`
	MediaGroupID                 *string                       `json:"media_group_id,omitempty"`
	AuthorSignature              *string                       `json:"author_signature,omitempty"`
	Text                         *string                       `json:"text,omitempty"`
	Entities                     []Entity                      `json:"entities,omitempty"`
	LinkPreviewOptions           *LinkPreviewOptions           `json:"link_preview_options,omitempty"`
	Animation                    *Animation                    `json:"animation,omitempty"`
	Audio                        *Audio                        `json:"audio,omitempty"`
	Document                     *Document                     `json:"document,omitempty"`
	Photo                        []PhotoSize                   `json:"photo,omitempty"`
	Sticker                      *Sticker                      `json:"sticker,omitempty"`
	Story                        *Story                        `json:"story,omitempty"`
	Video                        *Video                        `json:"video,omitempty"`
	VideoNote                    *VideoNote                    `json:"video_note,omitempty"`
	Voice                        *Voice                        `json:"voice,omitempty"`
	Caption                      *string                       `json:"caption,omitempty"`
	CaptionEntities              []Entity                      `json:"caption_entities,omitempty"`
	HasMediaSpoiler              *bool                         `json:"has_media_spoiler,omitempty"`
	Contact                      *Contact                      `json:"contact,omitempty"`
	Dice                         *Dice                         `json:"dice,omitempty"`
	Game                         *Game                         `json:"game,omitempty"`
	Poll                         *Poll                         `json:"poll,omitempty"`
	Venue                        *Venue                        `json:"venue,omitempty"`
	Location                     *Location                     `json:"location,omitempty"`
	NewChatMembers               []User                        `json:"new_chat_members,omitempty"`
	LeftChatMember               *User                         `json:"left_chat_member,omitempty"`
	NewChatTitle                 *string                       `json:"new_chat_title,omitempty"`
	NewChatPhoto                 []PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto              *bool                         `json:"delete_chat_photo,omitempty"`
	GroupChatCreated             *bool                         `json:"group_chat_created,omitempty"`
	SupergroupChatCreated        *bool                         `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated           *bool                         `json:"channel_chat_created,omitempty"`
	AutoDeleteTimerChanged       *AutoDeleteTimerChanged       `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID              *int64                        `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID            *int64                        `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                *MaybeInaccessibleMessage     `json:"pinned_message,omitempty"`
	Invoice                      *Invoice                      `json:"invoice,omitempty"`
	SuccessfulPayment            *SuccessfulPayment            `json:"successful_payment,omitempty"`
	UsersShared                  *UsersShared                  `json:"users_shared,omitempty"`
	ChatShared                   *ChatShared                   `json:"chat_shared,omitempty"`
	ConnectedWebsite             *string                       `json:"connected_website,omitempty"`
	WriteAccessAllowed           *WriteAccessAllowed           `json:"write_access_allowed,omitempty"`
	PassportData                 *PassportData                 `json:"passport_data,omitempty"`
	ProximityAlertTriggered      *ProximityAlertTriggered      `json:"proximity_alert_triggered,omitempty"`
	ForumTopicCreated            *ForumTopicCreated            `json:"forum_topic_created,omitempty"`
	ForumTopicEdited             *ForumTopicEdited             `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed             *ForumTopicClosed             `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened           *ForumTopicReopened           `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden      *GeneralForumTopicHidden      `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden    *GeneralForumTopicUnhidden    `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated              *GiveawayCreated              `json:"giveaway_created,omitempty"`
	Giveaway                     *Giveaway                     `json:"giveaway,omitempty"`
	GiveawayWinners              *GiveawayWinners              `json:"giveaway_winners,omitempty"`
	GiveawayCompleted            *GiveawayCompleted            `json:"giveaway_completed,omitempty"`
	VideoChatScheduled           *VideoChatScheduled           `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted             *VideoChatStarted             `json:"video_chat_started,omitempty"`
	VideoChatEnded               *VideoChatEnded               `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	WebAppData                   *WebAppData                   `json:"web_app_data,omitempty"`
	ReplyMarkup                  *InlineKeyboardMarkup         `json:"reply_markup,omitempty"`
}

// MarshalJSON to be JSON serializable, but only include non-empty fields.
func (u *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalJSON to be JSON unserializable
func (u *Message) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, u)
}

// String returns a string representation of this user.
func (u *Message) String() string {
	indented, _ := json.MarshalIndent(u, "", "  ")
	return fmt.Sprintf("Message{ID: %d, Chat: @%v}\n%s\n", u.ID, u.Chat.ID, indented)
}
