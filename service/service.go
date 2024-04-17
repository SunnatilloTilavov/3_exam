package service

import (
	"clone/3_exam/storage"
	"clone/3_exam/pkg/logger"
)

type IServiceManager interface {
	User() UserService
	Auth() authService
}

type Service struct {
	UserService UserService
	logger logger.ILogger
	auth            authService
}

func New(storage storage.IStorage,log logger.ILogger, redis storage.IRedisStorage) Service {
	services := Service{}
	services.UserService = NewUserService(storage,log,redis)
	services.logger=log
	services.auth=NewAuthService(storage, log, redis)

	return services
}

func (s Service) User() UserService {
	return s.UserService
}

func (s Service) Auth() authService {
	return s.auth
}

