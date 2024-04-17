package service

import (
	"clone/3_exam/api/models"
	"clone/3_exam/pkg"
	"clone/3_exam/pkg/jwt"
	"clone/3_exam/pkg/logger"
	"clone/3_exam/pkg/password"
	"clone/3_exam/pkg/smtp"
	"clone/3_exam/storage"
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

func (a authService) UserLogin(ctx context.Context, loginRequest models.UserLoginRequest) (models.UserLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Mail)
	User, err := a.storage.User().GetByLogin(ctx, loginRequest.Mail)
	if err != nil {
		a.log.Error("error while getting User credentials by login", logger.Error(err))
		return models.UserLoginResponse{}, err
	}

	if err = password.CompareHashAndPassword(User.Password, loginRequest.Password); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return models.UserLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = User.Id
	m["user_role"] = "User"

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for User login", logger.Error(err))
		return models.UserLoginResponse{}, err
	}

	return models.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}



func (a authService) UserRegister(ctx context.Context, loginRequest models.UserRegisterRequest) error {

	_, err := a.storage.User().GetGmail(ctx, loginRequest.Mail)
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
		a.log.Error("error while setting otpCode to redis User register", logger.Error(err))
		return err
	}

	err = smtp.SendMail(loginRequest.Mail, msg)
	if err != nil {
		a.log.Error("error while sending otp code to User register", logger.Error(err))
		return err
	}


	return nil
}



func (u authService) UserRegisterCreateConfirm(ctx context.Context, User models.LoginUser) (models.UserLoginResponse, error) {
	
	resp := models.UserLoginResponse{}


OTPCODE:=u.storage.Redis().Get(context.Background(),User.Mail)
OTPCODEStr, ok := OTPCODE.(string)
	if !ok {
		u.log.Error("ERROR in service layer while creating car", logger.Error(errors.New("Failed to convert OTP code to string")))
		return resp, errors.New("Failed to convert OTP code to string")
	}

if OTPCODEStr!=User.MailCode{
	u.log.Error("ERROR in service layer while creating car", logger.Error(errors.New("The code you entered is not the same as the code sent to your gmail address")))
	return resp,errors.New("The code you entered is not the same as the code sent to your gmail address")
}

id, err := u.storage.User().UserRegisterCreateConfirm(ctx,User)
if err != nil {
	u.log.Error("ERROR in service layer while creating car", logger.Error(err))
	return resp, err
}
	var m = make(map[interface{}]interface{})

	m["user_id"] = id
	m["user_role"] = "User"

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		u.log.Error("error while generating tokens for User register confirm", logger.Error(err))
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}


func (a authService) UserLoginOtp(ctx context.Context, loginRequest models.UserRegisterRequest) error {

	_, err := a.storage.User().GetGmail(ctx, loginRequest.Mail)
	if err == nil {
		a.log.Error("gmail address is don't registered", logger.Error(err))
		return  errors.New("gmail address is don't registered")
	}

	otpCode := pkg.GenerateOTP()

	msg := fmt.Sprintf("Your otp code is: %v, for registering RENT_CAR. Don't give it to anyone", otpCode)

	fmt.Printf("Your otp code is: %v, for registering RENT_CAR. Don't give it to anyone", otpCode)
	
	fmt.Println(loginRequest.Mail,"----------",otpCode)

	err = a.redis.SetX(ctx, loginRequest.Mail, otpCode, time.Minute*2)
	if err != nil {
		a.log.Error("error while setting otpCode to redis User register", logger.Error(err))
		return err
	}

	err = smtp.SendMail(loginRequest.Mail, msg)
	if err != nil {
		a.log.Error("error while sending otp code to User register", logger.Error(err))
		return err
	}


	return nil
}




func (u authService) UserLoginOtp2(ctx context.Context, User models.UserLoginOTP) (models.UserLoginResponse, error) {
	
	resp := models.UserLoginResponse{}


OTPCODE:=u.storage.Redis().Get(context.Background(),User.Mail)
OTPCODEStr, ok := OTPCODE.(string)
	if !ok {
		u.log.Error("ERROR in service layer while login user", logger.Error(errors.New("Failed to convert OTP code to string")))
		return resp, errors.New("Failed to convert OTP code to string")
	}

if OTPCODEStr!=User.Optcode{
	u.log.Error("ERROR in service layer while login user", logger.Error(errors.New("The code you entered is not the same as the code sent to your gmail address")))
	return resp,errors.New("The code you entered is not the same as the code sent to your gmail address")
}

id, err := u.storage.User().GetGmail(ctx,User.Mail)
if err == nil {
	u.log.Error("ERROR in service layer while login", logger.Error(err))
	return resp, err
}
	var m = make(map[interface{}]interface{})

	m["user_id"] = id
	m["user_role"] = "User"

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		u.log.Error("error while generating tokens for User register confirm", logger.Error(err))
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}


func (u authService) Forgetpassword2(ctx context.Context, User models.Forgetpassword2) (id string,err error) {
	
	resp := models.Forgetpassword2{}


OTPCODE:=u.storage.Redis().Get(context.Background(),User.Mail)
OTPCODEStr, ok := OTPCODE.(string)
	if !ok {
		u.log.Error("ERROR in service layer while login user", logger.Error(errors.New("Failed to convert OTP code to string")))
		return id, errors.New("Failed to convert OTP code to string")
	}

if OTPCODEStr!=User.Optcode{
	u.log.Error("ERROR in service layer while login user", logger.Error(errors.New("The code you entered is not the same as the code sent to your gmail address")))
	return id,errors.New("The code you entered is not the same as the code sent to your gmail address")
}

id, err = u.storage.User().UpdatePasswordForget(ctx,resp)
if err == nil {
	u.log.Error("ERROR in service layer while login", logger.Error(err))
	return id, err
}
	

	return id, nil
}

