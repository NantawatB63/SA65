package main

import (
	"github.com/NantawatB63/sa-65-example/controller"
	"github.com/NantawatB63/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		//router.Use(middlewares.Authorizes())
		{
			// Employee Routes
			router.GET("/users", controller.ListUsers)
			router.GET("/user/:id", controller.GetUser)
			router.PATCH("/users", controller.UpdateUser)
			router.DELETE("/users/:id", controller.DeleteUser)

			// Shelf Routes
			router.GET("/shelfs", controller.ListShelf)
			router.GET("/shelf/:id", controller.GetShelf)
			router.POST("/shelfs", controller.CreateShelf)
			router.PATCH("/shelfs", controller.UpdateShelf)
			router.DELETE("/shelfs/:id", controller.DeleteShelf)

			// Lot Routes
			router.GET("/Lots", controller.ListLot)
			router.GET("/lot/:id", controller.GetLot)
			router.POST("/lots", controller.CreateLot)
			router.PATCH("/lots", controller.UpdateLot)
			router.DELETE("/lots/:id", controller.DeleteLot)

			// Stock Routes
			router.GET("/Stocks", controller.ListStocks)
			router.GET("/Stock/:id", controller.GetStock)
			router.POST("/Stocks", controller.CreateStock)
			router.PATCH("/Stocks", controller.UpdateStock)
			router.DELETE("/Stocks/:id", controller.DeleteStock)
		}
	}

	// Signup User Route
	//r.POST("/signup", controller.CreateUser)
	// login User Route
	//r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
