package telebot

import "fmt"

type UserProfilePhotos struct {
	TotalCount int          `json:"total_count"`
	Photos     []PhotoSizes `json:"photos"`
}

func (u *UserProfilePhotos) ReflectType() string { return fmt.Sprintf("%T", u) }
func (u *UserProfilePhotos) Type() string        { return "UserProfilePhotos" }
