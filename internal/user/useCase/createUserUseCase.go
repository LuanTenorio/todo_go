package usecase

import (
	"context"

	"github.com/LuanTenorio/todo_go/internal/user/dto"
	"github.com/LuanTenorio/todo_go/internal/user/entity"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, userDto *dto.CreateUserDTO) (*entity.User, error)
}
