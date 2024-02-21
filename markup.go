package telebot

import (
	"encoding/json"
	"errors"
)

type ReplyMarkup interface {
	ReplyMarkup()

	AddRow(row ...Button)
	AddRows(rows ...Row)

	Inline()
	Keyboard(opts ...any)
	Remove(opts ...any)
	ForceReply(opts ...any)

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error

	deepEqual(other any) bool
}

func NewMarkup() ReplyMarkup {
	return &baseMarkUp{}
}

func (*baseMarkUp) ReplyMarkup() {}

type baseMarkUp struct {
	markupType markupType
	inline     *InlineKeyboardMarkup
	keyboard   *ReplyKeyboardMarkup
	remove     *ReplyKeyboardRemove
	forceReply *ForceReplyMarkup
}

func (b *baseMarkUp) Remove(opts ...any) {
	isSelective := false
	if len(opts) > 0 {
		for _, opt := range opts {
			switch opt.(type) {
			case Option:
				switch opt.(Option) {
				case Selective:
					isSelective = true
				default:
					panic("telebot: unsupported option for RemoveKeyboard")
				}
			default:
				panic("telebot: unsupported option for RemoveKeyboard")
			}

		}
	}
	b.markupType = markupTypeRemoveKeyboard
	b.remove = &ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      &isSelective,
	}
}
func (b *baseMarkUp) ForceReply(opts ...any) {
	isSelective := false
	input := ""

	if len(opts) > 0 {
		for _, opt := range opts {
			switch opt.(type) {
			case string:
				input = opt.(string)

			case Option:
				switch opt.(Option) {
				case Selective:
					isSelective = true
				default:
					panic("telebot: unsupported option for forceReplyKeyboard")
				}
			default:
				panic("telebot: unsupported option for forceReplyKeyboard")
			}

		}
	}
	b.markupType = markupTypeForceReply
	b.forceReply = &ForceReplyMarkup{
		ForceReply:            true,
		InputFieldPlaceholder: &input,
		Selective:             &isSelective,
	}
}
func (b *baseMarkUp) Inline() {
	b.markupType = markupTypeInline
	b.inline = &InlineKeyboardMarkup{
		InlineKeyboard: make([]Row, 0), // basically a Nil Slice of Rows with no Keys
	}
}
func (b *baseMarkUp) Keyboard(opts ...any) {
	isSelective := false
	input := ""
	resize := false
	oneTime := false
	persistent := false

	if len(opts) > 0 {
		for _, opt := range opts {
			switch opt.(type) {
			case string:
				input = opt.(string)

			case Option:
				switch opt.(Option) {
				case ResizeKeyboard:
					resize = true
				case OneTimeKeyboard:
					oneTime = true
				case PersistentKeyboard:
					persistent = true
				case Selective:
					isSelective = true
				default:
					panic("telebot: unsupported option for keyboard")
				}
			default:
				panic("telebot: unsupported option for keyboard")
			}

		}
	}
	b.markupType = markupTypeKeyboard
	b.keyboard = &ReplyKeyboardMarkup{
		Keyboard:              make([]Row, 0),
		IsPersistent:          &persistent,
		ResizeKeyboard:        &resize,
		OneTimeKeyboard:       &oneTime,
		InputFieldPlaceholder: &input,
		Selective:             &isSelective,
	}
}

func (b *baseMarkUp) MarshalJSON() ([]byte, error) {
	switch b.markupType {
	case markupTypeInline:
		return json.Marshal(b.inline)
	case markupTypeKeyboard:
		return json.Marshal(b.keyboard)
	case markupTypeRemoveKeyboard:
		return json.Marshal(b.remove)
	case markupTypeForceReply:
		return json.Marshal(b.forceReply)
	}
	return nil, nil
}

func (b *baseMarkUp) UnmarshalJSON(data []byte) error {
	var temp struct {
		InlineKeyboard any `json:"inline_keyboard"`
		Keyboard       any `json:"keyboard"`
		RemoveKeyboard any `json:"remove_keyboard"`
		ForceReply     any `json:"force_reply"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch {
	case temp.InlineKeyboard != nil:
		b.markupType = markupTypeInline
		b.inline = &InlineKeyboardMarkup{}
		return b.inline.UnmarshalJSON(data)

	case temp.Keyboard != nil:
		b.markupType = markupTypeKeyboard
		b.keyboard = &ReplyKeyboardMarkup{}
		return b.keyboard.UnmarshalJSON(data)

	case temp.RemoveKeyboard != nil:
		b.markupType = markupTypeRemoveKeyboard
		b.remove = &ReplyKeyboardRemove{}
		return json.Unmarshal(data, b.remove)

	case temp.ForceReply != nil:
		b.markupType = markupTypeForceReply
		b.forceReply = &ForceReplyMarkup{}
		return json.Unmarshal(data, b.forceReply)
	}

	return errors.New("telebot: unknown markup type")
}

func (b *baseMarkUp) AddRows(rows ...Row) {
	switch b.markupType {
	case markupTypeInline:
		b.inline.InlineKeyboard = append(b.inline.InlineKeyboard, rows...)
	case markupTypeKeyboard:
		b.keyboard.Keyboard = append(b.keyboard.Keyboard, rows...)

	default:
		panic("telebot: only Inline and Keyboard markups can have rows.")
	}
}
func (b *baseMarkUp) AddRow(row ...Button) {
	switch b.markupType {
	case markupTypeInline:
		b.inline.InlineKeyboard = append(b.inline.InlineKeyboard, row)
	case markupTypeKeyboard:
		b.keyboard.Keyboard = append(b.keyboard.Keyboard, row)

	default:
		panic("telebot: only Inline and Keyboard markups can have rows.")
	}
}
