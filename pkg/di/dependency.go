package dependency

import (
	http "usermanagement/pkg/api"
	"usermanagement/pkg/api/handler"
	"usermanagement/pkg/config"
	"usermanagement/pkg/db"
	"usermanagement/pkg/repository"
	usecase "usermanagement/pkg/useCase"
)

func InitializeDependencies(cnf config.Config) (*http.ServerHTTP, error) {
	db, err := db.ConnectToDb(cnf)
	if err != nil {
		return nil, err
	}

	println(db)

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)
	userServer := http.NewServerHTTP(userHandler)

	return userServer, nil

}
