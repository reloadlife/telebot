package telebot

func (b *bot) Close() error {
	_, err := b.sendMethodRequest(methodClose, nil)
	return err
}

func (b *bot) Logout() error {
	_, err := b.sendMethodRequest(methodLogOut, nil)
	return err
}
