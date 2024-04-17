package models

type CreateCustomer struct {
	Id          string  `json:"id"`
	Login     string  `json:"login"`
	First_name        string  `json:"first_name"`
	Last_name       string  `json:"last_name"`
	Gmail       string  `json:"gmail"`
	Phone     string  `json:"phone"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	Password string `json:"password"`
}

type PasswordCustomer struct{
	Phone     string  `json:"phone"`
	NewPassword string `json:"Newpassword"`
	OldPassword string `json:"Oldpassword"`
}

type GetPassword struct{
	Phone     string  `json:"phone"`
	Password string `json:"password"`
}

type GetCustomer struct {
	Id          string  `json:"id"`
	First_name        string  `json:"first_name"`
	Last_name       string  `json:"last_name"`
	Gmail       string  `json:"gmail"`
	Phone     string  `json:"phone"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}


type GetAllCustomersResponse struct {
	Customer  []GetAllCustomer `json:"customer"`
	Count int64 `json:"count"`
}

type GetAllCustomersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
type GetAllCustomer struct{
	Id          string  `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gmail       string  `json:"gmail"`
	Phone       string  `json:"phone"`
	Is_Blocked  bool     `json:"isblocked"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	OrderCount  int `json:"ordercount"`
	CarsCount   int `json:"carscount"`
	Order       GetOrder    `json:"order"`
	Car         Car          `json:"car"`
	Password        string  `json:"password,omitempty"`
}
type GetAllCustomerCars struct{
	Id          string  `json:"id"`
	Name       string   `json:"name"`
	CreatedAt  string    `json:"creatAt"`
	Amount     float32     `json:"amount"`
}

type GetAllCustomerCarsRequest struct {
    Search string `json:"search"`
	Id     string  `json:"id"`
	Page uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}

type GetAllCustomerCarsResponse struct {
	Customer []GetAllCustomerCars `json:"orders"`
	Count int16 `json:"count"`
}
