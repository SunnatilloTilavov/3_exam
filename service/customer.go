package service

import (
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/storage"
	"context"
	"clone/rent_car_us/pkg/logger"
	// "encoding/json"
)

type CustomerService struct {
	storage storage.IStorage
	logger  logger.ILogger
	redis   storage.IRedisStorage
}


func NewCustomerService(storage storage.IStorage,logger logger.ILogger,redis storage.IRedisStorage) CustomerService {
	return CustomerService{
		storage: storage,
		logger:  logger,
		redis:   redis,
	}
}
func (u CustomerService) Create(ctx context.Context, Customer models.CreateCustomer) (string, error) {

	pKey, err := u.storage.Customer().Create(ctx, Customer)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating Customer", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (u CustomerService) Update(ctx context.Context, Customer models.GetCustomer) (string, error) {

	pKey, err := u.storage.Customer().Update(ctx, Customer)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating Customer", logger.Error(err))
		return "", err
	}
	
	err = u.redis.Del(ctx, Customer.Id)
	if err != nil {
		u.logger.Error("error while setting otpCode to redis customer update", logger.Error(err))
		return "error redis update",err
	}

	return pKey, nil
}

func (u CustomerService) Delete(ctx context.Context, id string) error {

	err := u.storage.Customer().Delete(ctx, id)
	if err != nil {
		u.logger.Error("error service delete Customer", logger.Error(err))
		return err
	}

	err = u.redis.Del(ctx, id)
	if err != nil {
		u.logger.Error("error while setting otpCode to redis customer deleted", logger.Error(err))
		return err
	}

	return nil
}

func (u CustomerService) GetAllCustomers(ctx context.Context, Customer models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {

	pKey, err := u.storage.Customer().GetAllCustomers(ctx, Customer)
	if err != nil {
		u.logger.Error("ERROR in service layer while getalling Customer", logger.Error(err))
		return models.GetAllCustomersResponse{}, err
	}

	return pKey, nil
}

//////////

func (u CustomerService) GetByIDCustomer(ctx context.Context, Id string) (models.GetCustomer, error) {
	

	pKey, err := u.storage.Customer().GetByID(ctx, Id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyID Customer",logger.Error(err))
		return models.GetCustomer{}, err
	}

	return pKey, nil
}

func (u CustomerService) GetAllCustomerCars(ctx context.Context,Customer models.GetAllCustomerCarsRequest) (models.GetAllCustomerCarsResponse, error) {
	pKey, err := u.storage.Customer().GetAllCustomerCars(ctx, Customer)
	if err != nil {
		u.logger.Error("ERROR in service layer while getalling Customer", logger.Error(err))
		return models.GetAllCustomerCarsResponse{}, err
	}

	return pKey, nil
}


func (u CustomerService) UpdatePassword(ctx context.Context, Customer models.PasswordCustomer) (string, error) {

	pKey, err := u.storage.Customer().UpdatePassword(ctx, Customer)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating Customer", logger.Error(err))
		return "", err
	}

	return pKey, nil
}


func (u CustomerService) GetPassword(ctx context.Context, phone string) (string, error) {
	pKey, err := u.storage.Customer().GetPassword(ctx, phone)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyID Customer",logger.Error(err))
		return "Error", err
	}

	return pKey, nil
}




