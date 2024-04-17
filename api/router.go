package api

import (
	"clone/3_exam/api/handler"
	"clone/3_exam/service"
	// "errors"
	// "net/http"
	"clone/3_exam/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(services service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(services,log)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	r.POST("/User/login", h.UserLogin)
	r.POST("/User/register", h.UserRegister)
	r.POST("/User/auth/create", h.UserRegisterCreateConfirm)
	r.POST("/User/auth/loginotp", h.UserLoginOtp2)
	r.POST("/User/loginotp", h.UserLoginOtp)
	r.POST("/User/Forgetpassword", h.Forgetpassword)
	r.POST("/User/Forgetpassword2", h.Forgetpassword2)
	
	r.PATCH("/User/status/update/:id", h.UpdateStatus)

	r.PATCH("/User/password", h.UpdatePassword)

	r.POST("/User", h.CreateUser)
	r.GET("/User/:id", h.GetByIDUser)
	r.GET("/User", h.GetAllUsers)
	r.PUT("/User/:id", h.UpdateUser)
	r.DELETE("/User/:id", h.DeleteUser)



	return r
}





