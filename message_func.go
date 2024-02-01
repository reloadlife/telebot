package telebot

import "encoding/json"

// sendMessageParams represents the parameters for the sendMessage method.
type sendMessageParams struct {
	// ChatID is the unique identifier for the target chat or username of the target channel
	// (in the format @channelusername).
	ChatID interface{} `json:"chat_id" validate:"required"`

	// MessageThreadID is the unique identifier for the target message thread (topic) of the forum;
	// applicable for forum supergroups only.
	MessageThreadID *int `json:"message_thread_id,omitempty"`

	// Text is the text of the message to be sent, 1-4096 characters after entities parsing.
	Text string `json:"text" validate:"required,max=4096"`

	// ParseMode is the mode for parsing entities in the message text. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`

	// Entities is a JSON-serialized list of special entities that appear in message text,
	// which can be specified instead of ParseMode.
	Entities []Entity `json:"entities,omitempty"`

	// LinkPreviewOptions is the link preview generation options for the message.
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// DisableNotification sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`

	// ProtectContent protects the contents of the sent message from forwarding and saving.
	ProtectContent *bool `json:"protect_content,omitempty"`

	// ReplyParameters is the description of the message to reply to.
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// ReplyMarkup is additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to remove reply keyboard, or to force a reply from the user.
	ReplyMarkup *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (b *bot) SendMessage(recipient Recipient, text string, option ...Option) (*Message, error) {
	params := sendMessageParams{
		ChatID: recipient.Recipient(),
		Text:   text,
	}

	options := extractOptions(option)

	if options != nil {
		params.ParseMode = &options.ParseMode
		params.Entities = options.Entities
		params.LinkPreviewOptions = &options.LinkPreviewOptions
		params.DisableNotification = &options.DisableNotification
		params.ProtectContent = &options.ProtectContent
		params.ReplyParameters = &options.ReplyParameters
		params.ReplyMarkup = options.ReplyMarkup
	}

	var resp struct {
		Result *Message
	}

	req, err := b.sendMethodRequest(methodSendMessage, params)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, err
}
