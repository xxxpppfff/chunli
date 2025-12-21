package controllers

import (
	"chunli/models"
	"chunli/services"
	"chunli/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body models.UserModel true "用户信息"
// @Success 200 {object} response.Response
// @Router /api/users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	if err := uc.userService.CreateUser(&user); err != nil {
		response.InternalError(c, "Failed to create user: "+err.Error())
		return
	}

	response.Success(c, user)
}

// GetUserByID 根据ID获取用户
// @Summary 获取用户详情
// @Description 根据ID获取用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response
// @Router /api/users/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}

	user, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		response.NotFound(c, "User not found")
		return
	}

	response.Success(c, user)
}

// GetUserList 获取用户列表
// @Summary 获取用户列表
// @Description 分页获取用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(10)
// @Success 200 {object} response.Response
// @Router /api/users [get]
func (uc *UserController) GetUserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	users, total, err := uc.userService.GetUserList(page, pageSize)
	if err != nil {
		response.InternalError(c, "Failed to get user list: "+err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":     users,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// UpdateUser 更新用户
// @Summary 更新用户
// @Description 更新用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body models.UserModel true "用户信息"
// @Success 200 {object} response.Response
// @Router /api/users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}

	var user models.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, "Invalid request data: "+err.Error())
		return
	}

	if err := uc.userService.UpdateUser(uint(id), &user); err != nil {
		response.InternalError(c, "Failed to update user: "+err.Error())
		return
	}

	response.SuccessWithMessage(c, "User updated successfully", nil)
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 根据ID删除用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response
// @Router /api/users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}

	if err := uc.userService.DeleteUser(uint(id)); err != nil {
		response.InternalError(c, "Failed to delete user: "+err.Error())
		return
	}

	response.SuccessWithMessage(c, "User deleted successfully", nil)
}

