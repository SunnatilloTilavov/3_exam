package storage

import (
	"clone/3_exam/api/models"
	"context"
	"time"
)
type IStorage interface {
	CloseDB()
	User() IUserStorage
	Redis() IRedisStorage
}


type IUserStorage interface {

	UserRegisterCreateConfirm(ctx context.Context, User models.LoginUser) (string, error)

	Create(context.Context,models.CreateUser) (string, error)
	GetByIDUser(context.Context,string) (models.GetAllUser, error)
	GetAllUsers(context.Context,models.GetAllUsersRequest) (models.GetAllUsersResponse, error)
	Update(context.Context,models.UpdateUser) (string, error)
	Delete(context.Context,string) error
	GetPassword(ctx context.Context, phone string) (string, error)
	UpdatePassword(context.Context,models.PasswordUser) (string, error)
	GetByLogin(context.Context, string) (models.GetIdPassword, error)
	GetGmail (ctx context.Context, gmail string) (string, error)

	UpdatePasswordForget(ctx context.Context, User models.Forgetpassword2) (string, error)
	UpdateStatus(ctx context.Context, User models.UpdateStatus) (string, error)
}


type IRedisStorage interface {
	SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) interface{}
	Del(ctx context.Context, key string) error
}