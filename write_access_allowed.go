package telebot

import "fmt"

// WriteAccessAllowed represents a service message about a user allowing a bot to write messages
// after adding it to the attachment menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess.
type WriteAccessAllowed struct {
	// FromRequest is an optional field indicating if the access was granted after the user accepted an explicit request from a Web App sent by the method requestWriteAccess.
	FromRequest *bool `json:"from_request,omitempty"`

	// WebAppName is an optional field representing the name of the Web App if the access was granted when the Web App was launched from a link.
	WebAppName *string `json:"web_app_name,omitempty"`

	// FromAttachmentMenu is an optional field indicating if the access was granted when the bot was added to the attachment or side menu.
	FromAttachmentMenu *bool `json:"from_attachment_menu,omitempty"`
}

func (c *WriteAccessAllowed) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *WriteAccessAllowed) Type() string        { return "WriteAccessAllowed" }
