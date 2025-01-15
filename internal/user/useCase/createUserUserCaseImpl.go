package usecase

import (
	"context"

	"github.com/LuanTenorio/todo_go/internal/user/dto"
	"github.com/LuanTenorio/todo_go/internal/user/entity"
	"github.com/LuanTenorio/todo_go/internal/user/repository"
)

type userUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUseCaseImpl(ur repository.UserRepository) UserUseCase {
	return &userUseCaseImpl{userRepository: ur}
}

func (uc *userUseCaseImpl) CreateUser(ctx context.Context, userDto *dto.CreateUserDTO) (*entity.User, error) {
	id, err := uc.userRepository.CreateUser(ctx, userDto)

	if err != nil {
		return &entity.User{}, err
	}

	user := entity.NewUserByCreateDto(userDto, id)

	return user, nil
}
