package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "assignment2/docs"
	"assignment2/handler"
	"assignment2/model"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7000
// @BasePath /
// @schemes http
func main() {
	var port = ":7000"
	var ord []*model.Orders
	var ite []*model.Items

	service := handler.NewOrderService(ord, ite)

	router := gin.Default()

	p := router.Group("/go")
	{
		p.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
	}

	router.POST("/orders", service.CreateOrder)
	router.GET("/orders", service.GetOrders)
	router.PUT("/order/:id", service.UpdateOrder)
	router.DELETE("/order/:id", service.DeleteOrder)

	url := ginSwagger.URL("http://localhost:7000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(port)
	fmt.Println("Server running at port", port)
}
