package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Name  string
	EMail string
}

func orm() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("dsn type is wrong %s", err)
		return
	}
	fmt.Println(db)

	var userList []User
	db.Table("user").Find(&userList)
	fmt.Println(userList)
}
