// Func: String() for all types
// the reason for this, is that so if anything goes wrong with the String, outputs,
//we can just come here and know where to look.

package telebot

import (
	"encoding/json"
	"fmt"
)

func (u *Update) String() string {
	indented, _ := json.MarshalIndent(u, "", "  ")
	return fmt.Sprintf("%s{%d}\n%s\n", u.Type(), u.ID, indented)
}

func (u *User) String() string {
	indented, _ := json.MarshalIndent(u, "", "  ")
	isBot := ""
	if u.IsBot {
		isBot = "(Bot)"
	}
	return fmt.Sprintf("User{%sID: %d, User: @%v}\n%s\n", isBot, u.ID, u.Username, indented)
}
