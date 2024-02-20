package telebot

func (u *Update) Verify() error                      { return verify(u) }
func (r *MessageReaction) Verify() error             { return verify(r) }
func (r *MessageReactionCountUpdated) Verify() error { return verify(r) }

func (u *AccessibleMessage) Verify() error        { return verify(u) }
func (u *InaccessibleMessage) Verify() error      { return verify(u) }
func (u *MaybeInaccessibleMessage) Verify() error { return verify(u) }
