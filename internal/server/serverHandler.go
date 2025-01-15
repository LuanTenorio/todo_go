package server

import (
	userHandler "github.com/LuanTenorio/todo_go/internal/user/handler"
	userRepository "github.com/LuanTenorio/todo_go/internal/user/repository"
	userUseCase "github.com/LuanTenorio/todo_go/internal/user/useCase"
)

func (s *echoServer) bootHandlers() {
	bootUserHandler(s)
}

func bootUserHandler(s *echoServer) {
	userRepo := userRepository.NewUserPGRepository(s.db)
	userUc := userUseCase.NewUserUseCaseImpl(userRepo)
	userHand := userHandler.NewUserUseCaseImpl(userUc)

	const ApiRouterBase = ApiPrefix + "/users"
	userRoutes := s.app.Group(ApiRouterBase)

	userRoutes.POST("", userHand.CreateUser)
}
