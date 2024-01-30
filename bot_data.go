package telebot

import "encoding/json"

func (b *OldBot) SetMyDescription(description, languageCode string) error {
	_, err := b.Raw("setMyCommands", map[string]interface{}{
		"description":   description,
		"language_code": languageCode,
	})
	return err
}

type BotDescription struct {
	Description string `json:"description"`
}

func (b *OldBot) GetMyDescription(languageCode string) (*BotDescription, error) {
	params := map[string]string{
		"language_code": languageCode,
	}

	data, err := b.Raw("getMyCommands", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result BotDescription
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return &resp.Result, nil
}

func (b *OldBot) SetMyShortDescription(shortDescription, languageCode string) error {
	_, err := b.Raw("setMyShortDescription", map[string]interface{}{
		"short_description": shortDescription,
		"language_code":     languageCode,
	})
	return err
}

type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

func (b *OldBot) GetMyShortDescription(languageCode string) (*BotShortDescription, error) {
	data, err := b.Raw("getMyCommands", map[string]string{
		"language_code": languageCode,
	})
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result BotShortDescription
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return &resp.Result, nil
}

func (b *OldBot) SetMyName(name, languageCode string) error {
	_, err := b.Raw("setMyName", map[string]interface{}{
		"name":          name,
		"language_code": languageCode,
	})
	return err
}

type BotName struct {
	Name string `json:"name"`
}

func (b *OldBot) GetMyName(languageCode string) (*BotName, error) {
	data, err := b.Raw("getMyCommands", map[string]string{
		"language_code": languageCode,
	})
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result BotName
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return &resp.Result, nil
}
