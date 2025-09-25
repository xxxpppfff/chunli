package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func originSql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("dsn type is wrong %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("db connect failed %s", err)
	}
	fmt.Println(db)

	// res, err := db.Exec("INSERT INTO user (name, age, email) VALUES ('张三', 25, 'zhangsan@example.com'),('李四', 30, 'lisi@example.com'),('王五', 28, 'wangwu@example.com');")
	// if err != nil {
	// 	log.Fatalf("insert failed %s", err)
	// 	return
	// }
	// fmt.Println(res)

	rows, _ := db.Query("SELECT id, name FROM user;")
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		fmt.Println(id, name, err)
	}

	// var id int
	// var name string
	// err = db.QueryRow("SELECT id, name FROM user").Scan(&id, &name)
	// fmt.Println(id, name, err)
}
