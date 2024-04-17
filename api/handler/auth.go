package handler

import (
	_ "clone/3_exam/api/docs"
	"clone/3_exam/api/models"
	"clone/3_exam/pkg/check"
	"clone/3_exam/pkg/password"
	"strings"

	// "clone/3_exam/storage/postgres"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserLogin godoc
// @Router       /User/login [POST]
// @Summary      User login
// @Description  User login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.UserLoginRequest true "login"
// @Success      201  {object}  models.UserLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) UserLogin(c *gin.Context) {
	loginReq := models.UserLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ",loginReq)

	if err := check.ValidatePassword(loginReq.Password); err != nil {
		handleResponse(c,h.Log,"error while validating  old password,old password: "+loginReq.Password, http.StatusBadRequest, err.Error())
		return
	}
mail:=strings.ToLower(loginReq.Mail)
	loginReq.Mail=mail	
	
	loginResp, err := h.Services.Auth().UserLogin(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "unauthorized", http.StatusUnauthorized, err)
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, loginResp)

}



// UserRegister godoc
// @Router       /User/register [POST]
// @Summary      User register
// @Description  User register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        register body models.UserRegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) UserRegister(c *gin.Context) {
	loginReq := models.UserRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}

	mail:=strings.ToLower(loginReq.Mail)
	loginReq.Mail=mail

	if err := check.CheckEmail(loginReq.Mail); err != nil {
		handleResponse(c,h.Log,"Email address does not exist or is undeliverable "+loginReq.Mail, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Services.Auth().UserRegister(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.Log, "Otp sent successfull", http.StatusOK, "")
}





// UserCreateRegister godoc
// @Router       /User/auth/create [POST]
// @Summary      User password check and create 
// @Description  User password check and create
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.LoginUser true "login"
// @Success      201  {object}  models.LoginUser
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) UserRegisterCreateConfirm(c *gin.Context) {
	loginReq := models.LoginUser{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	if err := check.CheckEmail(loginReq.Mail); err != nil {
		handleResponse(c,h.Log,"Email address does not exist or is undeliverable "+loginReq.Mail, http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidatePassword(loginReq.Password); err != nil {
		handleResponse(c,h.Log,"error while validating  password, password: "+loginReq.Password, http.StatusBadRequest, err.Error())
		return
	}

	HashPassword,err:=password.HashPassword(loginReq.Password)
	if err!=nil{
		handleResponse(c, h.Log, "password hashed error", http.StatusUnauthorized, err)
	}

	loginReq.Password=HashPassword

	loginResp, err := h.Services.Auth().UserRegisterCreateConfirm(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "erorororor", http.StatusUnauthorized, err)
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, loginResp)

}



// UserRegister godoc
// @Router       /User/loginotp [POST]
// @Summary      User login otp1
// @Description  User login otp1
// @Tags         auth_otp
// @Accept       json
// @Produce      json
// @Param        register body models.UserRegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) UserLoginOtp(c *gin.Context) {
	loginReq := models.UserRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	
	mail:=strings.ToLower(loginReq.Mail)
	loginReq.Mail=mail

	if err := check.CheckEmail(loginReq.Mail); err != nil {
		handleResponse(c,h.Log,"Email address does not exist or is undeliverable "+loginReq.Mail, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Services.Auth().UserLoginOtp(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "error", http.StatusBadRequest, err)
		return 
	}

	handleResponse(c, h.Log, "Otp sent successfull", http.StatusOK, "")
}



// UserCreateRegister godoc
// @Router       /User/auth/loginotp [POST]
// @Summary      User otp code  check and login
// @Description  User otp code check and login
// @Tags         auth_otp
// @Accept       json
// @Produce      json
// @Param        login body models.UserLoginOTP true "login"
// @Success      201  {object}  models.UserLoginOTP
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) UserLoginOtp2(c *gin.Context) {
	loginReq := models.UserLoginOTP{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	if err := check.CheckEmail(loginReq.Mail); err != nil {
		handleResponse(c,h.Log,"Email address does not exist or is undeliverable "+loginReq.Mail, http.StatusBadRequest, err.Error())
		return
	}
	
	loginResp, err := h.Services.Auth().UserLoginOtp2(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "erorororor", http.StatusUnauthorized, err)
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, loginResp)

}




// UserRegister godoc
// @Router       /User/Forgetpassword [POST]
// @Summary      User Forgetpassword
// @Description  User Forgetpassword
// @Tags         Forgetpassword
// @Accept       json
// @Produce      json
// @Param        register body models.UserRegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) Forgetpassword(c *gin.Context) {
	loginReq := models.UserRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	
	mail:=strings.ToLower(loginReq.Mail)
	loginReq.Mail=mail

	if err := check.CheckEmail(loginReq.Mail); err != nil {
		handleResponse(c,h.Log,"Email address does not exist or is undeliverable "+loginReq.Mail, http.StatusBadRequest, err.Error())
		return
	}
	err := h.Services.Auth().UserLoginOtp(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "error", http.StatusBadRequest, err)
		return 
	}

	handleResponse(c, h.Log, "Otp sent successfull", http.StatusOK, "")
}


// UserCreateRegister godoc
// @Router       /User/Forgetpassword2 [POST]
// @Summary      User Forgetpassword 2
// @Description  User Forgetpassword 2
// @Tags         Forgetpassword
// @Accept       json
// @Produce      json
// @Param        login body models.Forgetpassword2 true "login"
// @Success      201  {object}  models.Forgetpassword2
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) Forgetpassword2(c *gin.Context) {
	loginReq := models.Forgetpassword2{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	if err := check.CheckEmail(loginReq.Mail); err != nil {
		handleResponse(c,h.Log,"Email address does not exist or is undeliverable "+loginReq.Mail, http.StatusBadRequest, err.Error())
		return
	}
	
	loginResp, err := h.Services.Auth().Forgetpassword2(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "erorororor", http.StatusUnauthorized, err)
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, loginResp)

}


