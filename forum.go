package telebot

import "fmt"

// ForumTopicCreated represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	// Name is the name of the topic.
	Name string `json:"name"`

	// IconColor is the color of the topic icon in RGB format.
	IconColor IconColor `json:"icon_color"`

	// IconCustomEmojiID is an optional unique identifier of the custom emoji shown as the topic icon.
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

func (c *ForumTopicCreated) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ForumTopicCreated) Type() string        { return "ForumTopicCreated" }

// ForumTopicEdited represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	// Name is optional. New name of the topic, if it was edited.
	Name string `json:"name,omitempty"`

	// IconCustomEmojiID is optional. New identifier of the custom emoji shown as the topic icon, if it was edited;
	// an empty string if the icon was removed.
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

func (c *ForumTopicEdited) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ForumTopicEdited) Type() string        { return "ForumTopicEdited" }

// ForumTopicClosed represents a service message about a forum topic closed in the chat.
// Currently holds no information.
type ForumTopicClosed struct{}

func (c *ForumTopicClosed) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ForumTopicClosed) Type() string        { return "ForumTopicClosed" }

// ForumTopicReopened represents a service message about a forum topic reopened in the chat.
// Currently holds no information.
type ForumTopicReopened struct{}

func (c *ForumTopicReopened) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ForumTopicReopened) Type() string        { return "ForumTopicReopened" }

// GeneralForumTopicHidden represents a service message about General forum topic hidden in the chat.
// Currently holds no information.
type GeneralForumTopicHidden struct{}

func (c *GeneralForumTopicHidden) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *GeneralForumTopicHidden) Type() string        { return "GeneralForumTopicHidden" }

// GeneralForumTopicUnhidden represents a service message about General forum topic unhidden in the chat.
// Currently holds no information.
type GeneralForumTopicUnhidden struct{}

func (c *GeneralForumTopicUnhidden) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *GeneralForumTopicUnhidden) Type() string        { return "GeneralForumTopicUnhidden" }
