package service

import (
	"clone/3_exam/api/models"
	"clone/3_exam/storage"
	"context"
	"clone/3_exam/pkg/logger"
	// "encoding/json"
)

type UserService struct {
	storage storage.IStorage
	logger  logger.ILogger
	redis   storage.IRedisStorage
}


func NewUserService(storage storage.IStorage,logger logger.ILogger,redis storage.IRedisStorage) UserService {
	return UserService{
		storage: storage,
		logger:  logger,
		redis:   redis,
	}
}
func (u UserService) Create(ctx context.Context, User models.CreateUser) (string, error) {

	pKey, err := u.storage.User().Create(ctx, User)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating User", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (u UserService) Update(ctx context.Context, User models.UpdateUser) (string, error) {

	pKey, err := u.storage.User().Update(ctx, User)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating User", logger.Error(err))
		return "", err
	}
	
	err = u.redis.Del(ctx, User.Id)
	if err != nil {
		u.logger.Error("error while setting otpCode to redis User update", logger.Error(err))
		return "error redis update",err
	}

	return pKey, nil
}


func (u UserService) UpdateStatus(ctx context.Context, User models.UpdateStatus) (string, error) {

	pKey, err := u.storage.User().UpdateStatus(ctx, User)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating User", logger.Error(err))
		return "", err
	}
	
	err = u.redis.Del(ctx, User.Id)
	if err != nil {
		u.logger.Error("error while setting otpCode to redis User update", logger.Error(err))
		return "error redis update",err
	}

	return pKey, nil
}

func (u UserService) Delete(ctx context.Context, id string) error {

	err := u.storage.User().Delete(ctx, id)
	if err != nil {
		u.logger.Error("error service delete User", logger.Error(err))
		return err
	}

	err = u.redis.Del(ctx, id)
	if err != nil {
		u.logger.Error("error while setting otpCode to redis User deleted", logger.Error(err))
		return err
	}

	return nil
}

func (u UserService) GetAllUsers(ctx context.Context, User models.GetAllUsersRequest) (models.GetAllUsersResponse, error) {

	pKey, err := u.storage.User().GetAllUsers(ctx, User)
	if err != nil {
		u.logger.Error("ERROR in service layer while getalling User", logger.Error(err))
		return models.GetAllUsersResponse{}, err
	}

	return pKey, nil
}

//////////

func (u UserService) GetByIDUser(ctx context.Context, Id string) (models.GetAllUser, error) {
	

	pKey, err := u.storage.User().GetByIDUser(ctx, Id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyID User",logger.Error(err))
		return models.GetAllUser{}, err
	}

	return pKey, nil
}


func (u UserService) UpdatePassword(ctx context.Context, User models.PasswordUser) (string, error) {

	pKey, err := u.storage.User().UpdatePassword(ctx, User)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating User", logger.Error(err))
		return "", err
	}

	return pKey, nil
}


func (u UserService) GetPassword(ctx context.Context, phone string) (string, error) {
	pKey, err := u.storage.User().GetPassword(ctx, phone)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyID User",logger.Error(err))
		return "Error", err
	}

	return pKey, nil
}




