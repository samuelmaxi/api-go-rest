package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelmaxi/api-go-rest/controllers"
)

func HandlerRequest() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		controllers.Home(ctx)

	})

	r.GET("/api/personalidades", func(ctx *gin.Context) {
		controllers.AllPersonalities(ctx)
	})

	r.GET("/api/personalidades/:id", func(ctx *gin.Context) {
		controllers.ReturnPersonalitie(ctx)

	})

	r.Run(":8000")
}
