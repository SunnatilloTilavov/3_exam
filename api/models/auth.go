package models

type UserLoginRequest struct {
	Mail    string `json:"mail"`
	Password string `json:"password"`
}

type UserLoginOTP struct {
	Mail    string `json:"mail"`
	Optcode string `json:"otp_code"`
}


type Forgetpassword2 struct {
	Mail    string `json:"mail"`
	Optcode string `json:"otp_code"`
	Password string `json:"password"`

}


type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthInfo struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
}

type UserRegisterRequest struct {
	Mail string `json:"mail"`
}



type LoginUser struct {
	Mail       string  `json:"Mail"`
	First_name        string  `json:"first_name"`
	Last_name       string  `json:"last_name"`
	Phone     string  `json:"phone"`
	Password string `json:"password"`
	Sex string `json:"sex"`
	Active bool `json:"active"`
	MailCode string `json:"mailcode"`
}
