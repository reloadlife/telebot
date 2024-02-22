package telebot

import "encoding/json"

func (b *bot) GetUserProfilePhotos(userID int64, offset, limit int) (*UserProfilePhotos, error) {
	var resp struct {
		Result UserProfilePhotos
	}

	params := getUserProfilePhotosRequest{
		UserID: userID,
		Offset: offset,
		Limit:  limit,
	}

	req, err := b.sendMethodRequest(methodGetUserProfilePhotos, params)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}
