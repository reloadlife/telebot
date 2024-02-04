package telebot

type getFileRequest struct {
	FileID string `json:"file_id"`
}

type getUserProfilePhotosRequest struct {
	UserID int64 `json:"user_id"`
	Offset int   `json:"offset,omitempty"`
	Limit  int   `json:"limit,omitempty"`
}
