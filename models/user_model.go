package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        uint
	Name      string `gorm:"size:16;not null;unique"`
	Age       int    `gorm:"default:18"`
	EMail     string
	CreatedAt time.Time
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Hook BeforeCreate")
	return
}
