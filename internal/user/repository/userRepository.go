package repository

import (
	"github.com/LuanTenorio/todo_go/internal/user/dto"
	"github.com/LuanTenorio/todo_go/internal/user/entity"
)

type UserRepository interface {
	CreateUser(user *dto.CreateUserDTO) (*entity.User, error)
}
