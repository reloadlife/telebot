package telebot

import (
	"encoding/json"
	"testing"
)

func TestRights(t *testing.T) {
	tests := []struct {
		name string
		r    Rights
		want string
	}{
		{
			name: "default rights",
			r:    Rights{},
			want: `{"anonymous":false,"can_be_edited":true,"can_change_info":false,"can_post_messages":false,"can_edit_messages":false,"can_delete_messages":false,"can_pin_messages":false,"can_invite_users":false,"can_restrict_members":false,"can_promote_members":false,"can_send_messages":true,"can_send_audios":true,"can_send_documents":true,"can_send_photos":true,"can_send_videos":true,"can_send_video_notes":true,"can_send_voice_notes":true,"can_send_polls":true,"can_send_other":true,"can_add_previews":true,"can_manage_video_chats":false,"can_manage_chat":false,"can_manage_topics":false,"can_post_stories":false,"can_edit_stories":false,"can_delete_stories":false}`,
		},
		{
			name: "admin rights",
			r:    AdminRights(),
			want: `{"anonymous":false,"can_be_edited":true,"can_change_info":true,"can_post_messages":true,"can_edit_messages":true,"can_delete_messages":true,"can_pin_messages":true,"can_invite_users":true,"can_restrict_members":true,"can_promote_members":true,"can_send_messages":true,"can_send_audios":true,"can_send_documents":true,"can_send_photos":true,"can_send_videos":true,"can_send_video_notes":true,"can_send_voice_notes":true,"can_send_polls":true,"can_send_other":true,"can_add_previews":true,"can_manage_video_chats":true,"can_manage_chat":true,"can_manage_topics":false,"can_post_stories":false,"can_edit_stories":false,"can_delete_stories":false}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := json.Marshal(test.r)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}
			if string(got) != test.want {
				t.Errorf("json.Marshal(%v) = %s, want %s", test.r, got, test.want)
			}
		})
	}
}
