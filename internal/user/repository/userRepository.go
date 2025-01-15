package repository

import (
	"context"

	"github.com/LuanTenorio/todo_go/internal/user/dto"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *dto.CreateUserDTO) (int, error)
}
