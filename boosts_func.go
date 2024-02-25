package telebot

import "encoding/json"

type getUserChatBoosts struct {
	ChatID any `json:"chat_id"`
	UserID any `json:"user_id"`
}

// GetUserChatBoosts returns a list of boosts added to a chat by a user.
// <a href="https://core.telegram.org/bots/api#getuserchatboosts">/bots/api#getuserchatboosts</a>
//
// chatID: Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
//
// userID: Unique identifier of the target user
//
// Returns: (*UserChatBoosts, error)
//
// Example:
// ```go
//
//	chatID := &Chat{ID:1234567890}
//	userID := &User{ID:1234567890}
//
//	chatBoosts, err := bot.GetUserChatBoosts(chatID, userID)
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(chatBoosts)
//
// ```
func (b *bot) GetUserChatBoosts(chatID Recipient, userID Userable) (*UserChatBoosts, error) {
	params := getUserChatBoosts{
		ChatID: chatID,
		UserID: userID,
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
