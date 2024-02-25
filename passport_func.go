package telebot

// todo: requires a full re-write and full implementation of TelegramPassport (https://core.telegram.org/passport)

func (b *bot) SetPassportDataErrors(userID Userable, errors []PassportElementError) error {
	params := struct {
		UserID any                    `json:"user_id"`
		Errors []PassportElementError `json:"errors"`
	}{
		UserID: userID,
		Errors: errors,
	}

	_, err := b.sendMethodRequest(methodSetPassportDataErrors, params)
	return err
}
