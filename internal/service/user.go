package service

import (
	"context"

	"github.com/nuanxinqing123/flying-bird/internal/model"
	"github.com/nuanxinqing123/flying-bird/internal/repository"
)

type UserService struct{}

// NewUserService 创建 UserService 实例
func NewUserService() *UserService {
	return &UserService{}
}

// GetAllUsers 获取所有用户
func (s *UserService) GetAllUsers() (interface{}, error) {
	return repository.Users.Find()
}

// GetUserByID 根据 ID 获取用户
func (s *UserService) GetUserByID(id int64) (interface{}, error) {
	return repository.Users.Where(repository.Users.ID.Eq(id)).First()
}

// CreateUser 创建用户
func (s *UserService) CreateUser(name, email string, age int) (interface{}, error) {
	ageInt := int32(age)
	// 使用 Gorm Gen 创建用户
	user := model.Users{
		Name:  name,
		Email: email,
		Age:   &ageInt,
	}

	return user, repository.Users.WithContext(context.Background()).Create(&user)
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(id int64, name string, age int) error {
	_, err := repository.Users.Where(repository.Users.ID.Eq(id)).
		Updates(map[string]any{
			"Name": name,
			"Age":  int32(age),
		})
	return err
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id int64) error {
	_, err := repository.Users.Where(repository.Users.ID.Eq(id)).Delete()
	return err
}
