package repository

import (
	"github.com/LuanTenorio/todo_go/internal/database"
	"github.com/LuanTenorio/todo_go/internal/user/dto"
	"github.com/LuanTenorio/todo_go/internal/user/entity"
)

type userPGRepository struct {
	db database.Database
}

func NewUserPGRepository(db database.Database) UserRepository {
	return &userPGRepository{db: db}
}

func (r *userPGRepository) CreateUser(user *dto.CreateUserDTO) (*entity.User, error) {
	return &entity.User{}, nil
}
