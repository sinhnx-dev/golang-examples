package main

import (
	"sinhnx.dev/restfulapi/gonic/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// router := gin.Default()
	// api := router.Group("/api")
	// {
	// 	api.GET("/test", func(ctx *gin.Context) {
	// 		ctx.JSON(200, gin.H{
	// 			"message": "test successful",
	// 		})
	// 	})
	// }
	router := gin.Default()
	controllers.SetupItemsRouter(router)
	controllers.SetupItemRouter(router)
	router.Run(":2011")
}
