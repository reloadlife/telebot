package telebot

import (
	"encoding/json"
	"time"
)

type getUpdatesRequest struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func (b *bot) GetUpdates(offset, limit int, timeout time.Duration, allowed ...UpdateType) (Updates, error) {
	if b.offlineMode {
		return nil, nil
	}
	params := getUpdatesRequest{
		Offset:  offset,
		Timeout: int(timeout / time.Second),
	}

	if len(allowed) > 0 {
		params.AllowedUpdates = make([]string, len(allowed))
		for i, v := range allowed {
			params.AllowedUpdates[i] = v.String()
		}
	}

	if limit != 0 {
		params.Limit = limit
	}

	res, err := b.sendMethodRequest(methodGetUpdates, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result []Update
	}

	if err = json.Unmarshal(res, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}
