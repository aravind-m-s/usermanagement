package dependency

import (
	http "usermanagement/pkg/api"
	"usermanagement/pkg/api/handler"
	"usermanagement/pkg/config"
	"usermanagement/pkg/db"
	"usermanagement/pkg/repository"
	"usermanagement/pkg/service"
	usecase "usermanagement/pkg/useCase"
)

func InitializeDependencies(cnf config.Config) (*http.ServerHTTP, error) {
	db, err := db.ConnectToDb(cnf)
	if err != nil {
		return nil, err
	}

	println(db)

	userService := service.NewUserService()
	userRepo := repository.NewUserRepository(db, userService)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)
	userServer := http.NewServerHTTP(userHandler)

	return userServer, nil

}
