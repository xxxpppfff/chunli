package main

import "chunli/global"

func main() {
	// //初始化，返回gin的引擎
	// r := gin.Default()
	// //加路由，参数是个回调
	// r.GET("/", func(context *gin.Context) {
	// 	context.String(http.StatusOK, "hello,gin")
	// })
	// //加中间件
	// //r.Use(gin.Logger())
	// //绑定端口
	// r.Run(":8080")

	// originSql()
	// orm()
	global.Connect()
	// migrate()
	singleOperation()
}
