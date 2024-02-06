package telebot

import "encoding/json"

func (b *bot) GetUserChatBoosts(chatId Recipient, userId Recipient) (*UserChatBoosts, error) {
	params := getUserChatBoosts{
		ChatId: chatId.Recipient(),
		UserId: userId.Recipient(),
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
