package telebot

import "fmt"

type UserProfilePhotos struct {
	TotalCount int          `json:"total_count"`
	Photos     []PhotoSizes `json:"photos"`
}

func (c *UserProfilePhotos) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *UserProfilePhotos) Type() string        { return "UserProfilePhotos" }
