package service

import (
	"clone/rent_car_us/storage"
	"clone/rent_car_us/pkg/logger"
)

type IServiceManager interface {
	Car() carService
	Customer() CustomerService
	Order() OrderService
	Auth() authService
}

type Service struct {
	carService carService
	CustomerService CustomerService
	OrderService OrderService
	logger logger.ILogger
	auth            authService
}

func New(storage storage.IStorage,log logger.ILogger, redis storage.IRedisStorage) Service {
	services := Service{}
	services.carService = NewCarService(storage,log)
	services.CustomerService = NewCustomerService(storage,log,redis)
	services.OrderService = NewOrderService(storage,log)
	services.logger=log
	services.auth=NewAuthService(storage, log, redis)

	return services
}

func (s Service) Car() carService {
	return s.carService
}

func (s Service) Customer() CustomerService {
	return s.CustomerService
}


func (s Service) Order() OrderService {
	return s.OrderService
}

func (s Service) Auth() authService {
	return s.auth
}

