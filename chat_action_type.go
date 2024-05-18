package telebot

type ChatAction string

const (
	ChatTyping                ChatAction = "typing"
	ChatUploadPhoto           ChatAction = "upload_photo"
	ChatRecordVideo           ChatAction = "record_video"
	ChatUploadVideo           ChatAction = "upload_video"
	ChatRecordVoice           ChatAction = "record_voice"
	ChatUploadVoice           ChatAction = "upload_voice"
	ChatActionUploadDocument  ChatAction = "upload_document"
	ChatActionChooseSticker   ChatAction = "choose_sticker"
	ChatActionFindLocation    ChatAction = "find_location"
	ChatActionRecordVideoNote ChatAction = "record_video_note"
	ChatActionUploadVideoNote ChatAction = "upload_video_note"
)

type sendChatActionRequest struct {
	ChatID     string           `json:"chat_id"`
	Action     ChatAction       `json:"action"`
	BusinessID *BusinessID      `json:"business_connection_id,omitempty"`
	ThreadID   *MessageThreadID `json:"thread_id,omitempty"`
}
