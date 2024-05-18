package telebot

import (
	"encoding/json"
)

func (b *bot) GetBusinessConnection(BusinessConnectionID BusinessID) (*BusinessConnection, error) {
	params := getBusinessConnectionRequest{
		ID: BusinessConnectionID,
	}

	req, err := b.sendMethodRequest(methodGetBusinessConnection, params)
	var resp struct {
		Result *BusinessConnection
	}

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
