package service

import (
	"context"
	"fmt"
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/storage"
	"clone/rent_car_us/pkg/logger"
)

type carService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewCarService(storage storage.IStorage, logger logger.ILogger) carService {
	return carService{
		storage: storage,
		logger:  logger,
	}
}
func (u carService) Create(ctx context.Context, car models.Car) (string, error) {

	pKey, err := u.storage.Car().Create(ctx, car)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating car", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (u carService) Update(ctx context.Context, car models.Car) (string, error) {

	pKey, err := u.storage.Car().Update(ctx, car)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating car", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (u carService) Delete(ctx context.Context,id string) error {

	err := u.storage.Car().Delete(ctx, id)
	if err != nil {
		u.logger.Error("error service delete car", logger.Error(err))
		return err
	}

	return nil
}

func (u carService) GetAllCars(ctx context.Context, car models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {

	pKey, err := u.storage.Car().GetAllCars(ctx, car)
	if err != nil {
	 fmt.Println("ERROR in service layer while getalling car", err.Error())
	 return models.GetAllCarsResponse{}, err
	}
   
	return pKey, nil
   }

   func (u carService) GetByIDCar(ctx context.Context,Id string) (models.Car, error) {
	pKey, err := u.storage.Car().GetByID(ctx,Id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyID car", logger.Error(err))
	 return models.Car{}, err
	}
   
	return pKey, nil
   }



func (u carService) GetAllCarsFree(ctx context.Context, car models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {

	pKey, err := u.storage.Car().GetAllCarsFree(ctx, car)
	if err != nil {
		u.logger.Error("ERROR in service layer while getalling car", logger.Error(err))
	 return models.GetAllCarsResponse{}, err
	}
   
	return pKey, nil
   }
