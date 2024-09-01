package di

import (
	"repocon/infrastructure/postgres/repository"
	handler "repocon/interface/api"
	"repocon/usecase"

	"gorm.io/gorm"
)

func User(db *gorm.DB) handler.IUserHandler {
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)
	return uh
}
