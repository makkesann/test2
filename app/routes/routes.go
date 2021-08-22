package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/makkesann/test2/app/controller"
)

func GetRouter() *gin.Engine { // *gin.Engineの表記は返り値の型
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	router.GET("/", controller.IndexDisplayAction)

	return router
}
