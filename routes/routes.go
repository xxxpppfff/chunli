package routes

import (
	"chunli/controllers"
	"chunli/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 全局中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API 路由组
	api := r.Group("/api")
	{
		setupUserRoutes(api)
	}
}

func setupUserRoutes(rg *gin.RouterGroup) {
	userController := controllers.NewUserController()
	users := rg.Group("/users")
	{
		users.POST("", userController.CreateUser)      // 创建用户
		users.GET("", userController.GetUserList)     // 获取用户列表
		users.GET("/:id", userController.GetUserByID) // 获取用户详情
		users.PUT("/:id", userController.UpdateUser)   // 更新用户
		users.DELETE("/:id", userController.DeleteUser) // 删除用户
	}
}

