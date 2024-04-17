package handler

import (
	_ "clone/rent_car_us/api/docs"
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/pkg/check"	
	// "clone/rent_car_us/pkg/password"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// CreateCustomer godoc
// @Router 		/customer [POST]
// @Summary 	create a customer
// @Description This api is creates a new customer and returns it's id
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		customer body models.CreateCustomer true "customer"
// @Success		200  {object}  models.CreateCustomer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCustomer(c *gin.Context) {
	customer := models.CreateCustomer{}

	if err := c.ShouldBindJSON(&customer); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidatePassword(customer.Password); err != nil {
		handleResponse(c,h.Log,"error while validating  password, password: "+customer.Password, http.StatusBadRequest, err.Error())
		return
	}


	id, err := h.Services.Customer().Create(context.Background(),customer)
	if err != nil {
		handleResponse(c,  h.Log,"error while creating Customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// Updatecustomer godoc
// @Router 		/customer/{id} [PUT]
// @Summary 	update a customer
// @Description This api is update a  customer and returns it's id
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Param		customer body models.GetCustomer true "customer"
// @Success		200  {object}  models.GetCustomer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateCustomer(c *gin.Context) {
	Customer := models.GetCustomer{}

	if err := c.ShouldBindJSON(&Customer); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateEmail(Customer.Gmail); err != nil {
		handleResponse(c,  h.Log,"error while validating Customer Gmail, Gmail: "+Customer.Gmail, http.StatusBadRequest,err.Error())
		return
	}

	if err := check.ValidatePhone(Customer.Phone); err != nil {
		handleResponse(c, h.Log, "error while validating Customer Phone, Phone"+Customer.Phone, http.StatusBadRequest,err.Error())
		return
	}
	Customer.Id = c.Param("id")

	err := uuid.Validate(Customer.Id)
	if err != nil {
		handleResponse(c,  h.Log,"error while validating Customer id,id: "+Customer.Id, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.Customer().Update(context.Background(),Customer)
	if err != nil {
		handleResponse(c, h.Log, "error while updating Customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GETALLCustomerS godoc
// @Router 		/customer [GET]
// @Summary 	Get customer list
// @Description Get customer list
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		page path string false "page"
// @Param		limit path string false "limit"
// @Param		search path string false "search"
// @Success		200  {object}  models.GetCustomer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllCustomers(c *gin.Context) {
	var (
		request = models.GetAllCustomersRequest{}
	)

	request.Search = c.Param("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit
	Customers, err := h.Services.Customer().GetAllCustomers(context.Background(),request)
	if err != nil {
		handleResponse(c, h.Log, "error while gettign Customers", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, h.Log, "", http.StatusOK, Customers)
}

// @Security ApiKeyAuth
// Deletecustomer godoc
// @Router 		/customer/{id} [DELETE]
// @Summary 	delete a customer
// @Description This api is delete a  customer and returns it's id
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteCustomer(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.Customer().Delete(context.Background(),id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting Customer", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c,  h.Log,"", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GETBYIDCustomer godoc
// @Router 		/customer/{id} [GET]
// @Summary 	Get customer 
// @Description Get customer
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.GetCustomer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetByIDCustomer(c *gin.Context) {
 
	id := c.Param("id")
	fmt.Println("id: ", id)
   
	admin, err := h.Services.Customer().GetByIDCustomer(context.Background(),id)
	if err != nil {
	 handleResponse(c, h.Log, "error while getting admin by id", http.StatusInternalServerError, err)
	 return
	}
	handleResponse(c, h.Log, "", http.StatusOK, admin)
   }


// @Security ApiKeyAuth
// GETALLCustomerS godoc
// @Router 		/customercars [GET]
// @Summary 	Get user list
// @Description Get user list
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		Id path string false "Id"
// @Param		page path string false "page"
// @Param		limit path string false "limit"
// @Param		search path string false "search"
// @Success		200  {object}  models.GetAllCustomerCarsRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllCustomerCars(c *gin.Context) {
	var (
		request = models.GetAllCustomerCarsRequest{}
	)

	request.Search = c.Query("search")
	request.Id=c.Param("id")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c,  h.Log,"error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit
	Orders, err := h.Services.Customer().GetAllCustomerCars(context.Background(),request)
	if err != nil {
		handleResponse(c, h.Log, "error while gettign CustomerCars", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, h.Log, "", http.StatusOK, Orders)
}



// @Security ApiKeyAuth
// UpdatePassword godoc
// @Router 		/customer/password [PATCH]
// @Summary 	update password
// @Description This api is update password
// @Tags 		Password
// @Accept		json
// @Produce		json
// @Param		customer body models.PasswordCustomer true "customer"
// @Success		200  {object}  models.PasswordCustomer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdatePassword(c *gin.Context) {
	customer := models.PasswordCustomer{}

	if err := c.ShouldBindJSON(&customer); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if customer.NewPassword == customer.OldPassword {
		h.Log.Error("new and old passwords are the same, please change the new password")
		handleResponse(c, h.Log, "new and old passwords are the same, please change the new password "+customer.NewPassword, http.StatusBadRequest, errors.New("change new password"))
		return
	}

	if err := check.ValidatePhone(customer.Phone); err != nil {
		handleResponse(c, h.Log, "error while validating phone, phone: "+customer.Phone, http.StatusBadRequest,err.Error())
		return
	}

	if err := check.ValidatePassword(customer.NewPassword); err != nil {
		handleResponse(c,h.Log,"error while validating  new password, new password: "+customer.NewPassword, http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidatePassword(customer.OldPassword); err != nil {
		handleResponse(c,h.Log,"error while validating  old password,old password: "+customer.OldPassword, http.StatusBadRequest, err.Error())
		return
	}


	id, err := h.Services.Customer().UpdatePassword(context.Background(),customer)
	if err != nil {
		handleResponse(c, h.Log, "error while updating Customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}


