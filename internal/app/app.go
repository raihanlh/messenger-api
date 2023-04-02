package app

import (
	"gitlab.com/raihanlh/messenger-api/config"
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	userHandler "gitlab.com/raihanlh/messenger-api/internal/domain/user/delivery/handler"
	userRepository "gitlab.com/raihanlh/messenger-api/internal/domain/user/repository"
	userUsecase "gitlab.com/raihanlh/messenger-api/internal/domain/user/usecase"
	healthHandler "gitlab.com/raihanlh/messenger-api/internal/health/handler"
	"gitlab.com/raihanlh/messenger-api/pkg/postgres"
)

// Initiate databases
func NewDatabases(config *config.Config) *dependency.Databases {
	return &dependency.Databases{
		Main: postgres.New(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBTimezone),
	}
}

// Initiate repositories
func NewRepositories(db *dependency.Databases) *dependency.Repositories {
	return &dependency.Repositories{
		User: userRepository.New(db.Main),
	}
}

// Initiate Usecases
func NewUsecases(r *dependency.Repositories) *dependency.Usecases {
	return &dependency.Usecases{
		User: userUsecase.New(r),
	}
}

// Initiate repositories
func NewHandlers(u *dependency.Usecases) *dependency.Handlers {
	return &dependency.Handlers{
		User:   userHandler.New(u),
		Health: healthHandler.New(),
	}
}
