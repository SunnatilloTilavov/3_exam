package api

import (
	"clone/rent_car_us/api/handler"
	"clone/rent_car_us/service"
	// "errors"
	// "net/http"
	"clone/rent_car_us/pkg/logger"
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

	// r.Use(authMiddleware)
	r.POST("/customer/login", h.CustomerLogin)
	r.POST("/customer/register", h.CustomerRegister)
	r.POST("/customer/auth/create", h.CustomerRegisterCreate)


	r.POST("/car", h.CreateCar)
	r.GET("/car/:id", h.GetByIDCar)
	r.GET("/car", h.GetAllCars)
	r.GET("/car/free", h.GetAllCarsFree)


	r.PUT("/car/:id", h.UpdateCar)
	r.DELETE("/car/:id", h.DeleteCar)

	r.PATCH("/customer/password", h.UpdatePassword)

	r.POST("/customer", h.CreateCustomer)
	r.GET("/customer/:id", h.GetByIDCustomer)
	r.GET("/customer", h.GetAllCustomers)
	r.PUT("/customer/:id", h.UpdateCustomer)
	r.DELETE("/customer/:id", h.DeleteCustomer)
	r.GET("/customercars",h.GetAllCustomerCars)

	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetOne)
	r.GET("/order", h.GetAllOrders)
	r.PUT("/order/:id", h.UpdateOrder)
	r.DELETE("/order/:id", h.DeleteOrder)
	r.PATCH("/order/status/:id", h.UpdateOrderStatus)



	return r
}

// func authMiddleware(c *gin.Context) {
// 	auth := c.GetHeader("Authorization")
// 	if auth == "" {
// 		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
// 	}
// 	c.Next()
// }




