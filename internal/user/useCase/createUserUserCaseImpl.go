package usecase

import (
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

func (uc *userUseCaseImpl) CreateUser(userDto *dto.CreateUserDTO) (*entity.User, error) {
	user, err := uc.userRepository.CreateUser(userDto)

	if err != nil {
		return &entity.User{}, err
	}

	return user, nil
}
