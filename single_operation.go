package main

import (
	"chunli/global"
	"chunli/models"
)

func singleOperation() {
	// err := global.DB.Create(&models.UserModel{
	// 	Name:  "Rick",
	// 	Age:   66,
	// 	EMail: "rick@163.com",
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

	// var userList = []models.UserModel{
	// 	{Name: "Cammy", Age: 25, EMail: "Cammy@163.com"},
	// }
	// err := global.DB.Create(&userList).Error
	// if err != nil {
	// 	fmt.Println("Create failed", err)
	// 	return
	// }

	// var userList []models.UserModel
	// global.DB.Find(&userList)
	// fmt.Println(userList)

	// var user models.UserModel
	// global.DB.Debug().Take(&user, 4)
	// fmt.Println(user)

	// var user models.UserModel
	// user.Name = "Snow"
	// global.DB.Save(&user)
	// fmt.Println(user)

	// var user models.UserModel
	// user.ID = 9
	// user.Age = 0
	// user.Name = "Xue"
	// user.CreatedAt = time.Now()
	// global.DB.Debug().Save(&user)
	// fmt.Println(user)

	// var user = models.UserModel{ID: 9}
	// //global.DB.Model(&user).Update("Age", 22)
	// global.DB.Model(&user).UpdateColumn("Age", 19)
	// fmt.Println(user)

	// var user = models.UserModel{ID: 9}
	// // global.DB.Model(&user).Updates(models.UserModel{
	// // 	Name:      "Kong",
	// // 	Age:       29,
	// // 	EMail:     "kong@163.com",
	// // 	CreatedAt: time.Now(),
	// // })
	// global.DB.Model(&user).Updates(map[string]any{
	// 	"Name":      "Kongkong",
	// 	"Age":       2,
	// 	"EMail":     "kong@163.com",
	// 	"CreatedAt": time.Now(),
	// })
	// fmt.Println(user)

	// var user = models.UserModel{ID: 3}
	// global.DB.Model(&user).Update("Age", gorm.Expr("Age + ?", 1))
	// fmt.Println(user)

	// var user = models.UserModel{ID: 9}
	// global.DB.Model(&user).Delete(&user)

	// global.DB.Delete(&models.UserModel{}, 8)
	global.DB.Delete(&models.UserModel{}, []int{4, 5})
}
