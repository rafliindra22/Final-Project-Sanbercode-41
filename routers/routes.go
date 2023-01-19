package routers

import (
	"Final-Project-Sanbercode-Go-Batch-41/controllers"
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", database.DbConnection)
		c.Next()
	})

	router.POST("/register-admin", controllers.RegisterAdmin)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	router.GET("/phones", controllers.GetAllPhone)
	router.GET("/phones/:id", controllers.GetPhoneById)
	router.GET("/phones/:id/spec-comment", controllers.GetSpecCommentByPhoneID)

	router.GET("/brands", controllers.GetAllBrand)
	router.GET("/brands/:id", controllers.GetBrandByID)
	router.GET("/brands/:id/phones", controllers.GetPhonesByBrandID)

	router.GET("/specs", controllers.GetAllSpec)
	router.GET("/specs/:id", controllers.GetSpecByID)

	router.GET("/comments", controllers.GetAllComment)

	merkMiddlewareroute := router.Group("/brands")
	merkMiddlewareroute.Use(middlewares.AdminMiddleware())
	merkMiddlewareroute.POST("/create", controllers.CreateBrand)
	merkMiddlewareroute.PUT("/update/:id", controllers.UpdateBrand)
	merkMiddlewareroute.DELETE("/delete/:id", controllers.DeleteBrand)

	phoneMiddleware := router.Group("/phone")
	phoneMiddleware.Use(middlewares.AdminMiddleware())
	phoneMiddleware.POST("/create", controllers.CreatePhone)
	phoneMiddleware.PUT("/update/:id", controllers.UpdatePhone)
	phoneMiddleware.DELETE("/delete/:id", controllers.DeletePhone)

	specMiddleware := router.Group("/spec")
	specMiddleware.Use(middlewares.AdminMiddleware())
	specMiddleware.POST("/create", controllers.CreateSpec)
	specMiddleware.PUT("/update/:id", controllers.UpdateSpec)
	specMiddleware.DELETE("/delete/:id", controllers.DeleteSpec)

	commentMiddleware := router.Group("/comment")
	commentMiddleware.Use(middlewares.PublicMiddleware())
	commentMiddleware.POST("/create", controllers.CreateComment)
	commentMiddleware.DELETE("/delete/:id", controllers.DeleteComment)

	router.Run(":" + os.Getenv("PORT"))

	return router
}
