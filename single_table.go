package main

import (
	"chunli/global"
	"chunli/models"
	"fmt"
)

func migrate() {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		fmt.Println("Migrate failed", err)
		return
	}
	fmt.Println("Migrate success")
}
