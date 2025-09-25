package models

import "time"

type UserModel struct {
	ID        uint
	Name      string `gorm:"size:16;not null;unique"`
	Age       int    `gorm:"default:18"`
	EMail     string
	CreatedAt time.Time
}
