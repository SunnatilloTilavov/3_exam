package handler

import (
	"clone/rent_car_us/api/models"
	// "clone/rent_car_us/pkg/check"
	_ "clone/rent_car_us/api/docs"
	"context"
	"fmt"
	"net/http"
	// "strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// CreateCar godoc
// @Router 		/car [POST]
// @Summary 	create a car
// @Description This api is creates a new car and returns it's id
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		car body models.Car true "car"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCar(c *gin.Context) {
	car := models.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		handleResponse(c, h.Log,"error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	


	id, err := h.Services.Car().Create(context.Background(),car)
	if err != nil {
		handleResponse(c,h.Log, "error while creating car", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c,h.Log, "Created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// Updatecar godoc
// @Router 		/car/{id} [PUT]
// @Summary 	update a car
// @Description This api is update a  car and returns it's id
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Param		car body models.Car true "car"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateCar(c *gin.Context) {
	car := models.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		handleResponse(c, h.Log,"error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	car.Id = c.Param("id")

	err := uuid.Validate(car.Id)
	if err != nil {
		handleResponse(c, h.Log, "error while validating car id,id: "+car.Id, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.Car().Update(context.Background(),car)
	if err != nil {
		handleResponse(c, h.Log, "error while updating car", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GETALLCARS godoc
// @Router 		/car [GET]
// @Summary 	Get user list
// @Description Get user list
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		page path string false "page"
// @Param		limit path string false "limit"
// @Param		search path string false "search"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllCars(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)

	request.Search = c.Query("search")

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
	cars, err := h.Services.Car().GetAllCars(context.Background(),request)
	if err != nil {
		handleResponse(c, h.Log, "error while gettign cars", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, h.Log, "", http.StatusOK, cars)
}

// @Security ApiKeyAuth
// Deletecar godoc
// @Router 		/car/{id} [DELETE]
// @Summary 	delete a car
// @Description This api is delete a  car and returns it's id
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteCar(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.Car().Delete(context.Background(),id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting car", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GETBYIDcar godoc
// @Router 		/car{id}  [GET]
// @Summary 	Get user 
// @Description Get user
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetByIDCar(c *gin.Context) {
 
	id := c.Param("id")
	fmt.Println("id: ", id)
   
	admin, err := h.Services.Car().GetByIDCar(context.Background(),id)
	if err != nil {
	 handleResponse(c, h.Log, "error while getting admin by id", http.StatusInternalServerError, err)
	 return
	}
	handleResponse(c, h.Log, "", http.StatusOK, admin)
   }


// @Security ApiKeyAuth
// GETALLCARSFREE godoc
// @Router 		/car/free [GET]
// @Summary 	Get user list
// @Description Get user list
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		page path string false "page"
// @Param		limit path string false "limit"
// @Param		search path string false "search"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllCarsFree(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)
	request.Search = c.Query("search")

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
	cars, err := h.Services.Car().GetAllCarsFree(context.Background(),request)
	if err != nil {
		handleResponse(c, h.Log, "error while gettign cars", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c,  h.Log,"", http.StatusOK, cars)
}