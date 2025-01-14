package usecase

import (
	"github.com/LuanTenorio/todo_go/internal/user/dto"
	"github.com/LuanTenorio/todo_go/internal/user/entity"
)

type UserUseCase interface {
	CreateUser(user *dto.CreateUserDTO) (*entity.User, error)
}
