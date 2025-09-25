package main

import (
	"chunli/global"
	"chunli/models"
	"fmt"
)

func singleOperation() {
	// err := global.DB.Create(&models.UserModel{
	// 	Name:  "chunli",
	// 	Age:   25,
	// 	EMail: "chunli@163.com",
	// }).Error
	// if err != nil {
	// 	fmt.Println("Create failed", err)
	// 	return
	// }

	// user := models.UserModel{
	// 	Name:  "viper",
	// 	Age:   25,
	// 	EMail: "chunli@163.com",
	// }
	// err := global.DB.Create(&user).Error
	// if err != nil {
	// 	fmt.Println("Create failed", err)
	// 	return
	// }
	// fmt.Println(user.ID, user.CreatedAt, user.Name)

	var userList = []models.UserModel{
		{Name: "Manon", Age: 25, EMail: "Manon@163.com"},
		{Name: "Mike", Age: 25, EMail: "Mike@163.com"},
	}
	err := global.DB.Create(&userList).Error
	if err != nil {
		fmt.Println("Create failed", err)
		return
	}
}
