package telebot

import "encoding/json"

func (b *bot) GetUserChatBoosts(chatID Recipient, userID Userable) (*UserChatBoosts, error) {
	params := getUserChatBoosts{
		ChatId: chatID,
		UserId: userID,
	}

	req, err := b.sendMethodRequest(methodGetUserChatBoosts, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result UserChatBoosts `json:"result"`
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, err
}
