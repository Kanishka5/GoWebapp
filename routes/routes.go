package routes

import (
	"github.com/Kanishka5/GoWebapp/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouters sets all the routes
func SetupRouters() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// routes
	router.GET("/", controllers.Home)
	router.GET("/users", controllers.AllUser)
	router.GET("/users/:id", controllers.FindUser)
	router.POST("/users", controllers.CreateUser)
	router.GET("/userform", controllers.UserForm)

	return router
}
