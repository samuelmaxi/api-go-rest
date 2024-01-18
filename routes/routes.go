package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/samuelmaxi/api-go-rest/controllers"
	"github.com/samuelmaxi/api-go-rest/middleware"
)

func HandlerRequest() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))
	r.Use(middleware.ContentTypeMiddleware())

	basePath := "api/v1"
	v1 := r.Group(basePath)
	{
		r.GET("/", controllers.Home)
		v1.GET("/personalidades", controllers.AllPersonalities)
		v1.GET("/personalidades/:id", controllers.ReturnPersonalitie)
		v1.POST("/personalidades", controllers.CreateNewPersonalitie)
		v1.DELETE("/personalidades/:id", controllers.DeletePersonalitie)
		v1.PUT("/personalidades/:id", controllers.EditPersonalitie)
	}

	r.Run(":8000")
}
