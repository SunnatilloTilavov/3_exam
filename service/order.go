package service

import (
	"context"
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/storage"
	"clone/rent_car_us/pkg/logger"
)

type OrderService struct {
	storage storage.IStorage
	logger  logger.ILogger
}


func NewOrderService(storage storage.IStorage,logger logger.ILogger) OrderService {
	return OrderService{
		storage: storage,
		logger:  logger,
	}
}
func (u OrderService) Create(ctx context.Context, Order models.CreateOrder) (string, error) {

	pKey, err := u.storage.Order().CreateOrder(ctx, Order)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating Order", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (u OrderService) Update(ctx context.Context, Order models.GetOrder) (string, error) {

	pKey, err := u.storage.Order().UpdateOrder(ctx, Order)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating Order", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (u OrderService) Delete(ctx context.Context,id string) error {

	err := u.storage.Order().DeleteOrder(ctx, id)
	if err != nil {
		u.logger.Error("error service delete Order", logger.Error(err))
		return err
	}

	return nil
}

func (u OrderService) GetAllOrders(ctx context.Context, Order models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error) {

	pKey, err := u.storage.Order().GetAll(ctx, Order)
	if err != nil {
		u.logger.Error("ERROR in service layer while getalling Order",logger.Error(err))
	 return models.GetAllOrdersResponse{}, err
	}
   
	return pKey, nil
   }

   func (u OrderService) GetByIDOrder(ctx context.Context,Id string) (models.GetOrder, error) {
	pKey, err := u.storage.Order().GetOne(ctx,Id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyID Order", logger.Error(err))
	 return models.GetOrder{}, err
	}
   
	return pKey, nil
   }



   func (u OrderService) UpdateStatus(ctx context.Context, Order models.GetOrder) (string, error) {

	pKey, err := u.storage.Order().UpdateOrderStatus(ctx, Order)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating Order", logger.Error(err))
		return "", err
	}

	return pKey, nil
}
