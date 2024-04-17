package handler

import (
	_ "clone/3_exam/api/docs"
	"clone/3_exam/api/models"
	"clone/3_exam/pkg/check"	
	"clone/3_exam/pkg/password"
	// "clone/3_exam/pkg/password"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// CreateUser godoc
// @Router 		/User [POST]
// @Summary 	create a User
// @Description This api is creates a new User and returns it's id
// @Tags 		User
// @Accept		json
// @Produce		json
// @Param		User body models.CreateUser true "User"
// @Success		200  {object}  models.CreateUser
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateUser(c *gin.Context) {
	User := models.CreateUser{}

	if err := c.ShouldBindJSON(&User); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	mail:=strings.ToLower(User.Mail)
	User.Mail=mail

	if err := check.ValidatePassword(User.Password); err != nil {
		handleResponse(c,h.Log,"error while validating  password, password: "+User.Password, http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateEmail(User.Mail); err != nil {
		handleResponse(c,h.Log,"error while validating  password, password: "+User.Mail, http.StatusBadRequest, err.Error())
		return
	}

	HashPassword,err:=password.HashPassword(User.Password)
	if err!=nil{
		handleResponse(c, h.Log, "password hashed error", http.StatusUnauthorized, err)
	}

	User.Password=HashPassword


	id, err := h.Services.User().Create(context.Background(),User)
	if err != nil {
		handleResponse(c,  h.Log,"error while creating User", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// UpdateUser godoc
// @Router 		/User/{id} [PUT]
// @Summary 	update a User
// @Description This api is update a  User and returns it's id
// @Tags 		User
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Param		User body models.UpdateUser true "User"
// @Success		200  {object}  models.UpdateUser
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateUser(c *gin.Context) {
	User := models.UpdateUser{}

	if err := c.ShouldBindJSON(&User); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateEmail(User.Mail); err != nil {
		handleResponse(c,  h.Log,"error while validating User Mail, Mail: "+User.Mail, http.StatusBadRequest,err.Error())
		return
	}

	mail:=strings.ToLower(User.Mail)
	User.Mail=mail

	if err := check.ValidatePhone(User.Phone); err != nil {
		handleResponse(c, h.Log, "error while validating User Phone, Phone"+User.Phone, http.StatusBadRequest,err.Error())
		return
	}
	User.Id = c.Param("id")

	err := uuid.Validate(User.Id)
	if err != nil {
		handleResponse(c,  h.Log,"error while validating User id,id: "+User.Id, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.User().Update(context.Background(),User)
	if err != nil {
		handleResponse(c, h.Log, "error while updating User", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GETALLUserS godoc
// @Router 		/User [GET]
// @Summary 	Get User list
// @Description Get User list
// @Tags 		User
// @Accept		json
// @Produce		json
// @Param		page path string false "page"
// @Param		limit path string false "limit"
// @Param		search path string false "search"
// @Success		200  {object}  models.GetAllUser
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllUsers(c *gin.Context) {
	var (
		request = models.GetAllUsersRequest{}
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
	request.Page = page
	request.Limit = limit
	Users, err := h.Services.User().GetAllUsers(context.Background(),request)
	if err != nil {
		handleResponse(c, h.Log, "error while gettign Users", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, h.Log, "", http.StatusOK, Users)
}

// @Security ApiKeyAuth
// DeleteUser godoc
// @Router 		/User/{id} [DELETE]
// @Summary 	delete a User
// @Description This api is delete a  User and returns it's id
// @Tags 		User
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.User().Delete(context.Background(),id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting User", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c,  h.Log,"", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GETBYIDUser godoc
// @Router 		/User/{id} [GET]
// @Summary 	Get User 
// @Description Get User
// @Tags 		User
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.GetAllUser
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetByIDUser(c *gin.Context) {
 
	id := c.Param("id")
	fmt.Println("id: ", id)
   
	admin, err := h.Services.User().GetByIDUser(context.Background(),id)
	if err != nil {
	 handleResponse(c, h.Log, "error while getting admin by id", http.StatusInternalServerError, err)
	 return
	}
	handleResponse(c, h.Log, "", http.StatusOK, admin)
   }




// @Security ApiKeyAuth
// UpdatePassword godoc
// @Router 		/User/password [PATCH]
// @Summary 	update password
// @Description This api is update password
// @Tags 		Password
// @Accept		json
// @Produce		json
// @Param		User body models.PasswordUser true "User"
// @Success		200  {object}  models.PasswordUser
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdatePassword(c *gin.Context) {
	User := models.PasswordUser{}

	if err := c.ShouldBindJSON(&User); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if User.NewPassword == User.OldPassword {
		h.Log.Error("new and old passwords are the same, please change the new password")
		handleResponse(c, h.Log, "new and old passwords are the same, please change the new password "+User.NewPassword, http.StatusBadRequest, errors.New("change new password"))
		return
	}

	if err := check.ValidateEmail(User.Mail); err != nil {
		handleResponse(c, h.Log, "error while validating mail, mail: "+User.Mail, http.StatusBadRequest,err.Error())
		return
	}

	if err := check.ValidatePassword(User.NewPassword); err != nil {
		handleResponse(c,h.Log,"error while validating  new password, new password: "+User.NewPassword, http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidatePassword(User.OldPassword); err != nil {
		handleResponse(c,h.Log,"error while validating  old password,old password: "+User.OldPassword, http.StatusBadRequest, err.Error())
		return
	}


	id, err := h.Services.User().UpdatePassword(context.Background(),User)
	if err != nil {
		handleResponse(c, h.Log, "error while updating User", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}


// @Security ApiKeyAuth
// UpdateUser godoc
// @Router 		/User/status/update/{id} [PATCH]
// @Summary 	update a User
// @Description This api is update status
// @Tags 		User
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Param		User body models.UpdateStatus true "User"
// @Success		200  {object}  models.UpdateStatus
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateStatus(c *gin.Context) {
	User := models.UpdateStatus{}

	if err := c.ShouldBindJSON(&User); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	User.Id = c.Param("id")

	err := uuid.Validate(User.Id)
	if err != nil {
		handleResponse(c,  h.Log,"error while validating User id,id: "+User.Id, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.User().UpdateStatus(context.Background(),User)
	if err != nil {
		handleResponse(c, h.Log, "error while updating User", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}