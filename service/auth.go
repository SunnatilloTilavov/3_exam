package service

import (
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/pkg"
	"clone/rent_car_us/pkg/jwt"
	"clone/rent_car_us/pkg/logger"
	"clone/rent_car_us/pkg/password"
	"clone/rent_car_us/pkg/smtp"
	"clone/rent_car_us/storage"
	"context"
	"errors"
	"fmt"
	"time"
)

type authService struct {
	storage storage.IStorage
	log     logger.ILogger
	redis   storage.IRedisStorage
}

func NewAuthService(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) authService {
	return authService{
		storage: storage,
		log:     log,
		redis:   redis,
	}
}

func (a authService) CustomerLogin(ctx context.Context, loginRequest models.CustomerLoginRequest) (models.CustomerLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Login)
	customer, err := a.storage.Customer().GetByLogin(ctx, loginRequest.Login)
	if err != nil {
		a.log.Error("error while getting customer credentials by login", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	if err = password.CompareHashAndPassword(customer.Password, loginRequest.Password); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = customer.Id
	m["user_role"] = "customer"

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for customer login", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	return models.CustomerLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}



func (a authService) CustomerRegister(ctx context.Context, loginRequest models.CustomerRegisterRequest) error {

	_, err := a.storage.Customer().GetGmail(ctx, loginRequest.Mail)
	if err != nil {
		a.log.Error("gmail address is already registered", logger.Error(err))
		return  err
	}

	otpCode := pkg.GenerateOTP()

	msg := fmt.Sprintf("Your otp code is: %v, for registering RENT_CAR. Don't give it to anyone", otpCode)

	fmt.Printf("Your otp code is: %v, for registering RENT_CAR. Don't give it to anyone", otpCode)
	
	fmt.Println(loginRequest.Mail,"----------",otpCode)

	err = a.redis.SetX(ctx, loginRequest.Mail, otpCode, time.Minute*2)
	if err != nil {
		a.log.Error("error while setting otpCode to redis customer register", logger.Error(err))
		return err
	}

	err = smtp.SendMail(loginRequest.Mail, msg)
	if err != nil {
		a.log.Error("error while sending otp code to customer register", logger.Error(err))
		return err
	}


	return nil
}



func (u authService) CustomerRegisterCreate(ctx context.Context, customer models.LoginCustomer) (models.CustomerLoginResponse, error) {
	
	resp := models.CustomerLoginResponse{}


OTPCODE:=u.storage.Redis().Get(context.Background(),customer.Gmail)
OTPCODEStr, ok := OTPCODE.(string)
	if !ok {
		u.log.Error("ERROR in service layer while creating car", logger.Error(errors.New("Failed to convert OTP code to string")))
		return resp, errors.New("Failed to convert OTP code to string")
	}

if OTPCODEStr!=customer.GmailCode{
	u.log.Error("ERROR in service layer while creating car", logger.Error(errors.New("The code you entered is not the same as the code sent to your gmail address")))
	return resp,errors.New("The code you entered is not the same as the code sent to your gmail address")
}

id, err := u.storage.Customer().CustomerRegisterCreate(ctx,customer)
if err != nil {
	u.log.Error("ERROR in service layer while creating car", logger.Error(err))
	return resp, err
}
	var m = make(map[interface{}]interface{})

	m["user_id"] = id
	m["user_role"] = "customer"

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		u.log.Error("error while generating tokens for customer register confirm", logger.Error(err))
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}
