package services

import (
	"chunli/global"
	"chunli/models"
	"errors"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(user *models.UserModel) error {
	if err := global.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(id uint) (*models.UserModel, error) {
	var user models.UserModel
	if err := global.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserList 获取用户列表
func (s *UserService) GetUserList(page, pageSize int) ([]models.UserModel, int64, error) {
	var users []models.UserModel
	var total int64

	offset := (page - 1) * pageSize

	// 获取总数
	if err := global.DB.Model(&models.UserModel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	if err := global.DB.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(id uint, user *models.UserModel) error {
	var existingUser models.UserModel
	if err := global.DB.First(&existingUser, id).Error; err != nil {
		return errors.New("user not found")
	}

	if err := global.DB.Model(&existingUser).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id uint) error {
	if err := global.DB.Delete(&models.UserModel{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByName 根据名称获取用户
func (s *UserService) GetUserByName(name string) (*models.UserModel, error) {
	var user models.UserModel
	if err := global.DB.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

