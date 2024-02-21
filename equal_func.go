package telebot

import "reflect"

func (m *ForceReplyMarkup) deepEqual(other any) bool {
	if m == nil {
		return other == nil
	}
	otherMarkup, ok := other.(*ForceReplyMarkup)
	if !ok {
		return false
	}

	return m.ForceReply == otherMarkup.ForceReply &&
		m.InputFieldPlaceholder == otherMarkup.InputFieldPlaceholder &&
		m.Selective == otherMarkup.Selective
}
func (m *ReplyKeyboardRemove) deepEqual(other any) bool {
	if m == nil {
		return other == nil
	}
	otherMarkup, ok := other.(*ReplyKeyboardRemove)
	if !ok {
		return false
	}

	return m.RemoveKeyboard == otherMarkup.RemoveKeyboard &&
		m.Selective == otherMarkup.Selective
}
func (m *InlineKeyboardMarkup) deepEqual(other any) bool {
	if m == nil {
		return other == nil
	}
	otherMarkup, ok := other.(*InlineKeyboardMarkup)
	if !ok {
		return false
	}

	if len(m.InlineKeyboard) != len(otherMarkup.InlineKeyboard) {
		return false
	}

	for i, row := range m.InlineKeyboard {
		if !row.deepEqual(otherMarkup.InlineKeyboard[i]) {
			return false
		}
	}

	return true
}
func (m *ReplyKeyboardMarkup) deepEqual(other any) bool {
	if m == nil {
		return other == nil
	}
	otherMarkup, ok := other.(*ReplyKeyboardMarkup)
	if !ok {
		return false
	}

	if len(m.Keyboard) != len(otherMarkup.Keyboard) {
		return false
	}

	for i, row := range m.Keyboard {
		if !row.deepEqual(otherMarkup.Keyboard[i]) {
			return false
		}
	}

	return m.IsPersistent == otherMarkup.IsPersistent &&
		m.ResizeKeyboard == otherMarkup.ResizeKeyboard &&
		m.OneTimeKeyboard == otherMarkup.OneTimeKeyboard &&
		m.InputFieldPlaceholder == otherMarkup.InputFieldPlaceholder &&
		m.Selective == otherMarkup.Selective
}
func (r *Row) deepEqual(other any) bool {
	if r == nil {
		return other == nil
	}
	otherRow, ok := other.(Row)
	if !ok {
		return false
	}

	if len(*r) != len(otherRow) {
		return false
	}

	for i, button := range *r {
		if !button.deepEqual(otherRow[i]) {
			return false
		}
	}

	return true
}
func (i *InlineKeyboardButton) deepEqual(other any) bool {
	otherButton, ok := other.(*InlineKeyboardButton)
	if !ok {
		return false
	}

	return reflect.DeepEqual(i, otherButton)
}
func (i *KeyboardButton) deepEqual(other any) bool {
	otherButton, ok := other.(*KeyboardButton)
	if !ok {
		return false
	}

	return reflect.DeepEqual(i, otherButton)
}
func (b *baseMarkUp) deepEqual(other any) bool {
	if b == other {
		return true
	}

	otherMarkup, ok := other.(*baseMarkUp)
	if !ok {
		inlineMarkup, ok := other.(*InlineKeyboardMarkup)
		if ok {
			return b.inline.deepEqual(inlineMarkup)
		}

		keyboardMarkup, ok := other.(*ReplyKeyboardMarkup)
		if ok {
			return b.keyboard.deepEqual(keyboardMarkup)
		}

		removeMarkup, ok := other.(*ReplyKeyboardRemove)
		if ok {
			return b.remove.deepEqual(removeMarkup)
		}

		forceReplyMarkup, ok := other.(*ForceReplyMarkup)
		if ok {
			return b.forceReply.deepEqual(forceReplyMarkup)
		}

		return false
	}

	if b.markupType != otherMarkup.markupType {
		return false
	}

	switch b.markupType {
	case markupTypeInline:
		return b.inline.deepEqual(other.(*baseMarkUp).inline)
	case markupTypeKeyboard:
		return b.keyboard.deepEqual(other.(*baseMarkUp).keyboard)
	case markupTypeRemoveKeyboard:
		return b.remove.deepEqual(other.(*baseMarkUp).remove)
	case markupTypeForceReply:
		return b.forceReply.deepEqual(other.(*baseMarkUp).forceReply)

	default:
		panic("telebot: unknown markup type")
	}

	return false
}
