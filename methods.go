package telebot

import "encoding/json"

func (b *bot) Close() error {
	raw, err := b.sendMethodRequest(methodClose, nil)
	if err != nil {
		return err
	}

	var resp struct {
		Result bool
	}

	if err = json.NewDecoder(raw.Body).Decode(&resp); err != nil {
		return err
	}

	if !resp.Result {
		return ErrNotClosed
	}

	return nil
}

func (b *bot) Logout() error {
	raw, err := b.sendMethodRequest(methodLogOut, nil)
	if err != nil {
		return err
	}

	var resp struct {
		Result bool
	}

	if err = json.NewDecoder(raw.Body).Decode(&resp); err != nil {
		return err
	}

	if !resp.Result {
		return ErrNotClosed
	}

	return nil
}
