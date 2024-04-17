package models

type CreateUser struct {
	Mail       string  `json:"Mail"`
	First_name        string  `json:"first_name"`
	Last_name       string  `json:"last_name"`
	Phone     string  `json:"phone"`
	Password string `json:"password"`
	Sex string `json:"sex"`
	Active bool `json:"active"`
}
type UpdateUser struct {
	Id          string  `json:"id"`
	Mail       string  `json:"Mail"`
	First_name        string  `json:"first_name"`
	Last_name       string  `json:"last_name"`
	Phone     string  `json:"phone"`
	Sex string `json:"sex"`
	Active bool `json:"active"`
}

type UpdateStatus struct {
	Id          string  `json:"id"`
	Active bool `json:"active"`
}



type PasswordUser struct{
	Mail     string  `json:"phone"`
	NewPassword string `json:"Newpassword"`
	OldPassword string `json:"Oldpassword"`
}

type GetPassword struct{
	Mail     string  `json:"phone"`
	Password string `json:"password"`
}


type GetAllUsersResponse struct {
	User  []GetAllUser `json:"User"`
	Count int64 `json:"count"`
}

type GetAllUsersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllUser struct{
	Id          string  `json:"id"`
	Mail       string  `json:"Mail"`
	First_name        string  `json:"first_name"`
	Last_name       string  `json:"last_name"`
	Phone     string  `json:"phone"`
	Sex string `json:"sex"`
	Active bool `json:"active"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type GetIdPassword struct{
	Id          string  `json:"id"`
	Mail       string  `json:"Mail"`
	First_name        string  `json:"first_name"`
	Last_name       string  `json:"last_name"`
	Phone     string  `json:"phone"`
	Sex string `json:"sex"`
	Password string `json:"password"`
	Active bool `json:"active"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

