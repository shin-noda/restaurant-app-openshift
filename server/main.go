package main

import (
	"os"

	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// Use Gin framework
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	// These are the endpoints: CRUD operations
	//C: Create
	router.POST("/order/create", routes.AddOrder)
	//R: Read
	router.GET("/waiter/:waiter", routes.GetOrdersByWaiter)
	router.GET("/orders", routes.GetOrders)
	router.GET("/order/:id/", routes.GetOrderById)
	//U: Update
	router.PUT("/waiter/update/:id", routes.UpdateWaiter)
	router.PUT("/order/update/:id", routes.UpdateOrder)
	//D: Delete
	router.DELETE("/order/delete/:id", routes.DeleteOrder)

	//this runs the server and allows it to listen to requests.
	router.Run(":" + port)
}
