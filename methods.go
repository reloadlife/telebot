package telebot

import "go.mamad.dev/gtb/log"

func (b *bot) Close() error {
	if b.offlineMode {
		log.Infof("Bot is in offline mode, Close() is a no-op")
		return nil
	}
	_, err := b.sendMethodRequest(methodClose, nil)
	return err
}

func (b *bot) Logout() error {
	if b.offlineMode {
		log.Infof("Bot is in offline mode, Logout() is a no-op")
		return nil
	}
	_, err := b.sendMethodRequest(methodLogOut, nil)
	return err
}
